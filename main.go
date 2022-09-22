package gocombine

func pick(i, g, l int) int {
	return (i * l) % g
}
func Combinations(groups ...[]string) [][]string {
	total := 1
	for i, group := range groups {
		if len(group) == 0 {
			groups = append(groups[:i], groups[i+1:]...)
		} else {
			total *= len(group)
		}
	}
	// pre-allocate the results
	results := make([][]string, total)
	for i := range results {
		results[i] = make([]string, len(groups))
		for g, group := range groups {
			results[i][g] = group[pick(i, g, len(group))]
		}
	}
	return results
}
