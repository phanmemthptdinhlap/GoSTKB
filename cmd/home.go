package main
import (
	"fmt"
	"net/http"
)
page.Home=func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
