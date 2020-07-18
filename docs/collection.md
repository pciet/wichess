# COLLECTION

Choosing pieces for the army before a game starts has some nuance, described here.

On the index webpage the player is shown two boxes with special pieces, the left and right picks. These are randomly selected, different than each other, and are persistent for the player until used.

When clicked the piece is added to the army, and the pick slot is replaced with a new special piece. Once the game is completed then the picked piece is added to the player's collection. Both picks can be included in an army. If the game is conceded then the player doesn't get the piece.

Some modes allow use of collection pieces. There are 21 slots for these, and they are gotten from completing games with pick pieces. If a collection piece is taken in a game then the piece is lost from the player's collection. Collection pieces cannot be used in concurrent games; the index shows this by making the piece unselectable and greyed out somehow.

## Index Army Logic

The index webpage requests an army when a mode's play button is pressed. The array of 16 pieces are sent in this format:

- Indexing starts at 0 at the top left (pawn) and continues right through the pawn row to 7.
- 8 starts at the left rook, queen is 11, king 12, to 15 for the right rook.
- One integer describes the requested piece for each array index.
- NoSlot (0) indicates a basic piece request.
- LeftPick (-1) indicates use of the left random special piece.
- RightPick (-2) indicates use of the right random special piece.
- A collection slot index (starting at 1 which matches SQL/postgres array indexing) indicates use of the corresponding collection piece.
