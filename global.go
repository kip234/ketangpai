package main

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/JWT"
	"KeTangPai/services/DC/KetangpaiDB"
	"KeTangPai/services/DC/NetworkDisk"
	"KeTangPai/services/DC/TestBank"
	"KeTangPai/services/DC/UserCenter"
	"KeTangPai/services/Email"
	"KeTangPai/services/Filter"
	"KeTangPai/services/Forum"
	"KeTangPai/services/RankingList"
)

const Addr=":8080"

type Services struct{
	Filter Filter.FilterClient
	JWT JWT.JWTClient
	User UserCenter.UserCenterClient
	Exercise Exercise.ExerciseClient
	TestBank TestBank.TestBankClient
	NetworkDisk NetworkDisk.NetworkDiskClient
	Forum Forum.ForumClient
	RankingList RankingList.RankingListClient
	Email Email.EmailClient
	KetangpaiDB KetangpaiDB.KetangpaiDBClient
}
var services=Services{
	Filter:   Filter.NewFilterConn(),
	JWT:      JWT.NewJWTConn(),
	User:     UserCenter.NewUCSConn(),
	Exercise: Exercise.NewExerciseConn(),
	TestBank: TestBank.NewTestBankConn(),
	NetworkDisk: NetworkDisk.NewTestBankConn(),
	Forum:Forum.NewForumConn(),
	RankingList:RankingList.NewRankingListConn(),
	Email: Email.NewEmailConn(),
	KetangpaiDB: KetangpaiDB.NewKetangpaiDBConn(),
}