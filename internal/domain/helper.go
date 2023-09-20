package domain

func contains(v string, e []string) bool {
	for _, s := range e {
		if v == s {
			return true
		}
	}
	return false
}
