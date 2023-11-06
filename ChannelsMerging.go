package main

func Merge(a <-chan string, b <-chan string) <-chan string {
	c := make(chan string)
	done := make(chan bool)
	go func() {
		for aChan := range a {
			c <- aChan
		}
		done <- true
	}()
	go func() {
		for bChan := range b {
			c <- bChan
		}
		done <- true
	}()
	go func() {
		<-done
		<-done
		close(c)
	}()
	return c
}
