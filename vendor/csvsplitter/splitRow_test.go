package csvsplitter

import (
	"testing"
)

func TestSplitRowsWithHeaderRow(t *testing.T) {
	inputWithHeaders := [][]string{
		[]string{"column 1", "column 2", "column 3"},
		[]string{"key3", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key3", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key3", "value", "value"},
	}

	expectWithHeaders := []Split{
		Split{
			key:       "key3",
			headerRow: []string{"column 1", "column 2", "column 3"},
			rows: [][]string{
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
			},
		},
		Split{
			key:       "key1",
			headerRow: []string{"column 1", "column 2", "column 3"},
			rows: [][]string{
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
			},
		},
		Split{
			key:       "key2",
			headerRow: []string{"column 1", "column 2", "column 3"},
			rows: [][]string{
				[]string{"key2", "value", "value"},
				[]string{"key2", "value", "value"},
				[]string{"key2", "value", "value"},
			},
		},
	}

	testSplitRows(t, inputWithHeaders, true, 0, expectWithHeaders)
}

func TestSplitRowsWithoutHeaderRow(t *testing.T) {
	inputWithoutHeaders := [][]string{
		[]string{"key3", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key3", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key1", "value", "value"},
		[]string{"key2", "value", "value"},
		[]string{"key3", "value", "value"},
	}

	expectWithoutHeaders := []Split{
		Split{
			key:       "key3",
			headerRow: nil,
			rows: [][]string{
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
			},
		},
		Split{
			key:       "key1",
			headerRow: nil,
			rows: [][]string{
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
			},
		},
		Split{
			key:       "key2",
			headerRow: nil,
			rows: [][]string{
				[]string{"key2", "value", "value"},
				[]string{"key2", "value", "value"},
				[]string{"key2", "value", "value"},
			},
		},
	}

	testSplitRows(t, inputWithoutHeaders, false, 0, expectWithoutHeaders)
}

func testSplitRows(t *testing.T, input [][]string, hasHeaderRows bool, keyColumn int, expect []Split) {

	got := SplitRows(input, hasHeaderRows, keyColumn)

	if got == nil {
		t.Errorf("SplitRows([][]string) returned nil.  Expected a non-nil value.")
	} else if len(expect) != len(got) {
		t.Errorf("SplitRows return a result with length of %d.  Expected %d", len(got), len(expect))
	} else {
		for i, s := range expect {
			if s.key != got[i].key {
				t.Errorf(
					"SplitRows([][]string): expect[%d].key != got[%d].key;  expected: %q, got: %q",
					i, i, s.key, got[i].key)
			}

			if s.headerRow == nil && got[i].headerRow != nil {
				t.Errorf("SplitRows([][]string): got[%d].headerRow is not nil.  Expected nil.", i)
			} else if s.headerRow != nil && got[i].headerRow == nil {
				t.Errorf("SplitRows([][]string): got[%d].headerRow is nil.  Expected non-nil value.", i)
			} else if s.headerRow != nil && got[i].headerRow != nil {
				for h, c := range s.headerRow {
					if c != got[i].headerRow[h] {
						t.Errorf(
							"SplitRows([][]string): expect[%d].headerRow[%d] != got[%d].headerRow[%d]; expected: %q, got: %q",
							i, h, i, h, c, got[i].headerRow[h],
						)
					}
				}
			}

			if s.rows == nil && got[i].rows != nil {
				t.Errorf(
					"splitRows([][]string): got[%d].rows is not nil.  Expected nil.",
					i,
				)
			} else if s.rows != nil && got[i].rows == nil {
				t.Errorf(
					"splitRows([][]string): got[%d].rows is nil.  Expected non-nil value.",
					i,
				)
			} else if s.rows != nil && got[i].rows != nil {
				for j, r := range s.rows {
					if got[i].rows[j][keyColumn] != s.key {
						t.Errorf("SplitRows([][]string): Key Column Does not match expected key value.")
					}
					for k, c := range r {
						if c != got[i].rows[j][k] {
							t.Errorf(
								"splitRows([][]string): expect[%d].rows[%d][%d] != got[%d].rows[%d][%d]; expected: %q, got: %q",
								i, j, k, i, j, k, c, got[i].rows[j][k],
							)
						}
					}
				}
			}

		}
	}
}
