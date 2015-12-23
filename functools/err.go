package functools

import (
	"fmt"
)

type InternelErr struct {
	Pkg  string
	Info string
	Err  error
}

func (self *InternelErr) Error() string {
	if self.Err != nil {
		return fmt.Sprintf("%s: %s\n%v", self.Pkg, self.Info, self.Err)
	}
	return fmt.Sprintf("%s: %s\n", self.Pkg, self.Info)
}

func generateErr(err error, format string, args ...interface{}) *InternelErr {
	if len(args) > 0 {
		return &InternelErr{
			Pkg:  "functools",
			Info: fmt.Sprintf(format, args...),
			Err:  err,
		}
	}
	return &InternelErr{
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
