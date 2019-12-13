package csvsplitter

// Split is a representation of a sub file
type Split struct {
	Key       string
	HeaderRow []string
	Rows      [][]string
}

// Equals evaluates the value equality of a Split object
func (a *Split) Equals(b Split) bool {
	if a.Key != b.Key {
		return false
	}

	if a.HeaderRow == nil && b.HeaderRow != nil {
		return false
	} else if a.HeaderRow != nil && b.HeaderRow == nil {
		return false
	} else if a.HeaderRow != nil && b.HeaderRow != nil {
		if len(a.HeaderRow) != len(b.HeaderRow) {
			return false
		}

		for i, column := range a.HeaderRow {
			if column != b.HeaderRow[i] {
				return false
			}
		}
	}

	if a.Rows == nil && b.Rows != nil {
		return false
	} else if a.Rows != nil && b.Rows == nil {
		return false
	} else if a.Rows != nil && b.Rows != nil {
		if len(a.Rows) != len(b.Rows) {
			return false
		}

		for i, row := range a.Rows {
			if len(row) != len(b.Rows[i]) {
				return false
			}

			for j, column := range row {
				if column != b.Rows[i][j] {
					return false
				}
			}
		}
	}
	return true
}

// SplitRows sepearates rows into multiple files.
func SplitRows(rows [][]string, hasHeaderRows bool, keyColumn int) map[string]Split {
	if rows == nil {
		return nil
	}
	var headerRow []string
	if hasHeaderRows {
		headerRow = rows[0]
	}

	result := make(map[string]Split)

	for i, row := range rows {
		if hasHeaderRows && i == 0 {
			continue
		}

		s, found := result[row[keyColumn]]
		if found {
			s.Rows = append(s.Rows, row)
			result[row[keyColumn]] = s
		} else {
			result[row[keyColumn]] = Split{
				Key:       row[keyColumn],
				HeaderRow: headerRow,
				Rows:      [][]string{row},
			}
		}
	}

	return result
}
