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

	logrus.Infof("Score: %d", part1(b))
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
	for _, r := range s {
		if lo.Contains(search, r) {
			left = r
			break
		}
	}

	for _, r := range lo.Reverse(s) {
		if lo.Contains(search, r) {
			right = r
			break
		}
	}

	i,_ := strconv.Atoi(string(left) + string(right))

	return i
}

