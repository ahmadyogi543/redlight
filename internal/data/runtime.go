package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := strconv.Quote(fmt.Sprintf("%d mins", r))

	return []byte(value), nil
}
