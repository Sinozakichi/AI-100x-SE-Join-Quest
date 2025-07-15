package bdd

import (
	"ai100x-order/src"
	"fmt"

	"github.com/cucumber/godog"
)

var service *src.ChineseChessService

func theBoardIsEmptyExceptForARedGeneralAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.General, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedGuardAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Guard, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedRookAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Rook, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedHorseAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Horse, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedCannonAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Cannon, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedElephantAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Elephant, Color: src.Red}
	return nil
}
func theBoardIsEmptyExceptForARedSoldierAt(row, col int) error {
	service = src.NewChineseChessService()
	service.Board.Grid[row-1][col-1] = &src.Piece{Type: src.Soldier, Color: src.Red}
	return nil
}
func theBoardHas(table *godog.Table) error {
	service = src.NewChineseChessService()
	for i, row := range table.Rows {
		if i == 0 {
			continue // skip header
		}
		pieceName := row.Cells[0].Value
		pos := row.Cells[1].Value // (row, col)
		var pieceType src.PieceType
		var color src.Color

		// 解析棋子名稱和顏色
		if pieceName == "Red General" {
			pieceType = src.General
			color = src.Red
		} else if pieceName == "Black General" {
			pieceType = src.General
			color = src.Black
		} else if pieceName == "Red Guard" {
			pieceType = src.Guard
			color = src.Red
		} else if pieceName == "Black Guard" {
			pieceType = src.Guard
			color = src.Black
		} else if pieceName == "Red Rook" {
			pieceType = src.Rook
			color = src.Red
		} else if pieceName == "Black Rook" {
			pieceType = src.Rook
			color = src.Black
		} else if pieceName == "Red Horse" {
			pieceType = src.Horse
			color = src.Red
		} else if pieceName == "Black Horse" {
			pieceType = src.Horse
			color = src.Black
		} else if pieceName == "Red Cannon" {
			pieceType = src.Cannon
			color = src.Red
		} else if pieceName == "Black Cannon" {
			pieceType = src.Cannon
			color = src.Black
		} else if pieceName == "Red Elephant" {
			pieceType = src.Elephant
			color = src.Red
		} else if pieceName == "Black Elephant" {
			pieceType = src.Elephant
			color = src.Black
		} else if pieceName == "Red Soldier" {
			pieceType = src.Soldier
			color = src.Red
		} else if pieceName == "Black Soldier" {
			pieceType = src.Soldier
			color = src.Black
		}

		var r, c int
		fmt.Sscanf(pos, "(%d, %d)", &r, &c)
		service.Board.Grid[r-1][c-1] = &src.Piece{Type: pieceType, Color: color}
	}
	return nil
}
func redMovesTheGeneralFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.General || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red general"}
		return nil
	}
	// palace: row 1~3, col 4~6
	if toRow < 1 || toRow > 3 || toCol < 4 || toCol > 6 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "out of palace"}
		return nil
	}
	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)
	if dr+dc != 1 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "must move 1 step"}
		return nil
	}
	// 模擬移動
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece
	// 檢查將帥對面
	var redCol, blackCol, redRow, blackRow int
	for r := 0; r < 10; r++ {
		for c := 0; c < 9; c++ {
			p := service.Board.Grid[r][c]
			if p != nil && p.Type == src.General {
				if p.Color == src.Red {
					redRow, redCol = r, c
				} else {
					blackRow, blackCol = r, c
				}
			}
		}
	}
	if redCol == blackCol {
		clear := true
		for r := redRow + 1; r < blackRow; r++ {
			if service.Board.Grid[r][redCol] != nil {
				clear = false
				break
			}
		}
		if clear {
			service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "generals face each other"}
			return nil
		}
	}
	service.LastMoveResult = &src.MoveResult{Legal: true}
	return nil
}
func redMovesTheGuardFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Guard || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red guard"}
		return nil
	}

	// 宮內移動限制: row 1~3, col 4~6
	if toRow < 1 || toRow > 3 || toCol < 4 || toCol > 6 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "out of palace"}
		return nil
	}

	// 士只能走斜線
	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)
	if dr != 1 || dc != 1 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "must move diagonally"}
		return nil
	}

	// 檢查目標位置是否有己方棋子
	target := service.Board.Grid[toRow-1][toCol-1]
	if target != nil && target.Color == piece.Color {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
		return nil
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true}
	return nil
}
func redMovesTheRookFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Rook || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red rook"}
		return nil
	}

	// 車只能走直線
	if fromRow != toRow && fromCol != toCol {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "must move in straight line"}
		return nil
	}

	// 檢查路徑上是否有其他棋子
	if fromRow == toRow { // 水平移動
		start, end := fromCol, toCol
		if fromCol > toCol {
			start, end = toCol, fromCol
		}
		for c := start + 1; c < end; c++ {
			if service.Board.Grid[fromRow-1][c-1] != nil {
				service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "path blocked"}
				return nil
			}
		}
	} else { // 垂直移動
		start, end := fromRow, toRow
		if fromRow > toRow {
			start, end = toRow, fromRow
		}
		for r := start + 1; r < end; r++ {
			if service.Board.Grid[r-1][fromCol-1] != nil {
				service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "path blocked"}
				return nil
			}
		}
	}

	// 檢查目標位置是否有己方棋子
	target := service.Board.Grid[toRow-1][toCol-1]
	if target != nil && target.Color == piece.Color {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
		return nil
	}

	// 檢查是否吃掉對方將軍
	isWin := false
	if target != nil && target.Type == src.General && target.Color != piece.Color {
		isWin = true
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true, Win: isWin}
	return nil
}
func redMovesTheHorseFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Horse || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red horse"}
		return nil
	}

	// 馬走日字
	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)
	if !((dr == 1 && dc == 2) || (dr == 2 && dc == 1)) {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "invalid horse move"}
		return nil
	}

	// 檢查馬腿是否被卡住
	var legRow, legCol int
	if dr == 1 { // 水平移動較多
		legRow = fromRow
		if fromCol < toCol {
			legCol = fromCol + 1
		} else {
			legCol = fromCol - 1
		}
	} else { // 垂直移動較多
		legCol = fromCol
		if fromRow < toRow {
			legRow = fromRow + 1
		} else {
			legRow = fromRow - 1
		}
	}

	if service.Board.Grid[legRow-1][legCol-1] != nil {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "horse leg blocked"}
		return nil
	}

	// 檢查目標位置是否有己方棋子
	target := service.Board.Grid[toRow-1][toCol-1]
	if target != nil && target.Color == piece.Color {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
		return nil
	}

	// 檢查是否吃掉對方將軍
	isWin := false
	if target != nil && target.Type == src.General && target.Color != piece.Color {
		isWin = true
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true, Win: isWin}
	return nil
}
func redMovesTheCannonFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Cannon || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red cannon"}
		return nil
	}

	// 炮只能走直線
	if fromRow != toRow && fromCol != toCol {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "must move in straight line"}
		return nil
	}

	// 檢查目標位置是否有棋子
	target := service.Board.Grid[toRow-1][toCol-1]

	// 計算路徑上的棋子數量
	screenCount := 0
	if fromRow == toRow { // 水平移動
		start, end := fromCol, toCol
		if fromCol > toCol {
			start, end = toCol, fromCol
		}
		for c := start + 1; c < end; c++ {
			if service.Board.Grid[fromRow-1][c-1] != nil {
				screenCount++
			}
		}
	} else { // 垂直移動
		start, end := fromRow, toRow
		if fromRow > toRow {
			start, end = toRow, fromRow
		}
		for r := start + 1; r < end; r++ {
			if service.Board.Grid[r-1][fromCol-1] != nil {
				screenCount++
			}
		}
	}

	// 炮的移動規則：
	// 1. 如果目標位置沒有棋子，則路徑上不能有任何棋子
	// 2. 如果目標位置有棋子，則路徑上必須有且只有一個棋子作為炮架
	if target == nil {
		if screenCount > 0 {
			service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot move over pieces"}
			return nil
		}
	} else {
		if screenCount != 1 {
			service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "need exactly one screen to capture"}
			return nil
		}

		// 不能吃自己的棋子
		if target.Color == piece.Color {
			service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
			return nil
		}
	}

	// 檢查是否吃掉對方將軍
	isWin := false
	if target != nil && target.Type == src.General && target.Color != piece.Color {
		isWin = true
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true, Win: isWin}
	return nil
}
func redMovesTheElephantFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Elephant || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red elephant"}
		return nil
	}

	// 象走田字
	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)
	if dr != 2 || dc != 2 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "must move exactly 2 steps diagonally"}
		return nil
	}

	// 象不能過河
	if toRow > 5 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "elephant cannot cross river"}
		return nil
	}

	// 檢查象眼是否被塞住
	eyeRow := (fromRow + toRow) / 2
	eyeCol := (fromCol + toCol) / 2
	if service.Board.Grid[eyeRow-1][eyeCol-1] != nil {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "elephant eye blocked"}
		return nil
	}

	// 檢查目標位置是否有己方棋子
	target := service.Board.Grid[toRow-1][toCol-1]
	if target != nil && target.Color == piece.Color {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
		return nil
	}

	// 檢查是否吃掉對方將軍
	isWin := false
	if target != nil && target.Type == src.General && target.Color != piece.Color {
		isWin = true
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true, Win: isWin}
	return nil
}
func redMovesTheSoldierFromTo(fromRow, fromCol, toRow, toCol int) error {
	piece := service.Board.Grid[fromRow-1][fromCol-1]
	if piece == nil || piece.Type != src.Soldier || piece.Color != src.Red {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "not red soldier"}
		return nil
	}

	dr := toRow - fromRow
	dc := toCol - fromCol

	// 兵只能前進，不能後退
	if dr < 0 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "soldier cannot move backward"}
		return nil
	}

	// 兵每次只能走一步
	if abs(dr) > 1 || abs(dc) > 1 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "soldier can only move one step"}
		return nil
	}

	// 兵不能同時水平和垂直移動
	if dr != 0 && dc != 0 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "soldier must move straight"}
		return nil
	}

	// 兵未過河前只能前進
	if fromRow < 6 && dc != 0 {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "soldier can only move forward before crossing river"}
		return nil
	}

	// 檢查目標位置是否有己方棋子
	target := service.Board.Grid[toRow-1][toCol-1]
	if target != nil && target.Color == piece.Color {
		service.LastMoveResult = &src.MoveResult{Legal: false, Reason: "cannot capture own piece"}
		return nil
	}

	// 檢查是否吃掉對方將軍
	isWin := false
	if target != nil && target.Type == src.General && target.Color != piece.Color {
		isWin = true
	}

	// 移動棋子
	service.Board.Grid[fromRow-1][fromCol-1] = nil
	service.Board.Grid[toRow-1][toCol-1] = piece

	service.LastMoveResult = &src.MoveResult{Legal: true, Win: isWin}
	return nil
}
func theMoveIsLegal() error {
	if service.LastMoveResult == nil || !service.LastMoveResult.Legal {
		return fmt.Errorf("move should be legal, but was illegal")
	}
	return nil
}
func theMoveIsIllegal() error {
	if service.LastMoveResult == nil || service.LastMoveResult.Legal {
		return fmt.Errorf("move should be illegal, but was legal")
	}
	return nil
}
func redWinsImmediately() error {
	if service.LastMoveResult == nil || !service.LastMoveResult.Win {
		return fmt.Errorf("red should win, but did not")
	}
	return nil
}
func theGameIsNotOverJustFromThatCapture() error {
	if service.LastMoveResult == nil || service.LastMoveResult.Win {
		return fmt.Errorf("game should not be over, but it is")
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func InitializeChineseChessScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the board is empty except for a Red General at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedGeneralAt)
	ctx.Step(`^the board is empty except for a Red Guard at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedGuardAt)
	ctx.Step(`^the board is empty except for a Red Rook at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedRookAt)
	ctx.Step(`^the board is empty except for a Red Horse at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedHorseAt)
	ctx.Step(`^the board is empty except for a Red Cannon at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedCannonAt)
	ctx.Step(`^the board is empty except for a Red Elephant at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedElephantAt)
	ctx.Step(`^the board is empty except for a Red Soldier at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedSoldierAt)
	ctx.Step(`^the board has:$`, theBoardHas)
	ctx.Step(`^Red moves the General from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheGeneralFromTo)
	ctx.Step(`^Red moves the Guard from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheGuardFromTo)
	ctx.Step(`^Red moves the Rook from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheRookFromTo)
	ctx.Step(`^Red moves the Horse from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheHorseFromTo)
	ctx.Step(`^Red moves the Cannon from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheCannonFromTo)
	ctx.Step(`^Red moves the Elephant from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheElephantFromTo)
	ctx.Step(`^Red moves the Soldier from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheSoldierFromTo)
	ctx.Step(`^the move is legal$`, theMoveIsLegal)
	ctx.Step(`^the move is illegal$`, theMoveIsIllegal)
	ctx.Step(`^Red wins immediately$`, redWinsImmediately)
	ctx.Step(`^the game is not over just from that capture$`, theGameIsNotOverJustFromThatCapture)
}
