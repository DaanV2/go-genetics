package nucleic_acid_codes

const (
	Adenine                  Code = 'A' // Adenine
	Cytosine                 Code = 'C' // Cytosine
	Guanine                  Code = 'G' // Guanine
	Thymine                  Code = 'T' // Thymine
	Uracil                   Code = 'U' // Uracil
	Inosine                  Code = 'i' // inosine (non-standard)
	PuRine                   Code = 'R' // A or G (I)	puRine
	PYrimidines              Code = 'Y' // C, T or U	pYrimidines
	Ketones                  Code = 'K' // G, T or U	bases which are Ketones
	AdenineOrCytosine        Code = 'M' // A or C	bases with aMino groups
	CytosineOrGuanine        Code = 'S' // C or G	Strong interaction
	AdenineOrThymineOrUracil Code = 'W' // A, T or U	Weak interaction
	NotAdenine               Code = 'B' // not A (i.e. C, G, T or U)	B comes after A
	NotCytosine              Code = 'D' // not C (i.e. A, G, T or U)	D comes after C
	NotGuanine               Code = 'H' // not G (i.e., A, C, T or U)	H comes after G
	NeitherThymineOrUracil   Code = 'V' // neither T nor U (i.e. A, C or G)	V comes after U
	Any                      Code = 'N' // A C G T U	Nucleic acid
	Gap                      Code = '-' // gap of indeterminate length
)

// Codes returns an slice of all the codes that exists
func Codes() []Code {
	return []Code{
		Adenine,
		Cytosine,
		Guanine,
		Thymine,
		Uracil,
		Inosine,
		PuRine,
		PYrimidines,
		Ketones,
		AdenineOrCytosine,
		CytosineOrGuanine,
		AdenineOrThymineOrUracil,
		NotAdenine,
		NotCytosine,
		NotGuanine,
		NeitherThymineOrUracil,
		Any,
		Gap,
	}
}

// Valid checks if the given code is a valid nucleic acid code or not
func Valid(code Code) bool {
	switch code {
	case Adenine:
	case Cytosine:
	case Guanine:
	case Thymine:
	case Uracil:
	case Inosine:
	case PuRine:
	case PYrimidines:
	case Ketones:
	case AdenineOrCytosine:
	case CytosineOrGuanine:
	case AdenineOrThymineOrUracil:
	case NotAdenine:
	case NotCytosine:
	case NotGuanine:
	case NeitherThymineOrUracil:
	case Any:
	case Gap:
	default:
		return false
	}

	return true
}

// Description returns a string description of the code
func Description(code Code) string {
	switch code {
	case Adenine:
		return "Adenine"
	case Cytosine:
		return "Cytosine"
	case Guanine:
		return "Guanine"
	case Thymine:
		return "Thymine"
	case Uracil:
		return "Uracil"
	case Inosine:
		return "Inosine"
	case PuRine:
		return "Adenine or Guanine (Inosine),puRine"
	case PYrimidines:
		return "Cytosine or Thymine or U,pYrimidines"
	case Ketones:
		return "Guanine or Thymine or Uracil,bases which are Ketones"
	case AdenineOrCytosine:
		return "Adenine or Cytosine,bases with aMino groups"
	case CytosineOrGuanine:
		return "Cytosine or Guanine,Strong interaction"
	case AdenineOrThymineOrUracil:
		return "Adenine or Thymine or Uracil,Weak interaction"
	case NotAdenine:
		return "Not Adenine"
	case NotCytosine:
		return "Not Cytosine"
	case NotGuanine:
		return "Not Guanine"
	case NeitherThymineOrUracil:
		return "neither Thymine nor Uracil"
	case Any:
		return "Adenine or Cytosine or Guanine or Thymine or Uracil"
	case Gap:
		return "Gap of indeterminate length"
	}

	return "Unknown"
}
