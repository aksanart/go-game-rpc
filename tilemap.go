package main

import (
	"encoding/json"
	"os"
)

type TilemapLayerJSON struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TilemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`
}

func NewTileMapJSON(filepath string) (*TilemapJSON, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var tilemapJSON TilemapJSON
	err = json.Unmarshal(content, &tilemapJSON)
	if err != nil {
		return nil, err
	}
	return &tilemapJSON, nil
}
