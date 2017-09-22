#!/bin/bash

go build generate.go
./generate wqueen.pov bqueen.pov wking.pov bking.pov wrook.pov brook.pov wknight.pov bknight.pov wbishop.pov bbishop.pov wpawn.pov bpawn.pov wswap.pov bswap.pov block.pov wlock.pov brecon.pov wrecon.pov bdetonate.pov wdetonate.pov bghost.pov wghost.pov bsteal.pov wsteal.pov
cp img/*.png ../web/img/
