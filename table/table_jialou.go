package table

import (
	"bytes"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int  {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func caculateScore(input string, scorePoint int) string {
	var b bytes.Buffer
	var my, your int
	my = 0
	your = 0
	for _, ch := range input {
		switch ch {
		case 'E':
			b.WriteString(fmt.Sprintf("%d:%d\n", my, your))
			break
		case 'W':
			my ++
		case 'L':
			your ++
		}
		if max(my, your) >= scorePoint && abs(my - your) >= 2 {
			b.WriteString(fmt.Sprintf("%d:%d\n", my, your))
			my = 0
			your = 0
		}
	}
	return b.String()

}

func tableJialou(input string) string {
	var b bytes.Buffer
	b.WriteString(caculateScore(input, 11))
	b.WriteString("\n")
	b.WriteString(caculateScore(input, 21))
	return b.String()
}

