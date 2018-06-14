# glod-cli

**glod-cli** is a small command line tool that using [glod](https://github.com/dwarvesf/glod) to download music/video from multiple sources.

glod-cli is written in [Go](http://golang.org/) with support for multiple platforms. We currently provide pre-built binaries for Windows, Linux, FreeBSD and  OS X (Darwin) for x64, i386 and ARM architectures.

glod-cli may also be compiled from source wherever the Go compiler tool chain can run, e.g. for other operating systems including DragonFly BSD, OpenBSD, Plan 9 and Solaris

[![Video Walkthrough](/img/walkthrough.gif)](/img/walkthrough.gif)

## Installation

### Binary Install

If you want to use glod-cli, simply install the glod-cli binaries. The glod-cli binaries have no external dependencies.

Installation is very easy. Simply download the appropriate version for your platform from [glod-cli Releases](https://github.com/dwarvesf/glod-cli/releases). Once downloaded it can be run from anywhere. You don’t need to install it into a global location. This works well for shared hosts and other systems where you don’t have a privileged account.

Ideally, you should download and put it somewhere in your `$PATH` for easy use. `/usr/local/bin` is the most probable location.

On OS X, if you have [Homebrew](http://brew.sh/), installation is even easier:

```
$ brew tap dwarvesf/homebrew-tap
$ brew install glod-cli
```

[![Homebrew Tap](/img/homebrew-tap.png)](/img/homebrew-tap.png)

### Build and Install the Binaries from Source
  
Add glod-cli and its package dependencies to your go src directory.

```
go get -v github.com/dwarvesf/glod-cli
```

Once the get completes, you should find your new `glod-cli` (or `glod-cli.exe`) executable sitting inside `$GOPATH/bin/`.

To update glod-cli dependencies, use `go get` with the `-u` option.

```
go get -u -v github.com/dwarvesf/glod-cli
```

### Upgrading

Upgrading glod-cli is as easy as downloading and replacing the executable you’ve placed in your `$PATH`.

## Usage

Make sure either `glod-cli` is in your `$PATH` or provide a path to it.

``` shell

$ glod-cli help

A small cli written in Go to help download music/video from multiple sources.

Usage:
  glod [command]

Available Commands:
  download    download command
  help        Help about any command
  play        play command
  version     Print the version number of glod-cli

Use "glod [command] --help" for more information about a command.

```

##### Example:

To download song/video to current directory
```
glod-cli download -l=https://www.youtube.com/watch?v=6d6oq0zGGmw 
```

To download song/video to custom directory
```
glod-cli download -l=https://www.youtube.com/watch?v=6d6oq0zGGmw -o=youtube-download
```

To play song/video after downloaded(OSX support)
```
glod-cli play https://www.youtube.com/watch?v=6d6oq0zGGmw
```

# Supported sources & TODO

### Music

- [x] [Nhaccuatui](http://www.nhaccuatui.com/)
- [x] [Zing Mp3](http://mp3.zing.vn/)
- [x] [SoundCloud](https://soundcloud.com)
- [x] [Chiasenhac](http://chiasenhac.com)

### Video 

- [x] [YouTube](https://www.youtube.com/)
- [x] [Facebook](https://facebook.com/)
- [x] [Vimeo](https://vimeo.com/)
- [ ] Lynda
- [ ] Udemy
- [ ] Coursera

### Files

- [ ] Flickr
- [ ] Slideshare
- [ ] Dropbox

# License

Copyright 2016 Dwarves Foundation

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.