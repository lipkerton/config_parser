package rules
import "config_scanner/internal/models"

type BindRule struct{}

func (r *BindRule) Name() string { return "Unrestricted Bind Address" }

func (r *BindRule) Analyze(config map[string]interface{}) []models.Issue {
	var issues []models.Issue

	var checkBind func(cfg map[string]interface{})
	checkBind = func(cfg map[string]interface{}) {
		for _, v := range cfg {
			if strVal, ok := v.(string); ok {
				if strVal == "0.0.0.0" {
					issues = append(issues, models.Issue{
						Severity: models.Medium,
						RuleName: r.Name(),
						Description: "Использование 0.0.0.0 без ограничений.",
						Recommendation: "Укажите конкретный IP-адрес интерфейса (например, 127.0.0.1) для ограничения доступа.",
					})
				}
			} else if nextMap, ok := v.(map[string]interface{}); ok {
				checkBind(nextMap)
			}
		}
	}

	checkBind(config)
	return issues
}