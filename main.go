package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

/*
以下main関数では、GO言語の機能別の挙動確認関数を定義している
それぞれ、コメントを外して出力結果を確認する。
*/
func main() {
	// fmt.Println("Hello World")

	// リテラル
	// Lit()

	// 変数初期化
	// Var()

	// シャドーイング（スコープ）
	//Shadow()

	// データ型
	// DataType()

	// 配列・スライス
	// Array()

	// マップ（辞書）
	// Map()

	// 制御構文
	// Seigyo()

	// 関数
	// Allfunc()

	// 構造体
	// インスタンス化①（fieldは全てゼロ値）
	var akiyama Person
	fmt.Println(akiyama) // { 0 0}

	// インスタンス化②（field順に全て初期化しないとエラー）
	Baba := Person{"Baba", 20, 173}
	fmt.Println(Baba) // {Baba 20 173}

	// インスタンス化③（fieldを指定して初期化）
	yamamoto := Person{age: 22, name: "yamamoto", height: 176}
	fmt.Println(yamamoto) // {yamamoto 22 176}

	// 値の代入
	akiyama.age = 21
	akiyama.name = "akiyama"
	akiyama.height = 175
	fmt.Println(akiyama) // {akiyama 21 175}

	// JSON→構造体のfieldに代入
	// タグ：``（JSONタグに指定したキーをもとにfieldに代入）
	type Book struct {
		Title      string    `json:"title"`
		Author     string    `json:"author"`
		Publisher  string    `json:"publisher"`
		ReleasedAt time.Time `json:"released_at"`
	}

	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error : ", err)
	}
	d := json.NewDecoder(f)
	var b Book
	d.Decode(&b)
	fmt.Println(b.Title)      // クリエイターズ・ファイル
	fmt.Println(b.Author)     // 秋山 竜次
	fmt.Println(b.Publisher)  // ワニブックス
	fmt.Println(b.ReleasedAt) // 2019-10-25 00:00:00 +0000 UTC
}
