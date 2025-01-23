package errchecker

func ContainsError(errMsg, expectedSubstring string) bool {
	return errMsg != "" && len(expectedSubstring) > 0 && len(errMsg) >= len(expectedSubstring) && (errMsg[:len(expectedSubstring)] == expectedSubstring)
}
