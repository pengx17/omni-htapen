REM Makes sure previous build is removed
RMDIR /s /q dist

ECHO Building kubecon-sh-demo:build
docker build -t pengxiao/omni-htapen:build -f .\dockerfiles\build.Dockerfile .
docker container create --name omni-htapen-build-extract pengxiao/omni-htapen:build
docker container cp omni-htapen-build-extract:./server .\dist
docker container rm omni-htapen-build-extract

ECHO Building kubecon-sh-demo:latest
docker build --no-cache -t pengxiao/omni-htapen:latest -f .\dockerfiles\Dockerfile .

RMDIR /s /q dist
