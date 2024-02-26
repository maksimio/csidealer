package fs_reader

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type FSReader struct {
	filename   string
	filesize   int64
	readsize   int64
	file       *os.File
	reader     *bufio.Reader
	openStatus bool
	logPath    string
	bufSize16  []byte
	bufSize32  []byte
}

func NewFSReader(logPath string) *FSReader {
	return &FSReader{
		logPath:   logPath,
		bufSize16: make([]byte, 2),
		bufSize32: make([]byte, 4),
	}
}

func (f *FSReader) List() ([]string, error) {
	list := []string{}
	files, err := os.ReadDir(f.logPath)
	if err != nil {
		return list, err
	}
	for _, entry := range files {
		if !entry.IsDir() {
			list = append(list, entry.Name())
		}
	}

	return list, nil
}

func (f *FSReader) Start(filename string) error {
	if f.openStatus {
		return errors.New("предыдущий файл не закрыт")
	}

	file, err := os.Open(f.logPath + f.filename)
	if err != nil {
		return err
	}

	f.file = file
	f.openStatus = true
	f.reader = bufio.NewReader(f.file)

	fi, err := f.file.Stat()
	if err != nil {
		return err
	}
	f.filesize = fi.Size()
	f.readsize = 0

	return nil
}

func (f *FSReader) GetDataPackage() []byte {
	if !f.openStatus {
		return []byte{}
	}

	f.reader.Read(f.bufSize16)
	f.bufSize32[1], f.bufSize32[0] = f.bufSize16[0], f.bufSize16[1]
	bufSize := binary.BigEndian.Uint16(f.bufSize16)
	buf := make([]byte, bufSize)

	_, err := io.ReadFull(f.reader, buf)
	if err != nil {
		f.Stop()
		return []byte{}
	}

	f.readsize += int64(bufSize)
	return append(f.bufSize32, buf...)
}

func (f *FSReader) IsOpen() bool {
	return f.openStatus
}

func (f *FSReader) Stop() error {
	err := f.file.Close()
	if err != nil {
		return err
	}

	f.reader = &bufio.Reader{}
	f.openStatus = false
	return nil
}

func (f *FSReader) GetReadPercent() float64 {
	if !f.openStatus {
		return 0
	}

	return float64(f.readsize) / float64(f.filesize)
}
