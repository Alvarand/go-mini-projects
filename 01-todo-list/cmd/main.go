package main

import (
	"flag"
	"log"
	"todo-list/internal/models"
	"todo-list/internal/todo"
)

func init() {
	flag.Parse()
}

func main() {

	args := flag.Args()
	models.CheckZeroArgs(args)

	action := args[0]
	actionFunc, ok := todo.ActionToFunc[action]
	if !ok {
		log.Fatalf(models.ErrorUnknownAction.Error(), action)
	}

	argsAfterAction := args[1:]
	actionFunc(argsAfterAction)
}
