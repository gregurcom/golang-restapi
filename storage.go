package main

import (
	"errors"
	"sync"
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int `json:"age"`
	Discipline string `json:"discipline"`
}

type Storage interface {
	Insert(st *Student)
	Get(id int) (Student, error)
	Update(id int, st Student)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data map[int]Student
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Student),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(st *Student) {
	s.Lock()

	st.ID = s.counter
	s.data[st.ID] = *st

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Student, error) {
	s.Lock()
	defer s.Unlock()

	student, ok := s.data[id]
	if !ok {
		return student, errors.New("student not found")
	}

	return student, nil
}

func (s *MemoryStorage) Update(id int, st Student) {
	s.Lock()
	s.data[id] = st
	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}
