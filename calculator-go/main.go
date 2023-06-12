package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"
)

const (
	LOW = "Вывод ошибки, так как строка " +
		"не является математической операцией."
	HIGH = "Вывод ошибки, так как формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Вывод ошибки, так как используются " +
		"одновременно разные системы счисления."
	PAN_DIFF = "Отрицательный результат"
	PAN_NULL = "Ошибка числа"
	PAN_DIAP = "Внимание на вывод!"
)

func main() {
	fmt.Println("Введите ваше выражение")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(s)))

	}
}

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var p1, p2 *int

var operators = map[string]func() int{
	"+": func() int { return *p1 + *p2 },
	"-": func() int { return *p1 - *p2 },
	"/": func() int { return *p1 / *p2 },
	"*": func() int { return *p1 * *p2 },
}

var data []string

func base(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	results := make([]string, 0)
	resultsToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(HIGH)
	case len(operator) < 1:
		panic(LOW)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			results = append(results, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(SCALE)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			p1, p2 = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(PAN_DIAP)
		}
	case 2:
		for _, elem := range results {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				resultsToInt = append(resultsToInt, val)
			} else {
				fmt.Println("Калькулятор работает только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно")
				panic(PAN_DIAP)
			}
		}
		if val, ok := operators[operator]; ok {
			p1, p2 = &resultsToInt[0], &resultsToInt[1]
			intToResult(val())
		}
	}
}

func intToResult(intResult int) {
	var romanNum string
	if intResult == 0 {
		fmt.Println("В римской системе нет такого числа!")
		panic(PAN_NULL)
	} else if intResult < 0 {
		fmt.Println("В римской системе нет отрицательных чисел!")
		panic(PAN_DIFF)
	}
	for intResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= intResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						intResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
