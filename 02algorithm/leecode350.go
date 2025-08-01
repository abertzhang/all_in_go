package main

func main() {
	//intersect()
	print("hell---o")
}

func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2
}
