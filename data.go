package main
import(
	"errors"
	"fmt"
	"encoding/gob"
	"os"
)
var (
	ErrNotFound error=errors.New("not found")
	ErrAlreadyExists error=errors.New("already exists")
)
var C Container
type Container []Entry
type Entry struct{
	Name string
	Content string
}
func (C *Container)StartSync()(error){
	f,e:=os.OpenFile("backup.gob",os.O_RDWR,0700)
	if h(e){ return nil}
	e=gob.NewDecoder(f).Decode(&C)
	if h(e){ return nil}
	return nil
}
func (C Container)EndSync(){
	f,e:=os.OpenFile("backup.gob",os.O_RDWR|os.O_CREATE,0700)
	if h(e){ return }
	E:=gob.NewEncoder(f)
	e=E.Encode(C)
	if h(e){ return }
}
//var C Container
func (C *Container)String()string{
	var out string
	var first bool=true
	for _,v:=range *C{
		if first{
			first=false
		}else{
			out+="\n"
		}
		out+=fmt.Sprintf("Name:%v\tContent:%v",v.Name,v.Content)
	}
	return out
}
func (C *Container)Create(entry Entry)(error){
	_,e:=C.Read(entry.Name)
	if e!=nil&&e!=ErrNotFound{
		return e
	}
	if e==nil{
		return ErrAlreadyExists
	}
	*C=append(*C,entry)
	return nil
}
func (C *Container)Read(name string)(*Entry, error){
	for _,v:=range *C{
		if name==v.Name{
			return &v,nil
		}
	}
	return nil,ErrNotFound
}
func (C *Container)Update(entry Entry)(error){
	//finding,e:=c.Read(entry.Name)
	i:=0
	var found bool
	for ;i<len(*C);i++{
		v:=(*C)[i]
		if v.Name==entry.Name{
			found=true
		}
	}
	i--
	if !found{
		return ErrNotFound
	}
	//if e!=nil{
	//	return e
	//}
	(*C)[i].Content=entry.Content
	//finding.Content=entry.Content
	return nil
}
func (C *Container)Delete(entry string)(error){
	var i int
	var found bool
	for ;i<len(*C);i++{
		v:=(*C)[i]
		if v.Name==entry{
			found=true
			break
		}
	}
	if !found{
		return ErrNotFound
	}
	*C=append((*C)[:i],(*C)[i+1:]...)
	return nil
}
