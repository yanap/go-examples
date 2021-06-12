package calc

import "testing"

// TestSumは加算のテストする
// 引数には*testing.Tを渡す
// 必ずTestから始まる名前にする
// すると、go testでの実行対象になる

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		// t.Fatalはテストが失敗したことを返すAPI
		// 多くのGoのテストコードでは条件分岐とt.Fatalを組み合わせて書くことになる
		// t.Fatal以外にも、t.Fatalfもある
		// これらはテスト失敗時のエラーメッセージを加工するもの
		// 別の例で詳しく見ていく
		t.Fatal("sum(1, 2) should be 3, but doesn't match")
	}
}
