package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/Workiva/go-datastructures/threadsafe/err"
)

type MathData struct {
	A  int
	B  int
	as int
}

var listTmpl = template.Must(template.New("index.html").ParseFiles("src/mihohoi/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	filename := "text.txt"
	if err != listTmpl.Execute(w, cal(filename)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func filecheck(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func cal(filename string) []MathData {
	if err := os.Remove(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file is NG")
	}
	f, err := os.Create(filename)
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
	return cals
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
