

//---- create migrate 

migrate create -ext sql -dir migrations UsersCreationMigration

//----- run migrate 

migrate -path migrations -database "postgres://localhost:5436/postgres?sslmode=disable&user=postgres&password=qwerty" up