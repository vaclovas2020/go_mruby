# mruby module for Go
This Go package let you use mruby interpreter. We use mruby C librarary (clone from https://github.com/mruby/mruby). Mruby license is included and you can find it in file mruby/LICENSE (MIT license).
IMPORTANT: You need to copy file libmruby.a to your project root directory. Example bash command to run on your project root directory:
```sh
export GOPATH=$(go env GOPATH)
cp $GOPATH/pkg/mod/webimizer.dev/go_mruby@v1.0.0-beta6/libmruby.a ./libmruby.a
```
