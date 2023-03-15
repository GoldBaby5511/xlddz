echo "go proto"
@.\protoc.exe --go_out=.. gateway.proto
@.\protoc.exe --go_out=plugins=grpc:.. center.proto
@.\protoc.exe --go_out=.. types.proto
@.\protoc.exe --go_out=.. logger.proto
@.\protoc.exe --go_out=.. config.proto

pause
