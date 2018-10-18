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


func TestFanInOrdered(t *testing.T) {
	//Generates 2 channels from SimpleConversation
	josh := OrderedConversation("Josh")
	celine := OrderedConversation("Celine")

	chResult := FanInOrdered(josh, celine)

	//Here, the fanIn will return both Josh and Celine in sequence
	//And there will be 10 messages from each side
	for i:=0; i <= 10; i++ {
		fromJosh := <- chResult
		fmt.Println(fromJosh.message)

		fromCeline := <- chResult
		fmt.Println(fromCeline.message)

		fromJosh.wait <- true
		fromCeline.wait <- true
	}
	fmt.Println("Thanks Josh and Celine")
}