package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Write what you want to encrypt:")

  encrypter, _ := reader.ReadString('\n')
  encrypter = strings.TrimSpace(strings.ToLower(encrypter))

  fixers := []string{""}

  for _, char := range encrypter {

    if char < 'a' || char > 'z' {
      if char == ' ' {
        continue
      }
      fmt.Println("Only characters from a to z are allowed, currently theres no support for numbers or special characters.")
      return
    }

    if char == ' ' {
      fmt.Print(" ")
    }

    output := int(char - 'a' + 1)

    if output > 9 {
      fixers = append(fixers, "-")
    } else {
      fixers = append(fixers, ".")
    }

    fmt.Print(output)
  }

    for i := 0; i < len(fixers); i++ {
        fmt.Print(fixers[i])
    }
}

// a (1) = . // o (15) = -
// abc (123) = ... // ola (15121) = --.