package phone_number_mnemonic

import (
	"strconv"
)

func PhoneMnemonic(phoneNumber string) []string {
	mapping := map[int]string{
		0: "0",
		1: "1",
		2: "ABC",
		3: "DEF",
		4: "GHI",
		5: "JKL",
		6: "MNO",
		7: "PQRS",
		8: "TUV",
		9: "WXYZ",
	}

	return mnemonicHelper(phoneNumber, mapping)
}

func mnemonicHelper(phoneNumber string, mapping map[int]string) []string {
	if len(phoneNumber) == 0 {
		return []string{""}
	}

	mnemonics := mnemonicHelper(string(phoneNumber[1:]), mapping)
	currentDigit, err := strconv.Atoi(string(phoneNumber[0]))
	if err != nil {
		panic(err)
	}

	result := make([]string, 0)

	for _, m := range mnemonics {
		for _, c := range mapping[currentDigit] {
			result = append(result, string(c)+m)
		}
	}

	return result
}
