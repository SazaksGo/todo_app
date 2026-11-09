package task

import "errors"

var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrTaskNotFound = errors.New("Task Not Found")
