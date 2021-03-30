# FizzBuzzでテスト駆動開発を学ぶ
このドキュメントでは、FizzBuzz問題風の関数作成を通してテスト駆動開発を学ぶ。
テスト駆動開発では

> RED(テストに失敗する状態) > Green(テストは通る状態) > Refactoring(テストが通り、機能も意味を持つ状態)

のサイクルを繰り返し行うことで、プログラムの品質を向上させる開発モデルの1つである。

受入テストを先にするとか、結合テストを先にするとかではない。あくまでもプログラムユニットテストにおいての開発手法である。

# 実際にやってみる
テスト駆動開発では以下の手順でテストを作成する。

1. TODOListを作成する。
2. テストコードを書き始める。
3. 2で作ったテストコードが通るコードを書く
4. 3で作ったテストコードをリファクタリングする。

実際にGolangを用いて実施する。
(GoはDockerかChocolatyを使うといい)
> choco install -y golang

> docker run go

## TODOListを作成する。
今回は以下の関数3つを実装する。
- 渡された引数2つの範囲で3で割り切れる回数を返す
- 渡された引数2つの範囲で5で割り切れる回数を返す
- 渡された引数2つの範囲で3と5で割り切れる回数を返す

## テストコードを書き始める。

この場合関数が満たす必要がある条件は

- 数値型の引数を2つ受け取る
- 割り切れる回数を返却する

必要がある。
これをテストコードとして書く。

```go 
func TestFizzCount(t *testing.T) {
	start, end := 1, 10
	if fizzCount(start, end) != 3 {
		t.Errorf("fizzCount() = %v, want %v", fizzCount(1, 10), 3)
	}
}
```

このテストコードに対応する関数を作成する。

```go 
func fizzCount(start, end int) int {
	return 0;
}
```

0と10を```fizzCount```関数に渡すと、3で割り切れる数は**3回**登場するので期待結果は「3」が返却される。

## テストコードが通るコードを書く

現時点では```fizzCount```関数は問答無用で**0**を返却しているので、当然テストは失敗する。

```
> go test fizzBuzz_test.go
--- FAIL: TestFizzCount (0.00s)
    fizzBuzz_test.go:16: fizzCount() = 0, want 3
```

期待結果が3なので、今度は3を返すように修正する。

```go 
func fizzCount(start, end int) int {
	fizzCount := 3

	return fizzCount
}
```

これならテストは通るが、まったく意味の無い関数となってしまっている。

```
go test fizzBuzz_test.go
=== RUN   TestFizzCount
--- PASS: TestFizzCount (0.00s)
PASS
```

## テストコードをリファクタリングする。
ここからはGreenとなったテスト結果を壊さないようにリファクタリングしていく。
現在実装できていない機能は

- 3で割り切れる回数をカウントする

であるので、引数の範囲ループする処理と条件分岐を書く。

```go
func fizzCount(start, end int) int {
	fizzCount := 0
	for i := start; start < end; i++ {
		if i%3 == 0 {
			fizzCount++
		}
	}
	return fizzCount
}
```

結果は同じくグリーンのまま。とりあえずこれで機能を満たしつつテストも通る状態となった。

```
go test fizzBuzz_test.go
=== RUN   TestFizzCount
--- PASS: TestFizzCount (0.00s)
PASS
```

同じように5で割り切れる場合と3と5で割り切れる場合のテストケースを作成する。

```go
func TestBuzzCount(t *testing.T) {
	start, end := 1, 10
	if buzzCount(start, end) != 2 {
		t.Errorf("buzzCount() = %v, want %v", buzzCount(1, 10), 3)
	}
}

func TestFizzBuzzCount(t *testing.T) {
	start, end := 1, 10
	if fizzBuzzCount(start, end) != 0 {
		t.Errorf("fizzBuzzCount() = %v, want %v", fizzBuzzCount(1, 10), 3)
	}
}
```

同じように関数を実装する。

```go

func buzzCount(start, end int) int {
	buzzCount := 0
	for i := start; start < end; i++ {
		if i%3 == 0 {
			buzzCount++
		}
	}
	return buzzCount
}

func fizzBuzzCount(start, end int) int {
	fizzBuzzCount := 0
	for i := start; start < end; i++ {
		if i%3 == 0 {
			fizzBuzzCount++
		}
	}
	return fizzBuzzCount
}

```

## テストケースを減らす
これらの関数を実装したことで、いくつか条件のみが違うテストケースが作成されてしまっている。

この場合は、3/5/3 or 5で割り切れるかどうかという部分しか違わないため、1つの関数に集約してしまえそう。

まずはテストケースを1つ変更し、集約できそうかを試してみる。

```go
func TestFizzBuzzCount(t *testing.T) {
	start, end := 1, 10
	fizzResult := fizzBuzzCount(start, end)["fizz"] != 3
	buzzResult := fizzBuzzCount(start, end)["buzz"] != 2
	fizzBuzzResult := fizzBuzzCount(start, end)["fizzBuzz"] != 0
	asertMap := map[string]int{"fizz": 3, "buzz": 2, "fizzBuzz": 0}

	if fizzResult && buzzResult && fizzBuzzResult {
		t.Errorf("fizzBuzzCount() = %v, want %v", fizzBuzzCount(1, 10), asertMap)
	}

}
```

Mapで値を返却するように修正し```TestFizzCount()```、```TestBuzzCount()```を削除した。
当然通らないので、とりあえずGreenにする。

関数を合わせて修正する。
```go
func fizzBuzzCount(start, end int) map[string]int {
	resultHash := map[string]int{"fizz": 3, "buzz": 2, "fizzBuzz": 0}
	fizzCount, buzzCount, fizzBuzzCount := 0, 0, 0
	return resultHash
}

```

これでテスト結果はGreenになるので、リファクタリングする。

```
go test -v
=== RUN   TestFizzBuzzCount
--- PASS: TestFizzBuzzCount (0.00s)
```

テストケースを削除しているので、```fizzCount```,```buzzCount```を合わせて削除し```fizzBuzzCount```に集約する。

```go
func fizzBuzzCount(start, end int) map[string]int {
	resultHash := map[string]int{}
	fizzCount, buzzCount, fizzBuzzCount := 0, 0, 0
	for ; start < end; start++ {
		if start%3 == 0 && start%5 == 0 {
			fizzBuzzCount++
		} else if start%3 == 0 {
			fizzCount++
		} else if start%5 == 0 {
			buzzCount++
		}
	}
	return resultHash
}

```

これでもgreenのままなので、関数を作成完了とする。

**今回のように、実実装に入る前にコードの重複に気づくことができることもテスト駆動開発のメリットの1つ。**