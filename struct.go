package main

// 大文字開始ではないField名は、外部パッケージからアクセスできない
// 同じパッケージなら問題なし
type Person struct {
	name        string
	age, height int
}
