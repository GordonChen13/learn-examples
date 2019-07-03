package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const indexFile = "index.json"
const urlFormat = "https://xkcd.com/%d/info.0.json"

type Index struct {
	Infos    map[int]*Info
	FilePath string
}

type Info struct {
	Title      string
	Transcript string
	ImgURL     string `json:"img_url"`
}

func NewInfoFromURL(comicURL string) (*Info, error) {
	resp, err := http.Get(comicURL)
	if resp.StatusCode != http.StatusOK || err != nil {
		resp.Body.Close()
		return nil, err
	}

	var info Info
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &info, nil
}

func NewIndex(filePath string) *Index {
	var infos map[int]*Info
	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		infos = make(map[int]*Info)
	} else {
		json.Unmarshal(out, &infos)
	}

	return &Index{
		Infos:    infos,
		FilePath: filePath,
	}
}
