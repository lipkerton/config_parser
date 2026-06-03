package models

type Severity string

const (
	Low Severity = "LOW"
	Medium Severity = "MEDIUM"
	High Severity = "HIGH"
)

type Issue struct {
	Severity Severity `json:"severity" grpc:"severity"`
	RuleName string `json:"rule_name" grpc:"rule_name"`
	Description string `json:"description" grpc:"description"`
	Recommendation string `json:"recommendation" grpc:"recommendation"`
}

type Rule interface {
	Name() string
	Analyze(config map[string]interface{}) []Issue
}
