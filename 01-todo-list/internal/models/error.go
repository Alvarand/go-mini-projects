package models

import "errors"

var ErrorZeroLenArgs = errors.New("zero args")
var ErrorUnknownAction = errors.New("unknown action: %s")

var ErrorConvertStringToInt = errors.New("failed to convert string to int: %s")
var ErrorConvertStringToBool = errors.New("failed to convert string to bool: %s")
var ErrorConvertStringToTime = errors.New("failed to convert string to time: %s")

var ErrorFailedCreateFile = errors.New("failed to create file: %s")
var ErrorFailedOpenFile = errors.New("failed to open file: %s")
var ErrorFailedCloseFile = errors.New("failed to close file: %s")
var ErrorFailedWriteFile = errors.New("failed to write in file: %s")
var ErrorFailedReadFile = errors.New("failed to read from file: %s")
