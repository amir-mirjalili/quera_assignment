package chess

import (
	"bufio"
	"fmt"
	"os"
)

const n = 8

func Threats() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	c := int(s[0] - 'a')
	r := 7 - int(s[1]-'1')

	var board [n][n]byte

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'

		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == r || j == c || i+j == r+c || i-j == r-c {
				board[i][j] = 'X'
			}
		}
	}

	board[r][c] = 'Q'

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(out, "%c ", board[i][j])
		}
		fmt.Fprintln(out)
	}
}
