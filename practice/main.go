package main

import (
	"practice/arrays"
	"practice/conditions"
	"practice/constants"
	"practice/errors"
	"practice/functions"
	"practice/interfaces"
	"practice/maps"
	"practice/panic"
	"practice/pointer"
	"practice/slices"
	"practice/structs"
	"practice/variables"
)

func main() {
	variables.Execute()
	constants.Execute()
	functions.Execute()
	pointer.Execute()
	arrays.Execute()
	slices.Execute()
	conditions.Execute()
	maps.Execute()
	structs.Execute()
	interfaces.Execute()
	errors.Execute()
	panic.Execute()
}
