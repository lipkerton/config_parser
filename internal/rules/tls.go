package rules

import (
	"config_scanner/internal/models"
	"strings"
)

type TLSRule struct{}

func (r *TLSRule) Name() string { return "TLS disabled" }

func (r *TLSRule) Analyze(config map[string]interface{}) []models.Issue {
	var issues []models.Issue

	var checkTLS func(cfg map[string]interface{})
	checkTLS = func(cfg map[string]interface{}) {
		for k, v := range cfg {
			keyLower := strings.ToLower(k)
			
			if strings.Contains(keyLower, "tls") || strings.Contains(keyLower, "ssl") {
				if boolVal, ok := v.(bool); ok && !boolVal {
					issues = append(issues, models.Issue{
						Severity: models.High,
						RuleName: r.Name(),
						Description: "TLS-проверка отключена (" + k + ": false).",
						Recommendation: "Включите TLS для обеспечения безопасного соединения.",
					})
				}
				if strVal, ok := v.(string); ok {
					valLower := strings.ToLower(strVal)
					if valLower == "disabled" || valLower == "false" {
						issues = append(issues, models.Issue{
							Severity: models.High,
							RuleName: r.Name(),
							Description: "TLS-проверка отключена (" + k + ": " + strVal + ").",
							Recommendation: "Включите TLS для обеспечения безопасного соединения.",
						})
					}
				}
			}

			if nextMap, ok := v.(map[string]interface{}); ok {
				checkTLS(nextMap)
			}
		}
	}

	checkTLS(config)
	return issues
}