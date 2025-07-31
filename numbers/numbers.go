package numbers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Calculate() {
	scanner := bufio.NewScanner(os.Stdin)

	length := 0
	numbers := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		if length == 0 {
			length = num
			continue
		}
		numbers = append(numbers, num)
		if len(numbers) == length {
			break
		}
	}

	yekaan := []string{"", "yek", "do", "se", "chahaar", "panj", "shesh", "haft", "hasht", "noh"}
	dahgaan := []string{"", "", "bist", "si", "chehel", "panjaah", "shast", "haftaad", "hashtaad", "navad", "sad"}
	dahtabist := []string{"dah", "yaazdah", "davaazdah", "sizdah", "chahaardah", "paanzdah", "shaanzdah", "hifdah", "hijdah", "noozdah"}
	sadgaan := []string{"", "sad", "devist", "sisad", "chahaarsad", "paansad", "sheshsad", "haftsad", "hashtsad", "nohsad"}

	for _, val := range numbers {
		a := val % 10
		b := (val / 10) % 10
		c := (val / 100) % 10

		ans := sadgaan[c]
		if c != 0 && a+b != 0 {
			ans += "o "
		}

		if b == 1 {
			ans += dahtabist[a]
		} else {
			ans += dahgaan[b]
			if b != 0 && a != 0 {
				ans += "o "
			}
			ans += yekaan[a]
		}
		fmt.Println(ans)

	}
}
