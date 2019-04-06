package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemStore(database *os.File) (*FileSystemStore, error) {
	database.Seek(0,0)

	err := initialisePlayerDBFile(database)
	if err != nil {
		return nil, err
	}

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", err, database)
	}
	return &FileSystemStore{
		database:json.NewEncoder(&tape{database}),
		league:league,
	}, nil
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemStore) RecordWin(name string)  {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
	return
}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0,0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}
