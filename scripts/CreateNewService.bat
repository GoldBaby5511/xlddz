@echo off & setlocal enabledelayedexpansion
@set /p var=ÇëÊäÈëÃû³Æ:
@cd ..
@if not exist .\cmd\%var% mkdir .\cmd\%var%
@if not exist .\cmd\%var%\business mkdir .\cmd\%var%\business
@cd scripts
for /f "tokens=*" %%i in (main.txt) do (
    if "%%i"=="" (echo.) else (set "line=%%i" & call :chg)
)>>..\cmd\%var%\main.go

@cd ..
@copy .\scripts\business.txt .\cmd\%var%\business\
@cd .\cmd\%var%
@go fmt
@cd business
@ren business.txt business.go
pause
exit
:chg 
set "line=!line:template=%var%!"
echo !line!
goto :eof

