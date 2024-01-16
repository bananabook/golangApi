package main
import(
"net/http"
"fmt"
)

func HdlTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("id:",r.FormValue("id"),":")
	fmt.Fprintln(w,"hi")
}
