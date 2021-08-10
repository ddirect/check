package check

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func mayFail(i int) error {
	if i%3 == 1 {
		return fmt.Errorf("fail%d", i)
	}
	return nil
}

func mayPanic(i int) {
	if i%3 == 1 {
		panic(fmt.Errorf("panic%d", i))
	}
}

func TestDeferredReturningError(t *testing.T) {
	defer func() {
		res := recover().(error).Error()
		if !strings.HasPrefix(res, "fail1 - fail4 - fail7 - ") || !strings.HasSuffix(res, "final") {
			t.Errorf("unexpected error: %s", res)
		}
	}()
	for i := 0; i < 10; i++ {
		j := i
		defer DeferredE(func() error {
			return mayFail(j)
		})
	}
	E(errors.New("final"))
}

func TestDeferredPaniking(t *testing.T) {
	defer func() {
		res := recover().(error).Error()
		if !strings.HasPrefix(res, "panic1 - panic4 - panic7 - ") || !strings.HasSuffix(res, "final") {
			t.Errorf("unexpected error: %s", res)
		}
	}()
	for i := 0; i < 10; i++ {
		j := i
		defer Deferred(func() {
			mayPanic(j)
		})
	}
	E(errors.New("final"))
}
