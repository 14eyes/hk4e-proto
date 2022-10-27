// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(RechargeReq) },
		func() ProtoMessage { return new(RechargeRsp) },
		func() ProtoMessage { return new(OrderFinishNotify) },
		func() ProtoMessage { return new(CardProductRewardNotify) },
		func() ProtoMessage { return new(PlayerRechargeDataNotify) },
		func() ProtoMessage { return new(OrderDisplayNotify) },
		func() ProtoMessage { return new(ReportTrackingIOInfoNotify) },
		func() ProtoMessage { return new(TakeResinCardDailyRewardReq) },
		func() ProtoMessage { return new(TakeResinCardDailyRewardRsp) },
		func() ProtoMessage { return new(ResinCardDataUpdateNotify) },
	)
}

const (
	ProtoCommandRechargeReq                 ProtoCommand = 4126
	ProtoCommandRechargeRsp                 ProtoCommand = 4118
	ProtoCommandOrderFinishNotify           ProtoCommand = 4125
	ProtoCommandCardProductRewardNotify     ProtoCommand = 4107
	ProtoCommandPlayerRechargeDataNotify    ProtoCommand = 4102
	ProtoCommandOrderDisplayNotify          ProtoCommand = 4131
	ProtoCommandReportTrackingIOInfoNotify  ProtoCommand = 4129
	ProtoCommandTakeResinCardDailyRewardReq ProtoCommand = 4122
	ProtoCommandTakeResinCardDailyRewardRsp ProtoCommand = 4144
	ProtoCommandResinCardDataUpdateNotify   ProtoCommand = 4149
)

func (*RechargeReq) ProtoCommand() ProtoCommand         { return ProtoCommandRechargeReq }
func (*RechargeReq) ProtoMessageType() ProtoMessageType { return "RechargeReq" }

func (*RechargeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRechargeRsp }
func (*RechargeRsp) ProtoMessageType() ProtoMessageType { return "RechargeRsp" }

func (*OrderFinishNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOrderFinishNotify }
func (*OrderFinishNotify) ProtoMessageType() ProtoMessageType { return "OrderFinishNotify" }

func (*CardProductRewardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCardProductRewardNotify
}
func (*CardProductRewardNotify) ProtoMessageType() ProtoMessageType { return "CardProductRewardNotify" }

func (*PlayerRechargeDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerRechargeDataNotify
}
func (*PlayerRechargeDataNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerRechargeDataNotify"
}

func (*OrderDisplayNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOrderDisplayNotify }
func (*OrderDisplayNotify) ProtoMessageType() ProtoMessageType { return "OrderDisplayNotify" }

func (*ReportTrackingIOInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReportTrackingIOInfoNotify
}
func (*ReportTrackingIOInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ReportTrackingIOInfoNotify"
}

func (*TakeResinCardDailyRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeResinCardDailyRewardReq
}
func (*TakeResinCardDailyRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeResinCardDailyRewardReq"
}

func (*TakeResinCardDailyRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeResinCardDailyRewardRsp
}
func (*TakeResinCardDailyRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeResinCardDailyRewardRsp"
}

func (*ResinCardDataUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandResinCardDataUpdateNotify
}
func (*ResinCardDataUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "ResinCardDataUpdateNotify"
}
