# Dead Simple Tetris with the Dead Simple Game Engine

This project aims to implement a copy of Tetris Classic in Golang. It uses [Ebitengine](https://ebitengine.org/) as a game engine.

This project was originally developed as an assignement from Software Engineering II from the System's Information course in Federal University of Minas Gerais, under Professor Andre Hora (https://homepages.dcc.ufmg.br/~andrehora/teaching/es2.html)

### Participants
1 - Daniel Neiva da Silva

### About the program
The program implements the basic features of [Tetris Classic](https://en.wikipedia.org/wiki/Tetris_Classic). It has 7 different types of pieces, each consisting of 4 tiles.
The game allows rotating the pieces, moving them horizontally and accellerating their fall. The game goes on ultil there are pieces all the way up to the start of the board.

The program is implemented in Golang, and uses Ebitengine, a game engine for Golang game development. The game engine provides key features for game development, such as drawing images into the screen, game loop and more.