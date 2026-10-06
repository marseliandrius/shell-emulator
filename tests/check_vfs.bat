@echo off
setlocal
pushd "%~dp0.."
if errorlevel 1 exit /b 1

if not exist "tests/startup_vfs.txt" goto failed

for %%V in ("data/minimal.csv" "data/vfs.csv" "data/nested.csv") do (
    echo === VFS: %%~V ===
    call ".\run.bat" -vfs "%%~V" -script "tests/startup_vfs.txt"
    if errorlevel 1 goto failed
)

echo VFS tests passed.
popd
exit /b 0

:failed
echo VFS tests failed.
popd
exit /b 1