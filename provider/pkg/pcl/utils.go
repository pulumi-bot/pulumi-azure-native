package pcl

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/zclconf/go-cty/cty"
)

// Camel replaces the first contiguous string of upper case runes in the given string with its lower-case equivalent.
func Camel(s string) string {
	c, sz := utf8.DecodeRuneInString(s)
	if sz == 0 || unicode.IsLower(c) {
		return s
	}

	// The first rune is not lowercase. Iterate until we find a rune that is.
	var word []rune
	for {
		s = s[sz:]

		n, nsz := utf8.DecodeRuneInString(s)
		if nsz == 0 {
			word = append(word, unicode.ToLower(c))
			return string(word)
		}
		if unicode.IsLower(n) {
			if len(word) == 0 {
				c = unicode.ToLower(c)
			}
			word = append(word, c)
			return string(word) + s
		}
		c, sz, word = n, nsz, append(word, unicode.ToLower(c))
	}
}

// PlainLit returns an unquoted string literal expression.
func PlainLit(v string) *model.LiteralValueExpression {
	return &model.LiteralValueExpression{Value: cty.StringVal(v)}
}

// QuotedLit returns a quoted string literal expression.
func QuotedLit(v string) *model.TemplateExpression {
	return &model.TemplateExpression{Parts: []model.Expression{PlainLit(v)}}
}

// ObjectConsItem returns a new ObjectConsItem with the given key and value.
func ObjectConsItem(key string, value model.Expression) model.ObjectConsItem {
	var keyExpr model.Expression = PlainLit(key)
	if !hclsyntax.ValidIdentifier(key) {
		keyExpr = QuotedLit(key)
	}
	return model.ObjectConsItem{
		Key:   keyExpr,
		Value: value,
	}
}

// Invoke returns a new function call expression which invokes the specified function.
func Invoke(token string, inputs ...model.ObjectConsItem) *model.FunctionCallExpression {
	args := []model.Expression{QuotedLit(token)}
	if len(inputs) != 0 {
		args = append(args, &model.ObjectConsExpression{Items: inputs})
	}
	return &model.FunctionCallExpression{
		Name: "invoke",
		Args: args,
	}
}

// RelativeTraversal returns a new RelativeTraversalExpression that accesses the given attribute of the source
// expression.
func RelativeTraversal(source model.Expression, attr string) *model.RelativeTraversalExpression {
	return &model.RelativeTraversalExpression{
		Source:    source,
		Traversal: hcl.Traversal{hcl.TraverseAttr{Name: attr}},
		Parts:     []model.Traversable{model.DynamicType, model.DynamicType},
	}
}

// MakeLegalIdentifier deletes characters that are not allowed in HCL2 identifiers with underscores. No attempt is
// made to ensure that the result is unique.
func MakeLegalIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if isLegalIdentifierPart(c) {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	if builder.Len() == 0 {
		return "x"
	}
	return builder.String()
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of an HCL2 identifier.
func isLegalIdentifierStart(c rune) bool {
	return c == '$' || c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of an HCL2 identifier.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || unicode.In(c, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc)
}
