-- name: GetUser :one
SELECT * FROM users
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: SearchUsers :many
SELECT * FROM users;

-- name: CreateUser :execrows
INSERT INTO users ( id, name, created_at, updated_at )
VALUES ( sqlc.arg(id), sqlc.arg(name), NOW(), NOW() );


-- name: GetTask :one
SELECT * FROM tasks WHERE id = sqlc.arg(id) LIMIT 1;

-- name: SearchTasks :many
SELECT * FROM tasks
WHERE (sqlc.arg(id) = '' OR id = sqlc.arg(id))
	AND (sqlc.arg(title) = '' OR title = sqlc.arg(title))
	AND (sqlc.arg(text) = '' OR text = sqlc.arg(text))
	AND (sqlc.arg(user_id) = '' OR user_id = sqlc.arg(user_id))
	AND (sqlc.arg(priority_id) = '' OR priority_id = sqlc.arg(priority_id))
ORDER BY updated_at DESC;

-- name: CreateTask :execrows
INSERT INTO tasks ( id, title, text, user_id, priority_id, created_at, updated_at )
VALUES ( sqlc.arg(id), sqlc.arg(title), sqlc.arg(text), sqlc.arg(user_id), sqlc.arg(priority_id), NOW(), NOW() );

-- name: GetTaskPriority :one
SELECT * FROM task_priorities WHERE value = sqlc.arg(value) LIMIT 1;
