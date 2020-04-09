# INSTALLATION SETUP GUIDE

For now the only way to play Wisconsin Chess is to compile it from this source code, which means you'll need to use the command line.

This implementation has only been made to work with the macOS or GNU/Linux (Ubuntu Server) operating systems, using the Bash shell. This guide is for macOS but replacing the Homebrew commands with the Ubuntu equivalent will make it work there.

The game operates by displaying the user interface to players in a web browser by delivering it from a host program through an Internet Protocol network which is the same way all web browsing works. This webpage user interface continues to communicate with the host whenever a player makes a move.

When a web browser requests a game webpage the computer program that responds is the host program called ```wichess``` that also computes the result of chess moves and keeps a persistent memory of things like player win/loss/draw records and active games. This memory is done by PostgreSQL which is a separate program you'll need to get.

This software architecture should limit cheating; there's no obvious way to modify the web browser or webpage to cheat, and you can limit access to the host computer to just the HTTP ```wichess``` responds to.

Players also don't need to install yet another app, all that's needed is a modern web browser.

## Table of Contents

1. [Local macOS](#localmacos)

<a name="localmacos"></a>
## Local macOS

Here's an overview of the steps to install Wisconsin Chess on a macOS computer, followed by the details:

1. Compile ```wichess```.
2. Install POV-Ray and generate the piece images.
3. Install and configure PostgreSQL.

Install the Xcode app from the App Store, and open it once to accept the license and install the Xcode command line tools.

Install the Go programming languages tools as described at [golang.org](golang.org).

Using Terminal (Applications/Utilities/Terminal.app) clone this repository in the folder you'd like to be the workspace for ```wichess```. We'll be generating picture files into subfolders here and saving persistent database files, so the home folder probably makes the most sense for this. The default Go folder layout might be what you want to do.

```
git clone https://github.com/pciet/wichess ~/go/src/github.com/pciet/wichess
cd ~/go/src/github.com/pciet/wichess
```

I've made a Bash script to do the rest for you (including automatically installing POV-Ray and PostgreSQL) if you have also installed Homebrew as described at [brew.sh](https://brew.sh). Do that then call the install script.

```
./InstallLocalMacOS.sh
```

The last thing this script will do is generate the piece images which might take more than an hour and 100 MB of drive space.

Start ```wichess``` with ```RunLocal.sh``` which also starts and stops PostgreSQL for you automatically.

```
./RunLocal.sh
```

If you want other devices to be able to connect to your macOS computer then press "Allow" when asked if you want ```wichess``` to accept incoming network connections, or in System Preferences -> Security & Privacy -> Firewall -> Firewall Options set ```wichess``` to "Allow incoming connections".

You can stop ```RunLocal.sh``` safely (meaning PostgreSQL will be turned off correctly) with the control+c key combination.

While ```RunLocal.sh``` is active, with a web browser connect to ```http://localhost:8080``` then enter a username and password you'd like to use in the form that should appear. If this login page reloads when you try to continue then you might need a longer password, the username is already in use, or you entered the wrong password.

For other devices in the network you'll use the IP address of the host computer instead of ```localhost```. This can be found in System Preferences -> Network or in the output of the ```ifconfig``` command. The address you'll then use for the web browser will likely look something like ```http://192.168.2.20:8080```.

If you would prefer to do the installation steps manually then I've described that next.


### Script Details

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

If you are using PostgreSQL in a different configuration then the ```dbconfig.json``` file and ```func InitializeDatabaseConnection``` in ```database.go``` can be changed for it.

