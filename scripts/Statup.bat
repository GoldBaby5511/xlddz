@echo build and run
@cd ..

@echo call buildApp
@call:buildApp logger
@call:buildApp center
@call:buildApp config
@call:buildApp gateway
@call:buildApp login
@call:buildRT room /AppType 15  /AppID 1000 /CenterAddr 127.0.0.1:10001 /ListenOnAddr 0.0.0.0:11000
@call:buildRT table /AppType 14  /AppID 2000 /CenterAddr 127.0.0.1:10001 /ListenOnAddr 0.0.0.0:12000
@goto:eof

:buildApp
@set appName=%~1
@if not exist .\cmd\%appName%\configs mkdir .\cmd\%appName%\configs\%appName%
@xcopy .\configs\%appName%\ .\cmd\%appName%\configs\%appName%\ /s /f /h /y
@cd .\cmd\%appName%\
@go build
@start .\%appName%.exe
@cd ../..
@goto:eof

:buildRT
@set appName=%~1
@cd .\cmd\%appName%\
@go build
@start .\%appName%.exe %~2 %~3 %~4 %~5 %~6 %~7 %~8 %~9
@cd ../..
@goto:eof
