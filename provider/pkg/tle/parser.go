package tle

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/go/common/util/contract"
)

type TLEParseResult struct {
	LeftSquareBracketToken  *Token
	Expression              Value
	RightSquareBracketToken *Token
}

func Parse(s string) (*TLEParseResult, error) {
	return parseString(quote(s))
}

func parseString(quotedStringValue string) (*TLEParseResult, error) {
	contract.Assert(1 <= len(quotedStringValue)) // TLE strings must be at least 1 character
	// The first character in the TLE string to parse must be a quote character
	contract.Assert(isQuoteCharacter(quotedStringValue[0]))

	var leftSquareBracketToken *Token
	var expression Value
	var rightSquareBracketToken *Token
	if 3 <= len(quotedStringValue) && quotedStringValue[1:2] == "[[" {
		expression = NewStringValue(CreateQuotedString(0, quotedStringValue))
	} else {
		tokenizer := fromString(quotedStringValue)
		tokenizer.Next()

		if tokenizer.Current() == nil || tokenizer.Current().GetType() != LeftSquareBracket {
			// This is just a plain old string (no brackets). Mark its expression as being
			// the string value.
			expression = NewStringValue(CreateQuotedString(0, quotedStringValue))
		} else {
			leftSquareBracketToken = tokenizer.Current()
			tokenizer.Next()

			if tokenizer.Current() != nil &&
				tokenizer.Current().GetType() != Literal &&
				tokenizer.Current().GetType() != RightSquareBracket {
				return nil, fmt.Errorf("Expected a literal value: %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
			}

			var err error
			expression, err = parseExpression(tokenizer)
			if err != nil {
				return nil, fmt.Errorf("failed to parse expression '%s': %w", quotedStringValue, err)
			}

			for tokenizer.Current() != nil {
				if tokenizer.Current().GetType() == RightSquareBracket {
					rightSquareBracketToken = tokenizer.Current()
					tokenizer.Next()
					break
				} else {
					return nil, fmt.Errorf("nxpected end of string: %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
				}
			}

			if rightSquareBracketToken != nil {
				if tokenizer.Current() != nil {
					return nil, fmt.Errorf("nothing should exist after the closing ']' except for whitespace, %s:%v", quotedStringValue, tokenizer.Current().Span().getStartIndex())
				}
			} else {
				return nil, fmt.Errorf("expected a right square bracket (']'), %s:%v", quotedStringValue, len(quotedStringValue)-1)
			}

			if expression == nil {
				errorSpan := &leftSquareBracketToken.span
				if rightSquareBracketToken != nil {
					errorSpan = union(errorSpan, &rightSquareBracketToken.span)
				}
				return nil, fmt.Errorf("expected a function of property expression, %s:%v", quotedStringValue, errorSpan.getStartIndex())
			}
		}
	}

	return &TLEParseResult{
		LeftSquareBracketToken:  leftSquareBracketToken,
		Expression:              expression,
		RightSquareBracketToken: rightSquareBracketToken,
	}, nil
}

func parseExpression(tokenizer *Tokenizer) (Value, error) {
	var expression Value
	if tokenizer.Current() != nil {
		var rootExpression Value // Initial expression

		token := tokenizer.Current()
		tokenType := token.GetType()
		if tokenType == Literal {
			var err error
			rootExpression, err = parseFunctionCall(tokenizer)
			if err != nil {
				return nil, err
			}
		} else if tokenType == QuotedString {
			if !strings.HasSuffix(token.stringValue, string(token.stringValue[0])) {
				return nil, fmt.Errorf("a constant string is missing an end quote: %d", token.span.getStartIndex())
			}
			rootExpression = NewStringValue(*token)
			tokenizer.Next()
		} else if tokenType == Number {
			rootExpression = NewNumberValue(*token)
			tokenizer.Next()
		} else if tokenType != RightSquareBracket && tokenType != Comma {
			return nil, fmt.Errorf("template language expressions must start with a function: %d", token.span.getStartIndex())
		}

		if rootExpression == nil {
			return nil, nil
		}

		expression = rootExpression
	} else {
		return nil, nil
	}

	// Check for property or array accesses off of the root expression
	for tokenizer.Current() != nil {
		if tokenizer.Current().GetType() == Period {
			periodToken := tokenizer.Current()
			tokenizer.Next()

			var propertyNameToken *Token
			var errorSpan Span

			if tokenizer.Current() != nil {
				if tokenizer.Current().GetType() == Literal {
					propertyNameToken = tokenizer.Current()
					tokenizer.Next()
				} else {
					errorSpan = tokenizer.Current().Span()

					tokenType := tokenizer.Current().GetType()
					if tokenType != RightParenthesis &&
						tokenType != RightSquareBracket &&
						tokenType != Comma {
						tokenizer.Next()
					}
				}
			} else {
				errorSpan = periodToken.Span()
			}

			if propertyNameToken == nil {
				return nil, fmt.Errorf("expected a literal value in span at %d", errorSpan.getStartIndex())
			}

			expression = NewPropertyAccessValue(expression, *periodToken, *propertyNameToken)
		} else if tokenizer.Current().GetType() == LeftSquareBracket {
			leftSquareBracketToken := tokenizer.Current()
			tokenizer.Next()

			indexValue, err := parseExpression(tokenizer)
			if err != nil {
				return nil, err
			}

			var rightSquareBracketToken *Token
			if tokenizer.Current() != nil && tokenizer.Current().GetType() == RightSquareBracket {
				rightSquareBracketToken = tokenizer.Current()
				tokenizer.Next()
			}

			if leftSquareBracketToken == nil || rightSquareBracketToken == nil {
				return nil, fmt.Errorf("malformed array offset specification near: %d", indexValue.Span().getStartIndex())
			}
			expression = NewArrayAccessValue(expression, *leftSquareBracketToken, indexValue, *rightSquareBracketToken)
		} else {
			break
		}
	}

	return expression, nil
}

func parseFunctionCall(tokenizer *Tokenizer) (*FunctionCallValue, error) {
	contract.Assert(tokenizer != nil)
	contract.Assert(tokenizer.Current() != nil)
	contract.Assert(Literal == tokenizer.Current().GetType())

	var namespaceToken, nameToken, periodToken *Token

	firstToken := tokenizer.Current()
	tokenizer.Next()

	// Check for <namespace>.<functionname>
	if tokenizer.Current() != nil && tokenizer.Current().GetType() == Period {
		// It's a user-defined function because it has a namespace before the function name
		periodToken = tokenizer.Current()
		namespaceToken = firstToken
		tokenizer.Next()

		// Get the function name following the period
		if tokenizer.HasCurrent() && tokenizer.Current().GetType() == Literal {
			nameToken = tokenizer.Current()
			tokenizer.Next()
		} else {
			return nil, fmt.Errorf("Expected user-defined function name: %v", periodToken.Span().getStartIndex())
		}
	} else {
		nameToken = firstToken
	}

	var leftParenthesisToken, rightParanthesisToken *Token
	var commaTokens []*Token
	var argumentExpressions []Value

	if tokenizer.Current() != nil {
		for tokenizer.Current() != nil {
			if tokenizer.Current().GetType() == LeftParenthesis {
				leftParenthesisToken = tokenizer.Current()
				tokenizer.Next()
				break
			} else if tokenizer.Current().GetType() == RightSquareBracket {
				return nil, fmt.Errorf("Missing function argument list: %d", nameToken.Span().getStartIndex())
			} else {
				return nil, fmt.Errorf("Expected the end of the string: %d", nameToken.Span().getStartIndex())
			}
		}
	} else {
		return nil, fmt.Errorf("Missing function argument list: %d", nameToken.Span().getStartIndex())
	}

	if tokenizer.HasCurrent() {
		expectingArgument := true

		// tslint:disable-next-line: strict-boolean-expressions
		for tokenizer.Current() != nil {
			if tokenizer.Current().GetType() == RightParenthesis || tokenizer.Current().GetType() == RightSquareBracket {
				break
			} else if expectingArgument {
				expression, err := parseExpression(tokenizer)
				if err != nil {
					return nil, err
				}
				if expression == nil && tokenizer.HasCurrent() && tokenizer.Current().GetType() == Comma {
					return nil, errors.New("expected a constant string, function, or property expression")
				}
				argumentExpressions = append(argumentExpressions, expression)
				expectingArgument = false
			} else if tokenizer.Current().GetType() == Comma {
				expectingArgument = true
				commaTokens = append(commaTokens, tokenizer.Current())
				tokenizer.Next()
			} else {
				return nil, errors.New("expected a comma (',')")
			}
		}
	} else if leftParenthesisToken != nil {
		return nil, fmt.Errorf("expected a right parenthesis (')')")
	}

	// tslint:disable-next-line: strict-boolean-expressions
	if tokenizer.Current() != nil {
		switch tokenizer.current.GetType() {
		case RightParenthesis:
			rightParanthesisToken = tokenizer.Current()
			tokenizer.Next()
			break

		case RightSquareBracket:
			if leftParenthesisToken != nil {
				return nil, fmt.Errorf("expected a right parenthesis (')')")
			}
		}
	}

	return &FunctionCallValue{
		NamespaceToken:        namespaceToken,
		PeriodToken:           periodToken,
		NameToken:             *nameToken,
		LeftParenthesisToken:  leftParenthesisToken,
		CommaTokens:           commaTokens,
		ArgumentExpressions:   argumentExpressions,
		RightParenthesisToken: rightParanthesisToken,
	}, nil
}
