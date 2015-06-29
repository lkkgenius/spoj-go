package main
import "fmt"

func main(){
    var i int
    fmt.Scanf("%d", &i)
    for i != 42 {
        fmt.Println(i);
        fmt.Scanf("%d", &i)
    }

}
