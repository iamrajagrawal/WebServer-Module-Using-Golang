package main
import (
"fmt"
"net/http"
"./one"
)
func main(){
http.HandleFunc("/one/",one.HttpHandler)
fmt.Println("Server listening at : 6060")
http.ListenAndServe(":6060",nil)
}