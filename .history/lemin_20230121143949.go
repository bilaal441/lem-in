package main

// func ReduceInt(a []int, f func(int, int) int) int {
// 	occ := a[0]
// 	for _, curr := range a[1:] {
// 		occ = f(occ, curr)

// 	}
// 	return occ
// }

// func Chunk(slice []int, size int) [][]int {
// 	if size == 0 {
// 		return [][]int{}

// 	}
// 	val := len(slice) / size
// // 	s := [][]int{}

// 	for i := 0; i < val; i++ {
// 		s = append(s, slice[i*size:size*(i+1)])
// 	}

// 	if len(slice)%size != 0 {
// 		s = append(s, slice[size*val:])
// 	}

// 	return s

// }

// func mostPoints(questions [][]int) int64 {

// 	if len(questions) == 1 {
// 		fmt.Println(questions)
// 		return int64(questions[0][0])
// 	}

// 	val := int64(0)

// 	for i, curr := range questions {
// 		point := curr[0]

// 		brainPower := curr[1] + i + 1
// 		if len(questions) > brainPower {
// 			fmt.Println("point", curr, "next", questions[brainPower])
// 			val = int64(math.Max(float64(val), float64(point)+float64(questions[brainPower][0])))
// 		}

// 	}

// 	return val

// }

func findNonOverlappingPaths(paths [][]string) [][]string {
	
	nonOverlappingPaths := [][]string{}

	for _, path := range paths {
		overlap := false
		for _, node := range path {
			if _, ok := visitedNodes[node]; !ok {
				visitedNodes[node] = make(map[string]bool)
			}
			for _, visitedPath := range visitedNodes[node] {
				if visitedPath {
					overlap = true
					break
				}
			}
			if overlap {
				break
			}
			for _, n := range path {
		
				visitedNodes[n][node] = true
			}
		}
		if !overlap {
			nonOverlappingPaths = append(nonOverlappingPaths, path)
			for _, n := range path {
				visitedNodes[n][path[0]] = true
			}
		}
	}
	return nonOverlappingPaths
}
