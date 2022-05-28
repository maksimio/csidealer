package fs_logger

import (
	"errors"
	"fmt"
	"os"
)

type RawLogger struct {
	Filename   string
	file       *os.File
	openStatus bool
}

func NewFileLogger() *RawLogger {
	return &RawLogger{}
}

func (r *RawLogger) Start(filename string) error {
	if r.openStatus {
		return errors.New("предыдущий файл не закрыт")
	}

	r.Filename = filename
	file, err := os.OpenFile(r.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	r.file = file
	r.openStatus = true

	fmt.Println("Начата запись в файл", r.Filename)
	return nil
}

func (r *RawLogger) Stop() {
	defer r.file.Close()
	r.openStatus = false
	fmt.Println("Остановлена запись в файл", r.Filename)
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
