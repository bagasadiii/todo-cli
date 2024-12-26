package command

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bagasadiii/todo-app/model"
	"github.com/bagasadiii/todo-app/repository"
)

type TodoCmdImpl interface {
	CreateTodoCmd(args []string)
	CompleteTodoCmd(args []string)
}
type TodoCmd struct {
	repo repository.TodoRepoImpl
}
func NewTodoCmd(repo repository.TodoRepoImpl)TodoCmdImpl{
	return &TodoCmd{repo: repo}
}
func (c *TodoCmd) CreateTodoCmd(args []string){
	if len(args) < 1 {
		fmt.Println("Usage: add <todo name>")
		return
	}
	todoName := strings.Join(args, " ")
	newTodo := model.Todo{
		Name: todoName,
		Status: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	newTodo.AddID()
	err := c.repo.CreateTodoRepo(context.Background(), &newTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Task created")
}
func (c *TodoCmd)CompleteTodoCmd(args []string){
	if len(args) < 1 {
		fmt.Println("Usage: update <todo id>")
		return
	}
	todoID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Only accept todo ID in int")
		return
	}
	c.repo.CompleteTodoRepo(context.Background(), todoID)
	fmt.Println("Task is completed")
}