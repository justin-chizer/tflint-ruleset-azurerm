// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermEventgridDomainInvalidInputSchemaRule checks the pattern is valid
type AzurermEventgridDomainInvalidInputSchemaRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermEventgridDomainInvalidInputSchemaRule returns new rule with default attributes
func NewAzurermEventgridDomainInvalidInputSchemaRule() *AzurermEventgridDomainInvalidInputSchemaRule {
	return &AzurermEventgridDomainInvalidInputSchemaRule{
		resourceType:  "azurerm_eventgrid_domain",
		attributeName: "input_schema",
		enum: []string{
			"EventGridSchema",
			"CustomEventSchema",
			"CloudEventV01Schema",
		},
	}
}

// Name returns the rule name
func (r *AzurermEventgridDomainInvalidInputSchemaRule) Name() string {
	return "azurerm_eventgrid_domain_invalid_input_schema"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermEventgridDomainInvalidInputSchemaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermEventgridDomainInvalidInputSchemaRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermEventgridDomainInvalidInputSchemaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermEventgridDomainInvalidInputSchemaRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as input_schema`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
