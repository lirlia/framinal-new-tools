package main

import (
	"fmt"

	lo "github.com/samber/lo"
)

func main() {


	// map
	// Assign: map の結合
	t1 := map[string]int{"a": 1, "b": 2}
	t2 := map[string]int{"c": 3, "d": 4}
	r := lo.Assign(t1, t2)
	fmt.Println(r) // map[a:1 b:2 c:3 d:4]

	// slice
	// Associate: slice から map を作成
	// 今回の例では slice の要素を key に、要素の長さを value にする
	t3 := []string{"a", "bb", "ccc"}
	r = lo.Associate(t3, func(v string) (string, int) {
		return v, len(v)
	})
	fmt.Println(r) // map[a:1 bb:2 ccc:3]

	// // 非同期
	// // Async: 非同期処理
	// fmt.Println(lo.Async(func() string{
	// 	return "hello"
	// }))

	// retry
	// Attempt 指定した回数までリトライ
	retry := func(r int) (int, error) {
		return lo.Attempt(r, func(idx int) error {
			if idx < 2 {
				return fmt.Errorf("error :%d", idx)
			}
			return nil
		})
	}

	count, err := retry(2)
	fmt.Printf("%v %v\n", count, err) // 2 error :1

	count, err = retry(3)
	fmt.Printf("%v %v\n", count, err) // 3 <nil>

	// AttemptWhile 指定の条件が true かつ回数までリトライ
	retry2 := func(r int) (int, error) {
		return lo.AttemptWhile(r, func(idx int) (error, bool) {
			if idx < 2 {
				return fmt.Errorf("error :%d", idx), true
			}
			return nil, false
		})
	}

	count, err= retry2(2)
	fmt.Printf("%v %v\n", count, err) // 2 error :1

	count, err= retry2(3)
	fmt.Printf("%v %v\n", count, err) // 2 error :1

	// AttemptWithDelay
	// AttemptWhileWithDelay
	// Delay あり、exponential Backoff はない

	// バッチ
	// Batch / BatchWithTimeout は非推奨

	// バッファ
	// Buffer : チャネルからスライスを作成
	// BufferWithTimeout : チャネルからスライスを作成（タイムアウトあり）

	// ChannelDispatcher : 入力チャネルからのメッセージをN個の子チャネルに分配し，クローズイベントは子チャネルに伝搬されます． 基となるチャネルには，固定バッファ容量と，capが0であればバッファリングされないバッファ容量があります．
	// ChannelMerge : 複数の入力チャネルからのメッセージを単一のバッファ付きチャネルに集めます。 出力メッセージには優先順位がありません。すべてのアップストリームチャネルがEOFに達すると，ダウンストリームチャネルは閉じられます． 非推奨．代わりに lo.FanIn を使ってください。
	// ChannelToSlice : チャネルのアイテムから構築されたスライスを返します。チャネルが閉じるまでブロックされます

	// Chunk : スライスを指定したサイズのチャンクに分割します
	t5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(lo.Chunk(t5, 3)) // [[1 2 3] [4 5 6] [7 8 9]]
	fmt.Println(lo.Chunk(t5, 4)) // [[1 2 3 4] [5 6 7 8] [9]]

	// ChunkString: 文字列を指定したサイズのチャンクに分割します
	t6 := "123456789"
	fmt.Println(lo.ChunkString(t6, 3)) // [123 456 789]
	fmt.Println(lo.ChunkString(t6, 4)) // [1234 5678 9]

	// Clamp : 値を min ~ max の範囲に収める
	fmt.Println(lo.Clamp(200, 0, 20)) // 20
	fmt.Println(lo.Clamp(10, 0, 20)) // 10

	// Coalesce : 最初のからではない引数を返す
	fmt.Println(lo.Coalesce(1,2,3)) // 1 true

	// Compact : - ではないスライスの全要素を返す
	fmt.Println(lo.Compact([]int{1,2,0,3,0})) // [1 2 3]

	// Contains : ある要素がスライスにある場合に true を返す
	fmt.Println(lo.Contains([]int{1,2,0,3,0}, 1)) // true
	fmt.Println(lo.Contains([]int{1,2,0,3,0}, 4)) // false

	// ContainsBy : 述語関数が true を返す要素がスライスにある場合に true を返す
	fmt.Println(lo.ContainsBy([]int{1,2,0,3,0}, func(v int) bool { // true
		return v == 1
	}))

	fmt.Println(lo.ContainsBy([]int{1,2,0,3,0}, func(v int) bool { // false
		return v > 100
	}))

	// Count: スライス内で value と同じ値がいくつあるかを返す
	fmt.Println(lo.Count([]int{1,2,0,3,0}, 1)) // 1
	fmt.Println(lo.Count([]int{1,2,0,3,0}, 0)) // 2
	fmt.Println(lo.Count([]int{1,2,0,3,0}, 100)) // 0

	// CountBy : スライス内で述語関数が true を返す要素がいくつあるかを返す
	fmt.Println(lo.CountBy([]int{1,2,0,3,0}, func(v int) bool { // 2
		return v > 1
	}))

	// CountValues : コレクション内の各要素の数を map で返す
	fmt.Println(lo.CountValues([]int{1,2,0,3,0})) // map[0:2 1:1 2:1 3:1]

	// CountValuesBy : コレクション内の各要素の数を map で返す


	a := 10
	fmt.Println(test1(&a)) // 11
	fmt.Printf("%p", test1(&a))
	fmt.Println(a)


}

func test1(p *int) int {
	*p++
	return *p
}
