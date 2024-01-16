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
	C.Create(Entry{Name:"n",Content:"c"})
	log.Println(C)
	C.Create(Entry{Name:"n2",Content:"c"})
	log.Println(C)
	C.Create(Entry{Name:"n3",Content:"c"})
	log.Println(C)
	C.Delete(Entry{Name:"n",Content:"c2"})
	log.Println(C)
}

func server(){
	server:=http.NewServeMux()
	server.HandleFunc("/",HdlRoot)
	server.HandleFunc("/error",HdlError)
	server.HandleFunc("/redir",HdlRedir)
	server.HandleFunc("/api/",HdlApi)
	server.HandleFunc("/api/create",HdlApiCreate)
	server.HandleFunc("/api/read",HdlApiRead)
	server.HandleFunc("/api/update",HdlApiUpdate)
	server.HandleFunc("/api/delete",HdlApiDelete)
	go func(){ log.Fatal(http.ListenAndServe(":8080",server)) }()
	Interactive()
	EndSync()
}
