package main
import(
"testing"
"net/http"
"time"
"encoding/json"
"bytes"
)
func TestInternal(t *testing.T){
	C:=new(Container)
	e:=C.Create(Entry{Name:"n",Content:"c"})
	if e!=nil{
		t.Error("cannot create")
	}
	read,e:=C.Read("n")
	if e!=nil{
		t.Error("cannot read")
	}
	if read.Name!="n"{
		t.Error("create name not same")
	}
	if read.Content!="c"{
		t.Error("create content not same")
	}

	e=C.Update(Entry{Name:"n",Content:"cc"})
	if e!=nil{
		t.Error("cannot update")
	}
	if read.Name!="n"{
		t.Error("update name not same")
	}
	if read.Content!="c"{
		t.Error("update content not same")
	}
	e=C.Delete("n")
	if e!=nil{
		t.Error("cannot delete")
	}
	_,e=C.Read("n")
	if e!=ErrNotFound{
		t.Error("cannot delete correctly")
	}
}
func TestExternal(t *testing.T){
	go server()
	time.Sleep(time.Second*2)
	/* create */
	target:="http://localhost:8080/api/"
	entry,_:=json.Marshal(Entry{Name:"n",Content:"c"})
	req,e:=http.NewRequest("POST",target+"create",bytes.NewBuffer(entry))
	if e!=nil{
		t.Error("no get")
		t.Error(e)
	}
	c:=new(http.Client)
	resp,_:=c.Do(req)
	responseStruct:=struct{Status string `json:"status"`}{}
	json.NewDecoder(resp.Body).Decode(&responseStruct)
	if responseStruct.Status!="success"{
		t.Error("not successfull creation")
		t.Error(responseStruct.Status)
	}
	/* read */
	resp,e=http.Get(target+"read?name=n")
	readEntry:=Entry{}
	json.NewDecoder(resp.Body).Decode(&readEntry)
	if readEntry.Name!="n"{
		t.Error("incorrect read name")
	}
	if readEntry.Content!="c"{
		t.Error("incorrect read content")
	}
	/* update */
	entry,_=json.Marshal(Entry{Name:"n",Content:"cc"})
	req,_=http.NewRequest("PUT",target+"update",bytes.NewBuffer(entry))
	resp,_=c.Do(req)
	responseStruct=struct{Status string `json:"status"`}{}
	json.NewDecoder(resp.Body).Decode(&responseStruct)
	if responseStruct.Status!="updated"{
		t.Error("not successfull update")
		t.Error(responseStruct.Status)
	}
	
	resp,e=http.Get(target+"read?name=n")
	json.NewDecoder(resp.Body).Decode(&readEntry)
	if readEntry.Name!="n"{
		t.Error("incorrect read name")
	}
	if readEntry.Content!="cc"{
		t.Error("incorrect read content")
	}
	/* delete */
	req,_=http.NewRequest("DELETE",target+"delete?name=n",nil)
	resp,_=c.Do(req)
	json.NewDecoder(resp.Body).Decode(&responseStruct)
	if responseStruct.Status!="deleted"{
		t.Error("not updated")
	}
}
