package visitor

import "path/filepath"

type Reader interface {
    Visit(Compressor) string
}

type Compressor interface {
    Compress(string) string
}

type TxtFile struct {
    file string
}

func (txt *TxtFile) Visit(c Compressor) string {
    return c.Compress(txt.file)
}

type CsvFile struct {
    file string
}

func (csv *CsvFile) Visit(c Compressor) string {
    return c.Compress(csv.file)
}

func NewReader(file string) Reader {
    switch filepath.Ext(file) {
    case ".txt":
        return &TxtFile{file: file}
    case ".csv":
        return &CsvFile{file: file}
    default:
        return nil
    }
}

type FileCompressor struct{}

func (f *FileCompressor) Compress(file string) string {
    switch filepath.Ext(file) {
    case ".txt":
        return "txt"
    case ".csv":
        return "csv"
    default:
        return ""
    }
}

func NewCompressor() Compressor {
    return &FileCompressor{}
}
