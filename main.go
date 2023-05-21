package main

func main() {
	server := NewAPIServer(":8001")
	server.Run()
}