package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isupper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func islower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func copy(board [8][8]byte) [8][8]byte {
	var boardCopy [8][8]byte
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			boardCopy[i][j] = board[i][j]
		}
	}
	return boardCopy
}

func board2string(board [8][8]byte) string {
	s := ""
	for _, row := range board {
		s += string(row[:])
	}
	return s
}

func Move(i1, j1, i2, j2 int, b *Board) {
	fmt.Println("move:", i1, j1, i2, j2)

	var board [8][8]byte
	for i, p := range b.Content {
		board[i/8][i%8] = byte(p)
	}

	if i1 == 7 && j1 == 4 && i2 == 7 && j2 == 2 {
		if !b.WhitesTurn {
			return
		}
		if !canCastle(
			copy(board), i1, j1, i2, j2, b.WhitesTurn,
			b.LeftWhiteRookMoved, b.RightWhiteRookMoved, b.WhiteKingMoved,
			b.LeftBlackRookMoved, b.RightBlackRookMoved, b.BlackKingMoved,
			true,
		) {
			return
		}
		board[7][4] = ' '
		board[7][2] = 'k'
		board[7][0] = ' '
		board[7][3] = 'r'
		b.Content = board2string(board)
		b.WhitesTurn = !b.WhitesTurn
		b.WhiteKingMoved = true
		b.LeftWhiteRookMoved = true
		b.Checkmate = checkmate(copy(board), !b.WhitesTurn)
	}

	if i1 == 7 && j1 == 4 && i2 == 7 && j2 == 6 {
		if !b.WhitesTurn {
			return
		}
		if !canCastle(
			copy(board), i1, j1, i2, j2, b.WhitesTurn,
			b.LeftWhiteRookMoved, b.RightWhiteRookMoved, b.WhiteKingMoved,
			b.LeftBlackRookMoved, b.RightBlackRookMoved, b.BlackKingMoved,
			false,
		) {
			return
		}
		board[7][4] = ' '
		board[7][6] = 'k'
		board[7][7] = ' '
		board[7][5] = 'r'
		b.Content = board2string(board)
		b.WhitesTurn = !b.WhitesTurn
		b.WhiteKingMoved = true
		b.RightWhiteRookMoved = true
		b.Checkmate = checkmate(copy(board), !b.WhitesTurn)
	}

	if i1 == 0 && j1 == 4 && i2 == 0 && j2 == 2 {
		if !b.WhitesTurn {
			return
		}
		if !canCastle(
			copy(board), i1, j1, i2, j2, b.WhitesTurn,
			b.LeftWhiteRookMoved, b.RightWhiteRookMoved, b.WhiteKingMoved,
			b.LeftBlackRookMoved, b.RightBlackRookMoved, b.BlackKingMoved,
			true,
		) {
			return
		}
		board[0][4] = ' '
		board[0][2] = 'K'
		board[0][0] = ' '
		board[0][3] = 'R'
		b.Content = board2string(board)
		b.WhitesTurn = !b.WhitesTurn
		b.BlackKingMoved = true
		b.LeftBlackRookMoved = true
		b.Checkmate = checkmate(copy(board), !b.WhitesTurn)
	}

	if i1 == 0 && j1 == 4 && i2 == 0 && j2 == 6 {
		if !b.WhitesTurn {
			return
		}
		if !canCastle(
			copy(board), i1, j1, i2, j2, b.WhitesTurn,
			b.LeftWhiteRookMoved, b.RightWhiteRookMoved, b.WhiteKingMoved,
			b.LeftBlackRookMoved, b.RightBlackRookMoved, b.BlackKingMoved,
			false,
		) {
			return
		}
		board[0][4] = ' '
		board[0][6] = 'K'
		board[0][7] = ' '
		board[0][5] = 'R'
		b.Content = board2string(board)
		b.WhitesTurn = !b.WhitesTurn
		b.BlackKingMoved = true
		b.RightBlackRookMoved = true
		b.Checkmate = checkmate(copy(board), !b.WhitesTurn)
	}

	if !isValidMovePiece(board[i1][j1], i1, j1, i2, j2) {
		return
	}

	if !isValidMove(copy(board), i1, j1, i2, j2, b.WhitesTurn) {
		return
	}

	if !isValidCheck(copy(board), i1, j1, i2, j2, b.WhitesTurn) {
		return
	}

	if i1 == 0 && j1 == 0 {
		b.LeftBlackRookMoved = true
	} else if i1 == 0 && j1 == 4 {
		b.BlackKingMoved = true
	} else if i1 == 0 && j1 == 7 {
		b.RightBlackRookMoved = true
	} else if i1 == 7 && j1 == 0 {
		b.LeftWhiteRookMoved = true
	} else if i1 == 7 && j1 == 4 {
		b.WhiteKingMoved = true
	} else if i1 == 7 && j1 == 7 {
		b.RightWhiteRookMoved = true
	}

	if board[i1][j1] == 'p' && i2 == 0 {
		board[i2][j2] = 'q'
	} else if board[i1][j1] == 'P' && i2 == 7 {
		board[i2][j2] = 'Q'
	} else {
		board[i2][j2] = board[i1][j1]
	}

	board[i1][j1] = ' '

	b.Content = board2string(board)
	b.Checkmate = checkmate(copy(board), !b.WhitesTurn)
	b.WhitesTurn = !b.WhitesTurn
}

func isValidMovePiece(piece byte, i1, j1, i2, j2 int) bool {
	switch piece {
	case 'P':
		if i1 == 1 && i2 == 3 && j1 == j2 {
			return true
		} else if i2-i1 == 1 && j1 == j2 {
			return true
		} else if i2-i1 == 1 && abs(j2-j1) == 1 {
			return true
		} else {
			return false
		}
	case 'p':
		if i1 == 6 && i2 == 4 && j1 == j2 {
			return true
		} else if i2-i1 == -1 && j1 == j2 {
			return true
		} else if i2-i1 == -1 && abs(j2-j1) == 1 {
			return true
		} else {
			return false
		}
	case 'r', 'R':
		return i1 == i2 || j1 == j2
	case 'n', 'N':
		if abs(i2-i1) == 1 && abs(j2-j1) == 2 {
			return true
		} else if abs(i2-i1) == 2 && abs(j2-j1) == 1 {
			return true
		} else {
			return false
		}
	case 'b', 'B':
		return abs(i2-i1) == abs(j2-j1)
	case 'q', 'Q':
		return i1 == i2 || j1 == j2 || abs(i2-i1) == abs(j2-j1)
	case 'k', 'K':
		return abs(i2-i1) <= 1 && abs(j2-j1) <= 1
	default:
		return false
	}
}

func isValidMove(board [8][8]byte, i1, j1, i2, j2 int, whitesTurn bool) bool {
	piece := board[i1][j1]
	if (whitesTurn && isupper(piece)) || (!whitesTurn && islower(piece)) {
		return false
	} else if i1 == i2 && j1 == j2 {
		return false
	} else if piece == ' ' {
		return false
	} else if board[i2][j2] != ' ' && ((islower(board[i1][j1]) && islower(board[i2][j2])) || (isupper(board[i1][j1]) && isupper(board[i2][j2]))) {
		return false
	} else if (piece == 'p' || piece == 'P') && j1 != j2 && board[i2][j2] == ' ' {
		return false
	} else if (piece == 'p' || piece == 'P') && j1 == j2 && board[i2][j2] != ' ' {
		return false
	} else if !isValidMovePiece(piece, i1, j1, i2, j2) {
		return false
	} else {
		return true
	}
}

func isValidCheck(board [8][8]byte, i1, j1, i2, j2 int, whitesTurn bool) bool {
	if board[i1][j1] == 'p' && i2 == 0 {
		board[i2][j2] = 'q'
	} else if board[i1][j1] == 'P' && i2 == 7 {
		board[i2][j2] = 'Q'
	} else {
		board[i2][j2] = board[i1][j1]
	}

	board[i1][j1] = ' '

	return !inCheck(board, whitesTurn)
}

func inCheck(board [8][8]byte, whitesTurn bool) bool {
	ki, kj := -1, -1
	for i, row := range board {
		for j, piece := range row {
			if (whitesTurn && piece == 'k') || (!whitesTurn && piece == 'K') {
				ki = i
				kj = j
			}
		}
	}

	for i, row := range board {
		for j, piece := range row {
			if isValidMove(board, i, j, ki, kj, !whitesTurn) {
				if piece == 'n' || piece == 'N' {
					return true
				} else if piece == 'p' || piece == 'P' {
					return true
				} else if piece == 'k' || piece == 'K' {
					return true
				} else if piece == 'b' || piece == 'B' {
					king_safe := false

					if i < ki && j < kj {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 || j+d > 7 {
								break
							}
							if board[i+d][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i < ki && j > kj {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 || j-d < 0 {
								break
							}
							if board[i+d][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki && j < kj {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 || j+d > 7 {
								break
							}
							if board[i-d][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki && j > kj {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 || j-d < 0 {
								break
							}
							if board[i-d][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if !king_safe {
						return true
					}
				} else if piece == 'r' || piece == 'R' {
					king_safe := false

					if i < ki {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 {
								break
							}
							if board[i+d][j] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 {
								break
							}
							if board[i-d][j] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if j < kj {
						for d := 1; d < kj-j; d++ {
							if j+d > 7 {
								break
							}
							if board[i][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if j > kj {
						for d := 1; d < j-kj; d++ {
							if j-d < 0 {
								break
							}
							if board[i][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if !king_safe {
						return true
					}
				} else if piece == 'q' || piece == 'Q' {
					king_safe := false

					if i < ki && j < kj {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 || j+d > 7 {
								break
							}
							if board[i+d][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i < ki && j > kj {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 || j-d < 0 {
								break
							}
							if board[i+d][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki && j < kj {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 || j+d > 7 {
								break
							}
							if board[i-d][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki && j > kj {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 || j-d < 0 {
								break
							}
							if board[i-d][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i < ki {
						for d := 1; d < ki-i; d++ {
							if i+d > 7 {
								break
							}
							if board[i+d][j] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if i > ki {
						for d := 1; d < i-ki; d++ {
							if i-d < 0 {
								break
							}
							if board[i-d][j] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if j < kj {
						for d := 1; d < kj-j; d++ {
							if j+d > 7 {
								break
							}
							if board[i][j+d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if j > kj {
						for d := 1; d < j-kj; d++ {
							if j-d < 0 {
								break
							}
							if board[i][j-d] != ' ' {
								king_safe = true
								break
							}
						}
					}
					if !king_safe {
						return true
					}
				}
			}
		}
	}
	return false
}

func checkmate(board [8][8]byte, whitesTurn bool) bool {
	if !inCheck(board, whitesTurn) {
		return false
	}
	for i1 := 0; i1 < 8; i1++ {
		for j1 := 0; j1 < 8; j1++ {
			for i2 := 0; i2 < 8; i2++ {
				for j2 := 0; j2 < 8; j2++ {
					if isValidMovePiece(board[i1][j1], i1, j1, i2, j2) && isValidMove(copy(board), i1, j1, i2, j2, whitesTurn) && isValidCheck(copy(board), i1, j1, i2, j2, whitesTurn) {
						return false
					}
				}
			}
		}
	}
	return true
}

func canCastle(
	board [8][8]byte, i1, j1, i2, j2 int, whitesTurn,
	leftWhiteRookMoved, rightWhiteRookMoved, whiteKingMoved,
	leftBlackRookMoved, rightBlackRookMoved, blackKingMoved,
	castlingLeft bool,
) bool {
	if inCheck(board, whitesTurn) {
		return false
	}
	if whitesTurn && castlingLeft {
		if whiteKingMoved || leftWhiteRookMoved {
			return false
		}
		for j := 1; j < 4; j++ {
			if board[7][j] != ' ' {
				return false
			}
		}
	}
	if whitesTurn && !castlingLeft {
		if whiteKingMoved || rightWhiteRookMoved {
			return false
		}
		for j := 5; j < 7; j++ {
			if board[7][j] != ' ' {
				return false
			}
		}
	}
	if !whitesTurn && castlingLeft {
		if blackKingMoved || leftBlackRookMoved {
			return false
		}
		for j := 1; j < 4; j++ {
			if board[0][j] != ' ' {
				return false
			}
		}
	}
	if !whitesTurn && castlingLeft {
		if blackKingMoved || rightBlackRookMoved {
			return false
		}
		for j := 5; j < 7; j++ {
			if board[0][j] != ' ' {
				return false
			}
		}
	}
	return true
}
