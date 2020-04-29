package main
 
import (
    "fmt"
    "net/http"
)
 
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "home")
}
 
func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "users")
}
 
func main() {
    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/users", usersHandler)
    http.ListenAndServe(":8181", nil)
