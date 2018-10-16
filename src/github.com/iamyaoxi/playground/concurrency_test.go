package playground 

import "testing"
//import "fmt"

func TestSimpleConversation(t *testing.T) {
	c := make(chan string)
	limit := 5

	go SimpleConversation(c, limit)

	for i:=0; i < limit; i++ {
		//fmt.Println(<- c)
	}
	//fmt.Println("Ok. I'm leaving. Bye")
}