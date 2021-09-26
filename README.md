# tabasco
Simple Text Generator uses https://fish-text.ru, https://baconipsum.com/ services.

## Usage

```
Tabasco is a CLI tool to generate a placeholder text akka 'Lorem ipsum'.

Usage:
	 tabasco [--arguments]

The arguments are:
	 file 	 a file name
	 docx 	 a file output format DOCX (if not set default is TXT)
	 size 	 an expected file size (in bytes by default)
	 kb 	 flag set a size unit as KiB
	 mb 	 flag set a size unit as MiB

Use "tabasco --help" for more information

Tabasco uses https://fish-text.ru, https://baconipsum.com/ services to get a random text. It means the Internet connection is important.

MIT License
Copyright (c) 2021 PepperKit.
```

## Example

### Generate 1MiB size text file

```
$ tabasco --file lorem.txt --size 1 --mb
```

### Generate 2KiB size text file

```
$ tabasco --file lorem.txt --size 2 --kb
```

### Generate 512 bytes size text file

```
$ tabasco --file lorem.txt --size 512
```

### Generate 2KiB content size DOCX file

```
$ tabasco --file document.docx --docx --size 2 --kb
```

## License

The library is licensed under the terms of the MIT License.
