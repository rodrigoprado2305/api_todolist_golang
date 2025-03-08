# API TODO LIST com GOLANG
Ele é uma aplicação simples de gerenciamento de tarefas (TODO) usando o framework Gin para criar uma API RESTful e o MySQL como banco de dados.

## Database
MySQL 5.7.44.0 - Usando docker
```
docker run --name mysql5.7 -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Banco12345* mysql:5.7
```

## Config

File|Run|Description
-|-|-
go.mod|go get github.com/go-sql-driver/mysql|Biblioteca para o MYSQL
go.mod|go get github.com/gin-gonic/gin|Biblioteca para APIs RESTful em Go
go.sum|__|hashes criptográficos das dependências

## Methods - Example - POSTMAN

Type|Rote|Description
-|-|-
POST|http://localhost:8080/tasks|Insert taks
GET|http://localhost:8080/tasks/3|Search one task
GET|http://localhost:8080/tasks|List all task
DELETE|http://localhost:8080/tasks/1|Delete task
PUT|http://localhost:8080/tasks/3|Update task

## Example - JSON - POST AND PUT
```
{
    "title": "test test test",
    "description": "Description xyz",
    "completed": true
}
```
