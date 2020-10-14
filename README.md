![Splash](https://github.com/pciet/wichess/blob/master/docs/splash.jpg)

Welcome to the source code for Wisconsin Chess!

Chess can be described as a fun tactical board game played by two people against each other. It enables both competition and perhaps evokes metaphors of war, debate, or many kinds of stories, and it's a good brain exercise and peaceful way to socialize. Chess was played and incrementally designed by generations of people through centuries and is often seen in art, and many people will play chess today around the world.

This source code is used to build the software of my computer game variation of chess, Wisconsin Chess. It adds to the traditional modern rules while keeping the abstract theme and important predictability of the original game; pieces start on the same squares, and castling, en passant, promotion, check, checkmate, and draws are included. If you want to learn chess or improve your skill then this game will help you do that.

What's added is the trading card game idea (but without trading for now). Players each have a collection of pieces with new characteristics that can optionally replace the regular chess pieces, so emergent scenarios from a choice of many strategies keep gameplay fresh and exciting no matter how long you've played.

Wisconsin Chess is hosted in your local network using a computer, then people play by connecting to it with a web browser through the network. The latest versions of Chrome, Firefox, and Safari are supported, and no browser plugins are necessary. This local network website is designed to work well both on mobile smartphones and tablets alongside desktop computers, and any number of concurrent players can be playing up to the limits of the host computer's capabilities.

Technical installation instructions that use this source code repository are described in docs/install.md. An overview of the source code is in docs/source.md.

## Project Status

Wisconsin Chess is a work in progress. It was inspired by one of the ideas I researched after quitting my tech job to pursue solo entrepreneurship in spring 2014. Coding began partway through 2017 and a first version was completed during the start of 2018 (see the v0.1 git tag), but the project was shelved then.

During 2020 I've had the opportunity to work on it again. Version 0.2 is a massive improvement and is currently in an alpha state as I continue to make it more fun, unbreakable, and valuable.

Shortly I will have to shelve the project again because of lack of funds. If you or someone you know is interested in a funding partnership for this project I'd be happy to discuss such business details, or if you would like to freely contribute a small amount of money then buying my music at [matthewjuran.bandcamp.com](https://matthewjuran.bandcamp.com) is the best way to do that.

If you play Wisconsin Chess and find a mistake or think something should be better then any report is appreciated. Please [open an issue](https://github.com/pciet/wichess/issues) on GitHub for your feedback, and thanks.

## Copyright, Licenses, Credits

This source code is the work from just me, but its foundation is the free and open software projects of many other people, and I started with knowledge gained from university studies and industry experience. Like people have done for me, I hope the quality of this source code could make it a resource to help you or others make their ideas.

A copyright notice for Wisconsin Chess is in the COPYRIGHT file.

Content with copyright owned by others is included in this repository under license, described here:

The web/fonts folder has the unmodified Linux Biolinum font, which has the copyright notice and license in licenses/SIL Open Font License.txt.

The Gorilla WebSocket package is in vendor/github.com/gorilla/websocket.

A subset of the golang.org/x/crypto package is in vendor/golang.org/x/crypto.

The knight piece was copied from a set of chess piece design files. Credit, copyright notice, and the license are in img/knight.inc.