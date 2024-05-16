package conversion

import "strings"

type PostFix struct {
	expression []string
}

func (p *PostFix) String() string {
	return strings.Join(p.expression, "")
}

func (p *PostFix) Add(s string) {
	p.expression = append(p.expression, s)
}

func InfixToPostfix(s string) (PostFix, error) {
	temArray := []string{"1", "2.5", "3", "/", "4", "*", "+"}
	postfix := PostFix{expression: temArray}

	return postfix, nil
}

func Tokenization(s string) ([]string, error) {

	expression := strings.ReplaceAll(s, " ", "")

	var digit string
	var infixTokenized []string
	for _, v := range expression {
		if v >= '0' && v <= '9' || v == '.' {
			digit += string(v)
		} else {
			if digit != "" {
				infixTokenized = append(infixTokenized, digit)
				digit = ""
			}
			infixTokenized = append(infixTokenized, string(v))
		}
	}

	if digit != "" {
		infixTokenized = append(infixTokenized, digit)
	}

	return infixTokenized, nil
}
