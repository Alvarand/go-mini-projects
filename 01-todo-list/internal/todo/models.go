package todo

var ActionToFunc = map[string]func([]string){
	"add":      Add,
	"delete":   Delete,
	"list":     List,
	"complete": Complete,
}
