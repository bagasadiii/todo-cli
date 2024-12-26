package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bagasadiii/todo-app/command"
	"github.com/bagasadiii/todo-app/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Database error: ", err)
	}
	table, err := os.ReadFile("table.sql")
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
	_, err = db.Exec(context.Background(), string(table))
	if err != nil {
		log.Fatal("Failed to exec table: ", err)
	}
	fmt.Println("Database connected")

	repo := repository.NewTodoRepo(db)
	cmd := command.NewTodoCmd(repo)

	if len(os.Args) < 2 {
		fmt.Println("Unknown command")
	}
	
	command := os.Args[1]
	switch command {
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please select todo id")
		}
		id := os.Args[2:]
		cmd.CompleteTodoCmd(id)
		fmt.Println("Task updated")
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <task>")
		}
		task := os.Args[2:]
		cmd.CreateTodoCmd(task)
	default:
		fmt.Println("Please type Valid command")
	}
	os.Exit(1)
}