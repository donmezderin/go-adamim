package list

import "fmt"

func main()  {

    var first *Node;

    first = first.AddInOrder("ciao")
    first.Print()
    first = first.AddInOrder("come")
    first.Print()
    first = first.AddInOrder("stai")
    first.Print()
    first = first.AddInOrder("abele")
    first.Print()
    first = first.AddInOrder("cane")
    first.Print()
    first = first.AddInOrder("zanzara")
    first.Print()

    fmt.Println(first.length())
}
