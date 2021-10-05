package iteration

const repeatCount = 5

// Repeat returns the repeated (5 times) supplied char
func Repeat(char string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += char
	}

	return repeated
}
