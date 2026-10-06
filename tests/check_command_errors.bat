@echo off
setlocal
pushd "%~dp0.."
if errorlevel 1 exit /b 1

set "commandErrorLog="
if not exist "data/commands.csv" (
    echo Missing VFS fixture.
    goto failed
)

set "commandErrorLog=%TEMP%\shell-emulator-errors-%RANDOM%-%RANDOM%.log"

for %%S in (
    startup_ls_missing.txt
    startup_cd_file.txt
    startup_tac_no_args.txt
    startup_rev_utf8.txt
    startup_cal_month.txt
) do (
    call :check "tests/%%S"
    if errorlevel 1 goto failed
)

echo Command error tests passed.
del "%commandErrorLog%" >nul 2>&1
popd
exit /b 0

:check
if not exist "%~1" (
    echo Missing script: %~1
    exit /b 1
)

echo === %~1 ===
call ".\run.bat" -vfs "data/commands.csv" -script "%~1" >"%commandErrorLog%" 2>&1
if not errorlevel 1 (
    powershell -NoProfile -Command "Get-Content -LiteralPath $env:commandErrorLog -Encoding UTF8"
    echo Expected a nonzero exit code.
    exit /b 1
)

powershell -NoProfile -Command "Get-Content -LiteralPath $env:commandErrorLog -Encoding UTF8"
findstr /L /C:" 2:" "%commandErrorLog%" >nul
if errorlevel 1 (
    echo Expected an error on line 2.
    exit /b 1
)

findstr /L /C:"STOP_MARKER" "%commandErrorLog%" >nul
if not errorlevel 1 (
    echo Execution continued after the error.
    exit /b 1
)
exit /b 0

:failed
echo Command error tests failed.
if defined commandErrorLog del "%commandErrorLog%" >nul 2>&1
popd
exit /b 1