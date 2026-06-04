package analyzer
import "config_scanner/internal/models"

type Analyzer struct {
	rules []models.Rule
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		rules: make([]models.Rule, 0),
	}
}

func (a *Analyzer) RegisterRule(r models.Rule) {
	a.rules = append(a.rules, r)
}

func (a *Analyzer) Run(config map[string]interface{}) []models.Issue {
	var allIssues []models.Issue
	for _, rule := range a.rules {
		issues := rule.Analyze(config)
		if len(issues) > 0 {
			allIssues = append(allIssues, issues...)
		}
	}
	return allIssues
}