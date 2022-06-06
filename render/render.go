package render

import (
	"fmt"
	"time"

	"github.com/ZayaSansanch/Game/data"
	"github.com/ZayaSansanch/Game/render/funcs"
)

const (
	fpsWanted            = 60
	clear                = "\033[H\033[2J"
	sc                   = "\u001B7"
	rc                   = "\u001B8"
	colorRed             = "\u001b[31m"
	colorBrightRed       = "\u001b[31;1m"
	colorBrightYellow    = "\u001b[33;1m"
	colorBackgroundWhite = "\u001b[47m"
	colorReset           = "\u001b[0m"
)

func fpsrender() {
	var (
		frame = 1000000 / fpsWanted * time.Microsecond
		fps   = 1000000 / frame.Microseconds()

		fpsCounter = 0
		fpsTimer   time.Duration
	)

	// for {
	start := time.Now()

	// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	fpsCounter++

	fmt.Print("\033[H\033[2J")
	UpDate()
	funcs.PrintMap(data.GameMMap)
	funcs.PrintMap(data.GameMap)
	// fmt.Println(colorBrightRed + colorBackgroundWhite + "Time: " + start.String() + colorReset)
	// fmt.Println(colorBrightRed + "Time: " + start.String() + colorReset)
	workInterval := time.Since(start)

	fpsTimer += workInterval
	if fpsTimer >= time.Second {
		fps = int64(fpsCounter)
		fpsCounter = 0
		fpsTimer = 0
	}

	fmt.Println(workInterval, "fps: ", fps)
	if workInterval > 40*time.Millisecond {
		fmt.Print(colorBrightYellow + "Warning: time > 40ms, time: " + workInterval.String() + colorReset + "\n")
	}

	if workInterval < frame {
		time.Sleep(frame - workInterval)
	}

	// data.Fps = 1000 / 40

	// time.Sleep(40*time.Millisecond - workInterval)
	// }

	// funcs.MPrintMMap(data.GameMMap)
	// funcs.MPrintMap(data.GameMap)
}

func Render() {
	// fmt.Println("Render - start")
	data.Height, data.Width = funcs.Hw()

	fpsrender()
}
