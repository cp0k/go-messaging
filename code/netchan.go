import "golang.org/x/exp/old/netchan"

type msg1 struct {
	Addr Addr
	Who  string
}
type msg2 struct {
	Greeting string
	Who      string
}

func server() {
	c1 := make(chan msg1)
	l, err := netchan.Listen(":12345")
	// Announce a greeting service.
	addr, err := l.Publish(c1, "GreetingService")
	for {
		m1 := <-c1
		c2 := make(chan msg2)
		err = netchan.Dial(m1.Addr, c2)
		c2 <- msg2{"Hello", m1.Who}
		close(c2)
	}
}
