package go_mruby

//#cgo CFLAGS: -Imruby/include
//#cgo CFLAGS: -std=c99
//#cgo LDFLAGS: ./mruby/libmruby.a -lm
//#include "mruby_go.c"
import "C"

/* Run ruby from given string */
func Mruby_load_from_string(code string) {
	C.mruby_load_from_string(C.CString(code))
}

/* Load from .rb file */
func Mruby_load_from_file(filepath string) {
	C.mruby_load_from_file(C.CString(filepath))
}

/* load from compilder .mrb file (Compiled using mruby CLI command) */
func Mruby_load_irep_file(filepath string) {
	C.mruby_load_irep_file(C.CString(filepath))
}

/* Load from .rb file and run entryPoint function */
func Mruby_load_from_file_entry_point(filepath string, entryPoint string) {
	C.mruby_load_from_file_entry_point(C.CString(filepath), C.CString(entryPoint))
}

/* Open mruby interpreter. Call this funcction first before any other function */
func Mruby_open() {
	C.mruby_open()
}

/* Close mruuby interpreter. Call it in the end, when you donot use interpreter anymore */
func Mruby_close() {
	C.mruby_close()
}
