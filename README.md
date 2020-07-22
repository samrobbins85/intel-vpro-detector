# Intel vPro Detector
A program to detect is intel vPro is available on a machine, and if so, if it is enabled

## Usage
**Note**: This program can only be used on Windows.

Download both executables from the releases section of this repository, and run the file `vPro Detector.exe` this will then give a dialog box with your vPro status

Note that you may get warnings when running this that it comes from an unverified source. This is due to the fact that Microsoft requires a paid certificate in order to verify you as a source. The best assurance I can offer you is that ther executable is built in the open using GitHub actions, and so it is assured that the underlying code produces the output executable.

## Development

The code I have written is in golang. First install go using your package manager or from [here](https://golang.org/dl/). Then run the following commands to build it:

```shell
go get github.com/akavel/rsrc
```

```shell
rsrc -manifest Main.exe.manifest main.syso
```

```shell
go build -ldflags -H=windowsgui -o "vPro Detector.exe"
```

## License 
Note that this repository is stated with an MIT license, the file `PlatformDiscovery.exe` was not created by me, and so does not fall under this license. 
