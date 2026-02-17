package main

import (
	"errors"
	"strings"
)

type RomanConverter struct {
	values map[rune]int
}

func NewRomanConverter() *RomanConverter {
	return &RomanConverter{
		values: map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		},
	}
}

func (rc *RomanConverter) RomanToArabic(roman string) (int, error) {
	if roman == "" {
		return 0, errors.New("пустая строка")
	}

	roman = strings.ToUpper(roman)

	for _, char := range roman {
		if _, exists := rc.values[char]; !exists {
			return 0, errors.New("недопустимый символ: " + string(char))
		}
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := rc.values[rune(roman[i])]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	if !rc.isValidRoman(roman, result) {
		return 0, errors.New("некорректное римское число")
	}

	return result, nil
}

func (rc *RomanConverter) ArabicToRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", errors.New("число должно быть от 1 до 3999")
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	temp := num

	for i := 0; i < len(values); i++ {
		for temp >= values[i] {
			result += symbols[i]
			temp -= values[i]
		}
	}

	return result, nil
}

func (rc *RomanConverter) isValidRoman(roman string, expected int) bool {
	backConverted, err := rc.ArabicToRoman(expected)
	if err != nil {
		return false
	}
	return backConverted == roman
}

func (rc *RomanConverter) GetRomanExamples() map[string]int {
	return map[string]int{
		"I":    1,
		"IV":   4,
		"V":    5,
		"IX":   9,
		"X":    10,
		"XL":   40,
		"L":    50,
		"XC":   90,
		"C":    100,
		"CD":   400,
		"D":    500,
		"CM":   900,
		"M":    1000,
		"MMXX": 2020,
		"MCML": 1950,
	}
}
