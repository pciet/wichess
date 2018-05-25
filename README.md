Note: ```wichess``` is not designed for children and the descriptions here and elsewhere are meant for an adult audience. There isn't anything particularly adult about ```wichess``` though. The game is not age rated by any organization and this implementation is not licensed to you beyond the base GitHub license. There aren't any easter eggs, practical jokes, or intentional mistakes, so what you see should be what you get.

There are no lyrics in the music recordings and the game client currently doesn't play them, but you can listen to them by opening the files at ```github.com/pciet/wichess/web/music/``` with a music player. The music has some intensity at times. The music makes wichess bigger than GCC.

Note: no encryption is used between the server program and web browser, so don't use sensitive passwords when playing. Despite being tested on iOS and macOS there are now new issues related to Safari on iOS. Chrome on macOS appears to still work. Firefox on macOS has a layout bug but does work. Other operating systems like Windows or Android are untested. Issues are noted in the issue tracker. The client may degrade further as web browsers are improved.

Note: the cryptography implementation is a 2017 snapshot of ```golang.org/x/crypto``` which isn't the latest and may include serious security vulnerabilities that are patched in later versions.

The intention here is a demonstration as a start for a possible longer programming and business effort. For more background see the album description found through the bandcamp link at the bottom of this readme.

```wichess``` requires a PostgreSQL database matching the configuration you set in ```dbconfig.json```. The file ```postgres_tables.sql``` contains the SQL commands to create the necessary tables, and ```create_postgres_tables.sh``` may be configured to run the table create commands if you have ```psql``` installed.

![Screenshot1](https://github.com/pciet/wichess/blob/master/screenshots/Screen%20Shot%202018-01-09%20at%201.56.21%20PM.png)

![Art1](https://github.com/pciet/wichess/blob/master/graphics/art/album/album.png)

The music can be downloaded with track names at https://pciet.bandcamp.com/album/wisconsin-chess

```wichess``` includes a snapshot of ```github.com/gorilla/websocket```, see ```github.com/pciet/wichess/vendor/github.com/gorilla/websocket/LICENSE``` for the copyright notice and license.

```wichess``` includes a snapshot of ```github.com/lib/pq```, see ```github.com/pciet/wichess/vendor/github.com/lib/pq/LICENSE.md``` for the copyright notice and license.

```wichess``` includes a snapshot of ```golang.org/x/crypto```, see ```github.com/pciet/wichess/vendor/golang.org/x/crypto/LICENSE``` for the copyright notice and license.
