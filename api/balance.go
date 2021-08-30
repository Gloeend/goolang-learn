package api

import (
	errors "errors"
	"strconv"
)

type MateshaArgs struct {
	Symbols Symbols `json:"chars"`
	Numbers Numbers `json:"numbers"`
}

type Symbols struct {
	Char string `json:"char"`
}

type Numbers struct {
	First  int64 `json:"first"`
	Second int64 `json:"second"`
}

func (api *Api) Matesha(args *MateshaArgs, result *string) error {
	fn := args.Numbers.First
	sn := args.Numbers.Second
	char := args.Symbols.Char
	if fn == 0 || sn == 0 || char == "" {
		return errors.New("Something is null")
	}
	switch char {
	case "+":
		*result = strconv.FormatInt(fn+sn, 10)
	case "-":
		*result = strconv.FormatInt(fn-sn, 10)
	case "*":
		*result = strconv.FormatInt(fn*sn, 10)
	case "/":
		*result = strconv.FormatInt(fn/sn, 10)
	default:
		return errors.New("Ur separator is WHAT?")
	}
	return nil
}
