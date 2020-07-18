# WISCONSIN CHESS README

Welcome to the source code for Wisconsin Chess!

A match of chess is both a competition and a story of abstract unpredictable battle that can be thought of as a symbol of war or a metaphor for constructive argument or other human interactions. Chess has a history lasting centuries, it's a competitive intellectual sport, and it's often seen even in contemporary art.

This is a variation of chess that keeps the traditional modern rules. Pieces and their starting squares are the same, and castling, en passant, promotion, timers, check, checkmate, and draws are included. If you want to learn chess or improve your skill then this game will help you do that.

What's added is an optional ability to vary your army with new kinds of pieces before you start a match. This causes new scenarios that may need a different thought process for the choosing of moves than in regular chess.

## As Software

Wisconsin Chess is that game, but it is also software that lets you use a host computer and network to have many people compete with each other through personal computers and devices like smartphones and pads. An updated web browser is the only thing needed by the players. The host computer determines and communicates moves available each turn so players focus on their turn choice instead of being distracted by the emergent complexity of the variant pieces.

I began to explore making this game in Portland, Oregon coffee shops in 2014, and a first draft of variant pieces was implemented in a prototype later that year. Other projects and experiences then happened before I did this web version in the second half of 2017 in Wisconsin. In 2020 I've begun making new improvements and am putting in full-time work on it again, at least while the pandemic is active.

## Playing

Wisconsin Chess is intended for adults; despite in my understanding there not being anything particularly adult about the game perhaps there are effects from games like this on a person's psychology that is better handled with a mature mind. Wisconsin Chess is not age rated or approved by any organization like the ESRB.

Technical installation instructions that use this source code repository is described in docs/install.md. An overview of the source code is in docs/source.md. A copyright notice is in the COPYRIGHT file.

A "just works" installation, without the need for technical expertise, is planned but not done yet. If you would like to support this project with money then buying that installer will be the best way, but for now please consider purchasing music I recorded for the game, at [matthewjuran.bandcamp.com](https://matthewjuran.bandcamp.com/album/typing-choice-tree).

## By Others

Other computer programs are necessary to have installed on your host which you'll need to get separately. Those are described in the docs/install.md file.

Content with copyright owned by others is included in this repository under license, described here:

The web/fonts folder has the unmodified Linux Biolinum font, which has the copyright notice and license in licenses/SIL Open Font License.txt.

The Gorilla WebSocket package is in vendor/github.com/gorilla/websocket.

The PQ package is in vendor/github.com/lib/pq.

The golang.org/x/crypto package is in vendor/golang.org/x/crypto.

Part of the knight piece was copied from a set of chess piece design files. Credit, copyright notice, and the license are in img/knight.inc.
