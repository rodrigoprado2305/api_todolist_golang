# API TODO LIST com GOLANG

Sistema para gravação de pedidos e itens.
Orientação a objetos, MVC e camada de persistencia com RTTI

## Banco de dados
MySQL 5.7.44.0 - Usando docker
docker run --name mysql5.7 -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Banco12345* mysql:5.7

Config|Run|Description
-|-|-

go.mod|go get github.com/go-sql-driver/mysql|Biblioteca para o MYSQL
go.mod|github.com/gin-gonic/gin|Biblioteca para utilização do gin 



Config|Description
-|-
Configuracoes.ini|ficando ao lado do .exe.
libmysql.dll|versao 5.7.44.0 x86
Teste_WK_Rod.exe|Aplicação .exe principal
