@echo build and run
@cd ..

@echo call build
@call:build logger -Type=1 -Id=1 -CenterAddr="127.0.0.1:10050"
@call:build center -Type=2 -Id=50 -CenterAddr="127.0.0.1:10050"
@call:build config -Type=3 -Id=60 -CenterAddr="127.0.0.1:10050"
@call:build gateway -Type=4 -Id=100 -CenterAddr="127.0.0.1:10050"
@call:build login -Type=5 -Id=70 -CenterAddr="127.0.0.1:10050"
@call:build list -Type=6 -Id=80 -CenterAddr="127.0.0.1:10050"
@call:build property -Type=7 -Id=90 -CenterAddr="127.0.0.1:10050"
@call:build table -Type=8 -Id=1000 -CenterAddr="127.0.0.1:10050"
@call:build room -Type=9 -Id=2000 -CenterAddr="127.0.0.1:10050"
@call:build robot -Type=10 -Id=3000 -CenterAddr="127.0.0.1:10050"
@goto:eof

:build
@set appName=%~1
@cd .\cmd\%appName%\
@go build
@start .\%appName%.exe %~2=%~3 %~4=%~5 %~6=%~7
@cd ../..
@goto:eof
