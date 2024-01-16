package main
import(
"net/http"
//"encoding/json"
//"fmt"
"io"
"os"
_ "embed"
)
// //go:embed files/index.html
// var index string
func HdlRoot(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w,index)
	f,e:=os.Open("files/index.html")
	if h(e){ return }
	io.Copy(w,f)
}
