// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(WatcherAllDataNotify) },
		func() ProtoMessage { return new(WatcherChangeNotify) },
		func() ProtoMessage { return new(WatcherEventNotify) },
		func() ProtoMessage { return new(WatcherEventTypeNotify) },
		func() ProtoMessage { return new(WatcherEventStageNotify) },
		func() ProtoMessage { return new(PushTipsAllDataNotify) },
		func() ProtoMessage { return new(PushTipsChangeNotify) },
		func() ProtoMessage { return new(PushTipsReadFinishReq) },
		func() ProtoMessage { return new(PushTipsReadFinishRsp) },
		func() ProtoMessage { return new(GetPushTipsRewardReq) },
		func() ProtoMessage { return new(GetPushTipsRewardRsp) },
	)
}

const (
	ProtoCommandWatcherAllDataNotify    ProtoCommand = 2272
	ProtoCommandWatcherChangeNotify     ProtoCommand = 2298
	ProtoCommandWatcherEventNotify      ProtoCommand = 2212
	ProtoCommandWatcherEventTypeNotify  ProtoCommand = 2235
	ProtoCommandWatcherEventStageNotify ProtoCommand = 2207
	ProtoCommandPushTipsAllDataNotify   ProtoCommand = 2222
	ProtoCommandPushTipsChangeNotify    ProtoCommand = 2265
	ProtoCommandPushTipsReadFinishReq   ProtoCommand = 2204
	ProtoCommandPushTipsReadFinishRsp   ProtoCommand = 2293
	ProtoCommandGetPushTipsRewardReq    ProtoCommand = 2227
	ProtoCommandGetPushTipsRewardRsp    ProtoCommand = 2294
)

func (*WatcherAllDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWatcherAllDataNotify }
func (*WatcherAllDataNotify) ProtoMessageType() ProtoMessageType { return "WatcherAllDataNotify" }

func (*WatcherChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWatcherChangeNotify }
func (*WatcherChangeNotify) ProtoMessageType() ProtoMessageType { return "WatcherChangeNotify" }

func (*WatcherEventNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWatcherEventNotify }
func (*WatcherEventNotify) ProtoMessageType() ProtoMessageType { return "WatcherEventNotify" }

func (*WatcherEventTypeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWatcherEventTypeNotify }
func (*WatcherEventTypeNotify) ProtoMessageType() ProtoMessageType { return "WatcherEventTypeNotify" }

func (*WatcherEventStageNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWatcherEventStageNotify
}
func (*WatcherEventStageNotify) ProtoMessageType() ProtoMessageType { return "WatcherEventStageNotify" }

func (*PushTipsAllDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPushTipsAllDataNotify }
func (*PushTipsAllDataNotify) ProtoMessageType() ProtoMessageType { return "PushTipsAllDataNotify" }

func (*PushTipsChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPushTipsChangeNotify }
func (*PushTipsChangeNotify) ProtoMessageType() ProtoMessageType { return "PushTipsChangeNotify" }

func (*PushTipsReadFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandPushTipsReadFinishReq }
func (*PushTipsReadFinishReq) ProtoMessageType() ProtoMessageType { return "PushTipsReadFinishReq" }

func (*PushTipsReadFinishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPushTipsReadFinishRsp }
func (*PushTipsReadFinishRsp) ProtoMessageType() ProtoMessageType { return "PushTipsReadFinishRsp" }

func (*GetPushTipsRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetPushTipsRewardReq }
func (*GetPushTipsRewardReq) ProtoMessageType() ProtoMessageType { return "GetPushTipsRewardReq" }

func (*GetPushTipsRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetPushTipsRewardRsp }
func (*GetPushTipsRewardRsp) ProtoMessageType() ProtoMessageType { return "GetPushTipsRewardRsp" }
