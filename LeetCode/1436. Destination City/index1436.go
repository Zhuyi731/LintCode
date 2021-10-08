func destCity(paths [][]string) string {
	m1 := map[string]bool{}
	m2 := map[string]bool{}
	for _, v := range paths {
		m1[v[0]] = true
		m2[v[1]] = true
	}

	for v := range m2 {
		if !m1[v] {
			return v
		}
	}
	return ""
}