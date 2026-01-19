package main

import "errors"

// можно, но легко сломать API
var ErrNotFound = errors.New("not found")
