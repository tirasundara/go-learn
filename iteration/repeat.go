package iteration

// Repeat repeats the string s by n times, and return it
func Repeat(s string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += s
	}

	return repeated
}
