// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMssqlDatabaseInvalidSampleNameRule checks the pattern is valid
type AzurermMssqlDatabaseInvalidSampleNameRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermMssqlDatabaseInvalidSampleNameRule returns new rule with default attributes
func NewAzurermMssqlDatabaseInvalidSampleNameRule() *AzurermMssqlDatabaseInvalidSampleNameRule {
	return &AzurermMssqlDatabaseInvalidSampleNameRule{
		resourceType:  "azurerm_mssql_database",
		attributeName: "sample_name",
		enum: []string{
			"AdventureWorksLT",
			"WideWorldImportersStd",
			"WideWorldImportersFull",
		},
	}
}

// Name returns the rule name
func (r *AzurermMssqlDatabaseInvalidSampleNameRule) Name() string {
	return "azurerm_mssql_database_invalid_sample_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMssqlDatabaseInvalidSampleNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMssqlDatabaseInvalidSampleNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMssqlDatabaseInvalidSampleNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMssqlDatabaseInvalidSampleNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sample_name`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
