package dto

import "sync"

type Music struct {
	Id       string
	Name     string
	Author   string
	Album    string
	Cover    string
	Source   string
	Duration int
	EndTime  int
}

type VoiceInstance struct {
	QueueMutex sync.Mutex
	ChannelId  string
	NowPlaying Music
	Queue      []Music
	PlayedTime int
}

type PlayerSignal struct {
	Command string
	Channel string
	Music   *Music
}
