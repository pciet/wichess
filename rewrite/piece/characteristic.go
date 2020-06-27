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
	{Reveals, NoCharacteristic},
	{Asserts, NoCharacteristic},
	{Enables, NoCharacteristic},
}

var CharacteristicDescriptions = []string{
	"",
	`When this is captured all adjacent pieces and the capturing piece are also captured.`,
	`This automatically moves itself to capture when the opponent moves adjacent.`,
	`Your adjacent pieces get added moves.`,
	`Your adjacent pieces get an added move across this one.`,
}

var CharacteristicNames = []string{
	"",
	"Neutralizes",
	"Asserts",
	"Enables",
	"Reveals",
}
