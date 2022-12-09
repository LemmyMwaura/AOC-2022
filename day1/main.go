package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func getMostCalories(calories []int) int {
	max_value := calories[0]

	for i := 0; i < len(calories); i++ {
		if calories[i] > max_value {
			max_value = calories[i]
		}
	}

	return max_value
}

func getTopThreeCalories(calories []int) int {
	count := 0

	sort.SliceStable(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	for i := 0; i < 3; i++ {
		count += calories[i]
	}

	return count
}

func main() {
	file, err := os.Open("day1/calories.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	caloriesPerElf := []int{}
	sumOfValues := 0

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		calories, _ := strconv.Atoi(string(line))
		sumOfValues += calories

		if calories == 0 {
			// calories is equal to 0 when there's an empty space
			caloriesPerElf = append(caloriesPerElf, sumOfValues)
			sumOfValues = 0
		}
	}

	fmt.Println("Day One Solutions:")
	fmt.Printf("- Most Calories = %d\n", getMostCalories(caloriesPerElf))
	fmt.Printf("- Top Three Calories = %d\n", getTopThreeCalories(caloriesPerElf))
}
