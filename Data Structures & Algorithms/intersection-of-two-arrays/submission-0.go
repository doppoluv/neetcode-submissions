func intersection(nums1 []int, nums2 []int) []int {
    numbers := make(map[int]bool)
    result := make([]int, 0)


    for _, num := range nums1 {
        numbers[num] = true
    }

    for _, num := range nums2 {
        if numbers[num] {
            result = append(result, num)
            numbers[num] = false
        }
    }

    return result
}