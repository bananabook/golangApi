package main
import(
"fmt"
"net/http"
"encoding/json"
//"io/ioutil"
)
type response struct{
		Status string `json:"status"`
}
func NewResponse(status string)*response{
	return &response{Status:status}
}
func (C *Container) WriteError(w http.ResponseWriter, message string, code int){
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(*NewResponse(message))
}
func (C *Container) HdlApi(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(*NewResponse("connected"))
}
func (C *Container) HdlApiCreate(w http.ResponseWriter, r *http.Request){
	if r.Method!="POST"{
		fmt.Println(r.Method)
		C.WriteError(w,"invalid method",400)
		return
	}
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		C.WriteError(w,"invalid json",400)
		return
	}
	e=C.Create(u)
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("success"))
	if h(e){ http.Error(w,"error decoding",http.StatusInternalServerError);return }
}
func (C *Container) HdlApiRead(w http.ResponseWriter, r *http.Request){
	if r.Method!="GET"{
		C.WriteError(w,"invalid method",400)
		return
	}
	//u:=Entry{}
	//e:=json.NewDecoder(r.Body).Decode(&u)
	//if h(e){
	//	WriteError(w,"invalid json",400)
	//	return
	//}
	name:=r.FormValue("name")
	entry,e:=C.Read(name)
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*entry)
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
	}
}
func (C *Container) HdlApiUpdate(w http.ResponseWriter, r *http.Request){
	if r.Method!="PUT"{
		C.WriteError(w,"invalid Method",400)
		return
	}
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		C.WriteError(w,"invalid json",400)
		return
	}
	e=C.Update(u)
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("updated"))
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
	}
}
func (C *Container) HdlApiDelete(w http.ResponseWriter, r *http.Request){
	if r.Method!="DELETE"{
		C.WriteError(w,"invalid Method",400)
		return
	}
	//u:=Entry{}
	//e:=json.NewDecoder(r.Body).Decode(&u)
	//if h(e){
	//	WriteError(w,"invalid json",400)
	//	return
	//}
	name:=r.FormValue("name")
	e:=C.Delete(name)
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("deleted"))
	if h(e){
		C.WriteError(w,fmt.Sprint(e),500)
	}
}
