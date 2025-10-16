package localrepo

func Truncate(s string, maxLen int, suffix string) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + suffix
}
