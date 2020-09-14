// Package wichess is the entry point to the Wisconsin Chess host process by players connecting
// with web browsers.
//
// This package is mostly net/http handlers that use other packages for any complex functionality.
// The paths' URL string constants and JSON structures are exported for access by testing programs.
// HTML template names and data structs are exported for documentation when developing the webpages.
//
// The cmd folder has package main that imports package wichess.
//
// See README.md for more explanation of the Wisconsin Chess source code repository.
package wichess
