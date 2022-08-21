package main

// matchAst reports where a c*regexp matches at the beggining
// of the text
func matchAst(c byte, re string, txt string) bool {
	for txt != "" && (txt[0] == c || c == '.') {
		if matchHere(re, txt) {
			return true
		}

		txt = txt[1:]
	}

	return matchHere(re, txt)
}

// matchHere reports whether the regexp matches at the
// beginning of the text
func matchHere(re string, txt string) bool {
	switch {
	case re == "":
		return true
	case re == "$":
		return txt == ""
	case len(re) >= 2 && re[1] == '*':
		return matchAst(re[0], re[2:], txt)
	case txt != "" && (re[0] == '.' || re[0] == txt[0]):
		return matchHere(re[1:], txt[1:])
	default:
		return false
	}
}

// Match reports if the regexp matches
func Match(re string, txt string) bool {
	if re != "" && re[0] == '^' {
		return matchHere(re[1:], txt)
	}

	for txt != "" {
		if matchHere(re, txt) {
			return true
		}

		txt = txt[1:]
	}

	return matchHere(re, txt)
}
