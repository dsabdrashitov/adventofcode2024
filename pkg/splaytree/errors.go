package splaytree

import "fmt"

type ErrorKeyNotFound struct {
	key string
}

func (e ErrorKeyNotFound) Error() string {
	return fmt.Sprintf("key [%v] not found", e.key)
}
