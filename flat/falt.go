package flat

func Execute(arr [][]int) []int {

	// Initialize the flattened array
	flatArr := []int{}

	// Iterate through the rows of the array
	for _, row := range arr {

		// Append the current row to the flattened array
		flatArr = append(flatArr, row...)
	}
	return flatArr
}
