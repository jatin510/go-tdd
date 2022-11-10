package iteration

func Repeat(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}
