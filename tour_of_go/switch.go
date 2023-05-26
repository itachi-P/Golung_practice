package main

import (
	"fmt"
	"runtime"
)

//switchステートメント

/*
GolangのswitchはCやC++、Java、JavaScript、PHPと違い、
選択されたcaseのみ実行され、それに続くcaseは実行されない。
これらの言語では各caseの最後に記述が必要なbreakが、Goでは自動的に入る。
（逆に意図的に次のcaseに続けたい場合、 fallthrough キーワードを記述）
また、Goのswitch文のcaseの値は定数・整数値である必要がない。
*/

func main() {
	fmt.Print("Go runs on ")

	//if文と同じく、条件式の前に簡易文を設定可能
	//runtime.GOOSの値は実行環境による
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, solaris,
		// windows, android, ios, etc.
		fmt.Printf("%s.\n", os)
	}

	//逆にswitchのみで各caseにif文のような条件式を書くことも可能
	n := 100
	switch {
	case n == 100:
		fmt.Print("Excellent!! ")
		//fallthroughは次のcaseと共通しない特別な前処理などに使える
		fallthrough
	case n >= 60 && n <= 99:
		fmt.Println(n, "点：合格です")
	case n > 0 && n < 60:
		fmt.Println(n, "点：頑張りましょう")
	default:
		fmt.Println("Oh...")
	}

	//switch文を使った型による処理分け
	v := interface{}("3.14159")

	switch v.(type) {
	case string:
		//修飾子%qだと数字や空文字の場合も""付きでわかりやすい
		fmt.Printf("%q is string.", v.(string))
	case int, int32, int64:
		fmt.Println(v.(int) * 256)
	case bool:
		fmt.Println("Lies or Truth.")
	default:
		fmt.Printf("Type: %T", v)
	}
}
