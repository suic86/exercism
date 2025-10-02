package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if isSublist(l1, l2) {
		if len(l1) == len(l2) {
			return RelationEqual
		}
		return RelationSublist
	}

	if isSublist(l2, l1) {
		return RelationSuperlist
	}

	return RelationUnequal
}

func isSublist(l1, l2 []int) bool {
	ll1 := len(l1)
	ll2 := len(l2)
	if ll1 > ll2 {
		return false
	}
	if ll1 == 0 {
		return true
	}

	for i, v := range l2 {
		if v == l1[0] {
			if i+ll1 > ll2 {
				break
			}
			sublist := true
			for j, v2 := range l1 {
				if l2[i+j] != v2 {
					sublist = false
					break
				}
			}
			if sublist {
				return true
			}
		}
	}

	return false
}
