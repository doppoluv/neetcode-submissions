func groupAnagrams(strs []string) [][]string {
	Map := make(map[[26]int][]string)
	for _, str := range strs {
		chars := [26]int{}
		for n := 0; n < len(str); n++ {
			chars[str[n]-'a']++
		}

		Map[chars] = append(Map[chars], str)
	}

	Arr := make([][]string, len(Map))
	idx := 0
	for _, angrm := range Map {
		Arr[idx] = angrm
		idx++
	}

	return Arr
}