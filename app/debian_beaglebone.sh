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

apt-get update

adduser --system --group --no-create-home --disabled-login wichess


##### Install and configure PostgreSQL to save to the card and allow access by the wichess user.

apt-get -y install postgresql

pg_ctlcluster 11 main stop
pg_dropcluster 11 main

mkdir /media/sd/pgsql
chown postgres /media/sd/pgsql

pg_createcluster -d /media/sd/pgsql/data -l /media/sd/pgsql/pgsql.log --start 11 wichess -- \
    --auth-local peer --auth-host md5 -U postgres

su -c "createdb wichess" postgres
su -c "psql -d wichess -f /media/sd/postgres_tables.sql" postgres

echo "enter the password 'wichess' at the following prompt:"
su -c "createuser -P wichess" postgres

su -c "psql -c 'GRANT CONNECT ON DATABASE wichess TO wichess;'" postgres

su -c "psql -d wichess \
    -c 'GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE players, games TO wichess;'" postgres

su -c "psql -d wichess \
    -c 'GRANT USAGE, SELECT ON SEQUENCE games_id_seq, players_id_seq TO wichess;'" postgres

systemctl start postgresql@11-wichess


##### Configure Wisconsin Chess to run as a background systemd service using the wichess user.

chown -R wichess:wichess /media/sd/dbconfig.json /media/sd/html /media/sd/web /media/sd/wichess

cp /media/sd/wichess.service /etc/systemd/system/wichess.service

systemctl daemon-reload
systemctl start wichess.service


##### Remove unnecessary Debian packages and folders.

# TODO: gcc-8-base would be ideal to remove, but purge gets to many essential packages

apt-get -y purge bb-bbai-firmware bb-wl18xx-firmware bluez bsdmainutils btrfs-progs \
    cloud-guest-utils crda dirmngr firmware-atheros firmware-brcm80211 firmware-iwlwifi \
    firmware-libertas firmware-misc-nonfree firmware-realtek firmware-zd1211 gdbm-l10n \
    gnupg-l10n gnupg-utils gpg-agent gpg gpgconf hostapd iputils-ping iw nano-tiny nano \
    patch perl-modules-5.28 perl pinentry-curses rfkill wget whiptail \
    wireless-regdb wireless-tools wpasupplicant ca-certificates sudo

rm -r /lib/firmware/brcm /lib/firmware/rtlwifi /etc/X11 /etc/sudoers.d

apt autoremove


## TODO: a complete no-debug app testing configuration would also remove these packages:
## openssh-client
## openssh-server
## openssh-sftp-server
## openssl

## TODO: add firewall rules to only allow HTTP and WebSocket traffic to wichess
## https://wiki.debian.org/DebianFirewall

## TODO: remove the debian user and any other unnecessary users
