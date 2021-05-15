:: Made by SirSAC for DisPro.
echo off
:: Clear Screen.
cls
:: Change Directory.
chdir /d %~dp0
:: Build.
cls && set "goos=linux" && set "goarch=amd64" && go.exe clean -cache -testcache -modcache && del .\resource.syso && del .\Lin_64\DisPro.bin && syso.exe -c .\resource.json -o .\resource.syso && attrib.exe -r +a -i .\resource.syso /d /l && compact.exe /c /a /i .\resource.syso && go.exe build -o .\Lin_64\DisPro.bin -i -p 4 -x -buildmode pie -compiler gc -gcflags="-dwarf=false -e=true -l=false -pack=false -w=true -wb=false" -ldflags="-aslr=true -benchmark cpu -compressdwarf=true -d=false -extld gcc -f=false -g=false -linkmode internal -n=false -race=false -s=true -strictdups 1 -w=true" .\main.go .\socks.go .\linux.go .\const.go && attrib.exe -r +a -i .\Lin_64\DisPro.bin /d /l && compact.exe /c /a /i .\Lin_64\DisPro.bin
cls && set "goos=darwin" && set "goarch=amd64" && go.exe clean -cache -testcache -modcache && del .\resource.syso && del .\Dar_64\DisPro.bin && syso.exe -c .\resource.json -o .\resource.syso && attrib.exe -r +a -i .\resource.syso /d /l && compact.exe /c /a /i .\resource.syso && go.exe build -o .\Dar_64\DisPro.bin -i -p 4 -x -buildmode pie -compiler gc -gcflags="-dwarf=false -e=true -l=false -pack=false -race=false -spectre all -w=true -wb=false" -ldflags="-aslr=true -benchmark cpu -compressdwarf=true -d=false -extld gcc -f=false -g=false -linkmode internal -n=false -race=false -s=true -strictdups 1 -w=true" .\main.go .\socks.go .\darwin.go .\const.go && attrib.exe -r +a -i .\Dar_64\DisPro.bin /d /l && compact.exe /c /a /i .\Dar_64\DisPro.bin
cls && set "goos=windows" && set "goarch=amd64" && go.exe clean -cache -testcache -modcache && del .\resource.syso && del .\Win_64\DisPro.exe && syso.exe -c .\resource.json -o .\resource.syso && attrib.exe -r +a -i .\resource.syso /d /l && compact.exe /c /a /i .\resource.syso && go.exe build -o .\Win_64\DisPro.exe -i -p 4 -x -buildmode pie -compiler gc -gcflags="-c 4 -dwarf=false -e=true -l=false -pack=false -race=false -spectre all -w=true -wb=false" -ldflags="-aslr=true -benchmark cpu -compressdwarf=true -d=false -extld gcc -f=false -g=false -linkmode external -msan=true -n=false -race=false -s=true -strictdups 1 -w=true" .\main.go .\socks.go .\windows.go .\const.go && attrib.exe -r +a -i .\Win_64\DisPro.exe /d /l && compact.exe /c /a /i .\Win_64\DisPro.exe
:: Timeout.
timeout.exe /t 1
:: Echoing.
echo:
