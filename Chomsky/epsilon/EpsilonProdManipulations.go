package epsilon

import (
	structure "awesomeProject/mystruct"
	utils "awesomeProject/utils"
)

func RemoveEpsilonProduction(myArray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate {
	var arrayToRemove []structure.MapsWithDuplicate
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].Values == "ε" {
			arrayToRemove = append(arrayToRemove, (*myArray)[i])
			makeSubstitution(myArray, findEpsilonProductionsSymbols(myArray))
			utils.DeleteMultipleElements1(myArray, arrayToRemove)
		}
	}
	return myArray
}

func makeSubstitution(myArray *[]structure.MapsWithDuplicate, stringsArray []string) *[]structure.MapsWithDuplicate {
	for i := 0; i < len(*myArray); i++ {
		for j := 0; j < len(stringsArray); j++ {
			/*fmt.Printf("%s--->%s==%t\n",myArray[i].values,stringsArray[j],contains(myArray[i].values,stringsArray[j]))*/
			if utils.Contains((*myArray)[i].Values, stringsArray[j]) {
				*myArray = append(*myArray,
					structure.MapsWithDuplicate{Symbols: (*myArray)[i].Symbols, Values: utils.MyTrimFunc((*myArray)[i].Values, stringsArray[j])})
			}
		}
	}
	return myArray
}

func findEpsilonProductionsSymbols(myArray *[]structure.MapsWithDuplicate) []string {
	var epsilonArray []string
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].Values == "ε" {
			epsilonArray = append(epsilonArray, (*myArray)[i].Symbols)
		}
	}
	return epsilonArray

}
