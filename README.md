# Muclean

A simple music file renamer

# Installation

```sh
go install github.com/CJ-Jackson/muclean@latest
```
# Usage Example

```
$ ls
ls
01.N.S.U..m4a                       08.Rollin' And Tumblin'.m4a
02.Sleepy Time Time.m4a             09.I'm So Glad.m4a
03.Dreaming.m4a                     10.Toad.m4a
04.Sweet Wine.m4a                   11.Wrapping Paper - Bonus Track.m4a
05.Spoonful.m4a                     12.I Feel Free - Bonus Track.m4a
06.Cat's Squirrel.m4a               13.Coffee Song - Bonus Track.m4a
07.Four Until Late.m4a

muclean "*.m4a"

ls
01 - Cream - Fresh Cream - N.S.U..m4a
02 - Cream - Fresh Cream - Sleepy Time Time.m4a
03 - Cream - Fresh Cream - Dreaming.m4a
04 - Cream - Fresh Cream - Sweet Wine.m4a
05 - Cream - Fresh Cream - Spoonful.m4a
06 - Cream - Fresh Cream - Cat's Squirrel.m4a
07 - Cream - Fresh Cream - Four Until Late.m4a
08 - Cream - Fresh Cream - Rollin' And Tumblin'.m4a
09 - Cream - Fresh Cream - I'm So Glad.m4a
10 - Cream - Fresh Cream - Toad.m4a
11 - Cream - Fresh Cream - Wrapping Paper  Bonus Track.m4a
12 - Cream - Fresh Cream - I Feel Free  Bonus Track.m4a
13 - Cream - Fresh Cream - Coffee Song  Bonus Track.m4a
```

# Motivation

Software like [Renamer](https://renamer.com/) and [Mp3Tag](https://www.mp3tag.de/en/) are nice tools, but I
needed something that is specialised that I can run from the terminal and not having to rely on regexp to sanitise the
file name from Windows System Reserved Characters.