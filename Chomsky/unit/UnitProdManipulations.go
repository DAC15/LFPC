package unit

import (
	structure "awesomeProject/mystruct"
	utils "awesomeProject/utils"
)

func MakeUnitSubstitution(myarray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate {
	var arrayToDelete []structure.MapsWithDuplicate
	for i := len(*myarray) - 1; i > 0; i-- {
		if isUnitProduction((*myarray)[i]) {
			arrayToDelete = append(arrayToDelete,
				structure.MapsWithDuplicate{Symbols: (*myarray)[i].Symbols, Values: (*myarray)[i].Values})
			for j := 0; j < len(getProductionBySymbol(*myarray, (*myarray)[i].Values)); j++ {
				*myarray = append(*myarray,
					structure.MapsWithDuplicate{Symbols: (*myarray)[i].Symbols, Values: getProductionBySymbol(*myarray, (*myarray)[i].Values)[j]})
			}
			utils.DeleteMultipleElements1(myarray, arrayToDelete)
			continue
		}
	}
	return myarray
}

func isUnitProduction(myarray structure.MapsWithDuplicate) bool {
	if (len(myarray.Symbols) == 1 && len(myarray.Values) == 1) &&
		(utils.IsNonTerminal(myarray.Symbols[0]) && utils.IsNonTerminal(myarray.Values[0])) {
		return true
	}

	return false
}

func getProductionBySymbol(myarray []structure.MapsWithDuplicate, symbol string) []string {
	var arrayProduction []string
	for i := 0; i < len(myarray); i++ {
		if myarray[i].Symbols == symbol {
			arrayProduction = append(arrayProduction, myarray[i].Values)
		}
	}
	return arrayProduction
}
