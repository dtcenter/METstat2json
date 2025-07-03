package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSortedKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected []string
	}{
		{
			name:     "empty map",
			input:    map[string]string{},
			expected: []string{},
		},
		{
			name:     "single item",
			input:    map[string]string{"a": "value"},
			expected: []string{"a"},
		},
		{
			name: "multiple items unsorted",
			input: map[string]string{
				"c": "value3",
				"a": "value1",
				"b": "value2",
			},
			expected: []string{"a", "b", "c"},
		},
		{
			name: "items with different cases",
			input: map[string]string{
				"Z": "valueZ",
				"a": "valuea",
				"B": "valueB",
			},
			expected: []string{"B", "Z", "a"},
		},
		{
			name: "special characters and numbers",
			input: map[string]string{
				"key-3": "value3",
				"key_1": "value1",
				"key2":  "value2",
			},
			expected: []string{"key-3", "key2", "key_1"},
		},
		{
			name: "MET header struct names",
			input: map[string]string{
				"MTD_3D_PS_header": "struct1",
				"MODE_OBJ_header":  "struct2",
				"TCST_PROBRIRW":    "struct3",
				"STAT_CTC":         "struct4",
				"GRID_SAL_header":  "struct5",
			},
			expected: []string{
				"GRID_SAL_header",
				"MODE_OBJ_header",
				"MTD_3D_PS_header",
				"STAT_CTC",
				"TCST_PROBRIRW",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getSortedKeys(test.input)
			assert.Equal(t, test.expected, result, "getSortedKeys() should return sorted keys")
		})
	}
}
