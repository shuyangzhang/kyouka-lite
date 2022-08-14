package global

import (
	"sync"

	"github.com/shuyangzhang/kyouka-light/domain/dto"
)

var (
	Mutex         sync.Mutex
	Instance      *dto.VoiceInstance
	PlayerChannel chan dto.PlayerSignal
)

func init() {
	Mutex = sync.Mutex{}
	Instance = &dto.VoiceInstance{
		QueueMutex: sync.Mutex{},
		ChannelId:  "",
		NowPlaying: dto.Music{},
		Queue:      []dto.Music{},
		PlayedTime: 0,
	}
	PlayerChannel = make(chan dto.PlayerSignal)
}
