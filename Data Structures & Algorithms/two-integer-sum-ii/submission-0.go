func twoSum(numbers []int, target int) []int {
    m := 0
    n := len(numbers)-1
    for m < len(numbers) && n >= 0 {
        if m == n {
            return []int{}
        }

        s := numbers[m] + numbers[n]
        
        if s > target {
            n--
            continue
        }

        if s < target {
            m++
            continue
        }

        return []int{m+1,n+1}
    }

    return []int{}
}