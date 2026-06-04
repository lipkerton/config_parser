package main
import (
	"flag"
	"fmt"
	"io"
	"os"
	"config_scanner/internal/analyzer"
	"config_scanner/internal/parsers"
	"config_scanner/internal/rules"
)

func main() {
	silent := flag.Bool("silent", false, "Не выходить с ошибкой при наличии проблем")
	shortSilent := flag.Bool("s", false, "Не выходить с ошибкой при наличии проблем (кратко)")
	stdinOpt := flag.Bool("stdin", false, "Прочитать конфигурацию из стандартного потока ввода")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование: %s [флаги] [путь_к_файлу]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	isSilent := *silent || *shortSilent
	var configData []byte
	var err error
	var parser parsers.Parser

	if *stdinOpt {
		configData, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения из stdin: %v\n", err)
			os.Exit(1)
		}
		parser = &parsers.JSONParser{}
	} else {
		if flag.NArg() < 1 {
			fmt.Println("Ошибка: не указан путь к файлу конфигурации.")
			flag.Usage()
			os.Exit(1)
		}
		filePath := flag.Arg(0)
		configData, err = os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения файла: %v\n", err)
			os.Exit(1)
		}

		parser, err = parsers.GetParserByExtension(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
	}

	parsedConfig, err := parser.Parse(configData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка парсинга конфигурации: %v\n", err)
		os.Exit(1)
	}

	cfgAnalyzer := analyzer.NewAnalyzer()
	cfgAnalyzer.RegisterRule(&rules.DebugRule{})
	cfgAnalyzer.RegisterRule(&rules.AlgoRule{})
	cfgAnalyzer.RegisterRule(&rules.BindRule{})
	cfgAnalyzer.RegisterRule(&rules.PasswordRule{})
	cfgAnalyzer.RegisterRule(&rules.TLSRule{})

	issues := cfgAnalyzer.Run(parsedConfig)

	if len(issues) > 0 {
		fmt.Printf("Найдено проблем: %d\n\n", len(issues))
		for _, issue := range issues {
			fmt.Printf("%s: %s %s\n", issue.Severity, issue.Description, issue.Recommendation)
		}

		if !isSilent {
			os.Exit(1)
		}
	} else {
		fmt.Println("Проблем безопасности не обнаружено.")
	}
}