package main

import "fmt"

//Map: KeyとValueの関連付け
//Structに似るが、Keyが必要。

type Vertex struct {
	Lat, Long float64 //一般的な「緯度・経度」を表す変数名
}

var m = map[string]Vertex{
	//米国のベル通信研究所（現在はNokiaの子会社）の座標（北緯＆西経）
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
