package libappimagego

func intToBool(x int) bool {
	if x == 0 {
		return false
	} else {
		return true
	}
}

func boolToInt(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}
