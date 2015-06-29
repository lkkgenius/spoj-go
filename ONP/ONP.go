package main
import (
    "fmt"
    "bytes"
)

func onp(s string) string {
    l := make([]int, 0, 50)
    var o bytes.Buffer

    for _,c := range s {
        switch c {
        case '(', '+', '-', '*', '/', '^':
            l = append(l, int(c))
        case ')':
            for l[len(l)-1] != '(' {
                o.WriteString(string(l[len(l)-1]))
                l = l[:len(l)-1]
            }
            l = l[:len(l)-1]
        default:
            o.WriteString(string(c))
        }
    }
    return o.String()
}

func main() {
    var r int
    var s string
    fmt.Scan(&r)
    for i := 1; i<= r; i++ {
        fmt.Scanln(&s)
        fmt.Println(onp(s))
    }
}
