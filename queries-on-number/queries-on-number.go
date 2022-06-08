package queries_on_point

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	result := []int{}
	for _, loc := range queries {
		x := float64(loc[0])
		y := float64(loc[1])
		r := float64(loc[2])
		count := 0

		for _, loc2 := range points {
			dist := math.Sqrt(math.Pow(float64(loc2[0])-x, 2) + math.Pow(float64(loc2[1])-y, 2))
			if r >= dist {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
