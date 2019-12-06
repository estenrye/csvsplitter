package csvsplitter

import (
	"testing"
)

func TestEqual(t *testing.T) {
	a := Split{
		key:       "key3",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}
	var b Split

	if a.Equals(b) {
		t.Errorf("Nil equality check failed.")
	}

	if !a.Equals(a) {
		t.Errorf("Equality check failed")
	}
}

func TestNotEqual(t *testing.T) {
	a := Split{
		key:       "key2",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}
	b := Split{
		key:       "key3",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(b) {
		t.Errorf("key equality check failure.")
	}

	c := Split{
		key:       "key3",
		headerRow: []string{"column 5", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(c) {
		t.Errorf("header row equality failure.")
	}

	d := Split{
		key:       "key3",
		headerRow: nil,
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(d) {
		t.Errorf("header row equality failure.")
	}

	e := Split{
		key:       "key3",
		headerRow: []string{"column 1", "column 2"},
		rows: [][]string{
			[]string{"key3", "value"},
			[]string{"key3", "value"},
			[]string{"key3", "value"},
		},
	}

	if a.Equals(e) {
		t.Errorf("header row equality failure.")
	}

	f := Split{
		key:       "key2",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(f) {
		t.Errorf("row equality failure.")
	}

	g := Split{
		key:       "key2",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows:      nil,
	}

	if a.Equals(g) {
		t.Errorf("row equality failure.")
	}

	h := Split{
		key:       "key2",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value1", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(h) {
		t.Errorf("row equality failure.")
	}

	i := Split{
		key:       "key2",
		headerRow: []string{"column 1", "column 2", "column 3"},
		rows: [][]string{
			[]string{"key3", "value", "value"},
			[]string{"key3", "value"},
			[]string{"key3", "value", "value"},
		},
	}

	if a.Equals(i) {
		t.Errorf("row equality failure.")
	}

}

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

	expectWithHeaders := map[string]Split{
		"key3": Split{
			key:       "key3",
			headerRow: []string{"column 1", "column 2", "column 3"},
			rows: [][]string{
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
			},
		},
		"key1": Split{
			key:       "key1",
			headerRow: []string{"column 1", "column 2", "column 3"},
			rows: [][]string{
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
			},
		},
		"key2": Split{
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

	expectWithoutHeaders := map[string]Split{
		"key3": Split{
			key:       "key3",
			headerRow: nil,
			rows: [][]string{
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
				[]string{"key3", "value", "value"},
			},
		},
		"key1": Split{
			key:       "key1",
			headerRow: nil,
			rows: [][]string{
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
				[]string{"key1", "value", "value"},
			},
		},
		"key2": Split{
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

func testSplitRows(t *testing.T, input [][]string, hasHeaderRows bool, keyColumn int, expect map[string]Split) {

	got := SplitRows(input, hasHeaderRows, keyColumn)

	if got == nil {
		t.Errorf("SplitRows([][]string) returned nil.  Expected a non-nil value.")
	} else if len(expect) != len(got) {
		t.Errorf("SplitRows return a result with length of %d.  Expected %d", len(got), len(expect))
	} else {
		for i, s := range expect {
			if !s.Equals(got[i]) {
				t.Errorf("expect[%s] != got[%s]", i, i)
			}
		}
	}
}
