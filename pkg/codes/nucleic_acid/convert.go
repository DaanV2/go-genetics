package nucleic_acid_codes

import xerrors "github.com/DaanV2/go-genetics/pkg/x/errors"

func Convert(data string) ([]Code, error) {
	return ConvertBytes([]byte(data))
}

func ConvertBytes(data []byte) ([]Code, error) {
	codes := make([]Code, len(data))
	var err error

	for i, b := range data {
		codes[i], err = ParseByte(b)
		if err != nil {
			return codes, xerrors.NewIndexedError(i, err)
		}
	}

	return codes, nil
}
