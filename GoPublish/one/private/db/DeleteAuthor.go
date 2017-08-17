package db
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
"net/http"
"io"
"strconv"
)

func DeleteAuthor(rw http.ResponseWriter,rq *http.Request){
db,err1:=Connect()
if err1!=nil {
panic(err1)
}
defer db.Close()
var err error
var rows *sql.Rows
code1:=rq.URL.Query()["code"][0]
vCode,err:=strconv.Atoi(code1)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
rows,err=db.Query("Select code from author where code=?",vCode)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
defer rows.Close()
var code int
for rows.Next(){
err=rows.Scan(&code)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
}
if code!=vCode{
io.WriteString(rw,"Invalid code")
} else {
_,err=db.Exec("Delete from author where code=?",vCode)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Unable to delete author")
return
}
io.WriteString(rw,"Author deleted.")
g:=fmt.Sprintf("Author deleted with code %d",vCode)
io.WriteString(rw,g)
fmt.Println("Done")
}
}
