# Wallpaper

A utility to set desktop wallpaper on Windows, (Mac & Linux WIP).


## Set wallpaper from local image or from internet!
It can set a wallpaper from an image's URL or local path to image file.

```bash
wallpaper -i URL # -i for Internet. passing the image url
# or
wallpaper -f C:/path/to/image.png # -f or --file
# For Windows the path looks weird, right? but it is working only with / only, \ or \\ in path is not working. idk...
```


## Building the wallpaper
```bash
go build .
cd example
go build -o wallpaper.exe .
# save the wallpaper.exe somewhere in %PATH% or $PATH so that it can be accessible from anywhere in terminal.
```

## want to contribute?
Please! help me writing the better code by doing the code review, suggestions by creating issue/PR.

## Todo
 - [x] Windows
 - [ ] Mac (I don't have)
 - [ ] Linux (so many flavours, so many environments)


Built with :heart: in India.
