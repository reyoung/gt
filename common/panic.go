package common

import "log"

func Panic2[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
