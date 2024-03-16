package raw_writer

import (
	"csidealer/internal/models"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type RawWriterService struct {
	WriteByteCount    uint64
	WritePackageCount uint64
	IsWriting         bool
	StartTime         int64

	filename string
	file     *os.File
	path     string
	in       <-chan models.RawPackage
}

func NewRawWriterService(in <-chan models.RawPackage, path string) *RawWriterService {
	return &RawWriterService{
		in:   in,
		path: path,
	}
}

func (r *RawWriterService) Start(filename string) error {
	if r.IsWriting {
		return errors.New("предыдущий файл не закрыт")
	}

	err := os.MkdirAll(r.path, os.ModePerm)
	if err != nil {
		return err
	}

	r.filename = filename
	// TODO: использовать JOIN
	file, err := os.OpenFile(filepath.Join(r.path, r.filename), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	r.file = file
	r.IsWriting = true
	r.WriteByteCount = 0
	r.StartTime = time.Now().UnixMilli()
	r.WritePackageCount = 0

	log.Print("начата запись в файл", r.filename)
	return nil
}

func (r *RawWriterService) Stop() error {
	if !r.IsWriting {
		return errors.New("сейчас запись в файл не происходит. Нечего останавливать")
	}

	err := r.file.Close()
	if err != nil {
		return err
	}

	r.IsWriting = false
	log.Print("остановлена запись в файл", r.filename)
	return nil
}

func (r *RawWriterService) write(data []byte) error {
	// log.Print("write data")
	if _, err := r.file.Write(data); err != nil {
		return err
	}
	r.WriteByteCount += uint64(len(data))
	return nil
}

func (r *RawWriterService) SetMark(id string, text string, isActive bool) error {
	if !r.IsWriting {
		return errors.New("невозможно установить метку, так как запись в файл не происходит")
	}

	markFileName := r.filename + ".json"
	pathToFile := filepath.Join(r.path, markFileName)
	fmt.Println(pathToFile)

	file, err := os.OpenFile(pathToFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	var marks []models.Mark
	json.Unmarshal(byteValue, &marks) // TODO: будет ошибка, если файл новый
	marks = append(marks, models.Mark{Id: id, Text: text, IsActive: isActive, Timestamp: time.Now().UnixMilli()})
	result, err := json.MarshalIndent(marks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(pathToFile, result, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r *RawWriterService) Run() {
	for {
		rawPackage := <-r.in
		if !r.IsWriting {
			continue
		}

		bufSize16 := make([]byte, 2)
		binary.BigEndian.PutUint16(bufSize16, rawPackage.Size)
		r.write(bufSize16)
		r.write(rawPackage.Data)
		r.WritePackageCount += 1
	}
}
