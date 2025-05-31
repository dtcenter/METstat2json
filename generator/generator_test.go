package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"FCST_VALID", "fcstValid"},
		{"VERSION", "version"},
		{"N_VALID", "nValid"},
		{"STORM_ID", "stormId"},
		{"LINE_TYPE", "lineType"},
		{"INIT_MASK", "initMask"},
		{"N_CAT", "nCat"},
		{"DIAG_N", "diagN"},
		{"", ""},
		{"A", "a"},
		{"FOO_BAR_BAZ", "fooBarBaz"},
		{"FOO__BAR", "fooBar"},
		{"FOO BAR", "fooBar"},
		{"FOO-BAR", "fooBar"},
		{"INIT_12345", "init12345"},
		{"F1_O2", "f1O2"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := toCamelCase(test.input)
			assert.Equal(t, test.expected, result, "toCamelCase(%q) should return %q", test.input, test.expected)
		})
	}
}

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
