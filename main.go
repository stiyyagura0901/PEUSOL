package main

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"

	"github.com/PEUSOL/helpers"
	"github.com/divan/num2words"

	strcase "github.com/iancoleman/strcase"
)

func main() {
	flag.Parse()
	problemStr := flag.Args()[0]
	str, _ := strconv.Atoi(problemStr)
	problemNum := num2words.Convert(str)
	problem := strcase.ToCamel(problemNum)

	funcs := helpers.GetFuncs()

	fmt.Println(`____________________Executing Problem: `, problem)

	reflect.ValueOf(funcs[problem]).Call([]reflect.Value{})
}
