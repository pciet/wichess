# TECHNICAL INSTALLATION GUIDE

For now the only way to play Wisconsin Chess is to compile it from this source code, which means you'll need to use the command line.

The use of the host program, ```wichess```, is limited by the operating system and processor architectures supported by the Go programming language, and I am only supporting a subset of those:
 
* macOS on x86-64 (Apple computers)
* FreeBSD on x86-64 (DIY PC)
* Debian GNU/Linux on ARMv7 (Texas Instruments BeagleBone Black)

I use my FreeBSD system to cross-compile for the BeagleBone, see docs/bbb_sd.md for information about that. The rest of this guide is for macOS using Homebrew, but there is also some generalized descriptions that applies to the other systems. See docs/freebsd.md for FreeBSD-specific instructions.

## Host Software Architecture

Wisconsin Chess operates by delivering a web browser user interface to players. The computer that delivers the interface, the host, also manages communication between players and does the calculations of requested chess board moves. The program on the host is called ```wichess``` and is made of all of the Go programming language files in the repository. ```wichess``` saves information that persists across host reboots to a folder called mem.

This software architecture should limit cheating because there's no obvious way to modify the web browser or webpage to cheat and you can limit access to the host computer. Players also don't need to install yet another app; all that's needed is a modern web browser like Chrome, Firefox, or Safari, and the IP address of the host.

## macOS Installation

Installing Wisconsin Chess is a two step process. The ```wichess``` program is built, and the piece images are made using POV-Ray. These are the details of that process on macOS.

Install the Xcode app from the App Store, and open it once to accept the license and install the Xcode command line tools.

Install the Go programming languages tools as described at [golang.org](golang.org).

Using Terminal (Applications/Utilities/Terminal.app) clone this repository in the folder you'd like to be the workspace for ```wichess```. We'll be generating picture files into subfolders here and saving persistent runtime files, so the home folder probably makes the most sense for this. The default Go folder layout might be what you want to do.

```
git clone https://github.com/pciet/wichess ~/go/src/github.com/pciet/wichess
cd ~/go/src/github.com/pciet/wichess
```

Build ```wichess``` with the Build.sh script.

```
./Build.sh
```

That script also makes the mem folder.

The host program is ready, but the images still need to be rendered. Install Homebrew as described at [brew.sh](https://brew.sh) then install POV-Ray ([povray.org](http://www.povray.org)).

```
brew install povray
```

This next rendering step can take more than an hour and hundreds of megabytes of disk space.

```
cd img
./Render.sh
```

Now the host is ready. Start it.

```
./wichess
```

If you want other devices to be able to connect to the host then you must press "Allow" when asked if you want ```wichess``` to accept incoming network connections. This can also be changed in System Preferences using the Security & Privacy -> Firewall -> Firewall Options menu by setting ```wichess``` to "Allow incoming connections".

To determine your host's IP address use the ```ifconfig``` command or in System Preferences click on Network.

```wichess``` will correctly save information with the control+c interrupt command or if the computer is shutdown normally, but don't suddenly remove power to the computer.