package amino_acid_codes_test

import (
	"testing"

	"github.com/DaanV2/go-genetics/pkg/amino_acid_codes"
	"github.com/stretchr/testify/require"
)

func Test_Convert(t *testing.T) {
	validCodes := amino_acid_codes.Codes()
	validBytes := make([]byte, 0)

	for _, code := range validCodes {
		validBytes = append(validBytes, code.Data())
	}

	str := string(validBytes)
	conv, err := amino_acid_codes.Convert(str)

	require.NoError(t, err)
	require.Equal(t, conv, validCodes)
}