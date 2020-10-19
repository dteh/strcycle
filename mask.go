package strcycle

var illegalChars = map[string]bool{".": true, ",": true}

// SetIllegalCharacters set chars to skip in mask
func SetIllegalCharacters(skip []string) {
	illegalChars = map[string]bool{}
	for _, val := range skip {
		illegalChars[val] = true
	}
}

func incrementMask(mask []int, limit int) []int {
	i := 0
	for {
		// Skip illegal character
		if mask[i] == -1 {
			i++
			continue
		}

		// Increment mask value
		mask[i]++

		// Check if > limit, if so move on
		if mask[i] == limit {
			mask[i] = 0
			i++

			// Reached end of mask, start again
			if i == len(mask) {
				break
			}
		} else {
			break
		}
	}
	return mask
}

// check if mask is all 0s or -1s
func maskHasReset(mask []int) bool {
	for _, val := range mask {
		if val > 0 {
			return false
		}
	}
	return true
}
