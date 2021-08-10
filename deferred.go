package check

import "fmt"

type wrapped struct {
	main   error
	nested error
}

func (w *wrapped) Error() string {
	return fmt.Sprintf("%s - %s", w.main, w.nested)
}

func (w *wrapped) Unwrap() error {
	return w.nested
}

/*
This function allows checking the error code of deferred functions and
panic while preserving any pending error, by wrapping it in the new error.
Note that recover() only works inside the deferred function (not if called in a
nested function).
*/
func DeferredE(deferred func() error) {
	if err := deferred(); err != nil {
		if prev := recover(); prev != nil {
			panic(&wrapped{err, prev.(error)})
		}
		panic(err)
	}
}

func Deferred(deferred func()) {
	prev := recover()
	defer func() {
		if err := recover(); err != nil {
			if prev != nil {
				panic(&wrapped{err.(error), prev.(error)})
			}
			panic(err)
		}
		if prev != nil {
			panic(prev)
		}
	}()
	deferred()
}
