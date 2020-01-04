Wisconsin Chess is a computer network program to play chess with others. Some time and ability is needed to install it and I've done my best to explain that in the Setup section of this readme. First, here's a description of the game.

![Screenshot1](https://github.com/pciet/wichess/blob/master/screenshots/Screen%20Shot%202018-01-09%20at%201.56.21%20PM.png)

---

Before starting a match you can replace standard chess pieces in your army with new kinds of pieces you get one of when completing each timed game. These extended-rules pieces are lost from your collection when taken by an opponent.

Modes include an easy practice computer/AI opponent, up to six parallel untimed friend games, or skill-based matchmaking with a 15 minute or 5 minute clock.

Each person has a username and password that keeps your win/draw/loss record, set of pieces, and lets you setup the friend games by typing their username.

If you close your web browser or your computer glitches you are still able to return to your active matches later.

I began to explore making this game in Portland, Oregon coffee shops in 2014, and the new piece rules were first implemented in an experimental C++/Ogre3D prototype later that year. Other projects and experiences then happened before I implemented this Go/web version in the second half of 2017 which was a full-time effort for six months.

The new piece rules aren't a carefully crafted game design, but I still think it's fun and worth playing if you aren't being too competitive. Normal chess rules are included so if you don't use the new pieces then perhaps this could just be a useful chess game. In the future I hope to innovate in this game design space of collectible pieces with up to two special characteristics and multiple movesets.

The program has had significant testing which helped me fix obscure mistakes, and the user interface adapts to aspect ratio making it useful in web browsers across devices like tablets, laptops and smartphones.

I believe chess is a good intellectual training game and I hope you and others get value from my work. Most importantly, have fun!

---

For now it's intended for adults; despite in my understanding there not being anything particularly adult about the game perhaps there are unknown or unexpected effects on a person. Wisconsin Chess is not age rated or approved by any organization like the ESRB.

This repository is not licenced beyond the base GitHub license. I don't think I included any easter eggs, practical jokes, backdoors, or intentional mistakes, I hope what you see is what you get.

If you decide to play Wisconsin Chess you should be aware that no network communication encryption is used, so you shouldn't use usernames or passwords you'd care about if they were seen by others. The passwords are encrypted in the database but the cryptography implementation is an old version of ```golang.org/x/crypto``` and may not be free of mistakes that have been fixed after 2017. **Because of this weak security I recommend that you don't run ```wichess``` on any computer that needs to be secure, and I recommend that you don't play Wisconsin Chess over the Internet.**

A video walkthrough of the source code is at [youtu.be/yZn2jU8eCUo](https://youtu.be/yZn2jU8eCUo).

## Setup

For now the only way to play Wisconsin Chess is to compile it from this source code, which means you'll need to use the command line. This implementation has only been made to work with the macOS or GNU/Linux (Ubuntu Server) operating systems, using the Bash shell.

A simple configuration has the web browser user interface you'll play from delivered by a separate host computer on your network. This host is also responsible for keeping memory of games so that you can come back to them later, for keeping player records, for communicating to multiple client web browsers concurrently, and for other things.

Here's a summary of what must be done to initialize the host, followed by the details:

1. Compile ```wichess```.
2. Install POV-Ray and render the piece images.
3. Install and configure PostgreSQL.

```main.go``` and the other Go-language program files at the top folder level is the program that manages the host, compiled simply by the ```go build``` command from the Go programming language tools which you must install separately ([golang.org](https://golang.org)). This will make a program file called ```wichess``` in the folder. The ```web``` directory is all that's needed by this program.

The chess pieces shown on the board are at a different perspective for each square, and these images must be rendered and saved to the ```web``` directory. POV-Ray must be installed ([povray.org](http://www.povray.org)); I installed it with Homebrew on macOS ([brew.sh](https://brew.sh)):

```
brew install povray
```

Then call my script:

```
cd graphics
./render.sh
```



The resulting images will use about 0.75 GB of drive space and may take more than two hours to generate, but the tradeoff of using your computer to render means this repository can be relatively small.

A PostgreSQL database holds all persistent information and must be separately installed  on the host ([postgresql.org](https://www.postgresql.org)) then initialized with the ```create_postgres_tables.sh``` script which you can change along with ```dbconf.json``` to match how you want to configure PostgreSQL.

If you compile and launch ```wichess``` and no error is shown then the host is ready. Connect to the host's network address at port 8080 with a web browser (it will be something like entering an address like ```http://192.168.1.75:8080``` into the web browser's address bar) then enter a username and password you'd like to use in the form that should appear. If this login page reloads when you try to continue then you might need a longer password, the username is already in use, or you entered the wrong password.

## Status

This user interface implementation has only been tested on macOS with the Chrome, Firefox, and Safari web browsers, and on iOS (iPhone and iPad) with the Safari web browser. A layout issue exists with Firefox but it did still work, and as web browsers are improved new problems may appear which I'm planning to keep up with.

I made electronic music for the game and these audio files were included in this repository for awhile. To make the download time more reasonable I've removed this very large directory from the repo history along with the rendered piece images. I'm currently applying my latest music mastering skills to improve the presentation of these recordings; before I'm done with that you can hear my favorites in my June 2019 "Typing Choice Tree" album at [matthewjuran.bandcamp.com/album/typing-choice-tree](https://matthewjuran.bandcamp.com/album/typing-choice-tree).

## Credits

```wichess``` includes a snapshot of ```github.com/gorilla/websocket```, see ```vendor/github.com/gorilla/websocket/LICENSE``` for the copyright notice and license.

```wichess``` includes a snapshot of ```github.com/lib/pq```, see ```vendor/github.com/lib/pq/LICENSE.md``` for the copyright notice and license.

```wichess``` includes a snapshot of ```golang.org/x/crypto```, see ```vendor/golang.org/x/crypto/LICENSE``` for the copyright notice and license.

The knight piece was copied from a set of POV-Ray chess piece designs by Ted Fisher and James Garner. The original license text is in ```graphics/knight.inc``` followed by an explanation of my changes.
