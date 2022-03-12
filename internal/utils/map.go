package utils

func Map(vs []string, f func(string) string) []string {
	// https://stackoverflow.com/questions/33726731/short-way-to-apply-a-function-to-all-elements-in-a-list-in-golang
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
