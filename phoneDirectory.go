package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Phone   string
	Name    string
	Address string
}

func Phone(dir, num string) string {
	slice := strings.Split(dir, "\n")
	var v Person
	countIncidence := 0
	for _, address := range slice {
		if strings.Contains(address, num) {
			contact := strings.Split(address, " ")
			for _, info := range contact {
				if len(info) >= 14 {
					info = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(info, "+", ""), "/", ""), ",", "")
					if strings.Compare(num, info) == 0 {
						countIncidence += 1
						v.Phone = fmt.Sprintf("Phone => %s", num)
					}
				} else if strings.Contains(info, "<") && !strings.Contains(info, ">") {
					v.Name = fmt.Sprintf("Name => %s ", info[1:])
				} else if strings.Contains(info, ">") && !strings.Contains(info, "<") {
					v.Name += fmt.Sprintf("%s", info[0:len(info)-1])
				} else if strings.Contains(info, "<") && strings.Contains(info, ">") {
					v.Name = fmt.Sprintf("Name => %s", info[1:len(info)-1])
				} else {
					v.Address += fmt.Sprintf("%s ", info)
				}
			}
			v.Address = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Join(strings.Fields(v.Address), " "), "_", " "), "/", ""), ",", "")

		}

	}
	if len(v.Phone) < 14 {
		return fmt.Sprintf("Error => Not found: %s", num)
	}
	if countIncidence > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	}
	return fmt.Sprintf("%s, %s, Address => %s", v.Phone, v.Name, v.Address)
}

var dr = "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n High Street CC-47209 <Peter O'Brien> +1-908-512-2222\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010\n" + "+1-541-984-3012 <P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209\n" + "+1-741-984-3090 <Peter Reedgrave> _Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "+1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB-47209,\n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017\n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal    Roma\n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel> *+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av.  Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"

func main() {
	fmt.Println(Phone(dr, "1-541-754-3010"))
	fmt.Println(Phone(dr, "48-421-674-8974"))
	fmt.Println(Phone(dr, "8-421-674-8974"))
	fmt.Println(Phone(dr, "1-921-512-2222"))
	fmt.Println(Phone(dr, "1-541-754-3010"))
	fmt.Println(Phone(dr, "1-098-512-2222"))
	fmt.Println(Phone(dr, "5-555-555-5555"))
	fmt.Println(Phone(dr, "1-908-512-2222"))
}

// 1-908-512-2222, Name => Peter O'Brien, Address => High Street CC-47209
