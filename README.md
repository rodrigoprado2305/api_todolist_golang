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

## Summary

1-Golang/Flutter/Docker com MYSQL - Visual Studio Code

Index|Version|Description
-|-|-
1.1|Versão 01|API num aquivo main.go, procedural, simples para ser tipo um arquivo macro, um unico arquivo.
1.2|Versão 02|Evoluir ele para MVC, dividindo as responsabilidades
1.3|Versão 03|1 método com uso de GO ROUTINE
1.3|Versão 04|Uso de Access_token padrão JWT
1.4|Versão 05|1 método com uso de micro serviço
1.5|Versão 06|1 teste unitário
1.6|Versão 07|Documentação da API
1.7|Versão 08|Usar o Flutter, para criar uma interface visual para consumir os métodos.
1.8|Versão 09|
1.9|Versão 10|

## Mapa do projeto

api_todolist_golang/

│── main.go
│── config/
│   ├── database.go
│── model/
│   ├── task.go
│── controller/
│   ├── task_controller.go
│   ├── token_controller.go
│── middleware/
│   ├── auth_middleware.go
│── routes/
│   ├── routes.go
│── .env
│── go.mod <br>
│── go.sum

