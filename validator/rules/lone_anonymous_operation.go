package validator

import (
	"github.com/Hongbo-Miao/gqlparser/v2/ast"
)

func init() {
	AddRule("LoneAnonymousOperation", func(observers *Events, addError AddErrFunc) {
		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			if operation.Name == "" && len(walker.Document.Operations) > 1 {
				addError(
					Message(`This anonymous operation must be the only defined operation.`),
					At(operation.Position),
				)
			}
		})
	})
}
