package main

//Uncomment each block of code under each solution to run and view

//Solution for problem 1
import "fmt"

func main(){
	var k,m int

	fmt.Print("Enter the K value: ")
	fmt.Scanf("%d",&k)

	fmt.Print("Enter the M value: ")
	fmt.Scanf("%d",&m)

	var size int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&size)

	numbers := make([]int, size)

	fmt.Printf("Enter %d numbers separated by spaces or newlines:\n", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i]) 
	}

	currentWindowSum := 0
	maxValidSum := 0
	bestStart := -1

	for i:=0; i < k; i++{
		currentWindowSum += numbers[i]
	}

	if currentWindowSum <= m{
		maxValidSum = currentWindowSum
		bestStart = 0
	}

	for i := k; i < len(numbers); i++{
		currentWindowSum = currentWindowSum + numbers[i] - numbers[i-k]

		currentStart := i - k + 1

		if currentWindowSum <= m && (bestStart == -1 || currentWindowSum > maxValidSum){
			maxValidSum = currentWindowSum
			bestStart = currentStart
		}
	}

	if bestStart == -1 {
		fmt.Printf("\nNo valid subarray found under threshold %d\n", m)
		fmt.Println("Maximum valid subarray sum: 0")
		fmt.Println("Subarray: []")
	} else {
		maxSubarray := numbers[bestStart : bestStart+k]
		fmt.Printf("\nMaximum valid subarray sum: %d\n", maxValidSum)
		fmt.Printf("Subarray: %v\n", maxSubarray)
	}
}
