package piece

// DetailsHTML returns an HTML string meant for the individual piece details webpages.
// Characteristic descriptions are not included.
func DetailsHTML(of Kind) string { return detailsHTML[of] }

var detailsHTML = []string{
	"",
	`<h1>KING</h1>
<p>Each player has one king. Capturing your opponent's king wins the game, and the king can move
to and capture any adjacent square.</p>
<p>A king is in check when threatened with a possible capturing move by an opponent's piece. Any 
of your moves that cause your king to become in check are not allowed (Wisconsin Chess 
automatically removes these moves from your choices), and you have won when your opponent has no 
moves that take their king out of check.</p>
<p>The kings start on the same file (column) for each player, which is the fourth from the left of 
the black player or the fifth from the left of the white player. The starting squares for the 
<a href="/details?p=queen">queen</a> and king are mirrored for each player, and other pieces can
be called "queenside" or "kingside" to refer to which side of the army they're on.</p>
<p>When your king and <a href="/details?p=rook">rook</a> both haven't been moved yet, the squares 
between them are empty and unthreatened by your opponent's pieces, and your king isn't in check, 
then the king can do the castle move. Depending on the side, the king moves multiple squares and
jumps over the also moved <a href="/details?p=rook">rook</a>; this is difficult to describe in text,
but it will look like an extra side move when available in Wisconsin Chess.</p>`,

	`<h1>QUEEN</h1>
<p>The queen starts at the square to the left of the <a href="/details?p=king">king</a>, four
 squares from the player's left.</p>
<p>The queen can move and capture any distance up-down, left-right, and along the corner vectors.
 In other words, the queen can move like the <a href="/details?p=bishop">bishop</a> and
 <a href="/details?p=rook">rook</a> combined.</p>`,

	`<h1>ROOK</h1>
<p>Rooks start in the back row at the furthest left and right squares.</p>
<p>Rooks can move and capture any distance along the up-down and left-right vectors.</p>`,

	`<h1>BISHOP</h1>
<p>Bishops start in the back row adjacent to the <a href="/details?p=queen">queen</a> and
 <a href="/details?p=king">king</a>.</p>
<p>A bishop can move and capture any distance along the corner vectors.</p>`,

	`<h1>KNIGHT</h1>
<p>Knights start in the back row between the <a href="/details?p=bishop">bishop</a> and
 <a href="/details?p=rook">rook</a>.</p>
<p>Unlike other pieces, the knight can move over other pieces. The move is two squares forward and
 one to the left or right. This move or capture is up-down or left-right, for up to eight possible
 moves.</p>`,

	`<h1>PAWN</h1>
<p>An army has eight pawns. These pawns all start on the first row ahead of the other pieces.</p>
<p>The pawn moves forward toward the opponent by one square. If the pawn hasn't been moved yet then
 it can move either one or two squares forward.</p>
<p>Unlike other pieces, the pawn capture is different than the move. Captures are done in the cross
 square forward to the left or right. Pawns can still be moved forward even when they can capture.
</p>
<p>En passant is a capture that can only be done by pawns. If a pawn is moved forward two squares
 and passes through a square an opponent pawn could have captured it at, then only on their next
 turn the opponent can capture it by moving their pawn to the empty square.</p>
<p>If a pawn is moved entirely across the board to the last row then it must be promoted to a
 <a href="/details?p=queen">queen</a>, <a href="/details?p=rook">rook</a>,
 <a href="/details?p=bishop">bishop</a>, or <a href="/details?p=knight">knight</a>.</p>`,

	`<h1>WAR</h1>
<p>The war is a <a href="/details?p=pawn">pawn</a> except that it can't move two squares forward 
for the first move.</p>`,

	`<h1>FORM</h1>
<p>The form is a <a href="/details?p=pawn">pawn</a>.</p>`,

	`<h1>CONSTRUCTIVE</h1>
<p>The constructive is a <a href="/details?p=knight">knight</a>.</p>`,

	`<h1>CONFINED</h1>
<p>The confined is a <a href="/details?p=pawn">pawn</a>.</p>`,

	`<h1>ORIGINAL</h1>
<p>The original is a <a href="/details?p=bishop">bishop</a> except it only moves one square.</p>`,

	`<h1>IRRELEVANT</h1>
<p>The irrelevant is a <a href="/details?p=rook">rook</a> except it can only move five squares.
</p>`,

	`<h1>EVIDENT</h1>
<p>The evident is a <a href="/details?p=pawn">pawn</a> that only moves one square on its first turn
 and captures backwards.</p>`,

	`<h1>LINE</h1>
<p>The line is a <a href="/details?p=knight">knight</a> that can't move over pieces and only moves
 toward the opponent.</p>`,

	`<h1>IMPOSSIBLE</h1>
<p>The impossible is a <a href="/details?p=rook">rook</a> that only moves four squares.</p>`,

	`<h1>CONVENIENT</h1>
<p>The convenient is a <a href="/details?p=bishop">bishop</a> that only moves two squares.</p>`,

	`<h1>APPROPRIATE</h1>
<p>The appropriate is a <a href="/details?p=knight">knight</a> that can't move over pieces.</p>`,

	`<h1>WARP</h1>
<p>The warp is a <a href="/details?p=rook">rook</a> that only moves five squares.</p>`,

	`<h1>BRILLIANT</h1>
<p>The brilliant is a <a href="/details?p=knight">knight</a> that only moves toward the 
 opponent.</p>`,

	`<h1>SIMPLE</h1>
<p>The simple is a <a href="/details?p=rook">rook</a>.</p>`,

	`<h1>EXIT</h1>
<p>The exit is a <a href="/details?p=bishop">bishop</a> that can move over other pieces but only 
 moves three squares.</p>`,

	`<h1>IMPERFECT</h1>
<p>The imperfect is a <a href="/details?p=pawn">pawn</a> that can only capture the square behind
itself.</p>`,

	`<h1>DERANGE</h1>
<p>The derange is a <a href="/details?p=pawn">pawn</a> that can only capture the square behind
itself. The derange can otherwise move to any of the three squares ahead.</p>`,
}
