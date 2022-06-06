package funcs

import (
	"fmt"
)

func MPrintMap(matrix [25][50]int) {
	for i := 0; i < 25; i++ {
		for j := 0; j < 50; j++ {
			fmt.Print(matrix[i][j])
			// if j == 0 && i != 0 {
			// if j == 0 {
			// 	fmt.Println(matrix[i][j])
			// } else {
			// 	fmt.Print(matrix[i][j])
			// }
		}
		fmt.Print("\n")
	}
	// fmt.Print("0 \n")
	fmt.Print("\n")
}

func PrintMap(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
			// if j == 0 && i != 0 {
			// if j == 0 {
			// 	fmt.Println(matrix[i][j])
			// } else {
			// 	fmt.Print(matrix[i][j])
			// }
		}
		fmt.Print("\n")
	}
	// fmt.Print("0 \n")
	fmt.Print("\n")
}

func MPrintMMap(matrix [10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				fmt.Println(matrix[i][j])
			} else {
				fmt.Print(matrix[i][j])
			}
		}
	}
	fmt.Print("0 \n \n")
}

// fmt.Printf("Width: %d\n", data.Width)
// fmt.Printf("Height: %d\n", data.Height)

// func MPrint(matrix [][]int, h, w int) {
// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
// 			if j == 0 {
// 				fmt.Println(matrix[i][j])
// 			} else {
// 				fmt.Print(matrix[i][j])
// 			}
// 		}
// 	}
// 	fmt.Print("0 \n")
// }
