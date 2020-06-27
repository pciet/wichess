package piece

var CodeNames = []string{
	"none",
	"king",
	"queen",
	"rook",
	"bishop",
	"knight",
	"pawn",
	"war",
	"formpawn", // 'pawn' added to name to avoid POV-Ray reserved word 'form'
	"constructive",
	"confined",
}

var Names = []string{
	"",
	"King",
	"Queen",
	"Rook",
	"Bishop",
	"Knight",
	"Pawn",
	"War",
	"Form",
	"Constructive",
	"Confined",
}

func KindForCodeName(n string) Kind {
	for i, name := range CodeNames {
		if name == n {
			return Kind(i)
		}
	}
	return NoKind
}
