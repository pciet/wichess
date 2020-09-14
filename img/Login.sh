#!/usr/bin/env bash

# Login.sh renders login.pov, the teaser image on the login webpage.

SHORT=false

if [ $# -eq 1 ]
then
    SHORT=true
fi

DIM=2048

if [ "$SHORT" = true ]
then
    povray +Ilogin.pov +H512 +W512 Quality=5 +FN +Ologin.png
else
    povray +Ilogin.pov +H$DIM +W$DIM Quality=8 +FN +A +Ologin.png
fi
