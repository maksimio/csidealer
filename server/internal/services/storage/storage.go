package storage

import (
	"csidealer/internal/models"
	"log"
)

type StorageService struct {
	data      []models.Package
	fullCount uint64
	maxCount  uint64
	in        <-chan models.Package
	outs      []chan<- models.Package // при добавлении пакета отправляет в эти каналы данные
	// сейчас нужно для отправки в WebSocket сглаженных данных. После декомпозиции нужно оформить
	// сглаживание как отдельный сервис
}

func NewStorageService(in <-chan models.Package, maxCount uint64) *StorageService {
	return &StorageService{
		maxCount: maxCount,
		data:     make([]models.Package, 0, maxCount),
		in:       in,
	}
}

func (s *StorageService) Run() {
	for {
		pack := <-s.in
		s.push(pack)
		log.Println("Push in storage", s.fullCount)
		for _, out := range s.outs {
			out <- pack
		}
	}
}

func (c *StorageService) push(csiPackage models.Package) {
	c.data = append(c.data, csiPackage)
	c.fullCount += 1
	if c.fullCount > c.maxCount {
		c.data = c.data[1:]
	}
}

func (c *StorageService) GetLastN(n int) []models.Package {
	// TODO: я убрал передачу по указателю
	// подумать, нужно ли ее возвращать - будет ли эффективнее?
	length := len(c.data)
	if n > length {
		n = length
	}

	return c.data[length-n:]
}
