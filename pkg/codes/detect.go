package codes

import (
	amino_acid_codes "github.com/DaanV2/go-genetics/pkg/codes/amino_acid"
	nucleic_acid_codes "github.com/DaanV2/go-genetics/pkg/codes/nucleic_acid"
)

type CodeType int

const (
	NEITHER CodeType = 0
	NUCLEIC CodeType = 1
	AMINO   CodeType = 2
)

func Detect(sequence string) CodeType {
	for i := range sequence {
		b := sequence[i]
		nucleic := nucleic_acid_codes.Valid(nucleic_acid_codes.From(b))
		amino := amino_acid_codes.Valid(amino_acid_codes.From(b))

		switch {
		case nucleic && amino:
			continue
		case !nucleic && amino:
			return AMINO
		case nucleic && !amino:
			return NUCLEIC
		case !nucleic && !amino:
			return NEITHER
		}
	}

	return NEITHER
}

func DetectBytes(sequence []byte) CodeType {
	for _, b := range sequence {
		nucleic := nucleic_acid_codes.Valid(nucleic_acid_codes.From(b))
		amino := amino_acid_codes.Valid(amino_acid_codes.From(b))

		switch {
		case nucleic && amino:
			continue
		case !nucleic && amino:
			return AMINO
		case nucleic && !amino:
			return NUCLEIC
		case !nucleic && !amino:
			return NEITHER
		}
	}

	return NEITHER
}
