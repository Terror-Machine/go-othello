// Copyright (c) 2025 HyHy. All rights reserved.
//
// go-Othello adalah permainan Othello interaktif yang dimainkan
// melalui terminal, dengan output visual berupa gambar PNG.

package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/fogleman/gg"
)

const (
	EMPTY        = 0
	PLAYER_BLACK = 1
	PLAYER_WHITE = 2
)

type GameState struct {
	Board      [64]int
	Turn       int
	ValidMoves []int
}

func generateOthelloImage(board [64]int, validMoves []int) {
	const squareSize = 60.0
	const boardSize = squareSize * 8
	const fontPath = "font/arial.ttf"
	dc := gg.NewContext(int(boardSize), int(boardSize))
	dc.SetColor(color.RGBA{0, 128, 55, 255})
	dc.Clear()
	colLabels := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	rowLabels := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			index := r*8 + c
			x := float64(c * squareSize)
			y := float64(r * squareSize)
			dc.SetColor(color.RGBA{0, 100, 43, 255})
			dc.SetLineWidth(1)
			dc.DrawRectangle(x, y, squareSize, squareSize)
			dc.Stroke()
			if board[index] != EMPTY {
				x_center := x + squareSize/2
				y_center := y + squareSize/2
				radius := squareSize * 0.4
				if board[index] == PLAYER_BLACK {
					dc.SetRGB(0, 0, 0)
				} else {
					dc.SetRGB(1, 1, 1)
				}
				dc.DrawCircle(x_center, y_center, radius)
				dc.Fill()
			}
		}
	}
	if err := dc.LoadFontFace(fontPath, 14); err == nil {
		dc.SetRGB(1, 1, 1)
		for i := 0; i < 8; i++ {
			row_x := float64(0*squareSize) + 5
			row_y := float64(i*squareSize) + 15
			dc.DrawStringAnchored(rowLabels[i], row_x, row_y, 0, 0)
			col_x := float64(i*squareSize) + squareSize - 5
			col_y := float64(0*squareSize) + 15
			dc.DrawStringAnchored(colLabels[i], col_x, col_y, 1, 0)
		}
	} else {
		log.Printf("Gagal memuat font dari '%s': %v\n", fontPath, err)
	}
	dc.SetColor(color.RGBA{255, 255, 255, 80})
	for _, moveIndex := range validMoves {
		r := moveIndex / 8
		c := moveIndex % 8
		x := float64(c*squareSize) + squareSize/2
		y := float64(r*squareSize) + squareSize/2
		dc.DrawCircle(x, y, squareSize*0.15)
		dc.Fill()
	}
	if err := dc.SavePNG("othello.png"); err != nil {
		log.Fatalf("Gagal menyimpan gambar: %v", err)
	}
}

func getFlips(board [64]int, player int, index int) []int {
	if board[index] != EMPTY {
		return nil
	}
	opponent := PLAYER_WHITE
	if player == PLAYER_WHITE {
		opponent = PLAYER_BLACK
	}
	directions := []int{-9, -8, -7, -1, 1, 7, 8, 9}
	_, c := index/8, index%8
	var allFlips []int
	for _, dir := range directions {
		if (dir == -1 || dir == -9 || dir == 7) && c == 0 { continue }
		if (dir == 1 || dir == 9 || dir == -7) && c == 7 { continue }
		var flipsInDir []int
		current := index + dir
		for current >= 0 && current < 64 {
			if board[current] == opponent {
				flipsInDir = append(flipsInDir, current)
			} else if board[current] == player {
				allFlips = append(allFlips, flipsInDir...)
				break
			} else {
				break
			}
			if (dir == -1 || dir == -9 || dir == 7) && (current%8) == 0 { break }
			if (dir == 1 || dir == 9 || dir == -7) && (current%8) == 7 { break }
			current += dir
		}
	}
	return allFlips
}

func getValidMoves(board [64]int, player int) []int {
	var moves []int
	for i := 0; i < 64; i++ {
		if board[i] == EMPTY {
			if flips := getFlips(board, player, i); len(flips) > 0 {
				moves = append(moves, i)
			}
		}
	}
	return moves
}

func makeMove(board *[64]int, player int, index int, flips []int) {
	board[index] = player
	for _, i := range flips {
		board[i] = player
	}
}

func parseCoord(s string) (int, bool) {
	if len(s) != 2 {
		return 0, false
	}
	col := int(s[0] - 'a')
	row, err := strconv.Atoi(string(s[1]))
	if err != nil || col < 0 || col > 7 || row < 1 || row > 8 {
		return 0, false
	}
	return (row-1)*8 + col, true
}

func newGame() GameState {
	var board [64]int
	board[27] = PLAYER_WHITE
	board[28] = PLAYER_BLACK
	board[35] = PLAYER_BLACK
	board[36] = PLAYER_WHITE
	validMoves := getValidMoves(board, PLAYER_BLACK)
	return GameState{ Board: board, Turn: PLAYER_BLACK, ValidMoves: validMoves }
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Selamat datang di Game Othello CLI!")
	game := newGame()
	generateOthelloImage(game.Board, game.ValidMoves)
	fmt.Println("Papan Othello telah dibuat dan disimpan sebagai 'othello.png'.")
	fmt.Println("Anda bermain sebagai Hitam (⚫️). Giliran Anda untuk bergerak.")
	fmt.Println("\nPerintah:")
	fmt.Println("  d3        -> Meletakkan keping di D3.")
	fmt.Println("  pass      -> Melewati giliran (jika tidak ada langkah valid).")
	fmt.Println("  new       -> Memulai game baru.")
	fmt.Println("  exit      -> Keluar dari permainan.")
	fmt.Print("> ")
	for scanner.Scan() {
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if len(input) == 0 {
			fmt.Print("> ")
			continue
		}
		switch input {
		case "exit", "keluar":
			fmt.Println("Terima kasih sudah bermain!")
			return
		case "new":
			game = newGame()
			fmt.Println("Game baru telah dimulai! Anda bermain sebagai Hitam. Gambar diperbarui.")
			generateOthelloImage(game.Board, game.ValidMoves)
			fmt.Print("> ")
			continue
		case "pass":
			if len(game.ValidMoves) > 0 {
				fmt.Println("Anda masih memiliki langkah valid, tidak bisa melewati giliran.")
				fmt.Print("> ")
				continue
			}
			game.Turn = PLAYER_WHITE
		default:
			index, ok := parseCoord(input)
			if !ok {
				fmt.Println("Format gerakan tidak valid. Contoh: d3")
				fmt.Print("> ")
				continue
			}
			isValidMove := false
			for _, move := range game.ValidMoves {
				if move == index {
					isValidMove = true
					break
				}
			}
			if !isValidMove {
				fmt.Println("Langkah tidak valid. Pilih salah satu kotak yang ditandai.")
				fmt.Print("> ")
				continue
			}
			flips := getFlips(game.Board, PLAYER_BLACK, index)
			makeMove(&game.Board, PLAYER_BLACK, index, flips)
			game.Turn = PLAYER_WHITE
		}
		for game.Turn == PLAYER_WHITE {
			fmt.Println("Bot sedang berpikir...")
			time.Sleep(1 * time.Second)
			botMoves := getValidMoves(game.Board, PLAYER_WHITE)
			if len(botMoves) == 0 {
				fmt.Println("Bot tidak memiliki langkah valid dan melewati giliran.")
				game.Turn = PLAYER_BLACK
				break
			}
			bestMove := -1
			maxFlips := -1
			for _, move := range botMoves {
				flips := getFlips(game.Board, PLAYER_WHITE, move)
				if len(flips) > maxFlips {
					maxFlips = len(flips)
					bestMove = move
				}
			}
			flips := getFlips(game.Board, PLAYER_WHITE, bestMove)
			makeMove(&game.Board, PLAYER_WHITE, bestMove, flips)
			r, c := bestMove/8, bestMove%8
			fmt.Printf("Bot meletakkan keping di %c%d.\n", 'a'+c, r+1)
			game.Turn = PLAYER_BLACK
		}
		game.ValidMoves = getValidMoves(game.Board, PLAYER_BLACK)
		botMoves := getValidMoves(game.Board, PLAYER_WHITE)
		generateOthelloImage(game.Board, game.ValidMoves)
		fmt.Println("Gambar telah diperbarui.")
		if len(game.ValidMoves) == 0 && len(botMoves) == 0 {
			blackScore, whiteScore := 0, 0
			for _, piece := range game.Board {
				if piece == PLAYER_BLACK { blackScore++ }
				if piece == PLAYER_WHITE { whiteScore++ }
			}
			fmt.Println("\n--- PERMAINAN SELESAI ---")
			fmt.Printf("Skor Akhir: Hitam %d - %d Putih\n", blackScore, whiteScore)
			if blackScore > whiteScore {
				fmt.Println("Selamat, Anda menang!")
			} else if whiteScore > blackScore {
				fmt.Println("Bot menang.")
			} else {
				fmt.Println("Permainan berakhir seri!")
			}
			fmt.Print("Ketik 'new' untuk main lagi atau 'exit' untuk keluar.\n> ")
		} else if len(game.ValidMoves) == 0 {
			fmt.Println("Anda tidak memiliki langkah valid. Giliran Anda dilewati. Ketik 'pass' untuk melanjutkan.")
			fmt.Print("> ")
		} else {
			fmt.Println("Giliran Anda.")
			fmt.Print("> ")
		}
	}
}