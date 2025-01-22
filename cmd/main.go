package cmd

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Retail pulse Image Processor is working")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
