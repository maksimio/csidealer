package file_writer

import "os"

type RawLogWriter struct {
	filename string
}

func NewFileWriter(filename string) (*RawLogWriter, error) {
	fw := &RawLogWriter{
		filename: filename,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fw, err
	}
	defer file.Close()

	return fw, nil
}

func (fw *RawLogWriter) Write(data []byte) error {
	file, err := os.OpenFile(fw.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(data); err != nil {
		panic(err)
	}

	return nil
}
