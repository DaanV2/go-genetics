package nucleic_acid_codes_test

import (
	"fmt"
	"slices"
	"testing"

	nucleic_acid_codes "github.com/DaanV2/go-genetics/pkg/nucleic_acid_codes"
	"github.com/stretchr/testify/require"
)

func Test_Code_Data(t *testing.T) {
	for _, c := range nucleic_acid_codes.Codes() {
		t.Run(fmt.Sprintf("Code(%s)", c), func(t *testing.T) {
			d := c.Data()
			nc := nucleic_acid_codes.From(d)

			require.Equal(t, nc, c)

			str := c.String()
			nc, err := nucleic_acid_codes.Parse(str)
			require.NoError(t, err)
			require.Equal(t, nc, c)

			require.True(t, c.Valid())
			require.NotEqual(t, c.Description(), "Unknown")
		})
	}
}

func Test_Code_All(t *testing.T) {
	validCodes := nucleic_acid_codes.Codes()
	validBytes := make([]byte, 0, len(validCodes))
	for _, c := range validCodes {
		validBytes = append(validBytes, c.Data())
	}

	for b := range 256 {
		valid := slices.Contains(validBytes, byte(b))
		code, err := nucleic_acid_codes.ParseByte(byte(b))

		if valid {
			require.NoError(t, err)
			require.Equal(t, code.Data(), byte(b))
			require.NotEqual(t, code.Description(), "Unknown")
		} else {
			require.Error(t, err)
			require.Equal(t, code.Description(), "Unknown")
		}
	}
}
