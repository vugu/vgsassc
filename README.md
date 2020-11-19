# vgsassc
Sass compiler using go-libsass

## Installation

The Go compiler and a GCC-compatible C compiler (for CGO) are required.  Assuming those are installed, vgsassc installation is as simple as:

```bash
go get github.com/vugu/vgsassc
```

If you get an error about a missing compiler, see next.

### Mac

Download and install “Command Line Tools for Xcode”. In more recent MacOS versions you can run `xcode-select --install`

### Linux

Refer to your distribution's documentation for how to install C development tools, but it's usually as simple as `sudo apt install build-essential` (Ubuntu/Debian/etc) or `sudo yum groupinstall 'Development Tools'` (RHEL/CentOS/etc)

### Windows

This compiler has been known to work well: https://jmeubank.github.io/tdm-gcc/

## Help

Once you've been able to run the `go get github.com/vugu/vgsassc` command, it should put the `vgsassc` executable in your path.  Running `vgsassc --help` gives:

```
Usage of vgsassc:
  -I string
    	Specify directory to use for resolving @import
  -m	Shorthand for -output-style=compressed and takes precedence
  -o string
    	Output file
  -output-style string
    	One of: nested, expanded, compact or compressed (default "nested")
```
