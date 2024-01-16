package main
import(
"net/http"
"log"
)
func main(){
	//test()
	server()
}
func test(){
	var C Container
	C.Create(Entry{Name:"n",Content:"c"})
	log.Println(C)
	C.Create(Entry{Name:"n2",Content:"c"})
	log.Println(C)
	C.Create(Entry{Name:"n3",Content:"c"})
	log.Println(C)
	C.Delete("n")
	log.Println(C)
}

func server(){
	server:=http.NewServeMux()
	C:=new(Container)
	C.StartSync()
	server.HandleFunc("/",HdlRoot)
	server.HandleFunc("/error",HdlError)
	server.HandleFunc("/test",HdlTest)
	server.HandleFunc("/test/",HdlTest)
	server.HandleFunc("/redir",HdlRedir)
	server.HandleFunc("/api/",C.HdlApi)
	server.HandleFunc("/api/create",C.HdlApiCreate)
	server.HandleFunc("/api/read",C.HdlApiRead)
	server.HandleFunc("/api/update",C.HdlApiUpdate)
	server.HandleFunc("/api/delete",C.HdlApiDelete)
	go func(){ log.Fatal(http.ListenAndServe(":8080",server)) }()
	C.Interactive()
	C.EndSync()
}
