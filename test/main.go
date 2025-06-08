package main

//Copilot キーボードを押して離したとき

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer screen.Fini()

	// キーの状態を記録するマップ
	pressedKeys := make(map[rune]bool)

	// イベント処理ループ
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				// Escape または Ctrl+C で終了
				return
			}

			// キー押下時の処理
			if ev.Key() == tcell.KeyRune {
				if !pressedKeys[ev.Rune()] {
					log.Printf("キーが押されました: %c", ev.Rune())
				}
				pressedKeys[ev.Rune()] = true
			}

			// キーリリースを擬似的に検出
			if ev.Modifiers()&tcell.ModNone == tcell.ModNone {
				if pressedKeys[ev.Rune()] {
					log.Printf("キーが離されました: %c", ev.Rune())
					pressedKeys[ev.Rune()] = false
				}
			}

			// switch ev.Key() {
			// case tcell.KeyUp:
			// keyStates[tcell.KeyUp] = true
			// case tcell.KeyEnter:
			// keyStates[tcell.KeyEnter] = true
			// case tcell.KeyCtrlC, tcell.KeyEscape:
			// return
			// }

			// 同時押しをチェック
			// if keyStates[tcell.KeyUp] && keyStates[tcell.KeyEnter] {
			// log.Println("Up and Enter keys were pressed together!")
			// keyStates[tcell.KeyUp] = false // 状態をリセット
			// keyStates[tcell.KeyEnter] = false
			// }

		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

//
