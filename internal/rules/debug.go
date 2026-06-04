package rules
import (
	"config_scanner/internal/models"
)

type DebugRule struct{}

func (r *DebugRule) Name() string { return "Debug Mode Enabled" }

func (r *DebugRule) Analyze(config map[string]interface{}) []models.Issue {
	var issues []models.Issue

	if val, ok := config["debug"]; ok {
		if b, ok := val.(bool); ok && b {
			issues = append(issues, models.Issue{
				Severity: models.Low,
				RuleName: r.Name(),
				Description: "Открытый debug mode.",
				Recommendation: "Поменяйте режим на более избирательный (info+).",
			})
		}
	}

	if logSub, ok := config["log"].(map[string]interface{}); ok {
		if level, ok := logSub["level"].(string); ok && level == "debug" {
				issues = append(issues, models.Issue{
				Severity: models.Low,
				RuleName: r.Name(),
				Description: "Логирование в debug-режиме.",
				Recommendation: "Поменяйте режим на более избирательный (info+).",
			})
		}
	}

	return issues
}