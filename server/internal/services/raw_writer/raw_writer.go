package raw_writer

import (
	"csidealer/internal/models"
	"encoding/binary"
	"errors"
	"log"
	"os"
	"time"
)

type RawWriterService struct {
	filename        string
	file            *os.File
	openStatus      bool
	path            string
	writeByteCount  uint64
	startTime       int64
	logPackageCount uint64
	in              <-chan models.RawPackage
}

func NewRawWriterService(in <-chan models.RawPackage, path string) *RawWriterService {
	return &RawWriterService{
		in:   in,
		path: path,
	}
}

func (r *RawWriterService) Start(filename string) error {
	if r.openStatus {
		return errors.New("предыдущий файл не закрыт")
	}

	err := os.MkdirAll(r.path, os.ModePerm)
	if err != nil {
		return err
	}

	r.filename = filename
	// TODO: использовать JOIN
	file, err := os.OpenFile(r.path+r.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	r.file = file
	r.openStatus = true
	r.writeByteCount = 0
	r.startTime = time.Now().UnixMilli()

	log.Print("Начата запись в файл", r.filename)
	return nil
}

func (r *RawWriterService) Stop() {
	defer r.file.Close()
	r.openStatus = false
	log.Print("Остановлена запись в файл", r.filename)
}

func (r *RawWriterService) write(data []byte) error {
	// log.Print("write data")
	if _, err := r.file.Write(data); err != nil {
		return err
	}
	r.writeByteCount += uint64(len(data))
	return nil
}

func (r *RawWriterService) IsOpen() bool {
	return r.openStatus
}

func (r *RawWriterService) GetWriteByteCount() uint64 {
	return r.writeByteCount
}

func (r *RawWriterService) GetStartTime() int64 {
	return r.startTime
}

func (r *RawWriterService) Log() {
	for {
		rawPackage := <-r.in
		if !r.openStatus {
			continue
		}

		bufSize16 := make([]byte, 2)
		binary.BigEndian.PutUint16(bufSize16, rawPackage.Size)
		r.write(bufSize16)
		r.write(rawPackage.Data)
		r.logPackageCount += 1
	}
}
