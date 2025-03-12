package main
import "fmt"

func main(){
    var masukan, hasil string
    fmt.Scan(&masukan)
   for i := len(masukan) - 1; i >= 0; i-- {
        hasil += string(masukan[i])
    }
    fmt.Print(hasil)
}
