package main

import "fmt"

func GetMayorMenor(arr []int) (int, int) {
	mayor := 0
	menor := 1000
	for _, r := range arr {
		if r > mayor {
			mayor = r
		}
		if r < menor {
			menor = r
		}
	}
	return mayor, menor
}

func CalcSmallestDifference(arr1, arr2 []int) int {
	mayor1, menor1 := GetMayorMenor(arr1)
	mayor2, menor2 := GetMayorMenor(arr2)
	difference1 := mayor1 - menor2
	difference2 := mayor2 - menor1
	if difference1 < 0 {
		return difference2
	} else if difference2 < 0 {
		return difference1
	} else if difference1 > difference2 {
		return difference2
	} else {
		return difference1
	}
}

func main() {
	fmt.Println("Smallest Difference Challenge")

	arr1 := []int{9, 8, 7, 6}
	arr2 := []int{7, 3, 2, 5}

	smallestDiff := CalcSmallestDifference(arr1, arr2)
	fmt.Println(smallestDiff)
}
