package main

import (
	"fmt"
)

// 変数宣言
func Var() {
	// 整数型の変数
	var num1 int = 123

	// 右辺で型が決まるなら型名は省略可能
	var num2 = 456

	// 省略記法（関数内でのみ利用可能）
	num3 := 789

	// 変数の初期化のみ（各データ型のゼロ値で自動初期化される）
	// 初期化のみの変数は、使用しないとエラーになる。
	var s string

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(s) //コメント化するとエラーになる。

}
