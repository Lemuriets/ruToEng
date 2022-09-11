package ruToEng

import (
	"strings"

	"github.com/Lemuriets/ruToEng/pkg/symbols"
)

func toUpperEng(symbol, engSymbol string) string {
	if strings.ToUpper(symbol) == symbol {
		return string(strings.ToUpper(engSymbol)[0]) + engSymbol[1:]
	}
	return engSymbol
}

func RuToEng(str string) string {
	var result string
	for _, s := range str {
		symbol := string(s)
		if engSymbol, ok := symbols.Symbols[strings.ToLower(symbol)]; ok {
			result += toUpperEng(symbol, engSymbol)
		} else {
			result += symbol
		}
	}
	return result
}
