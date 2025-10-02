package isbn

func IsValidISBN(isbn string) bool {
	if len(isbn) == 0 {
		return false
	}

	lc := isbn[len(isbn)-1]
	if (lc < '0' || lc > '9') && lc != 'X' {
		return false
	}

	i := 10
	res := 0

	for _, c := range isbn {
		if i < 1 {
			return false
		}
		switch {
		case '0' <= c && c <= '9':
			res += int(c-'0') * i
			i--
		case c == 'X':
			if i != 1 {
				return false
			}
			res += 10 * i
			i--
		case c == '-':
		default:
			return false
		}
	}
	if i != 0 {
		return false
	}

	return res%11 == 0
}
