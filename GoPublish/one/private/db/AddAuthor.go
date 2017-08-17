package db
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
"net/http"
"io"
)

func AddAuthor(rw http.ResponseWriter,rq *http.Request){
db,err1:=Connect()
if err1!=nil {
panic(err1)
}
defer db.Close()
var err error
var rows *sql.Rows
vName:=rq.URL.Query()["name"][0]
rows,err=db.Query("Select name from author where name=?",vName)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
defer rows.Close()
var name string
for rows.Next(){
err=rows.Scan(&name)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
}
if name==vName{
io.WriteString(rw,"Author already exists")
}
st,err:=db.Prepare("insert into author(name) values(?);")
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
var result sql.Result
var code int64
result,err=st.Exec(vName)
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
code,err=result.LastInsertId()
if err!=nil{
io.WriteString(rw,err.Error())
io.WriteString(rw,"Problem")
return
}
io.WriteString(rw,fmt.Sprintf("%d,%s",code,vName));
fmt.Println("Author added with code as : ",code)
fmt.Println("Done")
}
