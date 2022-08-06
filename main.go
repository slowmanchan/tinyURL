package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphaArr := strings.Split(alphabet, "")

	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}

	digits := []int{}
	for num > 0 {
		remainder := num % 62
		digits = append(digits, remainder)
		num /= 62
	}

	reverse(digits)
	log.Println(digits)

	res := []string{}
	for _, d := range digits {
		res = append(res, alphaArr[d])
	}

	log.Println("tinyURL", strings.Join(res, ""))

	resInt := 0
	for _, c := range res {
		resInt = resInt*62 + strings.Index(alphabet, string(c))
		fmt.Println(resInt)
	}

	log.Println("ID in DB", resInt)
}

func reverse(nums []int) {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] > nums[y]
	})
}
