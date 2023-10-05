package main

import (
	"fmt"
	"strings"
)

func ModifyMultiply(str string, loc, num int) string {
	slice := []string{}
	splited := strings.Fields(str)
	if num == 0 {
		slice = append(slice, string(splited[loc]))
	} else {

		for i := 0; i < num; i++ {
			slice = append(slice, string(splited[loc]))
		}
	}
	return strings.Join(slice, "-")
}

func main() {
	fmt.Println(ModifyMultiply("bjKitgiy IDBIehkClsRrtIzLZHoo nIxEgGqgrE TkklUDtQ RpRXTT NGMnfUUxDkWgDpspnDzzqwN oinouoFDYKdPV YMjMC JGQSLvZkZrRnNvZ DZkDWbBbZNB gLMOFPWpJObKAWeJtxWcUS JtXLDPwKhFLfdRDOkMqJQ JsYlP qLOijgmitlxfOdRR OltvR OzPDGeI pXM BCDdqApbIFEK iTPHIqQNVhrMLkJNOt LfUd tBFyNWeLYfSdVQ fUFMePkwQxRCxUTtL", 18, 0))
}
