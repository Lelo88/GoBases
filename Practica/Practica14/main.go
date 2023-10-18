// After becoming famous, the CodeBots decided to move into a new building together. Each of the rooms has a 
// different cost, and some of them are free, but there's a rumour that all the free rooms are haunted! Since
// the Codebots are quite superstitious, they refuse to stay in any of the free rooms, or any of the rooms below
// any of the free rooms.
// Given matrix, a rectangular matrix of integers, where each value represents the cost of the room, your task is
// to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all the values that don't
// appear below a 0).

package main

func main() {
	
	var matrix [][]int = [][]int{
		{0, 1, 1, 2}, 
		{0, 5, 0, 0}, 
		{2, 0, 3, 3},
	}

	var sum int = matrixElementsSum(matrix)

	println(sum)
}

func matrixElementsSum(matrix [][]int) int {
	var sum int = 0
	var flag bool = false

	for i := 0; i < len(matrix[0]); i++ {
		flag = false
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == 0 {
				flag = true
			} else if !flag {
				sum += matrix[j][i]
			}
		}
	}

	return sum
}