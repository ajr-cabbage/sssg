package serve

import (
	"fmt"
	"net/http"
)

func Serve(path string, port int) error {
	fs := http.FileServer(http.Dir(path))
	fmt.Printf("Serving '%s' on Port '%d'\nCtrl+c to exit...\n", path, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), fs)
	if err != nil {
		return err
	}
	return nil
}
