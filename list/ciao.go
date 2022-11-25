package list

import "fmt"

/*  to build from terminal use:
      go build ciao.go
      then use ./ciao to execute, if there were no errors
*/

func main(){
  var n, i int
  fmt.Println("Quante volte vuoi essere salutato? ")
  fmt.Scan(&n)
  for i=0; i<n;i++{
    fmt.Println("Ciao")
  }
}
