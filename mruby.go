package go_mruby

//#cgo CFLAGS: -Imruby/include
//#cgo CFLAGS: -std=c99
//#cgo LDFLAGS: ./mruby/libmruby.a -lm
//#include "mruby_go.c"
import "C"

func Mruby_load_from_string(code string) {
	C.mruby_load_from_string(C.CString(code))
}

func Mruby_load_from_file(filepath string, entryPoint string) {
	C.mruby_load_from_file(C.CString(filepath), C.CString(entryPoint))
}

func Mruby_open() {
	C.mruby_open()
}

func Mruby_close() {
	C.mruby_close()
}
