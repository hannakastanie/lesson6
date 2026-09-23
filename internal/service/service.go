package service

import (
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var textSymbols = "ЁЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮёйцукенгшщзхъфывапролджэячсмитьбю1234567890,:?\\/()\""
var workSymbols = ".-ЁЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮёйцукенгшщзхъфывапролджэячсмитьбю1234567890,:?\\/()\""

// convert text to morse and vice versa
func ConvertTextMorse(data string) string{
	if !strings.ContainsAny(data, workSymbols){
		log.Fatal()
	}
	if strings.ContainsAny(data, textSymbols){
		return morse.ToMorse(data)
	} 
	return morse.ToText(data)
}