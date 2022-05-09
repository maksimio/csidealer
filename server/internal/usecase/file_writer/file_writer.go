package file_writer

import "os"

type RawLogWriter struct {
	filename string
	file     *os.File
}

func NewFileWriter(filename string) *RawLogWriter {
	return &RawLogWriter{
		filename: filename,
	}
}

func (fw *RawLogWriter) Start() error {
	file, err := os.OpenFile(fw.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	fw.file = file
	return nil
}

func (fw *RawLogWriter) Stop() {
	defer fw.file.Close()
}

func (fw *RawLogWriter) Write(data []byte) error {
	if _, err := fw.file.Write(data); err != nil {
		return err
	}
	return nil
}
