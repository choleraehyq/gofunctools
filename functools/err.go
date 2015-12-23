package functools

import (
	"fmt"
)

// This way of error handling is inspired by
// https://github.com/reusee/socks5hs/blob/master/err.go

type internelErr struct {
	Pkg  string
	Info string
	Err  error
}

func (self *internelErr) Error() string {
	if self.Err != nil {
		return fmt.Sprintf("%s: %s\n%v", self.Pkg, self.Info, self.Err)
	}
	return fmt.Sprintf("%s: %s\n", self.Pkg, self.Info)
}

func generateErr(err error, format string, args ...interface{}) *internelErr {
	if len(args) > 0 {
		return &internelErr{
			Pkg:  "functools",
			Info: fmt.Sprintf(format, args...),
			Err:  err,
		}
	}
	return &internelErr{
		Pkg:  "functools",
		Info: format,
		Err:  err,
	}
}

func newErr(err error, format string, args ...interface{}) {
	if err != nil {
		panic(generateErr(err, format, args...))
	}
}

func getErr(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e
		} else {
			panic(r)
		}
	}
}
