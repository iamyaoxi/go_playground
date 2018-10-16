package playground

import "time"

//Taken from a boring conversation by Rob Pike
//Watch: https://youtu.be/f6kdp27TYZs
func SimpleConversation(converse chan string, timeLimit int) {
	for i := 0; ; i++{
		converse <- "I'm saying something"
		time.Sleep(time.Second)
	}
}