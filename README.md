# intel-vpro-detector
A program to detect is intel vPro is available on a machine, and if so, if it is enabled

Note that the initial state of this repo is none of my code, the executable is proprietary and the python script has been produced by someone else.

The code I have written is in golang. First install go using your package manager.

```shell
go get github.com/akavel/rsrc
```

```shell
rsrc -manifest Main.exe.manifest main.syso
```

```shell
go build -ldflags -H=windowsgui -o "vPro Detector.exe"
```
