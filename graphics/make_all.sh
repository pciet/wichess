#!/bin/bash

go build generate.go
./generate wqueen.pov bqueen.pov wking.pov bking.pov wrook.pov brook.pov wknight.pov bknight.pov wbishop.pov bbishop.pov wpawn.pov bpawn.pov
cp img/*.png ../web/img/
