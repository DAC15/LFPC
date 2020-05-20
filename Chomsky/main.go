package main

import (
	unitUnreacheble "awesomeProject/Unreacheble"
	chomsky "awesomeProject/chomskyForm"
	epsilonManipulations "awesomeProject/epsilon"
	myStructure "awesomeProject/mystruct"
	"awesomeProject/unit"
	"awesomeProject/utils"
	"fmt"
)

func main() {

	var myarray []myStructure.MapsWithDuplicate
	utils.CreateMyMap(utils.ReadLines("varianta21.txt"),&myarray)
	utils.AddNewStartingSymbol(&myarray)
	epsilonManipulations.RemoveEpsilonProduction(&myarray)
	unit.MakeUnitSubstitution(&myarray)
	unitUnreacheble.RemoveNonGeneratingSymbols(&myarray)
	unitUnreacheble.RemoveUnreachebleSymbols1(&myarray)
	chomsky.AdjustToChomskyNormalForm(utils.DuplicateCount(chomsky.GetProductionInWrongForm(&myarray)),&myarray)
	fmt.Print(myarray)

}
