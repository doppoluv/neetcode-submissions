func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var Map = make(map[rune]int)
	for _, c := range s {
		mapC := Map[c]
		if mapC == 0 {
			Map[c] = 1
			continue
		}

		Map[c] = mapC + 1
	}

	for _, c := range t {
		mapC := Map[c]
		if mapC == 0 {
			return false
		}

		Map[c] = mapC - 1
	}

	return true
}
