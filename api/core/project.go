package core

import (
	"fmt"
	goio "io"
	"net/http"
	"os"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
)

const GTDF_LATEST_URL = "https://github.com/akrck02/GTD-Framework/releases/download/latest/GTD-Framework.zip"
const DISTRO_DIR = "./dist"
const LATEST_ZIP = "gtdf-latest.zip"
const DEFAULT_PROJECT_NAME = "GTD-Framework"

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

	logger.Log("Name: " + name)

	url = url + "/"
	getLatestVersion(url, name)

}

func getLatestVersion(url string, name string) {

	distUrl := url + DISTRO_DIR + "/"
	latestZipUrl := distUrl + LATEST_ZIP
	defaultProjectUrl := distUrl + DEFAULT_PROJECT_NAME
	newProjectUrl := url + name

	io.SecureMkdir(distUrl, 0777)
	logger.Log("Created base directory")

	io.RemoveAnyway(distUrl + "GTD-Framework")
	logger.Log("Cleaned old files")

	io.RemoveAnyway(newProjectUrl)
	logger.Log("Cleaned old files")

	file, err := os.Create(distUrl + "gtdf-latest.zip")
	if err != nil {
		logger.Error("Cannot create transition file. Check your permissions and try again.")
		return
	}
	logger.Log("Fetched latest version")

	resp, err := http.Get(GTDF_LATEST_URL)
	logger.Log("Server response: " + resp.Status)
	if err != nil || resp.StatusCode != 200 {
		logger.Error("Cannot get latest version, check your internet connection and try again.")
		return
	}

	size, err := goio.Copy(file, resp.Body)
	if err != nil {
		logger.Error("Cannot extract latest version. Check your permissions and try again.")
		return
	}

	defer file.Close()

	io.Unzip(latestZipUrl, distUrl)
	logger.Log("Extracted latest version of GTD Framework.")

	io.SecureRemoveFile(latestZipUrl)
	io.SecureRemoveAll(defaultProjectUrl + "/.github")
	io.SecureRenameFile(defaultProjectUrl, newProjectUrl)
	io.SecureRemoveAll(distUrl)
	logger.Log("Removed transition files.	")

	logger.Log(fmt.Sprintf("Downloaded %d bytes", size))
	logger.Log("Latest version fetched.")
}
