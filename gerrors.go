package gerrors

import (
	"fmt"
	"strings"
)

type GroupError struct {
	Message  string
	Args     []interface{}
	children []error
}

func (gerr *GroupError) Printf(pattern string, args ...interface{}) {
	gerr.Message = pattern
	gerr.Args = args
}

func (gerr GroupError) Error() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, gerr.Message, gerr.Args...)
	for _, cerr := range gerr.children {
		fmt.Fprint(&sb, cerr.Error())
	}

	return sb.String()
}

func (gerr GroupError) IsNil() bool {
	if len(gerr.children) > 0 {
		return false
	}

	return true
}

func (gerr GroupError) Value() error {
	if gerr.IsNil() {
		return nil
	}

	return gerr
}

func (gerr *GroupError) Append(err error) {
	gerr.children = append(gerr.children, err)
}
