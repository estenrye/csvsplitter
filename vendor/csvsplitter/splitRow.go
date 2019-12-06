package csvsplitter

// Split is a representation of a sub file
type Split struct {
	key       string
	headerRow []string
	rows      [][]string
}

// Equals evaluates the value equality of a Split object
func (a *Split) Equals(b Split) bool {
	if a.key != b.key {
		return false
	}

	if a.headerRow == nil && b.headerRow != nil {
		return false
	} else if a.headerRow != nil && b.headerRow == nil {
		return false
	} else if a.headerRow != nil && b.headerRow != nil {
		if len(a.headerRow) != len(b.headerRow) {
			return false
		}

		for i, column := range a.headerRow {
			if column != b.headerRow[i] {
				return false
			}
		}
	}

	if a.rows == nil && b.rows != nil {
		return false
	} else if a.rows != nil && b.rows == nil {
		return false
	} else if a.rows != nil && b.rows != nil {
		if len(a.rows) != len(b.rows) {
			return false
		}

		for i, row := range a.rows {
			if len(row) != len(b.rows[i]) {
				return false
			}

			for j, column := range row {
				if column != b.rows[i][j] {
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
			s.rows = append(s.rows, row)
			result[row[keyColumn]] = s
		} else {
			result[row[keyColumn]] = Split{
				key:       row[keyColumn],
				headerRow: headerRow,
				rows:      [][]string{row},
			}
		}
	}

	return result
}
