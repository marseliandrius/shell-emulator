@echo off
setlocal
pushd "%~dp0.."
if errorlevel 1 exit /b 1

if not exist "tests/startup_vfs.txt" goto failed
if not exist "tests/invalid_vfs.csv" goto failed
if exist "data/not_found.csv" goto failed

for %%S in ("tests/startup_error.txt" "tests/startup_bad_quotes.txt" "tests/startup_exit_error.txt") do (
    if not exist "%%~S" goto failed
)

for %%V in ("data/minimal.csv" "data/vfs.csv" "data/nested.csv") do (
    if not exist "%%~V" goto failed
    echo === VFS: %%~V ===
    for %%S in ("tests/startup_error.txt" "tests/startup_bad_quotes.txt" "tests/startup_exit_error.txt") do (
        echo === Script: %%~S ===
        call ".\run.bat" -vfs "%%~V" -script "%%~S"
        if not errorlevel 1 goto failed
    )
)

echo === Missing VFS ===
call ".\run.bat" -vfs "data/not_found.csv" -script "tests/startup_vfs.txt"
if not errorlevel 1 goto failed

echo === Invalid VFS ===
call ".\run.bat" -vfs "tests/invalid_vfs.csv" -script "tests/startup_vfs.txt"
if not errorlevel 1 goto failed

echo VFS error tests passed.
popd
exit /b 0

:failed
echo VFS error tests failed.
popd
exit /b 1