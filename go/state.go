package mxnet

/*
// go preamble
#cgo CFLAGS: -I/root/golang/own/src/github.com/songtianyi/go-mxnet/src -I/root/MXNet/mxnet/include/
#cgo LDFLAGS: -Libmxnetwrap.so -lmxnet
*/
import "C"
import (
	"fmt"
)

func MXRandomSeed(seed int) error {
	success, err := C.MXRandomSeedWrap(C.int(seed))
	if err != nil {
		return err
	}
	if success < 0 {
		fmt.Println("fuck")
		return GetLastError()
	}
	return nil
}
