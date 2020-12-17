package main

import (
	"./goldengatedata"
	"fmt"
)

func getSetEfficiency(x []int) float64 {
	var efficiency = float64(1.0)
	for _, num := range x {
		nCorrect := goldengatedata.Mismatches[goldengatedata.Key{int(num), int(num)}]
		nTotal := 0
		for _, num2 := range x {
			nTotal = nTotal + goldengatedata.Mismatches[goldengatedata.Key{int(num), int(num2)}]
		}
		efficiency = efficiency * (float64(nCorrect) / float64(nTotal))
	}
	return efficiency

}

func addOverhang(x []int) []int {
	checkList := []int{}
	var occurs bool
	// Get i who do not exist in list
	for i := 0; i < 120; i++ {
		occurs = false
		for _, current := range x {
			if i == current {
				occurs = true
			}
		}
		if occurs == false {
			checkList = append(checkList, i)
		}
	}
	maxEfficiency := float64(0)
	var efficiency float64
	var newOverhang int
	for _, i := range checkList {
		efficiency = getSetEfficiency(append(x, i))
		if efficiency > maxEfficiency {
			maxEfficiency = efficiency
			newOverhang = i
		}
	}
	return append(x, newOverhang)
}

func Overhangs(x []int) []string {
	overhangs := []string{}
	for _, i := range x {
		overhangs = append(overhangs, goldengatedata.Overhangs[i])
	}
	return overhangs
}

func main() {
	s := []int{105, 32, 5, 10, 40}
	efficiency := float64(0.95)
	for i := 0; i < 50; i++ {
		s = addOverhang(s)
		newEfficiency := getSetEfficiency(s)
		fmt.Println(newEfficiency)
		if newEfficiency < efficiency {
			break
		}
	}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(Overhangs(s))
	fmt.Println(getSetEfficiency(s))

	fmt.Println(getSetEfficiency([]int{105, 32, 5, 10, 40, 0, 16, 21, 31, 36, 42, 47, 56, 59, 72, 77, 88, 91, 98, 113, 114, 101, 68, 63}))

	maxLen := 1
	var t []int
	for i := 0; i < 120; i++ {
		s := []int{105, 40}
		for i := 0; i < 50; i++ {
			s = addOverhang(s)
			newEfficiency := getSetEfficiency(s)
			if newEfficiency < efficiency {
				break
			}
		}
		if len(s) > maxLen {
			maxLen = len(s)
			t = s
		}
	}
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(getSetEfficiency(t))
	fmt.Println(Overhangs(t))
}
