package storage

import (
	"fmt"
	"github.com/aasidxxx/mules/cmd/models"
	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
	"log"
)

type UserRepository struct {
	Storage *Storage
	testu   string
}

var (
	tableUser string = "users"
)

//func (ur *UserRepository) SelectAll() ([]*models.User, error) {}

func NewUserInit(st *Storage) *UserRepository {
	fmt.Println("NewUserInit")
	return &UserRepository{
		testu:   "testu",
		Storage: st,
	}
}

func NewUserRepository() *UserRepository {
	fmt.Println("NewUserRepository")
	return &UserRepository{
		testu: "testu",
	}
}

func (ur *UserRepository) TestUser() string {
	return "-------------- TestUser"

	return "TestUser"
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	fmt.Println("----------- users1 ")

	fmt.Println(ur.Storage)
	fmt.Println("----------- users2 ")
	db := ur.Storage.Db
	fmt.Println(db)
	//
	rows, err := ur.Storage.Db.Query("select * from users")
	//
	////rows, err := db.Query
	//
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		a := models.User{}
		err := rows.Scan(&a.ID, &a.Name, &a.Email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(rows)
		users = append(users, &a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users, nil

}
