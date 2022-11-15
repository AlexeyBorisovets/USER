package main

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	//"github.com/AlexeyBorisovets/USER/internal/model"
// 	"github.com/jackc/pgx/v5"
// )

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5432/Test")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// if err == nil {
	// 	fmt.Println("connect to database: successfully")
	// }
	// defer conn.Close(context.Background())

	// testusers := []model.User{
	// 	{
	// 		ID:           "id1",
	// 		Name:         "Alexey",
	// 		Password:     "123",
	// 		Balance:      50,
	// 		RefreshToken: "refresh1",
	// 		UserType:     "consumer",
	// 	},
	// 	{
	// 		ID:           "id2",
	// 		Name:         "Vlad",
	// 		Password:     "123",
	// 		Balance:      55,
	// 		RefreshToken: "refresh2",
	// 		UserType:     "vendor",
	// 	},
	// }

	// ctx := context.Context(context.Background())

	// _, err = conn.Exec(ctx, "insert into users(id,name,password,balance,usertype) values($1,$2,$3,$4,$5)",
	// 	testusers[0].ID, testusers[0].Name, testusers[0].Password, testusers[0].Balance, testusers[0].UserType)
	// if err != nil {
	// 	fmt.Errorf("error with inserting")
	// }

}
