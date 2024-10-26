package amino_acid_codes

const (
	UNKNOWN_STRING = "Unknown"
)

const (
	Alanine              Code = 'A' // Alanine
	AsparticOrAsparagine Code = 'B' // Aspartic acid (D) or Asparagine (N)
	Cysteine             Code = 'C' // Cysteine
	Aspartic             Code = 'D' // Aspartic acid
	Glutamic             Code = 'E' // Glutamic acid
	Phenylalanine        Code = 'F' // Phenylalanine
	Glycine              Code = 'G' // Glycine
	Histidine            Code = 'H' // Histidine
	Isoleucine           Code = 'I' // Isoleucine
	LeucineOrIsoleucine  Code = 'J' // Leucine (L) or Isoleucine (I)
	Lysine               Code = 'K' // Lysine
	Leucine              Code = 'L' // Leucine
	Methionine           Code = 'M' // Methionine/Start codon
	Asparagine           Code = 'N' // Asparagine
	Pyrrolysine          Code = 'O' // Pyrrolysine (rare)
	Proline              Code = 'P' // Proline
	Glutamine            Code = 'Q' // Glutamine
	Arginine             Code = 'R' // Arginine
	Serine               Code = 'S' // Serine
	Threonine            Code = 'T' // Threonine
	Selenocysteine       Code = 'U' // Selenocysteine (rare)
	Valine               Code = 'V' // Valine
	Tryptophan           Code = 'W' // Tryptophan
	Tyrosine             Code = 'Y' // Tyrosine
	GlutamicOrGlutamine  Code = 'Z' // Glutamic acid (E) or Glutamine (Q)
	Any                  Code = 'X' // Any
	Stop                 Code = '*' // Translation stop
	Gap                  Code = '-' // Gap of indeterminate length
)

// Codes returns an slice of all the codes that exists
func Codes() []Code {
	return []Code{
		Alanine,
		AsparticOrAsparagine,
		Cysteine,
		Aspartic,
		Glutamic,
		Phenylalanine,
		Glycine,
		Histidine,
		Isoleucine,
		LeucineOrIsoleucine,
		Lysine,
		Leucine,
		Methionine,
		Asparagine,
		Pyrrolysine,
		Proline,
		Glutamine,
		Arginine,
		Serine,
		Threonine,
		Selenocysteine,
		Valine,
		Tryptophan,
		Tyrosine,
		GlutamicOrGlutamine,
		Any,
		Stop,
		Gap,
	}
}

// Valid checks if the given code is a valid amino acid code or not
func Valid(code Code) bool {
	switch code {
	case Alanine:
	case AsparticOrAsparagine:
	case Cysteine:
	case Aspartic:
	case Glutamic:
	case Phenylalanine:
	case Glycine:
	case Histidine:
	case Isoleucine:
	case LeucineOrIsoleucine:
	case Lysine:
	case Leucine:
	case Methionine:
	case Asparagine:
	case Pyrrolysine:
	case Proline:
	case Glutamine:
	case Arginine:
	case Serine:
	case Threonine:
	case Selenocysteine:
	case Valine:
	case Tryptophan:
	case Tyrosine:
	case GlutamicOrGlutamine:
	case Any:
	case Stop:
	case Gap:
	default:
		return false
	}

	return true
}

// Description returns a string description of the code
func Description(code Code) string {
	switch code {
	case Alanine:
		return "Alanine"
	case AsparticOrAsparagine:
		return "Aspartic acid (D) or Asparagine (N)"
	case Cysteine:
		return "Cysteine"
	case Aspartic:
		return "Aspartic acid"
	case Glutamic:
		return "Glutamic acid"
	case Phenylalanine:
		return "Phenylalanine"
	case Glycine:
		return "Glycine"
	case Histidine:
		return "Histidine"
	case Isoleucine:
		return "Isoleucine"
	case LeucineOrIsoleucine:
		return "Leucine (L) or Isoleucine (I)"
	case Lysine:
		return "Lysine"
	case Leucine:
		return "Leucine"
	case Methionine:
		return "Methionine/Start codon"
	case Asparagine:
		return "Asparagine"
	case Pyrrolysine:
		return "Pyrrolysine (rare)"
	case Proline:
		return "Proline"
	case Glutamine:
		return "Glutamine"
	case Arginine:
		return "Arginine"
	case Serine:
		return "Serine"
	case Threonine:
		return "Threonine"
	case Selenocysteine:
		return "Selenocysteine (rare)"
	case Valine:
		return "Valine"
	case Tryptophan:
		return "Tryptophan"
	case Tyrosine:
		return "Tyrosine"
	case GlutamicOrGlutamine:
		return "Glutamic acid (E) or Glutamine (Q)"
	case Any:
		return "Any"
	case Stop:
		return "Translation stop"
	case Gap:
		return "Gap of indeterminate length"
	}

	return UNKNOWN_STRING
}
