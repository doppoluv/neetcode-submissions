func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    if len(strs) == 1 {
        return strs[0]
    }

    prefix := ""
    for i, c := range strs[0] {
        if len(strs[1]) <= i || rune(strs[1][i]) != c {
            break
        }

        prefix += string(c) 
    }

    for j := 2; j < len(strs); j++ {
        if prefix == "" {
            break
        }

        newPrefix := ""
        for k, c := range prefix {
            if len(strs[j]) <= k || rune(strs[j][k]) != c {
                break
            }

            newPrefix += string(c)
        }

        prefix = newPrefix
    }

    return prefix
}