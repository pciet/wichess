Note: no encryption is used between the server program and web browser, so don't use sensitive passwords when playing. Despite being tested on iOS and macOS there are now new issues related to Safari on iOS. Chrome on macOS appears to still work. Issues are noted in the github issue tracker.
The intention here is a demonstration as a start for a possible longer programming effort. For more background see the album description found through the bandcamp link at the bottom of this readme.
```wichess``` requires a PostgreSQL database matching the configuration you set in ```dbconfig.json```. The file ```postgres_tables.sql``` contains the SQL commands to create the necessary tables, and ```create_postgres_tables.sh``` may be configured to run the table create commands if you have ```psql``` installed.
![Screenshot1](https://github.com/pciet/wichess/blob/master/screenshots/Screen%20Shot%202018-01-09%20at%201.56.21%20PM.png)
![Art1](https://github.com/pciet/wichess/blob/master/graphics/art/album/album.png)
Associated music is included in the repository at wichess/web/music/ or can be downloaded in other formats with track names at https://pciet.bandcamp.com/album/wisconsin-chess
