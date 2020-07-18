# WISCONSIN CHESS WEB BROWSER USER INTERFACE

The webpage look depends on window size, done with the layout.js engine and input specific to each webpage.
More layout calculation is done beyond layout.js to make the board or army picker maximum dimensions, and other layout features are added by each webpage for things like piece movement.
Communication with the host is done both with HTTP and WebSockets for piece moving request-response and asynchronous alerts for when an opponent moves.