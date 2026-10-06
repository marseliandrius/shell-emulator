@echo off
setlocal
pushd "%~dp0.."
if errorlevel 1 exit /b 1

echo === VFS parameter ===
call ".\run.bat" -vfs "data/vfs.csv" < "tests/startup_ok.txt"
if errorlevel 1 goto failed

echo === Script parameter ===
call ".\run.bat" -script "tests/startup_ok.txt"
if errorlevel 1 goto failed

echo === Both parameters ===
call ".\run.bat" -vfs "data/nested.csv" -script "tests/startup_ok.txt"
if errorlevel 1 goto failed

echo Parameter tests passed.
popd
exit /b 0

:failed
echo Parameter tests failed.
popd
exit /b 1