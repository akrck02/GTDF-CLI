package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var gtdfLatestUrl = "http://127.0.0.1/GTDF-Latest.zip" //"https://akrck02.com/#/project/gtdf/latest"

func New_project(url string, name string) {

	ShowLogGtdfTitle()
	Log("Creating new project...")

	if url == "" {
		current_url, err := os.Getwd()

		if err != nil {
			Error("Cannot get current directory. Check your permissions and try again.")
			return
		}

		url = current_url
	}

	Jump()
	Log("URL : " + url)
	Log("Name: " + name)
	Jump()

	url = url + "/"
	getLatestVersion(url, name)

}

func getLatestVersion(url string, name string) {

	Log("Creating base directory...")
	SecureMkdir(url, 0777)

	Log("Getting latest version...")

	file, err := os.Create(url + "gtdf-latest.zip")
	if err != nil {
		Error("Cannot create transition file. Check your permissions and try again.")
		return
	}

	resp, err := http.Get(gtdfLatestUrl)

	Log("Server response: " + resp.Status)
	if err != nil || resp.StatusCode != 200 {
		Error("Cannot get latest version, check your internet connection and try again.")
		return
	}

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		Error("Cannot extract latest version. Check your permissions and try again.")
		return
	}

	defer file.Close()

	Log("Cleaning possible old files...")
	RemoveAnyway(url + name)

	Log("Extracting latest version...")
	Unzip(url+"gtdf-latest.zip", url)

	Log("Removing transition file...")
	SecureRemoveFile(url + "gtdf-latest.zip")
	SecureRenameFile(url+"GTD-Framework-main", url+name)
	SecureRemoveAll(url + name + "/.github")

	Log(fmt.Sprintf("Downloaded %d bytes", size))
	Log("Latest version fetched.")
}
