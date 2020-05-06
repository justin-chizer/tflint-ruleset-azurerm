// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule checks the pattern is valid
type AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule() *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule {
	return &AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule{
		resourceType:  "azurerm_data_factory_linked_service_mysql",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[-\w\._\(\)]+$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule) Name() string {
	return "azurerm_data_factory_linked_service_mysql_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryLinkedServiceMysqlInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._\(\)]+$`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}