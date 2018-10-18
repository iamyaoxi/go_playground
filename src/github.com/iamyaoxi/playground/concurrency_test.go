package playground 

import "testing"
import "fmt"

func TestSimpleConversation(t *testing.T) {
	c := SimpleConversation("John")

	for i:=0; i < 5; i++ {
		fmt.Println(<- c)
	}
	fmt.Println("Thanks John")
}

func TestFanIn(t *testing.T) {
	//Generates 2 channels from SimpleConversation
	ch1 := SimpleConversation("Annie")
	ch2 := SimpleConversation("Joe")

	chResult := FanIn(ch1, ch2)

	//Whatever it is, we only want to receive 20 first values from the fanIn channel
	for i:=0; i < 20; i++ {
		fmt.Println(<- chResult)
	}
	fmt.Println("Thanks Annie and Joe")
}