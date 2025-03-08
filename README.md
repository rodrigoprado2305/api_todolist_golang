# API TODO LIST com GOLANG
xxxxxxxxxxxxxxxxxxxx

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
