# gobook

This repo includes my solutions to "The Go Programming Language" book exercises. 
I use __go version 1.24.1__.  

[Code examples](https://github.com/adonovan/gopl.io/) from the book are available to download at __gopl.io__. `go get` has since been deprecated for this use and `go install` is used instead.  

```bash
$ export GOBIN=$HOME/gobook/bin                 # choose workspace directory
$ go install gopl.io/ch1/helloworld@latest      # fetch, build, install
$ $GOBIN/helloworld                             # run
Hello, 世界
```
