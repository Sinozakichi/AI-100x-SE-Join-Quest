package src

type PieceType string

type Color string

const (
	Red   Color = "Red"
	Black Color = "Black"
)

const (
	General  PieceType = "General"
	Guard    PieceType = "Guard"
	Rook     PieceType = "Rook"
	Horse    PieceType = "Horse"
	Cannon   PieceType = "Cannon"
	Elephant PieceType = "Elephant"
	Soldier  PieceType = "Soldier"
)

type Piece struct {
	Type  PieceType
	Color Color
}

type Board struct {
	Grid [10][9]*Piece // 1-based: row 1~10, col 1~9
}

type MoveResult struct {
	Legal  bool
	Win    bool
	Reason string
}

type ChineseChessService struct {
	Board          Board
	Turn           Color
	LastMoveResult *MoveResult
}

func NewChineseChessService() *ChineseChessService {
	return &ChineseChessService{
		Board: Board{},
		Turn:  Red,
	}
}

func (s *ChineseChessService) Reset() {
	s.Board = Board{}
	s.Turn = Red
}

// 之後會依據 scenario 實作各種 setup 與 move 方法
