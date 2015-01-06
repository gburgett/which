# Which for Windows

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

I'm using [locomotivecms](http://locomotivecms.com/) for a project I'm working on, and it uses [ImageMagick](http://www.imagemagick.org/) for resizing pictures on the fly.  The [standard instructions](http://doc.locomotivecms.com/get-started/install-wagon#windows) say to use the bitnami Ruby stack installer to get everything you need in a nice neat bundle.

That worked for me for a while, until I needed to update the version of ruby I'm using.  So, I went through the long and painful process of uninstalling the ruby stack, installing each component individually, identifying and correcting a ton of minor problems, and finally squashed the last bug by writing this program.

Essentially, the command `bundle exec wagon serve` looks up the location of ImageMagick's `convert.exe` program by invoking the `which` command, which only works on linux systems, or if you're using a port like unxutils.  So, once I got ruby working and was self-hosting my site for development, none of my resized images were showing up.  The console showed this error message:

```sh
which: no convert in (...;C:\Program Files\ImageMagick-6.9.0-Q16;...)
```

