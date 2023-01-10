package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToArrayConversion(t *testing.T) {
	cases := []struct {
		names    string
		expected []string
	}{
		{names: "", expected: []string{""}},
		{names: "Antanoliy,Amlas,Savla", expected: []string{"Antanoliy", "Amlas", "Savla"}},
	}

	for _, item := range cases {
		t.Run("", func(t *testing.T) {
			actual, _ := StringToArrayConversion(item.names)

			assert.Equal(t, item.expected, actual)
		})
	}
}
