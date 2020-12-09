# Chess
Implementation of the game of chess. Just for fun/practice go.

Should know the rules, eventually will use to write a simple chess engine.


## TODO

### Game rules
- Check
- End of game/Checkmate
- Castling
- Track whether the king/rook has moved for castling
   Maybe a simple "castleQueensideAllowed" and "castleKingsideAllowed" variables that are set to false when those pieces move?
- Prevent moving if a piece is intervening for non-knights
- Pawn promotion
- En Passant


### Gameplay
- Translate algebraic notation to Move
- Be able to list valid moves for a given piece
- Allow player to resign
- Track "score" (how many pieces captured)

### Engine
- First pass: list all moves, pick a random one
- Second pass: Determine score after moves, pick one with largest advantage
- Minimax?
- Parallelize minimax?

### Testing
- Fix current failure in pieces on board
- Consistent validation code
- Walk through games and make sure they end as expected
