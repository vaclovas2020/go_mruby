package go_mruby

//#cgo CFLAGS: -Imruby/include
//#cgo CFLAGS: -std=c99
//#cgo LDFLAGS: ./mruby/libmruby.a -lm
//#include <mruby.h>
//#include <mruby/compile.h>
//#include <stdio.h>
//
//static void mruby_load_from_string(char * code)
//{
//  mrb_state *mrb = mrb_open();
//  if (!mrb) { /* handle error */ }
//  mrb_load_string(mrb, code);
//  mrb_close(mrb);
//}
//
//static void mruby_load_from_file(char * fname)
//{
//  FILE * file = fopen(fname, "r");
//  mrb_state *mrb = mrb_open();
//  if (!mrb) { /* handle error */ }
//  mrb_load_file(mrb, file);
//  mrb_close(mrb);
//}
import "C"

func Mruby_load_from_string(code string) {
	C.mruby_load_from_string(C.CString(code))
}

func Mruby_load_from_file(filepath string) {
	C.mruby_load_from_file(C.CString(filepath))
}
