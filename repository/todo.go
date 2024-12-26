package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/bagasadiii/todo-app/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepoImpl interface {
	CompleteTodoRepo(ctx context.Context, id int)error
	CreateTodoRepo(ctx context.Context, new *model.Todo)error
}
type TodoRepo struct {
	db *pgxpool.Pool
}
func NewTodoRepo(db *pgxpool.Pool)TodoRepoImpl{
	return &TodoRepo{db:db}
}
func(r *TodoRepo)CreateTodoRepo(ctx context.Context, new *model.Todo)error{
	query := `
		INSERT INTO tasks (todo_id, name, status, created_at, completed_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query,
		new.TodoID,
		new.Name,
		new.Status,
		new.CreatedAt,
		new.CompletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add %s: %v", new.Name, err)
	}
	return nil
}
func(r *TodoRepo)CompleteTodoRepo(ctx context.Context, id int)error{
	query := `
		UPDATE tasks SET status = true, completed_at = $1
		WHERE todo_id = $2
	`
	_, err := r.db.Exec(ctx, query,
		time.Now(),
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update: %v", err)
	}
	return nil
}
