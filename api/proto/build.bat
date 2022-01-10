echo "go proto"
@.\protoc.exe --go_out=.. gate.proto
@.\protoc.exe --go_out=.. center.proto
@.\protoc.exe --go_out=.. login.proto
@.\protoc.exe --go_out=.. types.proto
@.\protoc.exe --go_out=.. logger.proto
@.\protoc.exe --go_out=.. config.proto
@.\protoc.exe --go_out=.. room.proto
@.\protoc.exe --go_out=.. table.proto
@.\protoc.exe --go_out=.. gameddz.proto
@.\protoc.exe --go_out=.. list.proto
@.\protoc.exe --go_out=.. property.proto

pause
