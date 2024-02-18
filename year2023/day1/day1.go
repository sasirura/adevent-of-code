package main

import (
	"bytes"
	"os"
	"strconv"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func main() {
	b, err := os.ReadFile("day1.txt")

	if err != nil {
		logrus.Fatal(err)
	}
	b2, err := os.ReadFile("day1_part2.txt")

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Score: %d", part1(b))
	logrus.Infof("Score: %d", part1(b2))
}

func part1(s []byte) int {
	score := 0
	for _, line := range bytes.Split(s, []byte("\n")) {
		score += getScore(line)
	}
	return score
}

func getScore(s []byte) int {
	var left, right byte
	search := []byte("0123456789")
	words := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	l := len(s)
	for i, r := range s {
		if lo.Contains(search, r) {
			if left == 0 {
				left = r
			}
			right = r
			continue
		}

		for k, v := range words {
			if r == k[0] {
				if i+len(k)-1 > l-1 {
					continue
				}

				if !bytes.Equal(s[i:i+len(k)], []byte(k)) {
					continue
				}

				if left == 0 {
					left = v
				}
				right = v
				break
			}

		}

	}

	// for _, r := range lo.Reverse(s) {
	// 	if lo.Contains(search, r) {
	// 		right = r
	// 		break
	// 	}
	// }

	i, _ := strconv.Atoi(string(left) + string(right))

	return i
}
