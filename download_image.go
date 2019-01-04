package wallpaper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
)

// downloadImage function is used to download image
func downloadImage(imgURL string) (string, error) {
	// send http get request to fetch the resource (Image).
	resp, err := http.Get(imgURL)
	if err != nil {
		return "", err
	}
	// close the Body buffer memory when function ends
	defer resp.Body.Close()

	// current user
	usr, _ := user.Current()
	fmt.Println("the current user is: ", usr.Name)

	// Unescape the query to get the proper file name
	imgURL, err = url.QueryUnescape(imgURL)
	if err != nil {
		fmt.Println("Could not escape url", err.Error())
	}
	// filenmae wil full path joining the user's home dir/Pictures/image.extension
	filename := path.Join(usr.HomeDir, "Pictures", path.Base(imgURL))
	fmt.Println("The filename is: ", filename)
	// Create the image file at the given path
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error while creating file")
		return "", err
	}
	// close the file when function ends
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error while copying response contents to file")
		return "", err
	}
	fmt.Println("File downloaded successfully!")
	return filename, nil
}
