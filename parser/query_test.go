package parser

import (
	"testing"

	"github.com/Hongbo-Miao/gqlparser/v2/ast"
	"github.com/Hongbo-Miao/gqlparser/v2/parser/testrunner"
)

func TestQueryDocument(t *testing.T) {
	testrunner.Test(t, "query_test.yml", func(t *testing.T, input string) testrunner.Spec {
		doc, err := ParseQuery(&ast.Source{Input: input, Name: "spec"})
		return testrunner.Spec{
			Error: err,
			AST:   ast.Dump(doc),
		}
	})
}
