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
	`<p>When captured all adjacent pieces and the capturing piece are also captured.</p>`,
	`<p>Automatically moves to capture any opponent piece that's moved to an adjacent square. You then do your normal move.</p>
<p>If multiple pieces would assert then the first moves that's counter-clockwise starting at the middle left from the white player's perspective.</p>`,
	`<p>Your adjacent pieces (except the <a href="/details?p=king">king</a> and <a href="/details?p=queen">queens</a>) get added moves that can't be used to capture.</p>`,
	`<p>Your adjacent pieces get an added move across this one.</p>`,
	`<p>Adjacent opponent pieces (except <a href="/details?p=queen">queens</a> and the <a href="/details?p=king">king</a>) can't move.</p>`,
	`<p>Can't be captured by <a href="/details?p=pawn">pawn</a> pieces.</p>`,
	`<p>Can only be captured by <a href="/details?p=queen">queens</a> or the <a href="/details?p=king">king</a>.</p>`,
	`<p>When captured this piece instead returns to its starting square if empty.</p>`,
	`<p>Your adjacent pieces are immaterial (<a href="/details?p=pawn">pawn</a> pieces can't capture them).</p>`,
	`<p>Can't be captured if adjacent to another piece with protective, even if that other piece is your opponent's.</p>`,
	`<p>When your <a href="/details?p=king">king</a> is in check it may move onto this piece from anywhere. When that move happens this piece is captured by your opponent.</p>`,
	`<p>Adjacent pieces aren't affected by their characteristics.</p>
<p>If two normalizing pieces are adjacent then only they are normalized.</p>
<p>Characteristics conveyed to a piece by other pieces aren't affected by normalizes.</p>`,
	`<p>Adjacent pieces neutralize (when captured all adjacent pieces and the capturing piece are
 also captured).</p>`,
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
