package file_writer

import "os"

type RawLogger struct {
	Filename   string
	file       *os.File
	openStatus bool
}

func NewFileWriter() *RawLogger {
	return &RawLogger{}
}

func (r *RawLogger) Start(filename string) error {
	r.Filename = filename
	file, err := os.OpenFile(r.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	r.file = file
	r.openStatus = true
	return nil
}

func (r *RawLogger) Stop() {
	defer r.file.Close()
	r.openStatus = false
}

func (r *RawLogger) Write(data []byte) error {
	if _, err := r.file.Write(data); err != nil {
		return err
	}
	return nil
}

func (r *RawLogger) IsOpen() bool {
	return r.openStatus
}
