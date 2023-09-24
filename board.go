package main

var Boards = make(map[string]*Board)

type Board struct {
    ID string
    Content string
    WhitesTurn bool
    Checkmate bool
    LeftWhiteRookMoved bool
    RightWhiteRookMoved bool
    WhiteKingMoved bool
    LeftBlackRookMoved bool
    RightBlackRookMoved bool
    BlackKingMoved bool
}

func NewBoard() *Board {
	return &Board{
        ID: "default",
        Content: "RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr",
        WhitesTurn: true,
        Checkmate: false,
        LeftWhiteRookMoved: false,
        RightWhiteRookMoved: false,
        WhiteKingMoved: false,
        LeftBlackRookMoved: false,
        RightBlackRookMoved: false,
        BlackKingMoved: false,
    }
}
