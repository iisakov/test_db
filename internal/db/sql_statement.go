package db

const GetTasks = `SELECT id, title, done FROM tasks;`
const GetTask = `SELECT id, title, done FROM tasks WHERE id = ?;`
const CreateTask = `INSERT INTO tasks (title, done) VALUES (?, ?);`
const UpdateTask = `UPDATE tasks SET title = ?, done = ? WHERE id = ?;`
const DeleteTask = `DELETE FROM tasks WHERE id = ?;`
