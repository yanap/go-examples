## テスティング入門

go test

go test -v
-v フラグは、パッケージ内の個々のテストの名前と実行時間を表示します

go test -v -run TestSum

# Sum で終わるテストを実行する

go test -v -run "Sum$"

go test-v -run="French|Canal"
-run フラグは関数名が指定されたパターンに一致するテストだけを go test に実行させます

テストはバグの存在を示すが、バグが存在しないことは示さない

カバレッジ
go tool cover

概要だけが必要
go test -cover

Banchmark

\*testing.B

go test -bench=.

## ベンチマーク入門

func BenchmarkXXX(b \*testing.B) {...}

+=
bytes.Buffer

go test -bench .


