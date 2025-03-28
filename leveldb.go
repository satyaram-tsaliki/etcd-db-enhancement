package main
import (
	"fmt"
	"time"
	"log"
	"github.com/some/ledgerdb"
)
type ETCDStore struct {
	db *ledgerdb.DB
}


func NewETCDStore(path string) (*ETCDStore, error) {
	db, err := ledgerdb.Open(path, &ledgerdb.Options{})
	if err != nil {
		return nil, err
	}
	return &ETCDStore{db: db}, nil
}
func (store *ETCDStore) Insert(key, value string) error {
	start := time.Now()
	err := store.db.Put([]byte(key), []byte(value))
	duration := time.Since(start)
	log.Printf("Insertion time: %v µs\n", duration.Microseconds())
	return err
}
func (store *ETCDStore) Delete(key string) error {
	start := time.Now()
	err := store.db.Delete([]byte(key))
	duration := time.Since(start)
	log.Printf("Deletion time: %v µs\n", duration.Microseconds())
	return err
}
func (store *ETCDStore) Search(key string) (string, error) {
	start := time.Now()
	value, err := store.db.Get([]byte(key))
	duration := time.Since(start)
	log.Printf("Search time: %v µs\n", duration.Microseconds())
	if err != nil {
		return "", err
	}
	return string(value), nil
}
func (store *ETCDStore) Close() error {
	return store.db.Close()
}
func main() {
	etcdStore, err := NewETCDStore("ledgerdb_data")
	if err != nil {
		log.Fatalf("Failed to initialize ETCD store: %v", err)
	}
	defer etcdStore.Close()
	err = etcdStore.Insert("key1", "value1")
	if err != nil {
		log.Printf("Insertion error: %v", err)
	}

	value, err := etcdStore.Search("key1")
	if err != nil {
		log.Printf("Search error: %v", err)
	} else {
		fmt.Printf("Found value: %s\n", value)
	}

	err = etcdStore.Delete("key1")
	if err != nil {
		log.Printf("Deletion error: %v", err)
	}
}

