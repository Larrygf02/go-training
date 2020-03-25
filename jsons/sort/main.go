package main
import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{6,3,7,11,2}
	sort.Ints(numbers)
	fmt.Println(numbers)


}