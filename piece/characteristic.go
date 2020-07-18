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
	Tense
	Fantasy
	Keep
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
	{Tense, NoCharacteristic},
	{Fantasy, Immaterial},
	{Keep, NoCharacteristic},
}

var CharacteristicDescriptions = []string{
	"",
	`When this is captured all adjacent pieces and the capturing piece are also captured.`,
	`This automatically moves itself to capture when the opponent moves adjacent.`,
	`Your adjacent pieces (except the king and queens) get added moves that can't be used to capture.`,
	`Your adjacent pieces get an added move across this one.`,
	`Adjacent opponent pieces except queens and the king can't move.`,
	`Pawn pieces can't capture this.`,
	`Can only be captured by queens or the king.`,
	`When captured if its starting square is empty then it returns there.`,
	`Adjacent friendly pieces become immaterial (pawn pieces can't capture them).`,
}

var CharacteristicNames = []string{
	"",
	"Neutralizes",
	"Asserts",
	"Enables",
	"Reveals",
	"Stops",
	"Immaterial",
	"Tense",
	"Fantasy",
	"Keep",
}
