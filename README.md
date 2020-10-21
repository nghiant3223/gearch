# gearch

## Overview

`gearch` is a Go CLI application used for searching dependency by keywords.

## Installing

`go get -u github.com/nghiant3223/gearch`

## Usage

```
$ gearch uber fx

github.com/gen0cide/cfx
https://github.com/gen0cide/cfx

----------

github.com/reactivex/rxgo/fx
https://github.com/reactivex/rxgo/fx
Package fx provides predicate-like function types to be used with operators such as Map, Filter, Scan, and Start.

----------

go.uber.org/fx
https://go.uber.org/fx
Package fx is a framework that makes it easy to build applications out of reusable, composable modules.

```

## Flags

```
  -c int Maximum number of results showing (default 3)
```