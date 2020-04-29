package cgofusewrapper

import (
	"fmt"

	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

type FuseError interface {
	error
	ErrorCode() int
}

type ErrorNotFound struct {
	Path string
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("Path %s not found", e.Path)
}
func (e *ErrorNotFound) ErrorCode() int {
	return errno.ENOENT
}

type ErrorAccessDenied struct {
	Path string
}

func (e *ErrorAccessDenied) Error() string {
	return fmt.Sprintf("Access denied to path %s", e.Path)
}
func (e *ErrorAccessDenied) ErrorCode() int {
	return errno.EACCES
}

type ErrorNotADirectory struct {
	Path string
}

func (e *ErrorNotADirectory) Error() string {
	return fmt.Sprintf("%s is not a directory", e.Path)
}
func (e *ErrorNotADirectory) ErrorCode() int {
	return errno.ENOTDIR
}

type ErrorUnknown struct {
	Path    string
	Message string
}

func (e *ErrorUnknown) Error() string {
	return fmt.Sprintf("Error %s accessing path %s", e.Message, e.Path)
}
func (e *ErrorUnknown) ErrorCode() int {
	return errno.EUNKNOWN
}

type ErrorNotImplemented struct{}

func (e *ErrorNotImplemented) Error() string {
	return "Not implemented"
}
func (e *ErrorNotImplemented) ErrorCode() int {
	return errno.EOPNOTSUPP
}
