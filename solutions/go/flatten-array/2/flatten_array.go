package flatten

func Flatten(nested any) []any {
	res := make([]any, 0)
	switch v := nested.(type) {
	case []any:
		for _, e := range v {
			res = append(res, Flatten(e)...)
		}
	case any:
		res = append(res, v)
	}
	return res
}
