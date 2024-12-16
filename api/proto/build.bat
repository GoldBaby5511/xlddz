echo "go proto"

@.\protoc.exe --go_out=.. center.proto
@.\protoc.exe --go_out=.. config.proto
@.\protoc.exe --go_out=.. gameddz.proto
@.\protoc.exe --go_out=.. gateway.proto
@.\protoc.exe --go_out=.. list.proto
@.\protoc.exe --go_out=.. lobby.proto
@.\protoc.exe --go_out=.. logger.proto
@.\protoc.exe --go_out=.. property.proto
@.\protoc.exe --go_out=.. room.proto
@.\protoc.exe --go_out=.. table.proto


pause
