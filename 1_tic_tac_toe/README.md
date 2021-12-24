# Tic Tac Toe
はじめてのGo言語学習のため、三目並べを題材にする。

## 要件
- 通常の三目並べ同等の機能を有すること

```shell
$ go run main.go
--------------------------------------

           [1 2 3]
           [4 5 6]
           [7 8 9]

input the number: 2

           [- ○ -]
           [- - -]
           [- - -]


           [1 2 3]
           [4 5 6]
           [7 8 9]

input the number: 5

           [- ○ -]
           [- × -]
           [- - -]


           [1 2 3]
           [4 5 6]
           [7 8 9]

input the number:
```

## 学習内容
### 変数
#### 変数宣言
```go
// var 変数名 型
var num int
var title string
var isOpen bool
```

#### 初期化と暗黙的な型宣言
初期化子が与えられている場合、型を省略できる
```go
var isOpen = true
var num = 1
var title = "Hello world!"
```

型を省略することもできるが、関数内でしか使えない
```go
isOpen := true
num := 1
title := "Hello world!"
```

### 定数
```go
const YEAR = 2021
```

#### スライス
```go
// 初期化なしの宣言
var arr []int

// 初期化
var arr []int{1, 2, 3}

// 初期化による暗黙的型宣言
var arr = []int{1, 2, 3}
arr := []int{}
arr := []int{1, 2, 3}

// 多重スライス
var s [][]int
s := [][]int{
    {10, 20, 30},
    {10, 20, 30},
}
```

#### make を使った初期化
```go
// make([型/type], [長さ/len], [容量/cap])
s := make([]int, 0, 10)
fmt.Printf("%#v, len: %d, cap: %d\n", s, len(s), cap(s)) // => []int{}, len: 0, cap: 10

// 長さに 1 以上の数値を指定すると、その分がゼロ値(int型の場合は0)で初期化される
s := make([]int, 3, 10)
fmt.Printf("%#v, len: %d, cap: %d\n", s, len(s), cap(s))
// []int{0, 0, 0}, len: 3, cap: 10
```

### 構造体
```go
// 宣言
type Book struct {
	title            string
	description      string
	author           string
	publication_year int
}

func main() {
    // インスタンス生成
    book := new(Book)

    // 代入
    book.title = "hoge"
    book.author = "hogehoge"
}
```

## 参考サイト
[配列とスライス](https://golang.hateblo.jp/entry/golang-slice-array)
