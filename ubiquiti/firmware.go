package ubiquiti

import (
	"net/http"
)

var client = &http.Client{}

type Firmware struct {
	Rank            int    `json:"rank"`
	Featured        bool   `json:"featured"`
	CategroySlug    string `json:"category__slug"`
	SdkID           int    `json:"sdk__id"`
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

func GetAvailableFirmware(device string) {
}
