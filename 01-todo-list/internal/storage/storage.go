package storage

type Storage interface {
	Add()
	Delete()
	List()
	Complete()
}

type storage struct {
}

var s *storage

func init() {
	s = &storage{}
}

func GetStorage() Storage {
	return s
}
