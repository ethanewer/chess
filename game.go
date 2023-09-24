package main

import "fmt"

func Move(i1, j1, i2, j2 int, board *Board) {
	fmt.Println("move: ", i1, j1, i2, j2)
}

func updateBoard(
	board [8][8]byte, i1, j1, i2, j2 int, whitesTurn, 
	leftWhiteRookMoved, rightWhiteRookMoved, whiteKingMoved, 
	leftBlackRookMoved, rightBlackRookMoved, blackKingMoved,
	castlingLeft bool,
) {

}

func isValidMovePiece(piece byte, i1, j1, i2, j2 int) bool {
	return false
}

func isValidMove(board [8][8]byte, i1, j1, i2, j2 int, whitesTurn bool) bool {
	return false;
}

func isValidCheck(board [8][8]byte, i1, j1, i2, j2 int, whitesTurn bool) bool {
	return false
}

func inCheck(board [8][8]byte, whitesTurn bool) bool {
	return false
}

func checkmate(board [8][8]byte, whitesTurn bool) bool {
	return false
}

func canCastle(
	board [8][8]byte, i1, j1, i2, j2 int, whitesTurn, 
	leftWhiteRookMoved, rightWhiteRookMoved, whiteKingMoved, 
	leftBlackRookMoved, rightBlackRookMoved, blackKingMoved,
	castlingLeft bool,
) bool {
	return false
}

