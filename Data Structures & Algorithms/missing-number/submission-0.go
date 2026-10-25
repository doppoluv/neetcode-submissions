func missingNumber(nums []int) int {
    n := make(map[int]bool)
    for _, num := range nums {
        n[num] = true
    }

    for i := 0; i <= len(nums); i++ {
        if !n[i] {
            return i
        }
    }

    return 0
}