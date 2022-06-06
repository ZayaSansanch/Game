package render

import (
	"fmt"
	// "time"

	"github.com/ZayaSansanch/Game/data"
)

func UpDate() {
	var cam data.Camera
	var myLen int = 0
	// for i := 0; i < len(data.GameMMap2); i++ {
	// 	for j := 0; j < len(data.GameMMap2[i]); j++ {
	// 		if data.GameMMap2[i][j] == "C" {
	// 			cam.H = i
	// 			cam.W = j
	// 		}
	// 	}
	// }

	for i := 1; i < len(data.GameMMap2)-1; i++ {
		for j := 1; j < len(data.GameMMap2[i])-1; j++ {
			for ci := 1; ci < i; ci++ {
				if data.GameMMap2[i-ci][j] == "@" {
					myLen = ci
				}
			}
		}
	}

	fmt.Println(myLen)

	for i := 1; i < len(data.GameMMap2)-1; i++ {
		for j := 1; j < len(data.GameMMap2[i])-1; j++ {
			if data.GameMMap2[i][j] == "C" {
				cam.H = i
				cam.W = j
				for megaj := 0; megaj < len(data.GameMMap2); megaj++ {
					for ii := 0; ii < len(data.GameMap); ii++ {
						for jj := 0; jj < len(data.GameMap[ii]); jj++ {
							if 
						}
					}
				}
			}
		}
	}

	// for i := 1; i < len(data.GameMMap2)-1; i++ {
	// 	for j := 1; j < len(data.GameMMap2[i])-1; j++ {
	// 		if data.GameMMap2[i][j] == "C" {
	// 			cam.H = i
	// 			cam.W = j
	// 				for ii := 0; ii < len(data.GameMap); ii++ {
	// 					for jj := 0; jj < len(data.GameMap[ii]); jj++ {
	// 						if ii > myLen && ii < (len(data.GameMap)-myLen) {
	// 							fmt.Println(myLen)
	// 							data.GameMap[ii][jj] = "@"
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// var ci, cj, cii int = 0
	// var while, while2 bool = true
	// for i := 1; i < len(data.GameMMap2)-1; i++ {
	// 	for j := 1; j < len(data.GameMMap2[i])-1; j++ {
	// 		if data.GameMMap2[i][j] == "C" {
	// 			cam.H = i
	// 			cam.W = j
	// 			// for cj := 0; cj < j; cj++ {
	// 			for ci := 1; ci < i; ci++ {
	// 				if data.GameMMap2[i-ci][j] != "@" {
	// 					myLen = ci
	// 					for ii := 0; ii < len(data.GameMap); ii++ {
	// 						for jj := 0; jj < len(data.GameMap[ii]); jj++ {
	// 							if ii > myLen && ii < (len(data.GameMap)-myLen) {
	// 								fmt.Println(myLen)
	// 								data.GameMap[ii][jj] = "@"
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 			// }
	// 		}

	// 		// for ii := 0; ii < len(data.GameMap); ii++ {
	// 		// 	for jj := 0; jj < len(data.GameMap[i]); jj++ {
	// 		// 		// for while {
	// 		// 		// 	if data.GameMMap2[i+ci][j+cj] == "@" {

	// 		// 		// 	}
	// 		// 		// }
	// 		// 	}
	// 		// }
	// 	}
	// }

	fmt.Println(cam.H, cam.W)

	// time.Sleep(40 * time.Millisecond)
}
