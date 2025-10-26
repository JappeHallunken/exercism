package secret

func Handshake(code uint) []string {
	actions := []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
		"reverse",
	}

	if code < 1 || code > 31 {
		return nil
	}

	var hs []string
	for i := range 4 {
		if (code>>i)&1 == 1 {
			hs = append(hs, actions[i])
		}
	}

	if (code>>4)&1 == 1 {
		for i, j := 0, len(hs)-1; i < j; i, j = i+1, j-1 {
			hs[i], hs[j] = hs[j], hs[i]
		}
	}
	return hs
}
