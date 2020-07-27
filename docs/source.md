# SOURCE CODE WALKTHROUGH

Wisconsin Chess is developed in the macOS operating system on a MacBook Pro ([apple.com/macbook-pro](https://www.apple.com/macbook-pro/)). The source code organization is optimized for use with the macOS Terminal and the MacVim editor ([macvim-dev.github.io](https://macvim-dev.github.io/macvim/)) without plugins or syntax highlighting. Symbol names are intended to be easily connected by the developer to filenames navigated to with the macOS Finder. Sometimes the grep Terminal command line tool is helpful for finding instances of a symbol.

Developing on macOS requires downloading Xcode from the App Store and installing the developer command line tools by opening it once and accepting the license.

Testing is done with the Ubuntu Server ([ubuntu.com](https://ubuntu.com/download/server)) and FreeBSD ([freebsd.org](https://www.freebsd.org/)) operating systems. These have a similar terminal and filesystem language to macOS which makes interaction with the source code almost the same.

Wisconsin Chess uses the Git version control system ([git-scm.com](https://git-scm.com)) to keep details of changes made to the source code. The Wisconsin Chess folder is called a repository, and Git can be used to understand or get the history of every file in it. Whenever I complete some work, usually 100 to 1000 lines of code, I commit it to the repository with an explanation of the intention of the changes. You can view a timeline of these changes with the ```git log``` command when in the Wisconsin Chess folder.

The use of Git I find significantly improves ability to fix mistakes later when the code I'm looking at isn't fresh on my mind, or when a recent change causes a new mistake Git can show exactly what could have caused it.

## Programming Languages

A variety of languages are used for Wisconsin Chess.

The host program that serves players in the local network is in Go ([golang.org](https://golang.org)) which uses a dialect of SQL to interact with PostgreSQL ([postgresql.org](https://postgresql.org)) to save persistent information. Test programs are also in Go.

Players get a website from the host through the network. This program is downloaded by their web browser as a set of HTML and CSS ([developer.mozilla.org](https://developer.mozilla.org/en-US/docs/Web/HTML)) files that describe the look of the webpages along with JavaScript ([developer.mozilla.org](https://developer.mozilla.org/en-US/docs/Web/JavaScript)) files that size the webpage for the player's browser dimensions and responds to the player's interactions by causing the browser to communicate more with the host.

The website also downloads chess piece images from the host. The geometry and look of these pieces and the framing of the picture are described in the POV-Ray ([povray.org](https://povray.org)) language. A program to split a chess board image into images of the individual squares is in Go.

Build, configuration, and installation is done with Bash ([gnu.org](https://www.gnu.org/software/bash/)) sh scripts.

JSON ([json.org](https://www.json.org/json-en.html)) is used to describe the database configuration, it's used for test cases in the test folder, and whenever more than a string or two is needed interaction between the host and web browsers is done in JSON.

## Repository Organization

The top level of the repository folder is the Go source files used to build the host program wichess with ```go build```.

Bash scripts are also here. InstallLocalMacOS.sh automatically sets up the repository folder to host Wisconsin Chess on MacOS. CreateDB.sh creates the database folder for persistent information to be written to while Wisconsin Chess runs using RunLocal.sh. DB.sh is a simple command to connect to the database for manual interaction with saved information.

Some repository information is here, including the COPYRIGHT file and README.md with an overview description.

postgres_tables.sql is used in creation of the database folder to describe to PostgreSQL how saved information is organized. This file is an important reference for understanding the database and for making new SQL interactions in wichess.

### wichess Source Code

The .go files at this top folder level are all part of ```package main```, and main.go is where the host program starts. In main.go only database and HTTP interactions are setup, because that's all the host does.

A metaphorical table of contents for the HTTP paths is http.go, and the implementation of these are in files that start with http_.

The database is organized as a table of games and a table of players. The HTTP interactions often cause database interations which are in files that start with players_ or game_ depending on which table is being accessed.

If a function is complex or logically separate of the database or HTTP then often it will be in another file with a name that is intended to clearly describe the content.

Some folders contain wichess-related code: piece has ```package piece``` which describes the pieces for these programs and as sentences for players, rules has ```package rules``` which calculates possible moves and results of moves with the minimum information needed, test has ```package test``` which is used to create test cases to quickly recreate board positions that broke something and to verify new changes don't break old in-game features, and html has HTML templates that are filled out then sent to players when requested.

### Test

Some testing of the ```package rules``` logic is done in the test folder by running ```go test``` there. The cases for these tests can be viewed and added to with the Go web program in the test/builder folder through a web browser.

### Web

wichess hosts the entirety of the web folder for access by players. The HTML webpages get images, sound, fonts, CSS, and JavaScript modules from this folder.

layout.js is how the webpages are visually organized. With this JavaScript library the webpage resizes and varies for the aspect ratio and appears like a dashboard by not scrolling.

Each webpage has a primary JavaScript module file: game.js, index.js, match.js, and reward.js, and if there is more code needed then a folder with the same name is there. Code shared between webpages is files alongside these, such as pieceDefs.js with the code names for each of the special pieces.

The JavaScript is an attempt to use state of the art techniques, including fetch, modules, no line termination semicolons, and avoidance of libraries like jQuery.

The web/img folder is created when the images are rendered in the img folder.

### Img

The img folder has a set of sh scripts that call the POV-Ray command line tool to render pictures of all of the chess pieces.

Board.sh makes two perspective images of the board for a piece, one for white and one for black, which can then be cut into 64 individual squares with cut.go.

Look.sh makes the image used for a piece on its details webpage.

Pick.sh makes the image used in the collection, for rewards, and in the army.

Take.sh makes the side orthographic images for black and white that are shown as part of the list of captured pieces in a game.

Render.sh calls all of these scripts for all of the pieces then copies the results to the web/img folder. This script is useful to see exactly is done for a complete render of the Wisconsin Chess images. It can take hours to run this script.

Each basic piece has a .inc file that describes its geometry. The special pieces start by including the basic piece's file then adds to it. The scripts choose which .inc to use depending on the code name provided to it.

Other input files for POV-Ray are board.inc, look.inc, pick.inc, and take.inc that describe the scripts' scenes, and materials.inc describes the generic, black, and white textures.

The cyl.pov file is a reference for the maximum dimensions of pieces to fit on the board but is only a loose guideline, piece designs are mostly done by rendering them then adjusting until they look good. cyl.pov is also described in docs/cylinders.md.

### Other Folders

The database folder is created to hold the PostgreSQL data when CreateDB.sh is called during installation.

The docs folder has this document and any other articles describing the Wisconsin Chess implementation.

The licenses folder has any license documents needed for content not made by me.

The screenshots folder is a pictures of the game in development all the way back to September 2017 when the board was first seen.

The vendor folder is for Go to build wichess and includes code made by other people.