func romanToInt(s string) int {
    sum := 0
    lastNum := 1000
    for _, c := range s {
        num := 0
        switch string(c) {
            case "I":
                num = 1
            case "V":
                num = 5
            case "X":
                num = 10
            case "L":
                num = 50
            case "C":
                num = 100
            case "D":
                num = 500
            case "M":
                num = 1000
        }

        if lastNum < num {
            sum = sum - lastNum - lastNum
        }

        lastNum = num
        sum += num
    }

    return sum
}