package Unreacheble

import (
	structure "awesomeProject/mystruct"
	utils "awesomeProject/utils"
)

// non generating symbols
func RemoveNonGeneratingSymbols(myarray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate{
	var arrayOfNonGeneratingSymbols []structure.MapsWithDuplicate
	for i:=0;i<len(*myarray);i++ {
		if utils.HasTerminal(utils.GetProductionBySymbol( /*fmt.Println(getProductionBySymbol(*myarray,findStartingSymbol(myarray)))*/*myarray,(*myarray)[i].Symbols)) {
			continue
		}
		arrayOfNonGeneratingSymbols = append(arrayOfNonGeneratingSymbols,structure.MapsWithDuplicate{Symbols: (*myarray)[i].Symbols, Values: (*myarray)[i].Values})
	}
	return utils.DeleteMultipleElements1(myarray,arrayOfNonGeneratingSymbols)
}

//unreacheble symbols - это все сранный костыль так как я говнокодер и ничего лучше после дня попыток не смог придумать :)
func RemoveUnreachebleSymbols1(myarray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate{
	var arrayToDelete []structure.MapsWithDuplicate
	for k := range compareProductions(getMyArrayProductions(myarray),getStaringSymbolsProduction(myarray)){
		if compareProductions(getMyArrayProductions(myarray),getStaringSymbolsProduction(myarray))[k]==false {
			for _,n := range utils.GetProductionBySymbol(*myarray,k){
				arrayToDelete = append(arrayToDelete,structure.MapsWithDuplicate{k,n})
			}
		}
	}
	return utils.DeleteMultipleElements1(myarray,arrayToDelete)
}


func compareProductions(myMap map[string]bool,startingSymbolMap map[string]bool) map[string]bool{
	for k := range myMap{
		for i :=range startingSymbolMap{
			if k==i {
				myMap[k]=true
			}
		}
	}
	return myMap
}

func getMyArrayProductions(myarray *[]structure.MapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(*myarray);i++ {
		myMap[(*myarray)[i].Symbols] = false
	}
	return myMap
}


func getStaringSymbolsProduction(myarray *[]structure.MapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(utils.GetProductionBySymbol(*myarray,"W"));i++ {
		for _,v := range utils.GetProductionBySymbol(*myarray,"W")[i]{
			if utils.IsNonTerminal(uint8(v)) {
				myMap["W"] = true
				myMap[string(uint8(v))] = true
			}
		}
	}
	return myMap
}