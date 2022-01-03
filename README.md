# TodoList-API-With-Go

## Objecttive
  1. To practice my Golang and Docker knowledge.
  2. Build Todo list RESTful API with Golang, Gorilla Mux router, GORM and PostgreSQL
  3. Containerization with Docker and docker-compose
  4. Basic testing with Resty
  
 ## API Route
| Endpoint | HTTP Method | CRUD Method | Result |
| ----------- | ----------- | ---------| -------|
| /api/todo | GET | READ | Get all todo list |
| /api/todo:id | GET | READ | Get a single todo list |
| /api/todo | POST | CREATE | Add a todo list |
| /api/todo/:id | PUT | UPDATE | Update a todo list |
| /api/todo/:id | DELETE | DELETE | Delete a todo list |
| /health | GET | READ | Check API status |
  
 
## Run on your machine
Build the images:
```
$ docker-compose build
```
Run the containers:
```
$ docker-compose up -d
```
Bring down the containers:
```
$ docker-compose down
```

