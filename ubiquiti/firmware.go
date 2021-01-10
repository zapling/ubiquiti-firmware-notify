package ubiquiti

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

type Firmware struct {
	Rank            int    `json:"rank"`
	Featured        bool   `json:"featured"`
	CategroySlug    string `json:"category__slug"`
	SdkID           int    `json:"sdk__id"`
	Id              int    `json:"id"`
	Size            string `json:"size"`
	Mib             string `json:"mib"`
	Filename        string `json:"filename"`
	CategroyName    string `json:"category__name"`
	Version         string `json:"version"`
	Build           string `json:"build"`
	Thumbnail       string `json:"thumbnail"`
	Description     string `json:"description"`
	ThumbnailRetina string `json:"thumbnail_retina"`
	DatePublished   string `json:"date_published"`
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	Changelog       string `json:"changelog"`
	FilePath        string `json:"file_path"`
	Products        string `json:"products"`
	Architecture    string `json:"architecture"`
}

func GetAvailableFirmware(device string) ([]Firmware, error) {
	url := "https://www.ui.com/download?group=" + device

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("An error occurred when trying to create a new request")
	}

	req.Header.Add("X-Requested-With", "XMLHttpRequest") // Fool Ubiquiti API that we are a browser

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error when sending request")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Unable to read response body")
	}

	firmware, err := unmarshalFirmwareRequest(bytes)
	if err != nil {
		return nil, err
	}

	return firmware, nil
}

func unmarshalFirmwareRequest(body []byte) ([]Firmware, error) {
	type firmwareResponse struct {
		Downloads []Firmware
	}

	var response firmwareResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unable to convert response to json")
	}

	return response.Downloads, nil
}
