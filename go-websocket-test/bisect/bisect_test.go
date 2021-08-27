package bisect_test

import (
	"testing"

	"example.com/go-websocket-test/bisect"
)

// 1より小さいkeyのとき、index0番にちゃんと入るか
func TestOver_left(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := -1

	ans := bisect.BisectLeft(a, key)
	if ans != 0 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

// 8より大きいkeyのとき、index9番にちゃんと入るか
func TestOver_right(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 9

	ans := bisect.BisectLeft(a, key)
	if ans != 9 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

func TestNot_exist(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 6

	ans := bisect.BisectLeft(a, key)
	if ans != 7 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

// 以下keyの値だけが違うテストがたくさんある

func TestStop_block(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 4

	ans := bisect.BisectLeft(a, key)
	// 期待する出力とは違うansが得られたら
	// t.Errorfメソッドでテストが失敗するようになる
	if ans != 3 {
		t.Errorf("get wrong ans: %d", ans)
	}
}
