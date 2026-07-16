func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var Map = make(map[string]int)
	for _, c := range s {
		mapC := Map[string(c)]
		if mapC == 0 {
			Map[string(c)] = 1
			continue
		}

		Map[string(c)] = mapC + 1
	}

	for _, c := range t {
		mapC := Map[string(c)]
		if mapC == 0 {
			return false
		}

		Map[string(c)] = mapC - 1
	}

	return true
}