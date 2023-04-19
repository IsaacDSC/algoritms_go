package removedequals

import "sort"

func Execute(listArr []int) []int {
	var (
		acumulator []int
	)
	sort.Slice(listArr, func(i, j int) bool {
		return listArr[i] < listArr[j]
	})
	for index := range listArr {
		if index == 0 {
			acumulator = append(acumulator, listArr[index])
		} else if listArr[index-1] != listArr[index] {
			acumulator = append(acumulator, listArr[index])
		}
	}
	return acumulator
}
