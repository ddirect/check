package check

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func Recover(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}

func limitPath(x string) string {
	const maxParts = 3
	parts := strings.Split(x, string(filepath.Separator))
	if len(parts) <= maxParts {
		return x
	}
	return filepath.Join(parts[len(parts)-maxParts:]...)
}

func throw(err error) {
	if _, file, line, ok := runtime.Caller(2); ok {
		panic(fmt.Errorf("%s line %d: %w", limitPath(file), line, err))
	}
	panic(err)
}

func Efile(op string, file string, err error) {
	if err != nil {
		throw(fmt.Errorf("%s on '%s': %w", op, file, err))
	}
}

func E(err error) {
	if err != nil {
		throw(err)
	}
}

func El(err error) {
	if err != nil {
		log.Println(err)
	}
}

func IE(n int, err error) int {
	if err != nil {
		throw(err)
	}
	return n
}

func U8E(n uint8, err error) uint8 {
	if err != nil {
		throw(err)
	}
	return n
}

func I64E(n int64, err error) int64 {
	if err != nil {
		throw(err)
	}
	return n
}

func SE(s string, err error) string {
	if err != nil {
		throw(err)
	}
	return s
}
