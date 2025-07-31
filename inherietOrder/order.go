package inherietorder

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fixName(s string) string {
	s = strings.ReplaceAll(s, ",", "")
	return strings.ReplaceAll(s, ":", "")
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func Order() {
	in := bufio.NewScanner(os.Stdin)
	var lines []string
	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, "class") {
			lines = append(lines, s)
		}
	}

	extends := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		className := fixName(parts[1])
		var parents []string
		if len(parts) >= 4 {
			for i := 3; i < len(parts); i++ {
				parents = append(parents, fixName(parts[i]))
			}
		}
		extends[className] = parents
	}

	for i := 0; i < len(extends); i++ {
		for c, parents := range extends {
			for _, p := range parents {
				for _, np := range extends[p] {
					if !contains(extends[c], np) {
						extends[c] = append(extends[c], np)
					}
				}
			}
		}
	}

	possible := true
	for c, parents := range extends {
		if contains(parents, c) {
			possible = false
			break
		}
	}

	if possible {
		fmt.Println("possible")
	} else {
		fmt.Println("impossible")
	}
}
