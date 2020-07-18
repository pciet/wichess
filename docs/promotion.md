# PROMOTION

In chess promotion is when a pawn reaches the last rank on the other side of the board. When that happens the player chooses to replace it with a knight, bishop, rook, or queen.

The promotion game state/condition is a special case. Unlike check, which in Wisconsin Chess is just an indication to the player without any gameplay changes, promotion affects the turn structure by requiring a second action after moving. Additionally, for the guard pawn, reverse promotion is possible where a promotion choice is needed without a move being made by the player.

This document specifies the variety of communication ordering between the webpage and host for promotion.

An arrow ```->``` toward the host (the wichess program) indicates an HTTP request, and the arrow away from the host toward a player (game.js in a web browser or the opponent) is the HTTP response. A lone arrow away from the host is a WebSocket message which is asynchronous, unlike HTTP which requires a request for a response from the host.

## Normal Communication Sequence

```
game.js        host       opponent
        alert
          <-
          
        /moves
          ->
          <-
          
        /move
          ->
                     alert
                      ->
          <-
                     normal sequence
```

The alert is caused by an opponent move and has the resulting changes to the board. A request of all moves is followed, then a move is chosen by the player and sent to /move.

The /move response might be before or after the opponent alert, but the following opponent call to /moves is guaranteed to read the game updated in the database by the /move call.

## Promotion and Reverse Promotion

```
game.js        host        opponent
        alert
          <-
          
        /moves
          ->
          <-
          
        /move
          ->
                     alert
                     -> (wait)
          <- (prom)
          
        /move
          -> (prom)
                     alert
                      ->
          <-
                     normal sequence
```

### Reverse

```
game.js        host        opponent
        alert
          <-
          
        /moves
          ->
          <-
          
        /move
          ->
                      alert
                       ->
          <-
          
                      /moves
                       <-
                       -> (prom)
                       
                      /move
                       <- (prom)
        alert
          <- (wait)
                       -> (continue)
                     
                       normal sequence

```

The promotion signals in paranthesis are special cases that are looked for by the programs.

* When /move responds with a promotion needed signal then a promotion choice sent to /move follows instead of waiting for an alert.
* A second alert is waited for when an alert signals to wait instead of continuing to the next move.
* When /moves responds with a promotion needed signal then a promotion choice is sent to /move.
* When /move responds with a continue signal then a normal move process follows instead of waiting for an alert.

## Promotion from Initialization

Since players can leave and return to a game at any time the state of promotion must also be part of the webpage initialization via /moves signaling.

```
game.js        host        opponent
        /boards
          ->
          <-
          
        /moves
          ->
          <- (prom)
   
        /move
          -> (prom)
                     alert
                      ->
          <-
                     normal sequence
```

### Reverse


```
game.js        host        opponent
        /boards
          ->
          <-
          
        /moves
          ->
          <- (prom)
   
        /move
          -> (prom)
                     alert
                      -> (wait)
          <- (continue)
  normal sequence
```

No added webpage logic is needed for the initialization cases since the /move continue response and alert wait signal are reused.