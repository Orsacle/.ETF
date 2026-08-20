package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Orsacle/.ETF/pkg/etf"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage:")
	fmt.Fprintln(os.Stderr, "  etf encode <input> <output.etf>")
	fmt.Fprintln(os.Stderr, "  etf decode <input.etf> <output>")
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 3 {
		usage()
	}

	cmd, src, dst := args[0], args[1], args[2]

	var err error
	switch cmd {
	case "encode":
		err = runEncode(src, dst)
	case "decode":
		err = runDecode(src, dst)
	default:
		usage()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "etf: %v\n", err)
		os.Exit(1)
	}
}

func runEncode(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return etf.Encode(out, in)
}

func runDecode(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	return etf.Decode(out, in)
}
