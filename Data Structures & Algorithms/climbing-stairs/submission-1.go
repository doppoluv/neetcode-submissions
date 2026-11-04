func climbStairs(n int) int {
    cache := make(map[int]int, n)

    return ways(n, cache)
}

func ways(n int, cache map[int]int) int {
    if n == 1 {
        return 1
    }

    if n == 2 {
        return 2
    }

    if cache[n] == 0 {
        cache[n] = ways(n - 2, cache) + ways(n - 1, cache)
    }

    return cache[n]
}