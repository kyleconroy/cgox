# CGox - Simple CGO Cross Compilation

CGox is a simple, no-frills tool for CGo cross compilation that behaves a
lot like standard `go build`. CGox will parallelize builds for multiple
platforms. Gox will also build the cross-compilation toolchain for you.

## Installation

CGox requires a patched version of Go that allows cross compilation with CGO
enabled. 

```
hg clone -u release https://code.google.com/p/go
curl https://gist.github.com/steeve/6905542/raw/cross_compile_goos.patch | patch -p1
```

To install CGox, please use `go get`. We tag versions so feel free to
checkout that tag and compile.

```
$ go get github.com/kyleconroy/cgox
...
$ cgox -h
...
```


## Usage

Before you use CGox, you must build the cross-compilation toolchain. CGox can
do this for you. This only has to be done once (or whenever you update Go):

```
$ cgox -build-toolchain
...
```

Once that is done, you're ready to cross compile!

If you know how to use `go build`, then you know how to use Gox. For
example, to build the current package, specify no parameters and just
call `cgox`. CGox will parallelize based on the number of CPUs you have
by default and build for every platform by default:

```
$ gox
Number of parallel builds: 4

-->    darwin/amd64: github.com/kyleconroy/cgox
-->     linux/amd64: github.com/kyleconroy/cgox
-->   windows/amd64: github.com/kyleconroy/cgox
```

Or, if you want to build a package and sub-packages:

```
$ cgox ./...
...
```

Or, if you want to build multiple distinct packages:

```
$ cgox github.com/mitchellh/gox github.com/hashicorp/serf
...
```

Or if you want to just build for linux:

```
$ cgox -os="linux"
...
```

Or maybe you just want to build for 64-bit linux:

```
$ cgox -osarch="linux/amd64"
...
```

And more! Just run `cgox -h` for help and additional information.

## Supported host and target architectures

Currently, CGox supports 64-bit Linux hosts. As more toolchains are built,
supported hosts will include OS X and Windows.

CGox can target 64-bit Windows, OS X and Linux. 32-bit support is planned.
