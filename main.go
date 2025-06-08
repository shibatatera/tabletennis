package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Racket struct {
	x  int
	y  int
	dy int
	// cpu 右側をcpu
	is_npc   bool
	npclvl   int
	npcspeed int
	//npctick  int //毎tick 1 加算
}

type BallSpeed struct {
	x int
	y int
}
type Ball struct {
	x      int
	y      int
	mx     string //飛んでいく向き
	my     string
	hansya int //ラケット
	speed  BallSpeed
	//tickspeed int //毎tick 1 加算
	stop int // 0 になるまで停止
	gra  rune
}

type Diamond struct {
	x   int
	y   int
	gra rune
}

type Sikaku struct {
	x  int
	y  int
	dx int
	dy int
}

func handleKeyEvent(s tcell.Screen, ch chan<- int) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			// w,s キーと矢印矢印キー両方で操作できるように変更した

			case tcell.KeyEscape: // Esc キーで終了
				ch <- 1
			case tcell.KeyUp:
				//ch <- 4
				ch <- 2
				//time.Sleep(100 * time.Millisecond)
				//racket.y--
			case tcell.KeyDown:
				//ch <- 5
				ch <- 3
				//time.Sleep(100 * time.Millisecond)
				//racket.y++
			default:
				switch ev.Rune() {
				case 'w':
					ch <- 2
					//time.Sleep(100 * time.Millisecond)
					//racket.y--
				case 's':
					ch <- 3
					//time.Sleep(100 * time.Millisecond)
					//racket.y++
				}
			}
		}
	}
}

// minからmaxのランダムな数値
func randm(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min+1) + min
	return r
}

// 呼び出すとボールのスピード変更関数
// func ballSpeedRand(ball *Ball) {
// ball.speed.x = randm(5, 15)
// }

// 数字一桁ずつ gemini
func numToDigits(n int) []int {
	num := n
	s := strconv.Itoa(num) // 数値を文字列に変換

	var digits []int
	for _, r := range s {
		// 各文字を数値に変換
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			// エラーハンドリング（ここでは単純にスキップまたはログ出力など）
			fmt.Printf("Error converting character '%c' to int: %v\n", r, err)
			continue
		}
		digits = append(digits, digit)
	}

	return digits
}

// var pos_x, pos_y = 5, 5
//var keyuppress = false

type GameState int

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	//構造体
	ball := Ball{
		x: 40, y: 15, mx: "left", my: "up", hansya: 0,
		speed: BallSpeed{x: 8, y: 8}, stop: 120, gra: '●'}
	//diamond := Diamond{x: 30, y: 10, gra: tcell.RuneDiamond}
	sikaku := Sikaku{x: 0, y: 0, dx: 80, dy: 25}
	racket := []Racket{
		{x: 10, y: 3, dy: 4, is_npc: false},                                     //player
		{x: sikaku.dx - 10, y: 3, dy: 4, is_npc: true, npclvl: 0, npcspeed: 12}, //npc
	}

	nowtick := 0

	ch := make(chan int)
	go handleKeyEvent(screen, ch)

	//ticker := time.NewTicker(1 * time.Second)
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	screen_x, screen_y := screen.Size()

	//ゲーム全体
	const (
		Menu GameState = iota
		GameClear
		GameOver
	)
	var game GameState = GameClear

	//testaaa := 0
	for {
		select {
		case <-ticker.C:
			screen.Clear()
			//描画----------------------------------------------------
			//ダイアモンド
			//screen.SetContent(diamond.x, diamond.y, diamond.gra, nil, tcell.StyleDefault)

			//unicode 48 ~ 57 が 0 ~ 9
			//テスト用
			// if nowtick%40 == 0 {
			// testaaa = randm(5, 9)
			// }

			drawString := func(x, y int, str string) { // 文字列描画
				for i, r := range str {
					screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
				}
			}

			drawInt := func(x, y, num int, color ...tcell.Style) { // 数字を描画
				digits := numToDigits(num)

				for i := 0; i < len(digits); i++ {
					if len(color) == 0 {
						screen.SetContent(x+i, y, rune('0'+digits[i]), nil, tcell.StyleDefault)
					} else {
						screen.SetContent(x+i, y, rune('0'+digits[i]), nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed))
					}
					//screen.SetContent(x+i, y, rune('0'+digits[i]), nil, tcell.StyleDefault)
					//screen.SetContent(x+i, y, rune('0'+digits[i]), nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed))
				}
			}

			drawStringAndInt := func(x, y, num int, str string) {
				drawString(x, y, str)
				drawInt(x+len([]rune(str)), y, num)
			}

			enummitainayatu := func(n GameState) int {
				switch n {
				case Menu:
					return 0
				case GameOver:
					return 1
				default:
					return 0
				}
			}

			siky := sikaku.y + sikaku.dy

			drawInt(1, siky+6, enummitainayatu(Menu))
			switch game {
			case Menu:
				drawString(5, siky+8, "hoge")
			case GameClear:
				drawString(sikaku.x+sikaku.dx/2-5, sikaku.y+sikaku.dy/2, "Game Clear")
			case GameOver:

				drawString(5, siky+8, "Game Over")
			default:
				drawString(5, siky+8, "default")
			}

			// ボールのspeedが一桁ずつ int[] で帰る
			// drawInt(1, sikaku.y+sikaku.dy+1, ball.speed)
			drawStringAndInt(1, siky+1, ball.speed.x, "ball.speed.x: ")
			drawStringAndInt(1, siky+2, ball.speed.y, "ball.speed.y: ")
			//screen.SetContent(1, sikaku.y+sikaku.dy+1, rune('0'+1), nil, tcell.StyleDefault)

			//ボール
			screen.SetContent(ball.x, ball.y, ball.gra, nil, tcell.StyleDefault)

			//大きな四角形
			// ┌
			screen.SetContent(sikaku.x, sikaku.y, tcell.RuneULCorner, nil, tcell.StyleDefault)
			// └
			screen.SetContent(sikaku.x, sikaku.y+sikaku.dy, tcell.RuneLLCorner, nil, tcell.StyleDefault)
			// ─
			for col := sikaku.x + 1; col < sikaku.x+sikaku.dx; col++ {
				screen.SetContent(col, sikaku.y, tcell.RuneHLine, nil, tcell.StyleDefault)
			}
			// │
			for col := sikaku.y + 1; col < sikaku.y+sikaku.dy; col++ {
				screen.SetContent(sikaku.x, col, tcell.RuneVLine, nil, tcell.StyleDefault)
			}
			for col := sikaku.y + 1; col < sikaku.y+sikaku.dy; col++ {
				screen.SetContent(sikaku.x+sikaku.dx, col, tcell.RuneVLine, nil, tcell.StyleDefault)
			}
			// ─
			for col := sikaku.x + 1; col < sikaku.x+sikaku.dx; col++ {
				screen.SetContent(col, sikaku.y+sikaku.dy, tcell.RuneHLine, nil, tcell.StyleDefault)
			}
			// ┐
			screen.SetContent(sikaku.x+sikaku.dx, sikaku.y, tcell.RuneURCorner, nil, tcell.StyleDefault)
			// ┘
			screen.SetContent(sikaku.x+sikaku.dx, sikaku.y+sikaku.dy, tcell.RuneLRCorner, nil, tcell.StyleDefault)

			//ラケット
			for i := 0; i < len(racket); i++ {
				for col := racket[i].y + 1; col < racket[i].y+racket[i].dy; col++ {
					screen.SetContent(racket[i].x, col, tcell.RuneBlock, nil, tcell.StyleDefault)
				}
			}

			screen.Show()

			//処理------------------------------------------------------------
			ballSpeedRand := func(d rune) { // ボールの速度変更
				min, max := 3, 15
				if d == 'x' {
					ball.speed.x = randm(min, max)
				}
				if d == 'y' {
					ball.speed.y = randm(min, max)
				}
			}

			nowtick++
			if nowtick == math.MaxInt32 {
				nowtick = 0
			}

			screen_x, screen_y = screen.Size()
			//ボールを動かす
			// stop <= 0 まで止まっている
			if ball.stop > 0 {
				ball.stop--
			}
			//ball.tickspeed++
			// if ball.tickspeed == math.MaxInt32 {
			// ball.tickspeed = 0
			// }

			if (0 == nowtick%ball.speed.x) && (ball.stop <= 0) { // x 移動
				if ball.mx == "left" {
					ball.x--
					ball.hansya--
				}
				if ball.mx == "right" {
					ball.x++
					ball.hansya--
				}
			}
			if (0 == nowtick%ball.speed.y) && (ball.stop <= 0) { // x 移動
				if ball.my == "up" {
					ball.y--
				}
				if ball.my == "down" {
					ball.y++
				}
			}

			//ボールが四角形の壁に当たると跳ね返る
			//ボールをスクリーン外に出ないようにする
			if ball.x >= screen_x {
				ball.x = screen_x - 1
				ball.mx = "left"
			}
			if ball.x < 0 {
				ball.x = 0
				ball.mx = "right"
			}
			if ball.y >= screen_y {
				ball.y = screen_y - 1
				ball.my = "up"
			}
			if ball.y < 0 {
				ball.y = 0
				ball.my = "down"
			}
			//四角形の中
			if ball.x >= sikaku.dx {
				ball.x = sikaku.dx - 1
				ball.mx = "left"
				ballSpeedRand('x')
			}
			if ball.x <= sikaku.x {
				ball.x = sikaku.x + 1
				ball.mx = "right"
				ballSpeedRand('x')
			}
			if ball.y >= sikaku.dy {
				ball.y = sikaku.dy - 1
				ball.my = "up"
				ballSpeedRand('x')
			}
			if ball.y <= sikaku.y {
				ball.y = sikaku.y + 1
				ball.my = "down"
				ballSpeedRand('x')
			}
			//ボールがラケットに当たる
			for i := 0; i < len(racket); i++ {
				if ball.x == racket[i].x && (ball.y <= racket[i].y+racket[i].dy && ball.y >= racket[i].y) {
					if ball.mx == "left" && ball.hansya <= 0 {
						ball.mx = "right"
						ball.hansya = 5
						ballSpeedRand('x')
						ballSpeedRand('y')
					}
					if ball.mx == "right" && ball.hansya <= 0 {
						ball.mx = "left"
						ball.hansya = 5
						ballSpeedRand('x')
						ballSpeedRand('y')
					}
				}
			}

			//ラケットを四角形の中から出ないようにする
			for i := 0; i < len(racket); i++ {
				if racket[i].y+racket[i].dy >= sikaku.dy {
					racket[i].y = sikaku.dy - racket[i].dy
				}
				if racket[i].y <= sikaku.y {
					racket[i].y = sikaku.y
				}
			}

			// npc |||||||||||||||||||||||||||||||||||||||||
			//とりあえずボール追いかける
			for i := 0; i < len(racket); i++ {
				if racket[i].is_npc {
					//switch
					pos_y := racket[i].y + racket[i].dy/2

					// racket[i].npctick++
					// if racket[i].npctick == math.MaxInt32 {
					// racket[i].npctick = 0
					// }

					if 0 == (nowtick % racket[i].npcspeed) {
						if pos_y > ball.y {
							racket[i].y--
						}
						if pos_y < ball.y {
							racket[i].y++
						}
					}
				}
			}

			//keyuppress = false
		case i := <-ch:
			switch i {
			case 1:
				return
			case 2:
				racket[0].y--
			case 3:
				racket[0].y++
				// case 4:
				// racket[1].y--
				// case 5:
				// racket[1].y++
			}
		}
		/*ev := screen.PollEvent() // イベントを取得する
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() { // 何のキーが押されたか？を調べる
			case tcell.KeyEscape:
				return
			}
		default:
			screen.Clear()
			screen.SetContent(pos_x, pos_y, ball, nil, tcell.StyleDefault)
			screen.Show()

			pos_x += 1
			time.Sleep(1 * time.Second)
		}*/

		//screen.ShowCursor(pos_x, pos_y) //カーソル

	}
}

/*case tcell.KeyUp:
	pos_y--
	if pos_y <= 0 {
		pos_y = 0
	}
case tcell.KeyDown:
	pos_y++
	if pos_y >= screen_y-1 {
		pos_y = screen_y - 1
	}
case tcell.KeyLeft:
	pos_x--
	if pos_x <= 0 {
		pos_x = 0
	}
case tcell.KeyRight:
	pos_x++
	if pos_x >= screen_x-1 {
		pos_x = screen_x - 1
	}*/
//}
/*case *tcell.EventResize:
screen.Sync()
screen_x, screen_y = screen.Size()*/
