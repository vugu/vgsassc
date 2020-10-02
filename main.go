package main

import (
	"flag"
	"io"
	"log"
	"os"

	libsass "github.com/wellington/go-libsass"
)

func main() {

	oFlag := flag.String("o", "", "Output file")
	// TODO: include path, option to disable timestamp check
	flag.Parse()

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

	comp, err := libsass.New(out, in)
	if err != nil {
		log.Fatal(err)
	}

	if err := comp.Run(); err != nil {
		log.Fatal(err)
	}
}
