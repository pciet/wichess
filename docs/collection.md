# COLLECTION

Choosing pieces for the army before a game starts has some nuance, described here.

On the index webpage the player is shown two boxes with special pieces, the left and right picks. These are randomly selected, different than each other, and are persistent for the player until used.

When clicked the piece can be added to the army by a click on it. Once a game is started then that random pick slot will be replaced with a new piece.

If a game is completed (not conceded) then a reward screen is shown where the left and right pieces, if selected before the game started, can optionally be added to the collection. An additional random reward piece is also presented here. If a rewarded piece is not added to the collection then it's lost forever.

The collection has 21 slots and can be used to create custom armies before starting games. Pieces are unchanged unless the player decides to replace pieces with new rewards.

## Index Army Logic

The index webpage requests an army when a mode's play button is pressed. The array of 16 pieces are sent in this format:

- Indexing starts at 0 at the top left (pawn) and continues right through the pawn row to 7.
- 8 starts at the left rook, queen is 11, king 12, to 15 for the right rook.
- One integer describes the requested piece for each array index.
- NoSlot (0) indicates a basic piece request.
- LeftPick (-1) indicates use of the left random special piece.
- RightPick (-2) indicates use of the right random special piece.
- A collection slot index (starting at 1 which matches SQL/postgres array indexing) indicates use of the corresponding collection piece.
