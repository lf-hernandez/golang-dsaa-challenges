/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/
package main

import (
	"fmt"
	"os"
)

func twoSum(nums []int, target int) []int {
	var answer []int

	for cursorA := 0; cursorA < len(nums)-1; cursorA++ {
		for cursorB := 1; cursorB < len(nums); cursorB++ {
			if nums[cursorA]+nums[cursorB] == target {
				answer = append(answer, cursorA)
				answer = append(answer, cursorB)
			}
		}
	}

	return answer
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	answerSlice := twoSum(nums, target)
	answerSum := 0

	fmt.Println("target: ", target)
	fmt.Println("answer slice:", answerSlice)
	fmt.Println("nums array mapping:", nums[answerSlice[0]], ",", nums[answerSlice[1]])

	for _, currentValue := range answerSlice {
		answerSum += nums[currentValue]
	}

	if answerSum == target {
		fmt.Println("Pass: answer == target")
		os.Exit(0)
	} else {
		fmt.Println("Fail: answer != target")
		os.Exit(1)
	}
}
