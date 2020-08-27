package main

import "fmt"
import "strconv"
import "math/rand"
import "time"
import "strings"

func main() {
  defer pulih()                               // pulih() akan dijalankan terakhir karena ada defer
  rand.Seed(time.Now().UnixNano())            // rand.Seed untuk mendapatkan nomor baru ke variable
  max := 10
  min := 1
  angka1 := rand.Intn(max-min+1)+min
  angka2 := rand.Intn(max-min+1)+min

  var input string
  fmt.Printf("%d * %d : ", angka1, angka2)
  fmt.Scanln(&input)

  var hasil = angka1 * angka2
  var number int
  var err error
  number, err = strconv.Atoi(input)
  // mengubah string ke integer,
  // jika input dapat diubah ke integer, maka error akan kosong atau <nil>
  // jika tidak, maka error akan terisi

  if err == nil{
    // pengecekan apakah err terisi
    if number == hasil {
      // jika input/number sesuai dengan hasil
      fmt.Println("Benar")
      fmt.Println(err) // err == <nil>
    } else {
      fmt.Println("Salah")
      panic("jawaban salah")
      //membuat panic dengan value "Jawaban salah" yang nanti dikirim ke recover()
      //bertujuan agar menampilkan teks coba lagi
    }
  } else {
    fmt.Println("Salah")
    // fmt.Println(err) // invalid syntax
    panic(strings.TrimSpace(strings.Split(err.Error(), ":")[2]))
    //membuat panic dengan value Error/String *jika dengan .Error()* yang nanti dikirim ke recover()
  }

}

func pulih(){
  r := recover()
  // mengambil nilai dari panic
  if r != nil {

    fmt.Println(r)
    fmt.Println("Coba Lagi")
  } else {
    fmt.Println("Terima Kasih")
  }
}
