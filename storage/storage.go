package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db             *sql.DB
	UserRepository *UserRepository
	test           string
}

func New() *Storage {
	fmt.Println("New")
	return &Storage{
		test: "test",
		//UserRepository: NewUserRepository(),
	}
}

func (s *Storage) Open() error {
	//connStr := "host=localhost port=5436 user=postgres password=qwerty dbname=postgres sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s.db = db

	connect := "host=localhost port=5436 user=postgres password=qwerty dbname=postgres sslmode=disable"

	s.Db, _ = sql.Open("postgres", connect)

	fmt.Println("connect DB")

	return nil
}

func (s *Storage) Close() {
	s.Db.Close()
}

//func (s *Storage) User() *UserRepository {
//	if s.UserRepository != nil {
//		return s.UserRepository
//	}
//	s.UserRepository = &UserRepository{
//		Storage: s,
//	}
//	return s.UserRepository
//}

//func (s *Storage) AllUser() []*models.User {
//	fmt.Println("Storage AllUser")
//	//if s.userRepository != nil {
//	fmt.Println("AllUser")
//
//	us, _ := s.userRepository.SelectAll()
//
//	fmt.Println(us)
//
//	return us
//	//}
//	//s.userRepository = &UserRepository{
//	//	storage: s,
//	//}
//	//return nil
//}
