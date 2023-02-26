package updater

import (
	"fmt"
	"runtime"
)

type InvalidData struct {
	File string
	Line int
}

func NewInvalidData() InvalidData {
	var err InvalidData
	_, file, line, _ := runtime.Caller(1)
	err.File = file
	err.Line = line
	return err
}

func (i InvalidData) Error() string {
	return fmt.Sprintf("Data on %s:%d has invalid datatype", i.File, i.Line)
}
