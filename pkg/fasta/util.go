package fasta

var (
	// FASTA file extensions:
	//
	// - fasta, fas, fa[13] => generic.
	// - fna => nucleic acid.
	// - ffn => nucleotide of gene regions.
	// - faa => amino acid.
	// - mpfa => amino acids.
	// - frn => non-coding RNA.
	Extensions = []string{
		".fasta", ".fas", ".fa",
		".fna",
		".ffn",
		".faa",
		".mpfa",
		".frn",
	}
)
