package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strings"

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scan_line() string {
	reader := bufio.NewReader(os.Stdin)
    str, err := reader.ReadString('\n')
	check_error(err)
    str = strings.TrimSpace(str)
	return str
}

func main() {
	// Scan line as string
	fmt.Print("Input string (can contain whitespace): ")
	str := scan_line()
	fmt.Println("Your string is:", str)
}
