# Which for Windows

### If you're trying to run locomotivecms wagon, and you are getting `which: no convert in ...`, you need this on your path.

I wrote this super-simple which clone because the which in unxutils wasn't finding my .exe and the which in msys was a sh script that returns unix-style paths.  It searches the current directory, then every directory in the path to find the first instance of a command matching the given name.  A command is any `.exe` or `.bat` executable file.

### Usage
```sh
C:\Users\gordon> which convert
"C:\Program Files\ImageMagick-6.9.0-Q16\convert.exe"
```

```sh
C:\Users\gordon> which conv
Which: cannot find conv
searched path: C:\windows\system32;C:\Program Files\ImageMagick-6.9.0-Q16;...
```

```sh
C:\Users\gordon> which -prefix conv
"C:\Program Files\ImageMagick-6.9.0-Q16\convert.exe"
```

```sh
C:\Users\gordon> which -prefix con
"C:\Program Files\ImageMagick-6.9.0-Q16\conjure.exe"
```

```sh
C:\Users\gordon> which -all convert
"C:\Program Files\ImageMagick-6.9.0-Q16\convert.exe"
C:\windows\system32\convert.exe
```

### Installation

1. `go get github.com/gburgett/which`
2. `go install github.com/gburgett/which`
3. `where which` should show your first installed `which` command - move %GOPATH%/bin/which.exe into the correct location so it's first on your path
4. `which -version` should print `Which for Windows v1.0.0`


### Why did I write a which clone?

[Check out my blog for all the awful hackiness](http://www.gordonburgett.net/post/which%20which/)