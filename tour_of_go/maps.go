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
	//カリフォルニア州Google本社（愛称Googleplex）所在地
	//10^10^100を意味するグーゴルプレックス（googolplex）に由来
	//ちなみに1グーゴル=10^100で1の後に0が100個続く101桁の数値m
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)

	//Mapへの要素追加（と上書き）
	m["Tsutenkaku"] = Vertex{34.390913, 135.302277}
	m["Bell Labs"] = Vertex{40.6961793, -73.452068}
	fmt.Println(m)

	//要素の存在確認
	elem, ok := m["Tsutenkaku"]
	if ok { // if ok == trueの省略形
		//MapにKeyが存在しなければelemは要素の型のゼロ値になる
		fmt.Printf("%T %E\n", elem, elem)
	}

	//要素の削除
	delete(m, "Tsutenkaku")
	v, ok := m["Tsutenkaku"]
	fmt.Println("Value:", v, "Presence?:", ok)
	fmt.Println(m)
}
