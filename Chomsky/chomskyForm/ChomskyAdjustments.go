package chomskyForm

import (
	structure "awesomeProject/mystruct"
	utils "awesomeProject/utils"
)

func AdjustToChomskyNormalForm(mymap map[string]int,myarray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate{
	myFirstMap := make(map[string]string)
	myMap := make(map[string]string)
	for k :=range mymap{
		myFirstMap[k[0:len(k)-1]]=utils.RandomStringGenerator(1,utils.CreateAccessibleRunes(utils.GetStaringSymbolsProduction(myarray)))
		myMap[(*myarray)[utils.GetTerminalByProduction(*myarray,(*myarray)[utils.GetTerminalByProduction(*myarray,k)].Values)].Values]=
			(*myarray)[utils.GetTerminalByProduction(*myarray,(*myarray)[utils.GetTerminalByProduction(*myarray,k)].Values)].Values[0:len(k)-1]
		*myarray = append(*myarray,structure.MapsWithDuplicate{Symbols: utils.RandomStringGenerator(1,utils.CreateAccessibleRunes(utils.GetStaringSymbolsProduction(myarray))), Values: k[0:len(k)-1]})
	}

	for i:=0;i<len(*myarray);i++{
		for k,v := range createArrayOfNewSymbols(myMap,myFirstMap){
			if (*myarray)[i].Values == k{
				(*myarray)[i].Values = v
			}
		}

	}

	return myarray
}



func createArrayOfNewSymbols(myMap map[string]string,myFirstMap map[string]string) map[string]string{
	newMap := make(map[string]string)
	for key,value:=range myMap{
		for k,v := range myFirstMap{
			if value==k {
				newMap[key]=v+key[len(k):]
			}
		}
	}
	return newMap
}

func GetProductionInWrongForm(myarray *[]structure.MapsWithDuplicate) []string{
	var wrongFormArray []string
	for k := range utils.GetMyArrayProductions(myarray){
		for _,v :=range utils.GetProductionBySymbol(*myarray,k) {
			if (len(v) == 2 && utils.IsTerminalChar(v[0]) && utils.IsNonTerminal(v[1])) ||
				(len(v) == 2 && utils.IsNonTerminal(v[0]) && utils.IsTerminalChar(v[1])) || len(v) > 2 {
				wrongFormArray = append(wrongFormArray, v)
			}
		}
	}
	return wrongFormArray
}



