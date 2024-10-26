package nucleic_acids_codes_test

import (
	"testing"

	nucleic_acids_codes "github.com/DaanV2/go-genetics/pkg/nucleic_acid_codes"
	"github.com/stretchr/testify/require"
)

func Test_Convert(t *testing.T) {
	validCodes := nucleic_acids_codes.Codes()
	validBytes := make([]byte, 0)

	for _, code := range validCodes {
		validBytes = append(validBytes, code.Data())
	}

	str := string(validBytes)
	conv, err := nucleic_acids_codes.Convert(str)

	require.NoError(t, err)
	require.Equal(t, conv, validCodes)
}