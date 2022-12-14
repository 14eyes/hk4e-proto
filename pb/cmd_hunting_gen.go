// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(TakeHuntingOfferReq) },
		func() ProtoMessage { return new(TakeHuntingOfferRsp) },
		func() ProtoMessage { return new(GetCityHuntingOfferReq) },
		func() ProtoMessage { return new(GetCityHuntingOfferRsp) },
		func() ProtoMessage { return new(GetHuntingOfferRewardReq) },
		func() ProtoMessage { return new(GetHuntingOfferRewardRsp) },
		func() ProtoMessage { return new(HuntingStartNotify) },
		func() ProtoMessage { return new(HuntingRevealClueNotify) },
		func() ProtoMessage { return new(HuntingRevealFinalNotify) },
		func() ProtoMessage { return new(HuntingSuccessNotify) },
		func() ProtoMessage { return new(HuntingFailNotify) },
		func() ProtoMessage { return new(HuntingOngoingNotify) },
		func() ProtoMessage { return new(HuntingGiveUpReq) },
		func() ProtoMessage { return new(HuntingGiveUpRsp) },
	)
}

const (
	ProtoCommandTakeHuntingOfferReq      ProtoCommand = 4326
	ProtoCommandTakeHuntingOfferRsp      ProtoCommand = 4318
	ProtoCommandGetCityHuntingOfferReq   ProtoCommand = 4325
	ProtoCommandGetCityHuntingOfferRsp   ProtoCommand = 4307
	ProtoCommandGetHuntingOfferRewardReq ProtoCommand = 4302
	ProtoCommandGetHuntingOfferRewardRsp ProtoCommand = 4331
	ProtoCommandHuntingStartNotify       ProtoCommand = 4329
	ProtoCommandHuntingRevealClueNotify  ProtoCommand = 4322
	ProtoCommandHuntingRevealFinalNotify ProtoCommand = 4344
	ProtoCommandHuntingSuccessNotify     ProtoCommand = 4349
	ProtoCommandHuntingFailNotify        ProtoCommand = 4320
	ProtoCommandHuntingOngoingNotify     ProtoCommand = 4345
	ProtoCommandHuntingGiveUpReq         ProtoCommand = 4341
	ProtoCommandHuntingGiveUpRsp         ProtoCommand = 4342
)

func (*TakeHuntingOfferReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeHuntingOfferReq }
func (*TakeHuntingOfferReq) ProtoMessageType() ProtoMessageType { return "TakeHuntingOfferReq" }

func (*TakeHuntingOfferRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeHuntingOfferRsp }
func (*TakeHuntingOfferRsp) ProtoMessageType() ProtoMessageType { return "TakeHuntingOfferRsp" }

func (*GetCityHuntingOfferReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetCityHuntingOfferReq }
func (*GetCityHuntingOfferReq) ProtoMessageType() ProtoMessageType { return "GetCityHuntingOfferReq" }

func (*GetCityHuntingOfferRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetCityHuntingOfferRsp }
func (*GetCityHuntingOfferRsp) ProtoMessageType() ProtoMessageType { return "GetCityHuntingOfferRsp" }

func (*GetHuntingOfferRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHuntingOfferRewardReq
}
func (*GetHuntingOfferRewardReq) ProtoMessageType() ProtoMessageType {
	return "GetHuntingOfferRewardReq"
}

func (*GetHuntingOfferRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetHuntingOfferRewardRsp
}
func (*GetHuntingOfferRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GetHuntingOfferRewardRsp"
}

func (*HuntingStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingStartNotify }
func (*HuntingStartNotify) ProtoMessageType() ProtoMessageType { return "HuntingStartNotify" }

func (*HuntingRevealClueNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHuntingRevealClueNotify
}
func (*HuntingRevealClueNotify) ProtoMessageType() ProtoMessageType { return "HuntingRevealClueNotify" }

func (*HuntingRevealFinalNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHuntingRevealFinalNotify
}
func (*HuntingRevealFinalNotify) ProtoMessageType() ProtoMessageType {
	return "HuntingRevealFinalNotify"
}

func (*HuntingSuccessNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingSuccessNotify }
func (*HuntingSuccessNotify) ProtoMessageType() ProtoMessageType { return "HuntingSuccessNotify" }

func (*HuntingFailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingFailNotify }
func (*HuntingFailNotify) ProtoMessageType() ProtoMessageType { return "HuntingFailNotify" }

func (*HuntingOngoingNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingOngoingNotify }
func (*HuntingOngoingNotify) ProtoMessageType() ProtoMessageType { return "HuntingOngoingNotify" }

func (*HuntingGiveUpReq) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingGiveUpReq }
func (*HuntingGiveUpReq) ProtoMessageType() ProtoMessageType { return "HuntingGiveUpReq" }

func (*HuntingGiveUpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHuntingGiveUpRsp }
func (*HuntingGiveUpRsp) ProtoMessageType() ProtoMessageType { return "HuntingGiveUpRsp" }
