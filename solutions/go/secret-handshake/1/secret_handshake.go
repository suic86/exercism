package secret

func Handshake(code uint) []string {
	mapping := []string{"wink", "double blink", "close your eyes", "jump"}
	result := make([]string, 0, 4)
	if (code>>4)&1 != 1 {
		for i := 0; i < 4; i++ {
			if (code>>i)&1 == 1 {
				result = append(result, mapping[i])
			}
		}
	} else {
		for i := 3; i >= 0; i-- {
			if (code>>i)&1 == 1 {
				result = append(result, mapping[i])
			}
		}
	}
	return result
}
