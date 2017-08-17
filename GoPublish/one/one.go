package one
import (
"io"
"net/http"
"strings"
"os"
"html/template"
"./private/db"
)
type Error404 struct {
Resource string
}

func sendError404(rw http.ResponseWriter,rq *http.Request){
//http.ServeFile(rw,rq,"./one/Error404.html")
var error404 Error404
error404.Resource=rq.URL.Path
t,err:=template.ParseFiles("./one/Error404.html","./one/MT.html","./one/MB.html")
if err!=nil { 
panic(err)
}
t.Execute(rw,error404)
}


func HttpHandler(rw http.ResponseWriter,rq *http.Request) {
uri:=rq.URL.Path
resource:=uri[5:]
var fileToServe string
if len(resource)==0 || resource=="publish" || strings.HasPrefix(resource,"publish/") {
// static content
if len(resource)<=8 {
fileToServe="./one/public/index.html"
} else {
fileToServe="./one/public/"+resource[8:]
}
file,err:=os.Stat(fileToServe)
if err!=nil || file.IsDir(){
sendError404(rw,rq)
} else {
http.ServeFile(rw,rq,fileToServe)
}
} else {
if resource=="AddAuthor.go"{
db.AddAuthor(rw,rq)
}
if resource=="GetAuthors.go"{
db.GetAuthors(rw,rq)
}
if resource=="UpdateAuthor.go"{
db.UpdateAuthor(rw,rq)
}
if resource=="DeleteAuthor.go"{
db.DeleteAuthor(rw,rq)
}
io.WriteString(rw,"OK")
}
}