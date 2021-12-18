@echo build and run
@echo build logger
@cd ..
@if not exist .\cmd\logger\configs mkdir .\cmd\logger\configs\logger
@xcopy .\configs\logger\ .\cmd\logger\configs\logger\ /s /f /h /y
@cd .\cmd\logger\
@go build
@start .\logger.exe
@cd ../..

@echo build center
@if not exist .\cmd\center\configs mkdir .\cmd\center\configs\center
@xcopy .\configs\center\ .\cmd\center\configs\center\ /s /f /h /y
@cd .\cmd\center\
@go build
@start .\center.exe
@cd ../..

@echo build config
@if not exist .\cmd\config\configs mkdir .\cmd\config\configs\config
@xcopy .\configs\config\ .\cmd\config\configs\config\ /s /f /h /y
@cd .\cmd\config\
@go build
@start .\config.exe
@cd ../..

@echo build gateway
@if not exist .\cmd\gateway\configs mkdir .\cmd\gateway\configs\gateway
@xcopy .\configs\gateway\ .\cmd\gateway\configs\gateway\ /s /f /h /y
@cd .\cmd\gateway\
@go build
@start .\gateway.exe
@cd ../..

@echo build login
@if not exist .\cmd\login\configs mkdir .\cmd\login\configs\login
@xcopy .\configs\login\ .\cmd\login\configs\login\ /s /f /h /y
@cd .\cmd\login\
@go build
@start .\login.exe
@cd ../..
