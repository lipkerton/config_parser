package rules
import (
	"config_scanner/internal/models"
	"strings"
)

type PasswordRule struct{}

func (r *PasswordRule) Name() string { return "Plaintext password" }

func (r *PasswordRule) Analyze(config map[string]interface{}) []models.Issue {
	var issues []models.Issue

	var checkPasswords func(cfg map[string]interface{})
	checkPasswords = func(cfg map[string]interface{}) {
		for k, v := range cfg {
			keyLower := strings.ToLower(k)
			if strings.Contains("password") || strings.Contains("pass") || strings.Contains("secret") {
				if strVal, ok := v.(string); ok && strVal != "" {
					issues = append(issues, models.Issue{
						Severity: models.High,
						RuleName: r.Name(),
						Description: "Пароль или секрет (" + k + ") хранится в открытом виде.",
						Recommendation: "Вынесите секрет в переменные окружения или используйте защищенное хранилище (Secret Manager).",
					})
				}
			}

			if nextMap, ok := v.(map[string]interface{}); ok {
				checkPasswords(nextMap)
			}
		}
	}
	checkPasswords(config)
	return issues
}