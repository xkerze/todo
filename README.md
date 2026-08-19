# todo
> A simple CLI todo manager written in Go with TXT storage
### Features
- simple and clear appearance
- easy commands
- saving tasks to a file
### Commands
|Command|Usage|
|--|--|
|!add <task>|adds a task to the list|
|!done <number>|removes the task from the list|
|!tasks|shows all tasks|
|!exit|exit from application|
### Installation
```bash
git clone https://github.com/xkerze/todo.git
cd todo
go mod download
go run cmd/main.go
```
