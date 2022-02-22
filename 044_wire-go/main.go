package main

// 使用wire前
// func main() {

// 	message := provider.NewMessage("hello world (unuse wire!)")
// 	greeter := provider.NewGreeter(message)
// 	event := provider.NewEvent(greeter)

// 	event.Start()
// }

// 使用wire后
func main() {
	event := InitializeEvent("hello world (use wire!)")

	event.Start()
}
