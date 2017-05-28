Backend Server for Othello Web Game written in Go

Temp Client While Real Client is under Construction:
  https://gist.github.com/Rogibb111/f045c7ebc6faec9d83b1282a5f2ef38b
  
  To use: Place contents of gist in static directory of source (make one if necessary)

API Documentation
=================

1. Routes

* /newgame

  This is the route to hit when a player wants to join a new game (potentially could be called joingame if a lobby is implemented)
  
  Arguments: none

  Return: `{ Full: bool, Player: int }`
  * Full: Tells whether or not the game already has 2 players. If it does, Full will be true
  * Player: If the game is not full, Player will indicates which player the current joiner has been designated(1 or 2)

* /move

  This is the route to hit when a player wants to make a move 

  Arguments: As form data - `row={row}&col={col}`
  Return: `{ Valid: bool }`
  * Valid: Tells whethor not a move was valid 

