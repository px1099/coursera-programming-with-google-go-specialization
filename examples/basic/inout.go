package basic

import "bufio"
import "fmt"
import "os"
import "strings"

//lint:ignore U1000 (example)
func printing_example() {
	/*	block
		comment	*/
	x := "Minh"
	fmt.Println("Hello ", x)
	fmt.Printf("Hi %s\n", x)
}

//lint:ignore U1000 (example)
func scanning_example() {
	var appleNum, bananaNum int
	fmt.Print("Number of apple & banana: ")  // 4 5
	num, _ := fmt.Scan(&appleNum, &bananaNum)
	fmt.Println("Number of successful scan: ", num)
}

//lint:ignore U1000 (example)
func read_line_as_string() {
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)  // remove \n
	_ = text
}

//lint:ignore U1000 (example)
func read_all_line_as_string() {
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
    	line := scanner.Text()
		_ = line
	}
}
