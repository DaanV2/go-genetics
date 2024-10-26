package amino_acid_codes

func ValidSequence(sequence string) bool {
	for i := range sequence {
		if !Valid(From(sequence[i])) {
			return false
		}
	}

	return true
}

func ValidSequenceBytes(sequence []byte) bool {
	for _, b := range sequence {
		if !Valid(From(b)) {
			return false
		}
	}

	return true
}
