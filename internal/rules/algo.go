package rules
import (
	"config_scanner/internal/models"
	"strings"
)

type AlgoRule struct{}

func (r *AlgoRule) Name() string { return "Weak Algorithm" }

func (r *AlgoRule) Analyze(config map[string]interface{}) []models.Issue {
	var issues []models.Issue

	var checkWeakAlgo func(cfg map[string]interface{})
	checkWeakAlgo = func(cfg map[string]interface{}) {
		for _, v := range cfg {
			if strVal, ok := v.(string); ok {
				if strings.ToUpper(strVal) == "MD5" || strings.ToUpper(strVal) == "SHA1" {
					issues = append(issues, models.Issue{
						Severity: models.High,
						RuleName: r.Name(),
						Description: "Слишком слабый алгоритм хеширования - " + strVal + ".",
						Recommendation: "Замените его на более безопасный.",
					})
				}
			} else if nextMap, ok := v.(map[string]interface{}); ok {
				checkWeakAlgo(nextMap)
			}
		}
	}

	checkWeakAlgo(config)
	return issues
}
