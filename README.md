# gobook

This repo includes my solutions to the book exercises that are interspersed throughout the chapters. I use __go version 1.25.0__.  

[Code examples](https://github.com/adonovan/gopl.io/) from the book are available to download at __gopl.io__. `go get` has since been deprecated and `go install` is used instead.  

```bash
$ export GOBIN=$HOME/gobook                     # choose workspace directory
$ go install gopl.io/ch1/helloworld@latest      # fetch, build, install
$ $GOBIN/bin/helloworld                         # run
Hello, 世界
```
