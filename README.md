Backend Server for Othello Web Game written in Go

Temp Client While Real Client is under Construction:
  https://gist.github.com/Rogibb111/f045c7ebc6faec9d83b1282a5f2ef38b
  
  To use: Place contents of gist in static directory of source (make one if necessary)

API Documentation
=================

1. **Routes**
  
  These are standard http routes that the server uses to execute player commands.

* /newgame - GET

  This is the route to hit when a player wants to join a new game (potentially could be called joingame if a lobby is implemented)
  
  Arguments: none

  Return: `{ Full: bool, Player: int }`
  * Full: Tells whether or not the game already has 2 players. If it does, Full will be true
  * Player: If the game is not full, Player will indicates which player the current joiner has been designated(1 or 2)

* /move - POST

  This is the route to hit when a player wants to make a move 

  Arguments: As form data - `row={row}&col={col}`
  Return: `{ Valid: bool }`
  * Valid: Tells whethor not a move was valid 

* /events - GET

  This is the route for using the javascript EventSource object with. It provides an event
  stream to the browser which will push all of the games notifications to the players.

  Arguments: none
  Return: none

2. **Events**
  
  These are the different types of events that come over the event stream. 

* move

  Data: `{ Row: Number, Col, Number, Player: Number, Turn: Number, Board: Array}`
  * Row: The row the space the piece was placed on
  * Col: The column of the space the piece was place on
  * Player: The player who placed the piece
  * Turn: The player who's turn it now is
  * Board: An array of arrays that make up the board. Each spot in each array can either be 0,1,2,3. O means the space is empty, 1 or 2 means a player 1 or 2 piece occupies that space, and 3 means that the space is empty, but a valid spot for the current player to put his piece.

* start

  Data: `{ Turn: Number }`
  * Turn: The player who's turn it now is

* end

  Data: `{ Winner: Number }`
  * Winner: The player who won.

