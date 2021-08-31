package main

import (
	"KeTangPai/Models"
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/JWT"
	"KeTangPai/services/DC/NetworkDisk"
	"KeTangPai/services/DC/TestBank"
	"KeTangPai/services/DC/UserCenter"
	"KeTangPai/services/Email"
	"KeTangPai/services/Filter"
	"KeTangPai/services/Forum"
	"KeTangPai/services/Log"
	"KeTangPai/services/RankingList"
)

func main()  {
	go Filter.Run()
	go JWT.Run()
	go UserCenter.Run()
	go TestBank.Run()
	go Exercise.Run()
	go NetworkDisk.Run()
	go Forum.Run()
	go RankingList.Run()
	go Email.Run()
	Log.Run()

	rooms:=make(map[int32]*Models.Room)

	server:= BuildRouter(services,rooms)
	server.Run(Addr)
}
