package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
	otherWord + "hq",
}

func main() {
	generate(os.Stdin, os.Stdout)
}

func generate(r io.Reader, w io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Fprintln(w, strings.Replace(t, otherWord, s.Text(), -1))
	}
}
