package fasta

import "strings"

const (
	cutSetLeft = ";>"
)

// Trim the given strings from any white spaces, but also comment and description prefixes
func Trim(line string) string {
	line = strings.TrimLeft(line, cutSetLeft)
	line = strings.TrimSpace(line)

	return line
}
