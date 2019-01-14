package main

import (
	"fmt"
	"os"

	"github.com/hirenchauhan2/wallpaper"
)

func printInfo() {
	fmt.Printf(`
Wallpaper is a utility for setting the desktop wallpaper.
You can give the local filepath or any image url.

Usage:
	wallpaper -f C:\Users\your_user_name\Pictures\made-a-botw-vector-wallpaper-4k-2560×1600.jpg

	wallpaper -i https://i.redd.it/l1764nd9h3721.jpg

	--version, -v version of the utility.
	--help, -h for help.
	--file, -f [filepath] local filepath.
	-i [url] download image from internet and set it as wallpaper
`)
}

func main() {
	var err error
	args := os.Args

	if len(args) >= 2 {
		switch args[1] {
		case "--file", "-f":
			err = wallpaper.SetLocalWallpaper(args[2])
			break
		case "--i", "-i":
			err = wallpaper.SetWallpaperFromURL(args[2])
			break
		case "--version":
		case "-v":
			fmt.Println("Wallpaper util version 0.1.1")
			break
		case "--help":
		case "-h":
			printInfo()
		default:
			fmt.Println("unknown flag!")
			printInfo()
		}

		if err != nil {
			fmt.Printf("\n Could not set wallpaper: %s\n", err.Error())
		}
	} else {
		printInfo()
	}
}
