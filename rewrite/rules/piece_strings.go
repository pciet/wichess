package rules

// TODO: strings and info defined in JSON shared between host and web so there's only one def

func (a PieceKind) Name() string {
	switch a {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	case War:
		return "War"
	case Form:
		return "Form"
	case Constructive:
		return "Constructive"
	}
	Panic("bad piece kind", a)
	return ""
}

const (
	CodeKing         = "king"
	CodeQueen        = "queen"
	CodeRook         = "rook"
	CodeBishop       = "bishop"
	CodeKnight       = "knight"
	CodePawn         = "pawn"
	CodeWar          = "war"
	CodeForm         = "formpawn"
	CodeConstructive = "constructive"

	Neutralizes = "Neutralizes"
	Asserts     = "Asserts"
	Enables     = "Enables"
	Reveals     = "Reveals"
)

func (a PieceKind) CodeName() string {
	switch a {
	case King:
		return CodeKing
	case Queen:
		return CodeQueen
	case Rook:
		return CodeRook
	case Bishop:
		return CodeBishop
	case Knight:
		return CodeKnight
	case Pawn:
		return CodePawn
	case War:
		return CodeWar
	case Form:
		return CodeForm
	case Constructive:
		return CodeConstructive
	}
	return ""
}

func CodeNameKind(a string) PieceKind {
	switch a {
	case CodeKing:
		return King
	case CodeQueen:
		return Queen
	case CodeRook:
		return Rook
	case CodeBishop:
		return Bishop
	case CodeKnight:
		return Knight
	case CodePawn:
		return Pawn
	case CodeWar:
		return War
	case CodeForm:
		return Form
	case CodeConstructive:
		return Constructive
	}
	return NoKind
}

func IsPieceCodeName(s string) bool {
	if CodeNameKind(s) == NoKind {
		return false
	}
	return true
}

func (a PieceKind) StartDescription() string {
	switch a {
	case King:
		return "The King starts in the back row five squares from the left."
	case Queen:
		return "The Queen starts in the back row four squares from the left."
	case Rook:
		return "Rooks start in the back row on the first and last squares."
	case Bishop:
		return "Bishops start in the back row adjacent to the King and Queen."
	case Knight:
		return "Knights start in the back row between the Rook and Bishop on both sides."
	case Pawn:
		return "Pawns all start on the first row with eight total."
	}
	return ""
}

func (a PieceKind) MovesDescription() string {
	switch a {
	case King:
		return `The King can move to any adjacent square. When the King and a friendly Rook both
 haven't moved and the squares between them are empty and unthreatened then the King can do the
 castle move. The King is in check when threatened by an opponent piece's move. Any moves that
 cause your King to be in check are not allowed, and you have won when the opponent King cannot
 leave check.`
	case Queen:
		return "The Queen can move any distance along adjacent vectors."
	case Rook:
		return "The Rook can move any distance along the up-down or left-right vectors."
	case Bishop:
		return "The Bishop can move any distance along the corner vectors."
	case Knight:
		return `The Knight can move over other pieces to get to it's move, which is two squares up
 then one square to the left or right. This move is up-down or left-right for up to eight possible
 moves.`
	case Pawn:
		return `The Pawn can only take forward toward the opponent and adjacent to it's current
 square. The first move is one or two squares forward, then moves after that are only one square
 forward. The en passant move allows a Pawn to take an opponent's Pawn that makes it's initial
 two square move through where this Pawn would normally be able to capture. If a Pawn moves to
 the final row on the other side of the board then it is promoted to a Queen, Rook, Bishop, or
 Knight. A Pawn does not have to take and can move forward instead.`
	case War:
		return "The War can only move one square forward for it's first move."
	}
	return ""
}

func (a PieceKind) CharacteristicAString() string {
	switch a {
	case Form:
		return Reveals
	case War:
		return Neutralizes
	case Constructive:
		return Asserts
	}
	return ""
}

func (a PieceKind) CharacteristicBString() string {
	switch a {
	case Form:
		return Enables
	}
	return ""
}

func CharacteristicDescription(c string) string {
	switch c {
	case Neutralizes:
		return `When this piece is captured the capturing piece and all adjacent pieces from both
 sides are also captured.`
	case Asserts:
		return `When an opponent piece moves adjacent then this piece automatically moves itself
 to capture it.`
	case Enables:
		return `Friendly pieces adjacent to this one gain additional moves. These moves can't be
 used to capture.`
	case Reveals:
		return `Friendly pieces in the three adjacent squares behind this piece can move without
 capturing to the square ahead.`
	}
	return ""
}
