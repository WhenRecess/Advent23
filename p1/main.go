package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	calibrate()
}

func calibrate() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)

	// fmt.Println(str)

	temp := strings.Split(str, "\n")

	var allNums []int

	for _, line := range temp {
		var firstNum byte
		var secondNum byte
		for i, character := range line {
			if unicode.IsDigit(character) {
				fmt.Printf("\n\nThe digit is %c", character)
				firstNum = line[i]
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit([]rune(line)[i]) {
				secondNum = line[i]
				fmt.Printf("The last digit is %c\n\n", []rune(line)[i])
				break
			}
		}
		first, err := strconv.Atoi(string(firstNum))
		if err != nil {
			fmt.Println(err)
		}
		second, err := strconv.Atoi(string(secondNum))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("THE AWESOME SAUCE", first)
		fmt.Println("THE AWESOME SEQUEL", second)
		// fmt.Printf("\nThe second is: %v\n", int(secondNum))
		// strInt := firstNum + secondNum
		// fmt.Printf("The string int is %v", strInt)
		combined := strconv.Itoa(first) + strconv.Itoa(second)
		fmt.Printf("\nThe combined number is: %d\n", firstNum+secondNum)
		toInt, err := strconv.Atoi(combined)
		if err != nil {
			fmt.Println(err)
		}
		allNums = append(allNums, toInt)
	}
	var total int
	for _, num := range allNums {
		total += num
	}
	fmt.Println("The number is ", total)
}
