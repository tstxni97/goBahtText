package gobahttext

import (
	"fmt"
	"strconv"
	"strings"
)

var thaiNumbersMap = map[string]string{
	"0": "ศูนย์",
	"1": "หนึ่ง",
	"2": "สอง",
	"3": "สาม",
	"4": "สี่",
	"5": "ห้า",
	"6": "หก",
	"7": "เจ็ด",
	"8": "แปด",
	"9": "เก้า",
}

var thaiMultipliersList = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

func splitMillionSequence(seq string) []string {
	mlength := 6
	seq = reverseString(seq)
	var splitedSeq []string
	for i := 0; i < len(seq); i += mlength {
		end := i + mlength
		if end > len(seq) {
			end = len(seq)
		}
		subSeq := reverseString(seq[i:end])
		splitedSeq = append(splitedSeq, subSeq)
	}
	return reverseStringSlice(splitedSeq)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func convertSpecialTwoDigits(intString string) string {
	bahtString := ""

	if intString[0] == '2' {
		bahtString += "ยี่"
	} else if intString[0] == '3' || intString[0] == '4' || intString[0] == '5' || intString[0] == '6' || intString[0] == '7' || intString[0] == '8' || intString[0] == '9' {
		bahtString += thaiNumbersMap[string(intString[0])]
	}

	if intString[0] != '0' {
		bahtString += "สิบ"
	}

	firstDigit, err := strconv.ParseInt(string(intString[1]), 10, 64)

	if err != nil {
		return ""
	}

	if firstDigit == 0 {
		// do nothing
	} else if intString[0] == '0' && (firstDigit > 0) {
		bahtString += thaiNumbersMap[string(intString[1])]
	} else if firstDigit == 1 {
		bahtString += "เอ็ด"
	} else {
		bahtString += thaiNumbersMap[string(intString[1])]
	}

	return bahtString
}

func convertMultipleMillions(intString string) string {
	bahtString := ""

	// integer part
	if len(intString) == 1 {
		bahtString += thaiNumbersMap[intString]
	} else if len(intString) == 2 {
		bahtString += convertSpecialTwoDigits(intString)
	} else if len(intString) >= 3 && len(intString) <= 6 {
		for numberIndex := 0; numberIndex < len(intString)-2; numberIndex++ {
			multiplier, err := strconv.Atoi(string(intString[numberIndex]))
			if err != nil {
				return ""
			}
			if multiplier != 0 {
				bahtString += thaiNumbersMap[string(intString[numberIndex])]
				bahtString += thaiMultipliersList[len(intString)-numberIndex-1]
			}
		}
		bahtString += convertSpecialTwoDigits(intString[len(intString)-2:])
	}

	return bahtString
}

func BahtText(inputFloat float64) string {
	bahtString := ""

	if inputFloat < 0 {
		inputFloat *= -1
		bahtString += "ลบ"
	}

	floatString := fmt.Sprintf("%.2f", inputFloat)
	inputString := strings.Split(floatString, ".")

	intString := inputString[0]
	decimalString := inputString[1]

	millionSequenceList := splitMillionSequence(intString)

	millionIndex := len(millionSequenceList) - 1

	for _, millionGroupString := range millionSequenceList {

		bahtString += convertMultipleMillions(millionGroupString)

		if millionIndex > 0 {
			bahtString += "ล้าน"
			millionIndex -= 1
		}
	}

	bahtString += "บาท"

	// decimal part
	decimalInt, _ := strconv.Atoi(decimalString)
	if decimalInt == 0 {
		bahtString += "ถ้วน"
	} else {
		if len(decimalString) == 1 {
			bahtString += thaiNumbersMap[string(rune(decimalInt))]
		} else if len(decimalString) == 2 {
			bahtString += convertSpecialTwoDigits(decimalString)
		}

		bahtString += "สตางค์"
	}

	return bahtString
}
