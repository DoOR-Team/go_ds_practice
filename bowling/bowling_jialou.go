package table

import (
	"bytes"
	"fmt"
	"strings"
)

func bowlingJialou(input string) string {
	var b bytes.Buffer
	var score []int

	input = strings.Replace(input, " ", "", -1)
	j := 0
	//round := 1

	for j < len(input) {
		now := 0
		// 当前轮全中
		if input[j] == '/' {
			// 基础分10
			now = 10
			if j+1 == len(input) {
				break
			}
			// 下一轮全中
			if input[j+1] == '/' {
				now += 10
				// 下第二轮全中
				if j+2 == len(input) {
					break
				}

				if input[j+2] == '/' {
					now += 10
				} else {
					now += int(input[j+2] - '0');
				}
			} else {
				// 第二轮非全中，第二轮采用方式2计分
				// 如果未补中，则为10分
				if j+2 == len(input) {
					break
				}

				if input[j+2] == '/' {
					now += 10
				} else {
					//非补中，打了几个，就是计分
					now += int(input[j+1]-'0') + int(input[j+2]-'0')
				}
			}
			// 因为当前全中，没有补中轮，只向后1轮
			j += 1
		} else {
			if j+1 == len(input) {
				break
			}

			if input[j+1] == '/' {

				if j+2 == len(input) {
					break
				}
				if input[j+2] == '/' {
					now = 20
				} else {
					now = 10 + int(input[j+2]-'0')
				}
			} else {
				now = int(input[j]-'0') + int(input[j+1]-'0')
			}
			// 多1个补中轮，向后2轮
			j += 2
		}
		score = append(score, now)
		//round++
	}

	for i, s := range score {
		b.WriteString(fmt.Sprintf("%d", s))
		if i < len(score)-1 {
			b.WriteString(" ")
		}
	}
	total := 0
	b.WriteString("\n")
	for i, s := range score {
		b.WriteString(fmt.Sprintf("%d", s+total))
		if i < len(score)-1 {
			b.WriteString(" ")
		}
		total += score[i]
	}
	return b.String()
}
