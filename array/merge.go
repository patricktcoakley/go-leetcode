package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	end := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[end] = nums1[i]
			i -= 1
		} else {
			nums1[end] = nums2[j]
			j -= 1
		}
		end -= 1
	}
}
