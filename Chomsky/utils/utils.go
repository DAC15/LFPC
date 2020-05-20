package utils

import (
	structure "awesomeProject/mystruct"
	"bufio"
	"math/rand"
	"os"
)

func CreateMyMap(lines []string, myarray *[]structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '-' {
				*myarray = append(*myarray, structure.MapsWithDuplicate{Symbols: lines[i][:j], Values: lines[i][j+1:]})
			}
		}
	}
	return myarray
}


func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func DeleteMultipleElements1(baseArray *[]structure.MapsWithDuplicate, arrayToDelete []structure.MapsWithDuplicate) *[]structure.MapsWithDuplicate {
	for i := 0; i < len(*baseArray); i++ {
		url := (*baseArray)[i]
		for _, rem := range arrayToDelete {
			if url == rem {
				*baseArray = append((*baseArray)[:i], (*baseArray)[i+1:]...)
				i--
				break
			}
		}
	}
	return baseArray

}

func MyTrimFunc(word string, charToTrim string) string {

	myRunes := []rune(word)
	for i := 0; i < len(myRunes); i++ {
		if string(myRunes[i]) == charToTrim && len(string(myRunes)) > 1 {
			myRunes := append(myRunes[:i], myRunes[i+1:]...)
			return string(myRunes)
		} else if word == charToTrim { //be carefull here also can be len(word)==1
			myRunes[i] = 949
			return string(myRunes)
		}
	}
	return "WORD DOES NOT CONTAIN CHARACTER"
}

func AddNewStartingSymbol(myMap *[]structure.MapsWithDuplicate) {
	for i := 0; i < len(*myMap); i++ {
		if Contains((*myMap)[i].Values, FindStartingSymbol(myMap)) {
			*myMap = append(*myMap, structure.MapsWithDuplicate{Symbols: "W", Values: FindStartingSymbol(myMap)})
			break
		}
	}
}

func FindStartingSymbol(myMap *[]structure.MapsWithDuplicate) string {
	return (*myMap)[0].Symbols
}

func IsNonTerminal(char uint8) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}

func Contains(word string, char string) bool {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			return true
		}
	}
	return false
}

func HasTerminal(stringArray []string) bool{
	for i:=0;i<len(stringArray);i++ {
		if isTerminalString(stringArray[i]) {
			return true
			break
		}
	}
	return false
}

func GetStaringSymbolsProduction(myarray *[]structure.MapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(GetProductionBySymbol(*myarray,"W"));i++ {
		for _,v := range GetProductionBySymbol(*myarray,"W")[i]{
			if IsNonTerminal(uint8(v)) {
				myMap["W"] = true
				myMap[string(uint8(v))] = true
			}
		}
	}
	return myMap
}

func isTerminalString(string2 string) bool{
	if len(string2)==1 && isTerminal(string2[0]){
		return true
	}else if len(string2)==1 && !isTerminal(string2[0]){
		return false
	} else if len(string2)>1 {
		for j:=0;j<len(string2);j++ {
			if isTerminal(string2[j]) && j==len(string2)-1  {
				continue
			}else if !isTerminal(string2[j]) && j==len(string2)-1{
				return false
			}
		}
	}
	return true
}

func isTerminal(char uint8) bool{
	if char >=97 && char<=122{
		return true
	}
	return false
}

func RandomStringGenerator(length int,letterRunes []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateAccessibleRunes(arrayToDelete map[string]bool) []rune{
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i:=0;i<len(letterRunes);i++{
		url := letterRunes[i]
		for rem := range arrayToDelete{
			if string(url)==rem {
				letterRunes = append(letterRunes[:i],letterRunes[i+1:]...)
				i--
				break
			}
		}
	}
	return letterRunes
}

func GetProductionBySymbol(myarray []structure.MapsWithDuplicate, symbol string) []string {
	var arrayProduction []string
	for i := 0; i < len(myarray); i++ {
		if myarray[i].Symbols == symbol {
			arrayProduction = append(arrayProduction, myarray[i].Values)
		}
	}
	return arrayProduction
}


func GetTerminalByProduction(myArray []structure.MapsWithDuplicate,input string) int{
	for i:=0;i<len(myArray);i++{
		if myArray[i].Values == input {
			return i
		}
	}
	return 0
}


func DuplicateCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}

func GetMyArrayProductions(myarray *[]structure.MapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(*myarray);i++ {
		myMap[(*myarray)[i].Symbols] = false
	}
	return myMap
}

func IsTerminalChar(char uint8) bool{
	if char >=97 && char<=122{
		return true
	}
	return false
}
