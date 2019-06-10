@echo OFF

set "args=%*"
pushd "%~dp0"
setlocal ENABLEDELAYEDEXPANSION
set GOPATH="%~dp0vendor"
rem Set the GOPROXY environment variable
set GOPROXY=https://goproxy.io

if /i "%args%"=="install" goto install
if /i "%args%"=="build" goto build
if /i "%args%"=="run" goto run

goto DEFAULT_CASE
:install
    mkdir vendor
    GOTO END_CASE
:build
    go build -o main.exe main.go
    GOTO END_CASE
:run
    CALL go build -o main.exe main.go && CALL %~dp0\main.exe
    GOTO END_CASE
:DEFAULT_CASE
    echo Compiling
    CALL go build -ldflags "-s -w" -o main.exe main.go
    echo Compiled
    GOTO END_CASE
:END_CASE
    GOTO :EOF