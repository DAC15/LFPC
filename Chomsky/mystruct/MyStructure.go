package mystruct

type MapsWithDuplicate struct {
	Symbols string
	Values  string
}


func createMyMap(lines []string, myarray *[]MapsWithDuplicate) *[]MapsWithDuplicate {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '-' {
				*myarray = append(*myarray, MapsWithDuplicate{lines[i][:j], lines[i][j+1:]})
			}
		}
	}
	return myarray
}