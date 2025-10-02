package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	res := []rune{}
	s := rune(shiftKey)
	for _, c := range plain {
		switch {
		case 'a' <= c && c <= 'z':
			res = append(res, (c-'a'+s)%26+'a')
		case 'A' <= c && c <= 'Z':
			res = append(res, (c-'A'+s)%26+'A')
		default:
			res = append(res, c)
		}
	}
	return string(res)
}
