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
type Container []Entry
type Entry struct{
	Name string
	Content string
}
func init(){
	f,e:=os.OpenFile("backup.gob",os.O_RDWR,0700)
	if h(e){ return }
	e=gob.NewDecoder(f).Decode(&C)
	if h(e){ return }
}
func EndSync(){
	f,e:=os.OpenFile("backup.gob",os.O_RDWR|os.O_CREATE,0700)
	if h(e){ return }
	E:=gob.NewEncoder(f)
	e=E.Encode(C)
	if h(e){ return }
}
var C Container
func (c *Container)String()string{
	var out string
	var first bool=true
	for _,v:=range *c{
		if first{
			first=false
		}else{
			out+="\n"
		}
		out+=fmt.Sprintf("Name:%v\tContent:%v",v.Name,v.Content)
	}
	return out
}
func (c *Container)Create(entry Entry)(error){
	_,e:=C.Read(entry.Name)
	if e!=nil&&e!=ErrNotFound{
		return e
	}
	if e==nil{
		return ErrAlreadyExists
	}
	C=append(C,entry)
	return nil
}
func (c *Container)Read(name string)(*Entry, error){
	for _,v:=range *c{
		if name==v.Name{
			return &v,nil
		}
	}
	return nil,ErrNotFound
}
func (c *Container)Update(entry Entry)(error){
	//finding,e:=c.Read(entry.Name)
	i:=0
	var found bool
	for ;i<len(*c);i++{
		v:=(*c)[i]
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
	(*c)[i].Content=entry.Content
	//finding.Content=entry.Content
	return nil
}
func (c *Container)Delete(entry Entry)(error){
	var i int
	var found bool
	for ;i<len(*c);i++{
		v:=(*c)[i]
		if v.Name==entry.Name{
			found=true
			break
		}
	}
	if !found{
		return ErrNotFound
	}
	*c=append((*c)[:i],(*c)[i+1:]...)
	return nil
}
