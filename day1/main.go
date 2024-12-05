package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)
	
func calculateDistance(leftList []int, rightList []int) int {
	distance := 0
		for i := 0; i < len(leftList); i++ {
			if leftList[i] != rightList[i] {
				distance += int(math.Abs(float64(leftList[i] - rightList[i])))
			}
		}

		return distance
}


func main() {
		file, err := os.Open("./day1/list.txt")
		if err != nil {
			panic(err)
		}

		var leftList []int = []int{}
		var rightList []int = []int{}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			leftNum, err1 := strconv.Atoi(parts[0])
			rightNum, err2 := strconv.Atoi(parts[1])
			
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing numbers in line:", line)
				continue
			}

			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		}
	
		file.Close()

		slices.Sort(leftList)
		slices.Sort(rightList)
	
		println("Distance:", calculateDistance(leftList, rightList))
}
