package main
import "fmt"

func Encrypt(s string) string {
    enc := []rune(s)

    for i := 1; i < len(enc); i++ {
        enc[i] = (enc[i] + enc[i-1] - 128) % 26 + 64
    }

    return string(enc)
}

func Decrypt(s string) string {
    dec := []rune(s)

    for i := len(dec) - 1; i > 0; i-- {
        dec[i] = (dec[i] - dec[i-1] + 26) % 26 + 64
    }

    return string(dec)
}

func main() {
    var s string
    fmt.Scan(&s)

    // Encrypt the string
    fmt.Println(Decrypt(s))

}