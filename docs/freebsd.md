# Installation On FreeBSD

Wisconsin Chess can be installed from source onto the FreeBSD operating system. Here's the sequence that worked for me with a 2016 amd64 DIY desktop computer I assembled from parts ordered at [newegg.com](https://newegg.com).

## 1. Install FreeBSD

Starting on my macOS laptop, I downloaded the 11.3 memstick image from [freebsd.org](https://www.freebsd.org/where.html) and verified the image copy was error-free by comparing the output of shasum with the corresponding number in the [checksum file](https://www.freebsd.org/releases/11.3R/CHECKSUM.SHA512-FreeBSD-11.3-RELEASE-amd64.asc).

```shasum -a 512 FreeBSD-11.3-RELEASE-amd64-memstick.img```

The FreeBSD installation instructions suggest writing the USB drive with the dd command, but that didn't work for me; on boot there was a ```Missing Operating System``` message.

[balenaEtcher](https://www.balena.io/etcher/) was mentioned in Ubuntu USB drive imaging instructions I found. Etcher worked for me with the .img, so try that if you're using macOS.

With the bootable USB drive I then installed FreeBSD onto the target computer, by selecting the USB drive from the boot menu gotten to with a press of F12 right away at power on (look for "boot menu" when the computer first turns on for the right key to press). Follow the installer instructions.

An Internet connection will be required for the next steps, so be sure networking is configured. I use the root account to install software needed for Wisconsin Chess, then a user account is used to run Wisconsin Chess.

## 2. Install Go

I installed the latest version of Go (1.14.3 while I'm typing this guide). HTTPS is used to get it which requires installation of security certificates first.

Do the following as the root user.

Run the following command, and read the notice that you have to judge the trustworthiness of this package yourself.

```
pkg install ca_root_nss
```

Fetch the Go archive.

```
fetch https://dl.google.com/go/go1.14.3.freebsd-amd64.tar.gz
```

Unpack the archive into the installation location.

```
tar -C /usr/local -xzf go1.14.3.freebsd-amd64.tar.gz
```

Use ```vi``` to add the Go tool to all users' path.

```
vi /etc/profile
```

Type ```G``` to go to the end of the file, then press ```o``` to begin text insertion on a new line, then type in this line.

```
export PATH=$PATH:/usr/local/go/bin
```

Press ```esc``` then type ```ZZ``` to save and exit from ```vi```.

Go should now be installed. This can be verified by logging in as a non-root user and seeing if this command returns a filepath.

```
which go
```

The official Go installation instructions are at [golang.org](https://golang.org/doc/install) which also shows how to test the installation in more detail.

## 3. Install Other Application Dependencies

As root install the applications we'll use to make and run Wisconsin Chess from source.

```
pkg install git
```

```
pkg install bash
```

```
pkg install povray37
```

POV-Ray must be linked to the ```povray``` command.

```
ln /usr/local/bin/povray37 /usr/local/bin/povray
```

## 4. Install Wisconsin Chess

The previous is all that needs to be done as root. Login as a regular user for the following commands.

Get the Wisconsin Chess source code.

```
go get github.com/pciet/wichess
cd ~/go/src/github.com/pciet/wichess
```

Render the piece images. This can take tens of minutes or hours and take hundreds of megabytes of disk space.

```
cd img
./Render.sh
```

Build the host ```wichess``` program.

```
cd ..
./Build.sh
```

If everything went right then the Wisconsin Chess host is ready.

## 5. Connect To This Host

From devices on the same network as your FreeBSD host you should be able to connect via web browsers. Find the IP address.

```
ifconfig
```

There should be a number like ```192.168.0.50``` that you'll use on other devices.

Start the Wisconsin Chess host.

```
./wichess
```

Assuming the example IP address above is yours, then in a web browser on a device in the same network go to the following address in the address bar.

```
192.168.0.50:8080
```

A login page should appear, then you create a new Wisconsin Chess player by typing in a fresh name and password.

When starting ```wichess``` a different IP port can be picked with the ```-port``` option. Call ```./wichess -h``` to print usage information.