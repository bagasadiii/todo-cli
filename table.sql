CREATE TABLE IF NOT EXISTS tasks (
    todo_id INT PRIMARY KEY,
    name TEXT,
    status bool DEFAULT FALSE,
    created_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ
);