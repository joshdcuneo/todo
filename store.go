package main

import (
	"encoding/json"
	"os"
	"os/user"
	"strings"
)

type Store struct {
	path string
	data []Todo
}

func NewStore(c AppConfig) *Store {
	return &Store{
		path: c.StorePath(),
		data: []Todo{},
	}
}

func (s *Store) Path() string {
	if strings.HasPrefix(s.path, "~/") {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		dir := usr.HomeDir

		return strings.Replace(s.path, "~", dir, 1)
	}

	return s.path
}

func (s *Store) Load() error {
	file, err := os.Open(s.Path())
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&s.data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Save() error {

	file, err := os.Create(s.Path())
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(s.data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Add(t Todo) {
	s.data = append(s.data, t)
}

func (s *Store) Delete(id int) {
	for i, _ := range s.data {
		if i == id {
			s.data = append(s.data[:i], s.data[i+1:]...)
		}
	}
}

func (s *Store) List() []Todo {
	return s.data
}

func (s *Store) Get(id int) *Todo {
	for i, t := range s.data {
		if i == id {
			return &t
		}
	}

	return nil
}

func (s *Store) Move(id int, to int) {
	t := s.Get(id)
	if t == nil {
		return
	}

	s.Delete(id)

	s.data = append(s.data, Todo{})
	copy(s.data[to+1:], s.data[to:])
	s.data[to] = *t
}
