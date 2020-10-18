package storage

// Storage ...
type Storage interface {
	Save(string) (string, error)
	Get(string) (string, error)
}
