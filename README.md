Wisconsin Chess is a computer network program to play chess with others, where tens or hundreds of people can be playing via one computer. Some time and ability is needed to install it and I've done my best to explain that in the Installation Setup Guide section of this readme. First, here's a description of the game.

![Screenshot1](https://github.com/pciet/wichess/blob/master/screenshots/Screen%20Shot%202018-01-09%20at%201.56.21%20PM.png)

---

You might not be impressed by the look of Wisconsin Chess. It's a work in progress that I believe has potential which I hope to build it into. Despite lacking equal refinement to the games you might be used to I do think it is a worthy chess toolbox now and a complete early version.

Before starting a match you can replace standard chess pieces in your army with new kinds of pieces you get one of when completing each timed game. These extended-rules pieces are lost from your collection when taken by an opponent.

Examples are locking movement of adjacent opponent pieces, guards that automatically respond to opponent moves with an adjacent take, and detonators that take all surrounding pieces when taken. The game manages check of the king and showing available moves, so illegal moves are not possible even in complex situations you can find yourself in. If moves you expect aren't shown then look at your king and possible check and checkmates. I believe the computer calculation of available moves is necessary to play this kind of chess variation in a reasonable amount of time.

Modes include an easy practice computer/AI opponent, up to six parallel untimed friend games, or matchmaking with a 15 minute or 5 minute clock.

Each person has a username and password that keeps your win/draw/loss record, set of pieces, and lets you setup the friend games by typing their username.

If you close your web browser or your computer glitches you are still able to return to your active matches later.

I began to explore making this game in Portland, Oregon coffee shops in 2014, and the new piece rules were first implemented in an experimental C++/Ogre3D prototype later that year. Other projects and experiences then happened before I implemented this Go/web version in the second half of 2017 which was a full-time effort for six months.

The new piece rules aren't a carefully crafted game design (an example is the extended knight has a two move checkmate), but I still think it's fun and worth playing if you aren't being too competitive. Normal chess rules are included so if you don't use the new pieces then perhaps this could just be a useful chess game.

In the future I hope to be innovative with the game architecture of collectible pieces that have up to two characteristics and multiple movesets, and I hope to improve the look of the game and really eliminate imperfections.

The program has had significant but narrow testing which helped me fix obscure mistakes mostly in the host program's concurrent and chessboard logic, and the user interface adapts to aspect ratio making it useful in web browsers across devices like tablets, laptops and smartphones. Much more testing will be necessary to be sure it works on as many devices with web browsers as possible and for efficiently adding new kinds of pieces without introducing new mistakes.

I believe chess is a good intellectual training game and I hope you and others get value from my work. But most importantly, have fun!

---

For now Wisconsin Chess is intended for adults; despite in my understanding there not being anything particularly adult about the game perhaps there are unknown or unexpected effects on a person. Wisconsin Chess is not age rated or approved by any organization like the ESRB.

This repository is not licenced beyond the base GitHub license. I don't think I included any easter eggs, practical jokes, backdoors, or intentional mistakes, I hope what you see is what you get.

If you decide to play Wisconsin Chess you should be aware that no network communication encryption is used, so you shouldn't use usernames or passwords you'd care about if they were seen by others. The passwords are encrypted in the database but the cryptography implementation is an old version of ```golang.org/x/crypto``` and may not be free of mistakes that have been fixed after 2017. **Because of this weak security I recommend that you don't run ```wichess``` on any computer that needs to be secure, and I recommend that you don't play Wisconsin Chess outside a local network.**

A video walkthrough of the source code is at [youtu.be/yZn2jU8eCUo](https://youtu.be/yZn2jU8eCUo).

## Installation Setup Guide

For now the only way to play Wisconsin Chess is to compile it from this source code, which means you'll need to use the command line. This implementation has only been made to work with the macOS or GNU/Linux (Ubuntu Server) operating systems, using the Bash shell. This guide is for macOS but replacing the Homebrew commands with the Ubuntu equivalent will make it work there.

The game operates by displaying the user interface for you in a web browser by delivering it from a host program through an Internet Protocol network, the same way all web browsing works. This webpage user interface continues to communicate with the host whenever a chess move is made.

When a web browser requests a game webpage the computer program that responds is the host program called ```wichess```, which also computes the result of chess moves and keeps a persistent memory of things like player win/loss/draw records and active games. This memory is done by PostgreSQL which is a separate program you'll need to get.

Maybe an illustration of the value of this software architecture for this game is that if a player doesn't have access to the computer running ```wichess``` then there is no obvious way for them to modify their browser or webpage user interface program to cheat, but they can still play.

Here's an overview of the steps to install Wisconsin Chess on a macOS computer, followed by the details:

1. Compile ```wichess```.
2. Install POV-Ray and generate the piece images.
3. Install and configure PostgreSQL.

### Local macOS

Install the Xcode app from the App Store, and open it once to accept the license and install the Xcode command line tools.

Install the Go programming languages tools as described at [golang.org](golang.org).

Using Terminal (```Applications/Utilities/Terminal.app```) clone this repository in the folder you'd like to be the workspace for ```wichess```. We'll be generating picture files into subfolders here and saving persistent database files and the home folder probably makes the most sense for this. The default Go folder layout might be what you want to do.

```
git clone https://github.com/pciet/wichess ~/go/src/github.com/pciet/wichess
cd ~/go/src/github.com/pciet/wichess
```

I've made a Bash script to do the rest for you (including automatically installing POV-Ray and PostgreSQL) if you have also installed Homebrew as described at [brew.sh](https://brew.sh). Do that then call the install script.

```
./InstallLocalMacOS.sh
```

The last thing this script will do is generate the piece images which might take more than an hour and 100 MB of drive space.

Start ```wichess``` with this run script which also starts and stops PostgreSQL for you automatically:

```
./RunLocal.sh
```

If you want other devices to be able to connect to your macOS computer then press "Allow" when asked if you want ```wichess``` to accept incoming network connections, or in System Preferences -> Security & Privacy -> Firewall -> Firewall Options set ```wichess``` to "Allow incoming connections".

You can stop the ```RunLocal.sh``` script safely (meaning PostgreSQL will be turned off correctly) with the control+c key combination.

While ```RunLocal.sh``` is active, with a web browser connect to ```http://localhost:8080``` then enter a username and password you'd like to use in the form that should appear. If this login page reloads when you try to continue then you might need a longer password, the username is already in use, or you entered the wrong password.

For other devices in the network you'll use the IP address of the host computer instead of ```localhost```. This can be found in System Preferences -> Network or in the output of the ```ifconfig``` command. The address you'll then use for the web browser will likely look something like ```http://192.168.2.20:8080```.

If you would prefer to do the installation steps manually then I've described that next.


#### Script Details

Build the ```wichess``` program file with the Go tool.

```
go build
```

Generate the chess piece images with POV-Ray ([povray.org](http://www.povray.org)). If you don't want to also make that from source code then a way is with the Homebrew system which I'll assume you're using. Install Homebrew from [brew.sh](https://brew.sh) then install POV-Ray.

```
brew install povray
```

Execute my generating script.

```
cd graphics
./render.sh
```

These may take over an hour to generate, but the tradeoff is that this repository is relatively small without a history of these image files included.

While the images are generating you can install PostgreSQL ([postgresql.org](https://www.postgresql.org)).

```
brew install postgresql
```

Initialize a database folder and start the PostgreSQL program.

```
cd ..
initdb database/data
pg_ctl -D database/data -l database/log.txt start
```

Configure the database with the data layout ```wichess``` uses.

```
createdb test
psql -d test -f postgres_tables.sql -h localhost -p 5432
```

You can start ```wichess``` now. Instead of just calling ```./wichess``` a better way to run it is the ```RunLocal.sh``` script which will manage the PostgreSQL program for you. First stop PostgreSQL.

```
pg_ctl -D database/data stop
```

Then call the script to start Wisconsin Chess.

```
./RunLocal.sh
```

If no error is shown then it's ready, follow the last instructions in the previous section.

If you are using PostgreSQL in a different configuration then the ```dbconfig.json``` file and ```func initializeDatabaseConnection``` in ```database.go``` can be changed for it.

## Work Status

This user interface implementation has only been tested on macOS with the Chrome, Firefox, and Safari web browsers, and on iOS (iPhone and iPad) with the Chrome and Safari web browsers. An army selection display problem exists with Firefox but it did still work. I haven't been able to get my Android phone to connect yet for some reason but it might work in your network, and the Chrome web browser is the most likely to work.

I made electronic music for the game and these audio files were included in this repository for awhile. To make the download time more reasonable I've removed that very large directory from the repo history along with the rendered piece images. I'm currently applying my latest music mastering skills to improve the presentation of these recordings; before I'm done with that you can hear my favorites in my June 2019 "Typing Choice Tree" album at [matthewjuran.bandcamp.com/album/typing-choice-tree](https://matthewjuran.bandcamp.com/album/typing-choice-tree).

A rewrite is being done with improvements like making the SQL easier to read, choosing more accurate identifiers, simplifying HTTP functionality and improving error handling, reducing code length and making smaller source files with easier to understand filenames, removing unnecessary packages like match, splitting the web page JavaScript into multiple files, and more. The ```rewrite``` folder has a partial rough draft of this where I started with the HTTP handlers.

## Credits

```wichess``` includes a snapshot of ```github.com/gorilla/websocket```, see ```vendor/github.com/gorilla/websocket/LICENSE``` for the copyright notice and license.

```wichess``` includes a snapshot of ```github.com/lib/pq```, see ```vendor/github.com/lib/pq/LICENSE.md``` for the copyright notice and license.

```wichess``` includes a snapshot of ```golang.org/x/crypto```, see ```vendor/golang.org/x/crypto/LICENSE``` for the copyright notice and license.

The knight piece was copied from a set of POV-Ray chess piece designs by Ted Fisher and James Garner. The original license text is in ```graphics/knight.inc``` followed by an explanation of my changes.
