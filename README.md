# Proposed architecture

A proposal of the architecture to be used

## API
A way to initialize a game, with the game type and other needed information. This should return a session id that should be used to create a websocket connection to the server

Could be made using fiber or we could roll our own using go's http library

## Websocket connection
Real time communication between the server and clients, determines what session the belong to using a session id and game type

Could be handled by the session handler or a connection manager

## Session handler
Handles a running game and all data associated with it, there are multiple version of them one for each game, all implementing the same interface.

Could be handled by a session manager that determines what session need to be killed and other needed functionality.