package file_writer

import "os"

type RawLogger struct {
	Filename string
	file     *os.File
}

func NewFileWriter(filename string) *RawLogger {
	return &RawLogger{
		Filename: filename,
	}
}

func (r *RawLogger) Start() error {
	file, err := os.OpenFile(r.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	r.file = file
	return nil
}

func (r *RawLogger) Stop() {
	defer r.file.Close()
}

func (r *RawLogger) Write(data []byte) error {
	if _, err := r.file.Write(data); err != nil {
		return err
	}
	return nil
}
