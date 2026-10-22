func romanToInt(s string) int {
	roman := map[byte]int{
        'I': 1,
		'V': 5,
		'X': 10,
        'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
    }

    sum := 0
    lastNum := 1000
    
	for _, c := range s {
        num := roman[byte(c)]

        if lastNum < num {
            sum = sum - lastNum - lastNum
        }

        lastNum = num
        sum += num
    }

    return sum
}