# WISCONSIN CHESS README

![Splash](https://github.com/pciet/wichess/blob/master/docs/splash.jpg)

Welcome to the source code for Wisconsin Chess!

The game of chess can enable both competition and the perception of metaphors for war, debate, dog training, or other endeavors with two-sided choice making. Chess was played and designed by generations of people through centuries and is often seen in art. Many people will play chess today.

This is my variation of chess which builds on the traditional modern pieces and rules while maintaining the predictability and abstraction that makes chess unique. Pieces start on the same squares, and castling, en passant, promotion, check, checkmate, and draws are included. If you want to learn chess or improve your skill then this game will help you do that.

Added are concepts from trading card games: players each have a collection of new pieces that can optionally replace their regular chess pieces. This addition causes new emergent scenarios that keep gameplay fresh and exciting.

Piece and army variations make determining moves and check conditions complex, so computer calculations inform players of the moves they can make to enable playing at a good pace. Wisconsin Chess is a computer game played in web browsers through a local network host computer that relays move results, determines the moves, and holds players' collections.

Technical installation instructions that use this source code repository is described in docs/install.md. An overview of the source code is in docs/source.md.

## COPYRIGHT, LICENSES, CREDITS

This source code is the work from just me, but its foundation is the free and open software projects of many other people, and I started with knowledge gained from university studies and industry experience. Like people have done for me, I hope the quality of this source code will help you or others make their ideas.

A copyright notice for Wisconsin Chess is in the COPYRIGHT file.

Content with copyright owned by others is included in this repository under license, described here:

The web/fonts folder has the unmodified Linux Biolinum font, which has the copyright notice and license in licenses/SIL Open Font License.txt.

The Gorilla WebSocket package is in vendor/github.com/gorilla/websocket.

The PQ package is in vendor/github.com/lib/pq.

The golang.org/x/crypto package is in vendor/golang.org/x/crypto.

Part of the knight piece was copied from a set of chess piece design files. Credit, copyright notice, and the license are in img/knight.inc.



## Status

Wisconsin Chess is a work in progress.

In 2018 a first version was completed which I'm calling v0.1 (there's a git tag for it), and now I'm a few months into making v0.2 with improvements in every aspect of the game and source code.

The intention of v0.2 is a stable program that I'll be able to sell as a complete bundle on a separate website. It will be fun, valuable, and deliver the promises above, and I want it to be the start for great future updates leading to a version 1.

If you would like to support this project with money then buying it will be the best way, but for now please consider purchasing music I recorded for the game, at [matthewjuran.bandcamp.com](https://matthewjuran.bandcamp.com).

If you play Wisconsin Chess and find a mistake or think something should be better then any report is definitely appreciated. Please [open an issue](https://github.com/pciet/wichess/issues) on GitHub for your feedback, and thanks.

Matt