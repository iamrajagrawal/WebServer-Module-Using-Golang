package db
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
"net/http"
"io"
)

func GetAuthors(rw http.ResponseWriter,rq *http.Request){
db,err1:=Connect()
if err1!=nil {
panic(err1)
}
defer db.Close()
var err error
var rows *sql.Rows
rows,err=db.Query("Select * from author")
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
defer rows.Close()
var comma=false
var code int
var name string
for rows.Next(){
err=rows.Scan(&code,&name)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
if comma==true{
io.WriteString(rw,",")
}
io.WriteString(rw,fmt.Sprintf("%d,%s",code,name));
comma=true
}
err=rows.Err()
if err!=nil{
panic(err.Error())
}
fmt.Println("Done")
}
