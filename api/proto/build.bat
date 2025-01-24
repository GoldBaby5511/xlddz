echo "go proto"

@.\protoc.exe --go_out=../../.. types.proto
@.\protoc.exe --go_out=../../.. center.proto
@.\protoc.exe --go_out=../../.. config.proto
@.\protoc.exe --go_out=../../.. gameddz.proto
@.\protoc.exe --go_out=../../.. gateway.proto
@.\protoc.exe --go_out=../../.. list.proto
@.\protoc.exe --go_out=../../.. lobby.proto
@.\protoc.exe --go_out=../../.. logger.proto
@.\protoc.exe --go_out=../../.. property.proto
@.\protoc.exe --go_out=../../.. room.proto
@.\protoc.exe --go_out=../../.. table.proto

@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf types.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf center.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf config.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf gateway.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf list.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf lobby.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf logger.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf property.proto
@.\protoc.exe --cpp_out=../../../mango-cpp/Common/protobuf room.proto

@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol types.proto
@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol center.proto
@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol list.proto
@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol lobby.proto
@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol gateway.proto
@.\protoc.exe --csharp_out=../../../mango-client/Assets/Scripts/Game/Protocol room.proto

pause
