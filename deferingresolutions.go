// We can delay a function call to the end of the current scope
// by using the defer keyword. defer tells Go to run a function,
// but at the end of the current function. 
// This is useful for logging, file writing, and other utilities.


package main
import "fmt"

func queryDatabase(query string) string {
  var result string
  connectDatabase()
  // Add deferred call to disconnectDatabase() here
  defer disconnectDatabase() //what defering does is that it lets the code that you are defering execute after a function has been called   
  if query == "SELECT * FROM coolTable;" {
    result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
  }  
  fmt.Println(result)
  return result
}

func connectDatabase() {
  fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
  fmt.Println("Disconnecting from the database.")
}

func main() {
  queryDatabase("SELECT * FROM coolTable;")
}