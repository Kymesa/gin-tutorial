mkdir go-books-api && cd go-books-api
go mod init go-books-api

# Instalar Gin y GORM con PostgreSQL
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres