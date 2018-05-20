package iteration

func Repeat(charachter string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += charachter
	}

	return repeated
}
