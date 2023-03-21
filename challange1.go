package main

import "fmt"

func main() {
    i := 21
    j := true

    // menampilakan nilai i : 21
    fmt.Printf("%v \n", i)

    // menampilkan tipe data dari variabel i
    fmt.Printf("%T \n", i)

    // menampilkan tanda %
    fmt.Printf("%% \n")

    // menampilkan nilai boolean j : true
    fmt.Printf("%t \n\n", j)

    // menampilkan nilai boolean j : true
    fmt.Printf("%b \n", i) 

    // menampilkan unicode russia : Я (ya)
    fmt.Printf("%c \n", '\u042F')

    // menampilkan nilai base 10 : 21
    fmt.Printf("%d \n", i)

    // menampilkan nilai base 8 :25
    fmt.Printf("%o \n", i)

    // menampilkan nilai base 16 : f
    fmt.Printf("%x \n", 15)

    // menampilkan nilai base 16 : F
    fmt.Printf("%X \n", 15)

    // menampilkan unicode karakter Я : U+042F
    fmt.Printf("%U \n\n", 'Я')

    k := 123.456
    // menampilkan float : 123.456000
    fmt.Printf("%f \n", k)

    // menampilkan float scientific : 1.234560E+02
    fmt.Printf("%E \n", k)
}