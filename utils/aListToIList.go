package utils

import "strconv"

func AListToIList(list []string) []int {
	var intList []int
	for _, item := range list {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		intList = append(intList, num)
	}
	return intList
}
