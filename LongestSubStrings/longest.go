package main

import (
	"fmt"
)

func main() {
	//ml := lengthOfLongestSubstring("jbpnbwwd")
	//fmt.Println("maxlength:", ml)

	n1 := []int{1, 4}
	n2 := []int{2, 3}
	r := findMedianSortedArrays(n1, n2)
	fmt.Println("median is: ", r)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := nums1[:]
	arr = append(arr, nums2...)

	for _, v := range arr {
		fmt.Println(v)
	}
	return 1.0
}

func lengthOfLongestSubstring(s string) int {

	strMap := make(map[string]int)
	maxLength := 0

	for i := 0; i < len(s); i++ {
		ss := s[i:]
		for _, ch := range ss {

			if _, ok := strMap[string(ch)]; ok == false {
				strMap[string(ch)]++
				//maxLength = len(strMap)

			} else {
				// clear the map
				for key, _ := range strMap {
					delete(strMap, key)
				}
				break

			}
			if length := len(strMap); length > maxLength {
				maxLength = length
			}

		}

	}

	return maxLength

}
