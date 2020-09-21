package piece

// CodeName returns a piece kind name string used for image filenames or other similar uses.
func CodeName(of Kind) string { return codeNames[of] }

// Name returns a capitalized name for the piece kind.
func Name(of Kind) string { return names[of] }

// KindForCodeName takes a code name string for the kind and returns either the associated kind
// or NoKind if none match.
func KindForCodeName(n string) Kind {
	for i, name := range CodeNames {
		if name == n {
			return Kind(i)
		}
	}
	return NoKind
}

var codeNames = []string{
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
	"imperfect",
	"derange",
}

var names = []string{
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
	"Imperfect",
	"Derange",
}
