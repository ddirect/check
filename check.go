package check

import (
	"log"
)

func Recover(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}

func E(err error) {
	if err != nil {
		panic(err)
	}
}

func El(err error) {
	if err != nil {
		log.Println(err)
	}
}

func IE(n int, err error) int {
	if err != nil {
		panic(err)
	}
	return n
}

func I64E(n int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return n
}

func SE(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}
