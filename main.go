package main

import (
	"flag"
	"io"
	"log"
	"os"
	"reflect"

	libsass "github.com/wellington/go-libsass"
)

func main() {

	oFlag := flag.String("o", "", "Output file")
	oOutputStyle := flag.String("output-style", "nested", "One of: nested, expanded, compact or compressed")
	oMinify := flag.Bool("m", false, "Shorthand for -output-style=compressed and takes precedence")
	oInclude := flag.String("I", "", "Specify directory to use for resolving @import")
	// TODO: include path, option to disable timestamp check, minimize option (by default look for .min.css file ext)
	// TODO: take a look through here - https://godoc.org/github.com/wellington/go-libsass - some interesting stuff like SourceMap etc
	flag.Parse()

	if *oMinify {
		*oOutputStyle = "compressed"
	}

	var out io.Writer
	if *oFlag == "" {
		out = os.Stdout
	} else {
		fout, err := os.Create(*oFlag)
		if err != nil {
			log.Fatalf("error creating output file %q: %v", *oFlag, err)
		}
		defer fout.Close()
		out = fout
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("no input files")
	}

	var in io.Reader
	// TODO: timestamp check to not rebuild file if output is newer than all inputs
	for _, arg := range args {

		fin, err := os.Open(arg)
		if err != nil {
			log.Fatalf("error opening input file %q: %v", arg, err)
		}
		defer fin.Close()

		if in == nil {
			in = fin
		} else {
			in = io.MultiReader(in, fin)
		}
	}

	var newArgs = make([]reflect.Value, 0, 8)
	newArgs = append(newArgs, reflect.ValueOf(out))
	newArgs = append(newArgs, reflect.ValueOf(in))
	switch *oOutputStyle {
	case "nested":
		newArgs = append(newArgs, reflect.ValueOf(libsass.OutputStyle(libsass.NESTED_STYLE)))
	case "expanded":
		newArgs = append(newArgs, reflect.ValueOf(libsass.OutputStyle(libsass.EXPANDED_STYLE)))
	case "compact":
		newArgs = append(newArgs, reflect.ValueOf(libsass.OutputStyle(libsass.COMPACT_STYLE)))
	case "compressed":
		newArgs = append(newArgs, reflect.ValueOf(libsass.OutputStyle(libsass.COMPRESSED_STYLE)))
	default:
		log.Fatalf("unexpected output-style option: %s", *oOutputStyle)
	}

	if *oInclude != "" {
		newArgs = append(newArgs, reflect.ValueOf(libsass.IncludePaths([]string{*oInclude})))
	}

	newRet := reflect.ValueOf(libsass.New).Call(newArgs)
	errV := newRet[1]
	err, ok := errV.Interface().(error)
	if ok {
		log.Fatal(err)
	}
	compV := newRet[0]
	comp := compV.Interface().(libsass.Compiler)

	if err := comp.Run(); err != nil {
		log.Fatal(err)
	}
}
