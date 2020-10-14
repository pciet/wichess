#!/bin/bash

# Do these on the BeagleBone Black before executing this script:

# Change the debian and root users' passwords with passwd.
# Format and populate a microSD card as described in docs/bbb_sd.md, and have it inserted.
# Use these following commands to cause Debian to automatically mount the card during boot. The
# last command will cause a reboot:
#     su
#     mkdir /media/sd
#     echo "/dev/mmcblk0  /media/sd  ext4  noatime  0  2" >> /etc/fstab
#     shutdown -r now

# After the reboot completes then execute this script on the BeagleBone as root and with an
# Internet connection.

#####################

# This script adjusts an installation of
# bone-eMMC-flasher-debian-10.5-console-armhf-2020-08-12-1gb.img
# from
# https://rcn-ee.net/rootfs/bb.org/testing/2020-08-12/buster-console/
# via
# https://groups.google.com/g/beagleboard/c/ZK1qWQiGnEA
# for use with Wisconsin Chess.
#
# Another source for links to console images is
# https://elinux.org/Beagleboard:BeagleBoneBlack_Debian#Debian_Buster_Console_Snapshot

#####################

export LANG=C.UTF-8
adduser --system --group --no-create-home --disabled-login wichess

##### Configure Wisconsin Chess to run as a background systemd service using the wichess user.

chown -R wichess:wichess /media/sd/mem /media/sd/html /media/sd/web /media/sd/wichess
cp /media/sd/wichess.service /etc/systemd/system/wichess.service
systemctl daemon-reload
systemctl enable wichess.service

##### Remove unnecessary Debian packages and folders.

# TODO: gcc-8-base would be ideal to remove, but purge gets to many essential packages

apt-get update
apt-get -y purge bb-bbai-firmware bb-wl18xx-firmware bluez bsdmainutils btrfs-progs \
    cloud-guest-utils crda dirmngr firmware-atheros firmware-brcm80211 firmware-iwlwifi \
    firmware-libertas firmware-misc-nonfree firmware-realtek firmware-zd1211 gdbm-l10n \
    gnupg-l10n gnupg-utils gpg-agent gpg gpgconf hostapd iputils-ping iw nano-tiny nano \
    patch perl-modules-5.28 perl pinentry-curses rfkill wget whiptail \
    wireless-regdb wireless-tools wpasupplicant ca-certificates sudo

rm -r /lib/firmware/brcm /lib/firmware/rtlwifi /etc/X11 /etc/sudoers.d
apt autoremove
apt-get upgrade

## TODO: a complete no-debug app testing configuration would also remove these packages:
## openssh-client
## openssh-server
## openssh-sftp-server
## openssl

## TODO: add firewall rules to only allow HTTP and WebSocket traffic to wichess
## https://wiki.debian.org/DebianFirewall
## use nftables for this
## redirect port 80 to wichess port

## TODO: remove the debian user and any other unnecessary users

## TODO: write logs to SD card

## TODO: enable AppArmor for wichess
# https://wiki.debian.org/AppArmor/HowToUse

# shutdown -r now
