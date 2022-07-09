package core

import (
	"fmt"
	goio "io"
	"net/http"
	"os"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
)

var gtdfLatestUrl = "http://127.0.0.1/GTDF-Latest.zip?cli=" + CLI_VERSION //"https://akrck02.com/#/project/gtdf/latest/cli/v"

func NewProject(url string, name string) {

	logger.ShowLogGtdfTitle()
	logger.Log("Creating new project...")

	if url == "" {
		current_url, err := os.Getwd()

		if err != nil {
			logger.Error("Cannot get current directory. Check your permissions and try again.")
			return
		}

		url = current_url
	}

	logger.Jump()
	logger.Log("URL : " + url)
	logger.Log("Name: " + name)
	logger.Jump()

	url = url + "/"
	getLatestVersion(url, name)

}

func getLatestVersion(url string, name string) {

	logger.Log("Creating base directory...")
	io.SecureMkdir(url, 0777)
	logger.Log("Ok.")
	logger.Jump()

	logger.Log("Getting latest version...")
	file, err := os.Create(url + "gtdf-latest.zip")
	if err != nil {
		logger.Error("Cannot create transition file. Check your permissions and try again.")
		return
	}

	resp, err := http.Get(gtdfLatestUrl)

	logger.Log("Server response: " + resp.Status)
	if err != nil || resp.StatusCode != 200 {
		logger.Error("Cannot get latest version, check your internet connection and try again.")
		return
	}

	logger.Jump()

	size, err := goio.Copy(file, resp.Body)
	if err != nil {
		logger.Error("Cannot extract latest version. Check your permissions and try again.")
		return
	}

	defer file.Close()

	logger.Log("Cleaning possible old files...")
	io.RemoveAnyway(url + name)
	logger.Log("Ok.")
	logger.Jump()

	logger.Log("Extracting latest version...")
	io.Unzip(url+"gtdf-latest.zip", url)
	logger.Log("Ok.")
	logger.Jump()

	logger.Log("Removing transition file...")
	io.SecureRemoveFile(url + "gtdf-latest.zip")
	io.SecureRenameFile(url+"GTD-Framework-main", url+name)
	io.SecureRemoveAll(url + name + "/.github")
	logger.Log("Ok.")
	logger.Jump()

	logger.Log(fmt.Sprintf("Downloaded %d bytes", size))
	logger.Log("Latest version fetched.")
}
