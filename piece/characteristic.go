package piece

// A Characteristic is something special about a piece that affects gameplay with it, with a one
// word name for it. A piece can have zero, one, or two characteristics.
type Characteristic int

// These Characteristic constants are used by package rules to adjust calculations of moves to
// consider what the characteristic does for the piece, or they are used to read a description
// using the CharacteristicDescription function.
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
	Protective
	Extricates
	Normalizes
	Orders
)

// Characteristics returns up to two Characteristic that apply to a piece. If the piece has no
// characteristics then the first return will be NoCharacteristic.
func Characteristics(of Kind) (Characteristic, Characteristic) {
	c := characteristicList[of]
	return c.A, c.B
}

// CharacteristicDescription returns a sentence or two that describe the characteristic for players.
func CharacteristicDescription(of Characteristic) string { return characteristicDescriptions[of] }

// CharacteristicName returns a capitalized name string of the characteristic.
func CharacteristicName(of Characteristic) string { return characteristicNames[of] }

type characteristics struct {
	A, B Characteristic
}

var characteristicList = []characteristics{
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
	{Protective, NoCharacteristic},
	{Protective, NoCharacteristic},
	{Extricates, NoCharacteristic},
	{Tense, NoCharacteristic},
	{Normalizes, Orders},
	{NoCharacteristic, NoCharacteristic},
	{Asserts, NoCharacteristic},
	{Fantasy, NoCharacteristic},
}

var characteristicDescriptions = []string{
	"",
	`When captured all adjacent pieces and the capturing piece are also captured.`,
	`Automatically moves to capture any opponent piece that's moved adjacent. You then do your normal move.`,
	`Your adjacent pieces (except the king and queens) get added moves that can't be used to capture.`,
	`Your adjacent pieces get an added move across this one.`,
	`Adjacent opponent pieces (except queens and the king) can't move.`,
	`Pawn pieces can't capture this.`,
	`Can only be captured by queens or the king.`,
	`When captured if its starting square is empty then it returns there instead.`,
	`Your adjacent pieces are immaterial (pawn pieces can't capture them).`,
	`Can't be captured if adjacent to another piece with protective even if it's your opponent's.`,
	`When your king is in check it may move onto this piece from anywhere. When that move happens this piece is captured by your opponent.`,
	`Adjacent pieces aren't affected by their characteristics. If two normalizing pieces are adjacent then only they are normalized`,
	`Adjacent pieces neutralize (when captured all adjacent pieces and the capturing piece are
 also captured).`,
}

var characteristicNames = []string{
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
	"Protective",
	"Extricates",
	"Normalizes",
	"Orders",
}
