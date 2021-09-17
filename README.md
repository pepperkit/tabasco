# tabasco
Simple Text Generator uses https://fish-text.ru service

## Usage

```
Tabasco is CLI tool to generate a placeholder text akka 'Lorem ipsum' generator.

Usage:
	 tabasco [--arguments]

The argumnets are:
	 file 	 a file name
	 size 	 an expected file size (in bytes by default)
	 kb 	 a size will be read as KBytes
	 mb 	 a size will be read as MBytes

Use "tabasco --help" for more information

Tabasco uses https://fish-text.ru service to get a random text. It means the Internet connection is important.

MIT License
Copyright (c) 2021 PepperKit.
```

## Example

### Generate 1Mb size text file

```
$ tabasco --file lorem.txt --size 1 --mb
```

### Generate 2Kb size text file

```
$ tabasco --file lorem.txt --size 2 --kb
```

### Generate 512 bytes size text file

```
$ tabasco --file lorem.txt --size 512
```

## License

The library is licensed under the terms of the MIT License.
