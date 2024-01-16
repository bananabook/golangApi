package main
import(
"fmt"
"atomicgo.dev/keyboard"
"atomicgo.dev/keyboard/keys"
)
func (C *Container)Interactive(){
	fmt.Println("X")
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	switch key.Code{
		case keys.CtrlC,keys.CtrlQ,keys.CtrlO:
			return true, nil
	}
	switch key.String(){
		case "q","Q","x","X":
			return true, nil
		case "i","I":
			fmt.Println(*C)
	}
	
	return false, nil
	})
}
