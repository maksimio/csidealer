package storage

// import (
// 	entity "csidealer/internal/models"
// 	"testing"
// )

// func TestRepoMaxLength(t *testing.T) {
// 	repo := NewStorageService(3)
// 	pack := &entity.Package{}
// 	repo.Push(pack)
// 	repo.Push(pack)
// 	repo.Push(pack)

// 	if len(repo.data) != 3 {
// 		t.Errorf("Ожидается длина 3, а она равна %d", len(repo.data))
// 	}
// }

// func TestRepoGet(t *testing.T) {
// 	repo := NewStorageService(3)
// 	pack := &entity.Package{}
// 	repo.Push(pack)
// 	repo.Push(pack)
// 	repo.Push(pack)

// 	data := repo.GetLastN(10)

// 	if len(data) != 3 {
// 		t.Errorf("Ожидается длина 3, а она равна %d", len(repo.data))
// 	}
// }
