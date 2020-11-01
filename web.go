package main
import "net/http"
func helloHandler(writer http.ResponseWriter, request *http.Request) {
	// Note: unhandled error!
	writer.Write([]byte("<h1>Hello, web!</h1>"))
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	// Note: unhandled error!
	http.ListenAndServe("localhost:8080", nil)
}