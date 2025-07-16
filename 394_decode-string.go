import (
	"strconv"
)

type item struct {
	n     int
	bytes []byte
}

func decodeString(s string) string {
	num := 0
	stack := []item{{1, []byte{}}}

	for i := range s {
		c := s[i]

		if c <= '9' && c >= '0' {
			j, _ := strconv.Atoi(string(c))
			num = num*10 + j
			continue
		}

		if c == '[' {
			stack = append(stack, item{num, []byte{}})
			num = 0
			continue
		}

		if c == ']' {
			i := stack[len(stack)-1]

			for i.n > 0 {
				stack[len(stack)-2].bytes = append(stack[len(stack)-2].bytes, i.bytes...)
				i.n = i.n - 1
			}

			stack = stack[:len(stack)-1]
			continue
		}

		stack[len(stack)-1].bytes = append(stack[len(stack)-1].bytes, c)
	}

	return string(stack[0].bytes)
}
