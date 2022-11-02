package Iteration

func Repeat(character string, times int) (res string) {
	for i := 0; i < times; i++ {
		res += character
	}
	return
}
