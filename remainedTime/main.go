package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//コマンド引数の入力が二つ必要な為、なければ通知する
	if len(os.Args) < 3 {
		println("\nOops! Sorry!\n年/月/日と、半角スペース、予定or目標を再入力してください。\nex) 2045/1/1 シンギュラリティ")
		return
	}
	//コマンド引数の入力の仕方がerrorの場合
	pT, err := time.Parse("2006/1/2", os.Args[1])
	if err != nil {
		//閏年のメッセージ処理
		ost := os.Args[1]
		if strings.Contains(ost, "2/29") {
			fmt.Println("\nOops! Sorry!\n閏年ではない為2/29日は存在しません。\n年/月/日, 半角スペース, 予定や目標の順に再入力してください。\nex) 2045/1/1 プログラマになる")
			return
		}
		//それ以外のerror処理
		println("\nOops! Sorry!\n年/月/日, 半角スペース, 予定や目標の順に再入力してください。\nex) 2045/1/1 シンギュラリティ")
		return
	}

	jst := pT.Local()
	nT := time.Now()
	//jstからnTを引いた値が-の場合は過去
	dur := jst.Sub(nT)
	//time.ParseはUTCで、pT.time.local()する時に、JST変換の時差分９時間が足される為-9でJSTの0時に調整
	hrs := int(dur.Hours()) - 9
	//計算後に過去の場合の処理
	if hrs < 0 {
		fmt.Println("\n\n\n今を生きていない時ほど過去を振り返ってしまう...\n\n")
		time.Sleep(2 * time.Second)
		fmt.Println("\n年/月/日, 未来の予定や目標を再入力してください。\nex) 2045/1/1 シンギュラリティ")
		return
	}

	//1年＝8760
	years := hrs / 8760
	//hrsがyearに繰り上がった場合に、yearの時間分をhrsから引く処理
	func() int {
		if years > 0 {
			hrs = hrs - (8760 * years)
		}
		return hrs
	}()
	//残された時間を使って日付と時間を計算
	days := hrs / 24
	hours := hrs % 24
	//分と秒は影響がない為引く作業をせず、durから直接生成
	mins := int(dur.Minutes()) % 60
	secs := int(dur.Seconds()) % 60

	gl := os.Args[2]
	fmt.Printf("\n\n\n%sまでに貴方に残された時間は...\n\n", gl)
	time.Sleep(3 * time.Second)
	if years == 0 {
		fmt.Printf("\n%d日と%d時間%d分%d秒です。\n\n\n", days, hours, mins, secs)
	} else if years >= 1 {
		fmt.Printf("\n%d年と%d日、そして%d時間%d分%d秒です。\n\n\n", years, days, hours, mins, secs)
	} else {
		fmt.Println("無効な入力です")
	}
}

func leapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
