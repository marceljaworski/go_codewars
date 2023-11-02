func Merge(a <-chan string, b <-chan string) <-chan string {
	c := make(chan string)
	done := make(chan string)
	go func() {
		for aChan := range a {
			c <- aChan
		}
		done <- ""
	}()
	go func() {
		for bChan := range b {
			c <- bChan
		}
		done <- ""
	}()
	go func() {
		<-done
		<-done
		close(c)
	}()
	return c
}
  