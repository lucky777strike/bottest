package currency

import "errors"

var ErrCurNotFound = errors.New("currency unknown")
var ErrParsingErr = errors.New("error when parsing site")
