package serve

import (
	"fmt"
	"net/http"
)

func Serve(path string, port int) error {
	fs := http.FileServer(http.Dir(path))
	fmt.Printf("Serving '%s' on Port '%d'\n", path, port)
	fmt.Printf("Navigate to http://localhost:%d\n", port)
	fmt.Println("Ctrl+c to exit...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), fs)
	if err != nil {
		return err
	}
	return nil
}
