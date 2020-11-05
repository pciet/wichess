# Wisconsin Chess v0.2 Alpha on BeagleBone Black

### [github.com/pciet/wichess](https://github.com/pciet/wichess)

### Don't leave the BeagleBone unattended while powered.

It's for development only and not tested for regular consumer use by the company that makes it. It's a fire hazard if components fail.

### Don't unplug power while the lights are on.

Power off by pressing the small button next to the Ethernet plug and wait for the blue lights to turn off after a few seconds.

### The BeagleBone isn't perfect.

Occasionally the power will come back on after pressing the power button, or the Ethernet won't start up or takes a long time.

## Playing

Plug the BeagleBone into power and your local network with the Ethernet cable. It may take a moment for the game to become available.

Play by typing in the BeagleBone's IP address into a web browser's address bar. Safari, Chrome, and Firefox are supported. The Wisconsin Chess port is 8080, so the address will look something like ```192.168.0.110:8080```.

### Getting the IP address

A way to get the IP address is by also plugging in the BeagleBone via USB to a computer. There it will show up as a network device and your computer will be assigned an IP address like 192.168.6.1. The BeagleBone's address will be 2 (192.168.6.2), which can then be used with SSH to determine the main local network's address.

Assuming you have ```ssh``` use it from a command line.

```
ssh debian@192.168.6.2
```

There are two user accounts on the BeagleBone with these passwords:

* debian:
* root:

Enter the debian user's password, then use the ```ifconfig``` command to list all network connections. The ```eth0``` device will have your local network's address.

Remember to append the 8080 port number to the IP address when connecting, so if the address shown is ```192.168.5.150``` then in the web browser you'll type ```192.168.5.150:8080```.

### Accounts

The login webpage you'll be presented with when connecting is simpler than most. Just type in a name and password and a new account will be created. If the webpage goes white when you try to login then the password is wrong or the username is already in use.