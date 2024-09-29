package main

import "bufio"
import "fmt"
import "log"
import "os"

//lint:ignore U1000 (example)
func scan_line() string {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	line := scanner.Text()
	return line
}

//lint:ignore U1000 (example)
func scan_all_lines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return lines
}

func main() {
	// Scan line as string
	fmt.Print("Input string (can contain whitespace): ")
	str := scan_line()
	fmt.Println("Your string is:", str)
}
