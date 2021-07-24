package check

import (
	"errors"
	"fmt"
	"testing"
)

func mayFail(i int) error {
	if i%3 == 1 {
		return fmt.Errorf("fail%d", i)
	}
	return nil
}

func TestDeferred(t *testing.T) {
	defer func() {
		res := recover().(error).Error()
		if res != "fail1 - fail4 - fail7 - final" {
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
