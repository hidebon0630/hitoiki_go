package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type MathData struct {
	A  int
	B  int
	as int
}

func filecheck(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func cal(filename string) {
	if filecheck(filename) {
		fmt.Println("file is OK")
		err := os.Remove(filename)
		fmt.Println(err)
	} else {
		fmt.Println("file is NG")
	}
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	// 長さが 10 と決まっているので make に変更
	// cals := []MathData{}
	cals := make([]MathData, 0, 10)

	// 問題の作成。10問の掛け算を作成

	for i := 0; i < 10; i++ {
		firstA := rand.Intn(100)
		firstB := rand.Intn(100)
		calcal := firstA * firstB

		answer := fmt.Sprintf("%v x %v = %v \n", firstA, firstB, calcal)

		_, err := f.Write([]byte(answer))
		if err != nil {
			log.Fatal(err)
		}
		cals = append(cals, MathData{A: firstA, B: firstB, as: calcal})
	}
	defer f.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	filename := "text.txt"
	cal(filename)
}
