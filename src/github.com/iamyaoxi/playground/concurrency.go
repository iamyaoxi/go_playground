package playground

import "time"
import "fmt"

//Taken from a boring conversation by Rob Pike
//Watch: https://youtu.be/f6kdp27TYZs

//SimpleConversation follows Generator pattern that will generate... simple conversation from a character
func SimpleConversation(character string) <- chan string {
	//We make channel string
	converse := make(chan string) 

	//and we execute goroutine function
	go func() {
		//which will do these in a different thread

		//the point is to keep sending us with as many as possible
		//Until the program exited
		for i := 0; ; i++ { 
			//Send the message to the channel
			converse <- fmt.Sprintf("%v says: %v", character, i)
			//But can sleep for a second. so we can monitor what happen
			time.Sleep(time.Second)
		}
	}()

	//return the channel as receive only channel
	return converse
}

//FanIn follows Multiplexing pattern.
//It will combine 2 channels into one channel
func FanIn(ch1, ch2 <- chan string) <- chan string {
	//First we make channel that holds other channel
	fanIn := make(chan string)

	//Then here we execute goroutine function
	go func() { 
		//This line here means 
		//Keep receiving and sending any value from the channel 1 to the fanIn channel
		for {
			fanIn <- <-ch1
			}
	}()
	go func() { 
		//Keep receiving and sending any value from the channel 2 to the fanIn channel
		for {
			fanIn <- <- ch2
			}
	}()

	//Return fanIn channel
	return fanIn
}