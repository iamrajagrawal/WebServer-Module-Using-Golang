package db
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
"net/http"
"io"
"strconv"
)

func UpdateAuthor(rw http.ResponseWriter,rq *http.Request){
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
fmt.Println("Erron in conversion")
panic(err)
}
vName:=rq.URL.Query()["name"][0]
rows,err=db.Query("Select code from author where code=?",vCode)
if err!=nil{
fmt.Println("Code does not exist")
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
if code==vCode{
_,err=db.Exec("Update author set name=? where code=?",vName,code1)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
}
io.WriteString(rw,"Author updated successfully")
g:=fmt.Sprintf("Author updated with name %s:",vName)
io.WriteString(rw,g)
fmt.Println("Done")
}
