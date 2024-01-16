package main
import(
"fmt"
"net/http"
)
func h(e error)bool{
	if e!=nil{
		fmt.Println("ERROR->")
		fmt.Println(e)
		fmt.Println("<-ERROR")
		return true
	}
	return false
}
func HdlError(w http.ResponseWriter, r *http.Request){
	http.Error(w, "oops", http.StatusInternalServerError)
}
func HdlRedir(w http.ResponseWriter, r *http.Request){
	http.Redirect(w,r,"http://localhost:8081/error",http.StatusFound)
	//http.Redirect(w,r,"https://google.com",http.StatusMovedPermanently)
}
