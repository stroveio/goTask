// GO CONCURRENCY TASK

package main

// connection's holdUntil time is being set to 3 seconds when getting connected
// implement function, checking if connection's holdUntil time exceeded current time
func (c *connection) isActive() bool {
	return false
}

func main() {
	runMockServer()
	defer checkSuccess()

	// handle here all requests located in activeRequest channel
	// use `handleRequest(r request)` function
}
