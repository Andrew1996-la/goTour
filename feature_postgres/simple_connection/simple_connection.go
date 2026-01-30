package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	connectionString := "postgres://postgres:1111@localhost:5432/postgres"
	connection, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	if err := connection.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подлкючение к базе данных прошло успешно")
}
