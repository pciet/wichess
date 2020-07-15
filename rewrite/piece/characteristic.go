package piece

type Characteristic int

const (
	NoCharacteristic Characteristic = iota
	Neutralizes
	Asserts
	Enables
	Reveals
	Stops
	Immaterial
)

type Characteristics struct {
	A, B Characteristic
}

var CharacteristicList = []Characteristics{
	{}, {}, {}, {}, {}, {}, {},
	{Neutralizes, NoCharacteristic}, // War
	{Reveals, NoCharacteristic},
	{Asserts, NoCharacteristic},
	{Enables, NoCharacteristic},
	{Stops, Asserts},
	{Immaterial, NoCharacteristic},
}

var CharacteristicDescriptions = []string{
	"",
	`When this is captured all adjacent pieces and the capturing piece are also captured.`,
	`This automatically moves itself to capture when the opponent moves adjacent.`,
	`Your adjacent pieces (except the king and queens) get added moves that can't be used to capture.`,
	`Your adjacent pieces get an added move across this one.`,
	`Adjacent opponent pieces except queens and the king can't move.`,
	`Pawn pieces can't capture this.`,
}

var CharacteristicNames = []string{
	"",
	"Neutralizes",
	"Asserts",
	"Enables",
	"Reveals",
	"Stops",
	"Immaterial",
}
