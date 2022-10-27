// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerRoutineDataNotify) },
		func() ProtoMessage { return new(WorldAllRoutineTypeNotify) },
		func() ProtoMessage { return new(WorldRoutineTypeRefreshNotify) },
		func() ProtoMessage { return new(WorldRoutineChangeNotify) },
		func() ProtoMessage { return new(WorldRoutineTypeCloseNotify) },
	)
}

const (
	ProtoCommandPlayerRoutineDataNotify       ProtoCommand = 3526
	ProtoCommandWorldAllRoutineTypeNotify     ProtoCommand = 3518
	ProtoCommandWorldRoutineTypeRefreshNotify ProtoCommand = 3525
	ProtoCommandWorldRoutineChangeNotify      ProtoCommand = 3507
	ProtoCommandWorldRoutineTypeCloseNotify   ProtoCommand = 3502
)

func (*PlayerRoutineDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerRoutineDataNotify
}
func (*PlayerRoutineDataNotify) ProtoMessageType() ProtoMessageType { return "PlayerRoutineDataNotify" }

func (*WorldAllRoutineTypeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldAllRoutineTypeNotify
}
func (*WorldAllRoutineTypeNotify) ProtoMessageType() ProtoMessageType {
	return "WorldAllRoutineTypeNotify"
}

func (*WorldRoutineTypeRefreshNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldRoutineTypeRefreshNotify
}
func (*WorldRoutineTypeRefreshNotify) ProtoMessageType() ProtoMessageType {
	return "WorldRoutineTypeRefreshNotify"
}

func (*WorldRoutineChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldRoutineChangeNotify
}
func (*WorldRoutineChangeNotify) ProtoMessageType() ProtoMessageType {
	return "WorldRoutineChangeNotify"
}

func (*WorldRoutineTypeCloseNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldRoutineTypeCloseNotify
}
func (*WorldRoutineTypeCloseNotify) ProtoMessageType() ProtoMessageType {
	return "WorldRoutineTypeCloseNotify"
}