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
func WriteError(w http.ResponseWriter, message string, code int){
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(*NewResponse(message))
}
func HdlApi(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(*NewResponse("connected"))
}
func HdlApiCreate(w http.ResponseWriter, r *http.Request){
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		WriteError(w,"invalid json",400)
		return
	}
	e=C.Create(u)
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("success"))
	if h(e){ http.Error(w,"error decoding",http.StatusInternalServerError);return }
}
func HdlApiRead(w http.ResponseWriter, r *http.Request){
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		WriteError(w,"invalid json",400)
		return
	}
	entry,e:=C.Read(u.Name)
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*entry)
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
	}
}
func HdlApiUpdate(w http.ResponseWriter, r *http.Request){
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		WriteError(w,"invalid json",400)
		return
	}
	e=C.Update(u)
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("update"))
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
	}
}
func HdlApiDelete(w http.ResponseWriter, r *http.Request){
	u:=Entry{}
	e:=json.NewDecoder(r.Body).Decode(&u)
	if h(e){
		WriteError(w,"invalid json",400)
		return
	}
	e=C.Delete(u)
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
		return
	}
	e=json.NewEncoder(w).Encode(*NewResponse("success"))
	if h(e){
		WriteError(w,fmt.Sprint(e),500)
	}
}
