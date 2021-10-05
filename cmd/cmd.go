package cmd

import (
	"flag"
	"fmt"
	"github.com/pepperkit/tabasco/txt"
	"os"
	"strings"
)

type TabascoArgs struct {
	FileName string
	FileSize int
	UnitKiB  bool
	UnitMiB  bool
	Docx     bool
	Language txt.Lang
}

func Parse() *TabascoArgs {
	inputFileName := flag.String("file", "", "a file name (default format is TXT)")
	inputExpectedSize := flag.Int("size", 0, "expected a file content size (default bytes)")
	unitKb := flag.Bool("kb", false, "flag set a size unit as KiB")
	unitMb := flag.Bool("mb", false, "flag set a size unit as MiB")
	docx := flag.Bool("docx", false, "flag set an output format as DOCX")
	lang := flag.String("lang", "ru", "choose a language, supported: ru, latin")
	flag.Parse()
	lg := parseLanguage(lang)
	return &TabascoArgs{
		FileName: *inputFileName,
		FileSize: *inputExpectedSize,
		UnitKiB:  *unitKb,
		UnitMiB:  *unitMb,
		Docx:     *docx,
		Language: lg,
	}
}

func parseLanguage(lang *string) txt.Lang {
	if strings.EqualFold(*lang, "latin") {
		return txt.LT
	}
	return txt.RU
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
	if args.FileSize <= 0 && len(args.FileName) <= 0 {
		fmt.Println("Tabasco is a CLI tool to generate a placeholder text akka 'Lorem ipsum'.")
		fmt.Println("")
		fmt.Println("Usage: ")
		fmt.Println("\t tabasco [--arguments]")
		fmt.Println("")
		fmt.Println("The arguments are: ")
		fmt.Println("\t file \t a file name")
		fmt.Println("\t docx \t a file output format DOCX (if not set default is TXT)")
		fmt.Println("\t size \t an expected content size (in bytes by default)")
		fmt.Println("\t kb \t flag set a size unit as KiB")
		fmt.Println("\t mb \t flag set a size unit as MiB")
		fmt.Println("\t lang \t choose a language, supported: ru, latin (default \"ru\")")
		fmt.Println("")
		fmt.Println("Use \"tabasco --help\" for more information")
		fmt.Println("")
		fmt.Println("Tabasco uses https://fish-text.ru, https://baconipsum.com/ services to get a random text.\nIt means the Internet connection is important.")
		fmt.Println("")
		fmt.Println("MIT License")
		fmt.Println("Copyright (c) 2021 PepperKit.")
		fmt.Println("")
		os.Exit(0)
	}
}
