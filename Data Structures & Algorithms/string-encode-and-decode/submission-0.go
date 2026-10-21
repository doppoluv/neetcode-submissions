type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	lenghts := ""
	strings := ""

	for _, str := range strs {
		lenghts = lenghts + strconv.Itoa(len(str)) + " "
		strings += str
	}

	return lenghts + "|" + strings
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "|" {
		return []string{}
	}
	
	decoded := make([]string, 0)
	runes := []rune(encoded)
	lenghts := []rune{}
	strings := []rune{}

	for i, c := range runes {
		if string(c) == "|" {
			lenghts = runes[0:i]
			strings = runes[i+1:len(runes)]
			break
		}
	}

	for string(lenghts) != "" {
		lenght := 0
		for i, c := range lenghts {
			if string(c) == " " {
				lenght, _ = strconv.Atoi(string(lenghts[0:i]))
				lenghts = lenghts[i+1:len(lenghts)]

				break
			}
		}

		decoded = append(decoded, string(strings[0:lenght]))
		strings = strings[lenght:len(strings)]
	}

	return decoded
}
