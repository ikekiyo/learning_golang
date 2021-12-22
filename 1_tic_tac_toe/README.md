# Tic Tac Toe
はじめてのGo言語学習のため、三目並べを題材にする。

## 基本文法
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

### 定数


## 参考サイト
[配列とスライス](https://golang.hateblo.jp/entry/golang-slice-array)
