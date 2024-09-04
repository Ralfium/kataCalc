package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romToArab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := scan.Text()

	data := strings.Split(line, " ")
	if len(data) != 3 {
		err := errors.New("invalid input")
		panic(err)
	}
	aS, bS := data[0], data[2]
	oper := data[1]
	romCheckA := romCheck(aS)
	romCheckB := romCheck(bS)

	if !romCheckA && !romCheckB {
		a, err := strconv.Atoi(aS)
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(bS)
		if err != nil {
			panic(err)
		}

		resultAr, err := arabicCalculator(a, b, oper)
		if err != nil {
			panic(err)
		}
		fmt.Println(resultAr)

	} else if romCheckA && romCheckB {
		aAr := romanToArabic(aS)
		bAr := romanToArabic(bS)
		resultAr, err := arabicCalculator(aAr, bAr, oper)
		if err != nil {
			panic(err)
		}

		if resultAr >= 1 {
			fmt.Println(arabicToRoman(resultAr))
		} else {
			err := errors.New("result less than I")
			panic(err)
		}

	} else if romCheckA || romCheckB {
		err := errors.New("mismatched of number types or number out of range 1 to 10")
		panic(err)
	}
}

func arabicCalculator(a, b int, oper string) (int, error) {
	err := errors.New("invalid operation")
	err2 := errors.New("only 1 to 10 numbers can be calculated")
	if a > 10 || b > 10 || a < 1 || b < 1 {
		return 0, err2
	}
	switch oper {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, err
	}
}

func romCheck(rom string) bool {
	for i := range romToArab {
		if rom == i {
			return true
		}
	}
	return false
}

func romanToArabic(rom string) int {
	return romToArab[rom]
}

func arabicToRoman(ar int) string {
	rom := ""
	arabToRom := []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, v := range arabToRom {
		for ar >= v.value {
			rom += v.symbol
			ar -= v.value
		}
	}
	return rom
}
