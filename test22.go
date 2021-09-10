package main

import "math"

func main() {
	num := 3 //每次处理的数量
	//all := []int{1, 2, 3, 4, 5, 6, 7, 8}
	all := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	value := sliceCuttingPart(all, num)
	println(value)
}

func sliceCuttingPart(handleArr interface{}, num int) interface{} {
	switch handleArr.(type) {
	case []int:
		all := handleArr.([]int)
		newAll := [][]int{}
		if len(all) == 0 || num <= 0 {
			newAll = append(newAll, all)
			return newAll
		}
		if num >= len(all) {
			newAll = append(newAll, all)
			return newAll
		}
		for i := 1; i <= int(math.Floor(float64(len(all)/num)))+1; i++ {
			low := num * (i - 1)
			if low >= len(all) {
				break
			}
			high := num * i
			if high > len(all) {
				high = len(all)
			}
			newAll = append(newAll, all[low:high])
		}
		return newAll
	case []string:
		all := handleArr.([]string)
		newAll := [][]string{}
		if len(all) == 0 || num <= 0 {
			newAll = append(newAll, all)
			return newAll
		}
		if num >= len(all) {
			newAll = append(newAll, all)
			return newAll
		}
		for i := 1; i <= int(math.Floor(float64(len(all)/num)))+1; i++ {
			low := num * (i - 1)
			if low >= len(all) {
				break
			}
			high := num * i
			if high > len(all) {
				high = len(all)
			}
			newAll = append(newAll, all[low:high])
		}
		return newAll
	default:
		return nil
	}
}
