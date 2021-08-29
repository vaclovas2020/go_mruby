package go_mruby

//#cgo CFLAGS: -Iinclude
//#cgo CFLAGS: -std=c99
//#cgo LDFLAGS: libmruby.a -lm
//#include <mruby.h>
//#include <mruby/compile.h>
//
//static void mruby_load_from_string(char * code)
//{
//  mrb_state *mrb = mrb_open();
//  if (!mrb) { /* handle error */ }
//  mrb_load_string(mrb, code);
//  mrb_close(mrb);
//}
import "C"

func Mruby_load_from_string(code string) {
	C.mruby_load_from_string(C.CString(code))
}
