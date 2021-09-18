package cmd

import (
	"flag"
	"fmt"
	"os"
)

type TabascoArgs struct {
	FileName     string
	FileSize     int
	UnitKiloByte bool
	UnitMegaByte bool
}

func Parse() *TabascoArgs {
	inputFileName := flag.String("file", "", "a file name")
	inputExpectedSize := flag.Int("size", 0, "a file expected size in bytes")
	unitKb := flag.Bool("kb", false, "flag set a size unit as KBytes")
	unitMb := flag.Bool("mb", false, "flag set a size unit as MBytes")
	flag.Parse()
	return &TabascoArgs{
		FileName:     *inputFileName,
		FileSize:     *inputExpectedSize,
		UnitKiloByte: *unitKb,
		UnitMegaByte: *unitMb,
	}
}

func ValidateFileSize(args *TabascoArgs) {
	if args.FileSize <= 0 {
		fmt.Println("specify --size arg")
		fmt.Println("Use \"tabasco --help\" for more information")
		fmt.Println("")
		fmt.Println("PepperKit(c) 2021.")
		fmt.Println("")
		os.Exit(1)
	}
}

func Info(args *TabascoArgs) {
	if args.FileSize <= 0 && !args.UnitKiloByte && !args.UnitMegaByte && len(args.FileName) <= 0 {
		fmt.Println("Tabasco is a CLI tool to generate a placeholder text akka 'Lorem ipsum'.")
		fmt.Println("")
		fmt.Println("Usage: ")
		fmt.Println("\t tabasco [--arguments]")
		fmt.Println("")
		fmt.Println("The arguments are: ")
		fmt.Println("\t file \t a file name")
		fmt.Println("\t size \t an expected file size (in bytes by default)")
		fmt.Println("\t kb \t a size will be read as KBytes")
		fmt.Println("\t mb \t a size will be read as MBytes")
		fmt.Println("")
		fmt.Println("Use \"tabasco --help\" for more information")
		fmt.Println("")
		fmt.Println("Tabasco uses https://fish-text.ru service to get a random text. It means the Internet connection is important.")
		fmt.Println("")
		fmt.Println("MIT License")
		fmt.Println("Copyright (c) 2021 PepperKit.")
		fmt.Println("")
		os.Exit(0)
	}
}