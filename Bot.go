package main

import (
	"fmt"
	"math"
	"os"
)

const (
	Up	= 1
	Down	= 2
	Left	= 4
	Right	= 8
)

// Bot here goes Hal
type Bot struct {
	row, col int
}

func (my Bot) play(g Game) {
	var oprow, opcol int

	for row := 0; row < g.fieldHeight; row++ {
		for col := 0; col < g.fieldWidth; col++ {
			if g.field.cells[row][col] == g.myBotID {
				my.row = row
				my.col = col
			} else {
				if g.field.cells[row][col] == g.otherBotID {
					oprow = row
                                	opcol = col
				}
			}
		}
	}
	fmt.Fprintln(os.Stderr, "my position: row (y) = %d, col (x) = %d", my.row, my.col)
	fmt.Fprintln(os.Stderr, "opponent position: row (y) = %d, col (x) = %d", oprow, opcol)
	
	moves := my.getValidMoves(g.field)
	cdist := int(math.Abs(float64(my.row - oprow))) + int(math.Abs(float64(my.col - opcol)))
	//taxicab geometry
	switch {
	case moves & Up != 0:
		ndist := int(math.Abs(float64(my.row - 1 - oprow))) + int(math.Abs(float64(my.col - opcol)))
		if ndist < cdist {
			fmt.Println("up")
		}

	case moves & Left != 0:
                ndist := int(math.Abs(float64(my.row - oprow))) + int(math.Abs(float64(my.col - 1 - opcol)))
                if ndist < cdist {
                        fmt.Println("left")
                }

	case moves & Right != 0:
                ndist := int(math.Abs(float64(my.row - oprow))) + int(math.Abs(float64(my.col + 1 - opcol)))
                if ndist < cdist {
                        fmt.Println("right")
                }
 
	case moves & Down != 0:
		ndist := int(math.Abs(float64(my.row + 1 - oprow))) + int(math.Abs(float64(my.col - opcol)))
                if ndist < cdist {
                        fmt.Println("down")
                }
	
	case moves & Up != 0:
                ndist := int(math.Abs(float64(my.row - 1 - oprow))) + int(math.Abs(float64(my.col - opcol)))
                if ndist == cdist {
                        fmt.Println("up")
                }

        case moves & Left != 0:
                ndist := int(math.Abs(float64(my.row - oprow))) + int(math.Abs(float64(my.col - 1 - opcol)))
                if ndist == cdist {
                        fmt.Println("left")
                }
       
        case moves & Right != 0:
                ndist := int(math.Abs(float64(my.row - oprow))) + int(math.Abs(float64(my.col + 1 - opcol)))
                if ndist == cdist {
                        fmt.Println("right")
                }
     
        case moves & Down != 0:
                ndist := int(math.Abs(float64(my.row + 1 - oprow))) + int(math.Abs(float64(my.col - opcol)))
                if ndist == cdist {
                        fmt.Println("down")
                }
        
	
	case moves & Up != 0:
                fmt.Println("up")

        case moves & Left != 0:
                fmt.Println("left")
        
        case moves & Right != 0:
                fmt.Println("right")
        
        case moves & Down != 0:
                fmt.Println("down")
	
	case moves == 0:
		fmt.Println("up")

	}
}
func (my Bot) getValidMoves(f Field) (result int) {
	result = 0
	if f.isValid(my.row+1, my.col) {
		result |= Down
	}
	if f.isValid(my.row-1, my.col) {
		result |= Up
	}
	if f.isValid(my.row, my.col+1) {
		result |= Right
	}
	if f.isValid(my.row, my.col-1) {
		result |=  Left
	}
	return result
}


