package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := strconv.Quote(fmt.Sprintf("%d mins", r))

	return []byte(value), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquoetedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquoetedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)

	return nil
}
