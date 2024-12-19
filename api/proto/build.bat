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

@.\protoc.exe --cpp_out=.. center.proto
@.\protoc.exe --cpp_out=.. config.proto
@.\protoc.exe --cpp_out=.. gameddz.proto
@.\protoc.exe --cpp_out=.. gateway.proto
@.\protoc.exe --cpp_out=.. list.proto
@.\protoc.exe --cpp_out=.. lobby.proto
@.\protoc.exe --cpp_out=.. logger.proto
@.\protoc.exe --cpp_out=.. property.proto
@.\protoc.exe --cpp_out=.. room.proto
@.\protoc.exe --cpp_out=.. table.proto

@.\protoc.exe --csharp_out=.. center.proto
@.\protoc.exe --csharp_out=.. list.proto
@.\protoc.exe --csharp_out=.. gateway.proto

pause
