package csvsplitter

// Split is a representation of a sub file
type Split struct {
	key       string
	headerRow []string
	rows      [][]string
}

// SplitRows sepearates rows into multiple files.
func SplitRows(rows [][]string, hasHeaderRows bool, keyColumn int) []Split {
	if rows == nil {
		return nil
	}

	// headerRow := rows[0]

	result := []Split{}

	return result
}
