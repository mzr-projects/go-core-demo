package channels

import (
	"log"
	"net/http"
	"time"
)

/*
	When we start a go program actually we're running a single go routine, and it's an engine that executes code
*/

func appWithoutConcurrency() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
	}

	/*
		Here we create a variable of channel type string

		channel <- 5 			Here we send 5 into channel
		myNumber <- channel 	Wait for a value to be sent into the channel When we get one assign to the channel
								fmt.Println(<- channel)
	*/
	c := make(chan string)

	//for _, link := range links {
	/*
		#	This "go" key word will start checkLink in a new go routine By default go use only one core of CPU
			for routines.
		#	We only can use this keyword in front of the function calls.
		#	And of course it's not that easy to use just go keyword because the main routine will create all
			the routines and when there is no work it will exit, and we get no output from our code here we must
			use channel to communicate with the main and the other routines and channels are typed means we can
			send messages with a specific type
		#	To use channels we must them to the function like below
	*/
	//go checkLink(c, link)
	//}

	/*
		This is a blocking line of code the main routine will wait for a channel to send a message as soon as one
		of the channels finishes the job this will be run and the program will be closed, Here in our example one
		of the links fetches faster, then we will get teh message from that channel and every thing is done

		log.Printf(<-c)
	*/
	/*
		Here in this for loop we actually wait for all routines to get their job done from the fastest to the slowest
		will send their messages and after getting all responses the program will shut down
	*/
	//for i := 0; i < len(links); i++ {
	//	log.Printf(<-c)
	//}

	/*
		Work on heartbeat of links repeatedly

		for {
			go heartBeat(c, <-c)
		}

		Another form of loop in go here we wait for the channel in the range to return a value as soon as it gets
		the value we assign it to the l variable and do our job just like the above one

		for l := range c {
			go heartBeat(c, l)
		}
	*/

	for _, link := range links {
		go heartBeat(c, link)
	}

	for l := range c {
		//go heartBeat(c, l)
		/*
				To make some pauses between checking the links the beat approach is to use function literals in GO
				But in this for loop we're ready to get the data from range c and assign it to l variable but this l
				is going to be assigned in the main routine(memory) and the other go routines must have their own l
			    variables, so we must pass the l to the function literal as well
		*/
		go func(link string) {
			time.Sleep(time.Second * 3)
			heartBeat(c, link)
		}(l)
	}
}

func checkLink(c chan string, link string) {
	/*
		Here we are ignoring the get response just deal with the error
	*/
	/*
		this http.Get(link) the the blocking call
	*/
	_, err := http.Get(link)
	if err != nil {
		log.Printf("link :%s might be down", link)
		c <- "Might be down ;("
		return
	}

	log.Printf("Link %s is up ;)", link)
	c <- "Its Up ;)"
}

func heartBeat(c chan string, link string) {
	/*
		Here we are ignoring the get response just deal with the error
	*/
	/*
		this http.Get(link) is the blocking call, we're waiting to get response from the URL
	*/
	_, err := http.Get(link)
	if err != nil {
		log.Printf("link :%s is down", link)
		c <- link
		return
	}

	log.Printf("Link %s is up ;)", link)
	c <- link
}
