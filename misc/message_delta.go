// for 2 given sequences s, check how many characters have been changed.
// The original message is 3 characters long

package misc

func message_delta(s string) int32 {
	ref := "SOS"
	var changed int32 = 0
	for idx, jdx := 0, 0; jdx < len(s); idx, jdx = (idx+1)%len(ref), jdx+1 {
		if ref[idx] != s[jdx] {
			changed += 1
		}
	}
	return changed
}
