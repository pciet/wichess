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
	"original",
	"irrelevant",
	"evident",
	"line",
	"impossible",
	"convenient",
	"appropriate",
	"warprook", // 'rook' added to avoid reserved 'warp'
	"brilliant",
	"simple",
	"exit",
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
	"Original",
	"Irrelevant",
	"Evident",
	"Line",
	"Impossible",
	"Convenient",
	"Appropriate",
	"Warp",
	"Brilliant",
	"Simple",
	"Exit",
}

func KindForCodeName(n string) Kind {
	for i, name := range CodeNames {
		if name == n {
			return Kind(i)
		}
	}
	return NoKind
}
