package bisect_test

import (
	"testing"

	"example.com/go-websocket-test/bisect"
)

// keyの値と、そのときに期待する出力のセットを
// 構造体のリストにしてまとめておく
var tests = []struct {
	title    string
	key, ans int
}{
	{title: "overLeft", key: -1, ans: 0},
	{title: "stopEven", key: 2, ans: 1},
	{title: "stopOdd", key: 3, ans: 2},
	{title: "stopBlock", key: 4, ans: 3},
	{title: "overRight", key: 9, ans: 9},
	{title: "notExist", key: 6, ans: 7},
}

func TestTableDrivien(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}

	// 構造体リストで表現したテストケースを
	// for分を使って回していく
	for _, tt := range tests {

		// t.Run()でサブテストを作る
		t.Run(tt.title, func(t *testing.T) {
			if got := bisect.BisectLeft(a, tt.key); got != tt.ans {
				t.Errorf("%s get wrong ans %d: want %d", tt.title, got, tt.ans)
			}
		})
	}
}
