package piece

type Characteristic int

const (
	NoCharacteristic Characteristic = iota
	Neutralizes
	Asserts
	Enables
	Reveals
)

type Characteristics struct {
	A, B Characteristic
}

var CharacteristicList = []Characteristics{
	{}, {}, {}, {}, {}, {}, {},
	{Neutralizes, NoCharacteristic}, // War
	{Reveals, Enables},
	{Asserts, NoCharacteristic},
}

var CharacteristicDescriptions = []string{
	"",
	`When this piece is captured the capturing piece and all adjacent pieces from both sides
 are also captured.`,
	`When an opponent piece moves adjacent then this piece automatically moves itself to
 capture it.`,
	`Friendly pieces adjacent to this one gain additional moves. These moves can't be used
 to capture.`,
	`Friendly pieces in the three adjacent squares behind this piece can move without capturing
 to the square ahead.`,
}

var CharacteristicNames = []string{
	"",
	"Neutralizes",
	"Asserts",
	"Enables",
	"Reveals",
}
