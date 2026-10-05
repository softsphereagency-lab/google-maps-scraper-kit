@echo off
title Google Maps Scraper Engine
echo ========================================================
echo Starting Google Maps Scraper Engine on http://127.0.0.1:8080
echo Keep this window open while scraping!
echo Web UI: http://127.0.0.1:8080
echo ========================================================
set "PATH=%USERPROFILE%\go\bin;%USERPROFILE%\go_runtime\go\bin;%PATH%"
if not exist "gmapsdata" mkdir "gmapsdata"
google-maps-scraper.exe -web -addr 127.0.0.1:8080 -data-folder .\gmapsdata
pause
