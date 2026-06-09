package main

import (
	"errors"
	"fmt"
)

// ── Domain layer: interface defined here, implementations injected ─────────────
var ErrNotFound = errors.New("not found")

// Store is the abstraction. Business logic depends ONLY on this interface.
type Store interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// ── Infrastructure layer: in-memory implementation (used in tests) ─────────────
type MemStore struct {
	data map[string]string
}

func NewMemStore() *MemStore {
	return &MemStore{ data: make(map[string]string) }
}

func (m *MemStore) Set (key, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemStore) Get (key string) (string, error) {
	v, ok := m.data[key]
	if !ok { return "", ErrNotFound }
	return v, nil
}

func (m *MemStore) Delete (key string) error {
	delete(m.data, key)
	return nil
}


// ── Business logic: uses Store interface ──────────────────────────────────────
type UserService struct {
	store Store // accepts any Store implementation
}

func NewUserService(s Store) *UserService { return &UserService{store: s}}

func (s *UserService) SaveUser(id, name string) error {
	 // Prefix key with "user:" to namespace it in the store.
	 return s.store.Set("user:" + id, name)
}

func (s *UserService) GetUser(id string) (string, error){
	name, err := s.store.Get(id)
	if errors.Is(err, ErrNotFound){
		return "", fmt.Errorf("user %s does not exist", id)
	}
	return name, err
}



// ── main: wires concrete implementation ───────────────────────────────────────
func main(){
	// In tests: pass NewMemStore(). In production: pass NewRedisStore() or NewPostgresStore().
	svc := NewUserService(NewMemStore())

	svc.SaveUser("u-101", "Ayush")
	svc.SaveUser("u-102", "Bob")

	name, err := svc.GetUser("u-101")
	fmt.Printf("user: %s err: %v\n", name, err)  // user: Ayush err: <nil>

	_, err = svc.GetUser("u-12")
	fmt.Printf("err: %v\n", err) // err: user u-12 does not exist

}