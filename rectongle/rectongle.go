package rectongle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Generate() {
	var inputs string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputs = scanner.Text()
	}
	values := strings.Split(inputs, " ")
	rows, _ := strconv.Atoi(values[0])
	columns, _ := strconv.Atoi(values[1])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("B")
			} else {
				fmt.Print("W")
			}
		}
		fmt.Println()
	}
}
