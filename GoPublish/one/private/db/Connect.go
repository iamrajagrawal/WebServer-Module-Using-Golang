package db
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
)

func Connect() (*sql.DB,error){
db,err:=sql.Open("mysql","root:raj1994@tcp(127.0.0.1:3306)/j2eedb")
if err!=nil{
panic(err)
}
return db,err
}
