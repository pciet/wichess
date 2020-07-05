package piece

var DetailsHTML = []string{
	"",
	`<h1>KING</h1>
<p>Capturing your opponent's king is the goal of the game. Kings start in the back row five squares
 from the player's left.</p>
<p>The king can move to and capture any adjacent square.</p>
<p>When the king and a friendly <a href="/details?p=rook">rook</a> both haven't been moved yet,
 the squares between them are empty and unthreatened, and the king isn't in check, then the king
 can do the castle move.</p>
<p>A king is in check when threatened by capture from an opponent's piece. Any moves by any piece
 that cause your king to be in check are not allowed, and you have won when the opponent has no
 moves that takes their king out of check.</p>`,

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
<p>The war piece is a <a href="/details?p=pawn">pawn</a> except that it cannot move two squares
 forward for its first move.</p>`,

	`<h1>FORM</h1>
<p>The form piece is a <a href="/details?p=pawn">pawn</a>.</p>`,

	`<h1>CONSTRUCTIVE</h1>
<p>The constructive is a <a href="/details?p=knight">knight</a>.</p>`,

	`<h1>CONFINED</h1>
<p>The confined is a <a href="/details?p=pawn">pawn</a>.</p>`,
}