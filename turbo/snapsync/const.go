package snapsync

import (
	"errors"
)

var (
	ErrInvalidSnapshot = errors.New("this snapshot for this chainID not supported ")
)