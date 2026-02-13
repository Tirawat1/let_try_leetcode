// func bubbleSort(arr []int) {

// 	for i := 0; i < len(arr)-1; i++ {
// 		// -i เพราะเราเอา ตัวที่มาที่สุดไปตำแหน่งท้ายสุดแล้ว
// 		for j := 0; j < len(arr)-1-i; j++ {
// 			if arr[j] > arr[j+1] {
// 				var temp = arr[j]
// 				arr[j] = arr[j+1]
// 				arr[j+1] = temp
// 			}
// 		}
// 	}

// }

func findContentChildren(g []int, s []int) int {

	//sort arr g and s first
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	j := 0
	count := 0
	for {
		if i > len(g)-1 || j > len(s)-1 {
			break
		}
		// [1,2,3,4]
		// [1,1,3]
		if s[j] >= g[i] {
			i++
			j++
			count++
		} else {
			j++
		}

	}

	return count
}