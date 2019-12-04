
go test

go test -v
-vフラグは、パッケージ内の個々のテストの名前と実行時間を表示します

go test-v -run="French|Canal"
-runフラグは関数名が指定されたパターンに一致するテストだけをgo testに実行させます

テストはバグの存在を示すが、バグが存在しないことは示さない

カバレッジ
go tool cover 

概要だけが必要
go test -cover

Banchmark

*testing.B

go test -bench=.

