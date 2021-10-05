package writer

import (
	"bufio"
	"github.com/gingfrederik/docx"
	"log"
	"os"
)

type DocumentWriter interface {
	WriteText(content string)
	Flush()
	FileName() string
}

const docxSuffix = ".docx"

func NewTxtWriter(fileName string) DocumentWriter {
	if len(fileName) <= 0 {
		panic("file name must be specified")
	}

	f, err := os.Create(fileName)
	checkError(err)

	err = f.Sync()
	checkError(err)

	writer := bufio.NewWriter(f)

	return DocumentWriter(&TxtWriter{
		fileName: fileName,
		writer:   writer,
		file:     f,
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type TxtWriter struct {
	fileName string
	writer   *bufio.Writer
	file     *os.File
}

func (tw *TxtWriter) WriteText(content string) {
	_, err := tw.writer.WriteString(content)
	checkError(err)
}

func (tw *TxtWriter) Flush() {
	tw.writer.Flush()
	defer tw.file.Close()
}

func (tw *TxtWriter) FileName() string {
	return tw.fileName
}

type DocxWriter struct {
	fileName string
	docxFile *docx.File
}

func NewDocxWriter(fileName string) DocumentWriter {
	if len(fileName) <= 0 {
		panic("file name must be specified")
	}

	docFile := docx.NewFile()

	return DocumentWriter(&DocxWriter{
		fileName: fileName + docxSuffix,
		docxFile: docFile,
	})
}

func (dw *DocxWriter) WriteText(content string) {
	paragraph := dw.docxFile.AddParagraph()
	paragraph.AddText(content)
}

func (dw *DocxWriter) Flush() {
	dw.docxFile.Save(dw.fileName)
}

func (dw *DocxWriter) FileName() string {
	return dw.fileName
}
