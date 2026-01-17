package processingErrors

import (
	"errors"
	"fmt"
)

var ErrPermissionDenied = errors.New("permission denied")

func checkAccess(role string) error {
	if role != "admin" {
		return ErrPermissionDenied
	}
	return nil
}

func start() {
	err := checkAccess("boba")
	if err != nil && errors.Is(err, ErrPermissionDenied) {
		fmt.Println(err)
	}
}
