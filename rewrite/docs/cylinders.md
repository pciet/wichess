# PIECE BOUNDING CYLINDERS

The look of the Wisconsin Chess pieces is defined in the img folder's .inc files. These piece designs are constrained by the maximum bounds in the cyl.pov file, which renders as a board of stacked cylinders.

These cylinders are standard dimensions; if a piece fits in the cylinders then it will display correctly in the perspective view for all 64 squares. This view is defined in img/board.inc.

These are the dimensions of the bounding cylinders from the top one down:

```
height radius
 1.6    0.5
  1     0.8
  2     1.2
  2     2.2
  3      3
```

Pawn kinds use the first three cylinders, specialists (rook, bishop, knight) also use the fourth, and the two sovereign pieces also use the fifth.