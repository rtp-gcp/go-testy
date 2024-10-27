
  
go mod init dbtesty
go run dbmain.go
go get github.com/mattn/go-sqlite3
sqlite3 tables.db
.tables
ctrl-d will exit sqlite3
select * from users;




