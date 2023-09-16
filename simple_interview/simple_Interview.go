package simple_interview

func copyMap(a map[string]bool) map[string]bool {

	b := make(map[string]bool)

	for key, value := range a {
		b[key] = value
	}
	return b
}
