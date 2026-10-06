@echo off
setlocal
pushd "%~dp0.."
if errorlevel 1 exit /b 1

if not exist "tests/startup_error.txt" goto failed
if not exist "tests/startup_bad_quotes.txt" goto failed
if exist "tests/missing_script.txt" goto failed

echo === Unknown command ===
call ".\run.bat" -vfs "data/vfs.csv" -script "tests/startup_error.txt"
if not errorlevel 1 goto failed

echo === Missing script ===
call ".\run.bat" -vfs "data/vfs.csv" -script "tests/missing_script.txt"
if not errorlevel 1 goto failed

echo === Unclosed quotes ===
call ".\run.bat" -vfs "data/vfs.csv" -script "tests/startup_bad_quotes.txt"
if not errorlevel 1 goto failed

echo Expected-error tests passed.
popd
exit /b 0

:failed
echo Expected-error tests failed.
popd
exit /b 1