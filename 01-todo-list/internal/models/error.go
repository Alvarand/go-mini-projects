package models

import "errors"

var ErrorZeroLenArgs = errors.New("zero args")
var ErrorUnknownAction = errors.New("unknown action: %s")
