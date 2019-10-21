#!/bin/bash

set -e

cd /client
cp LICENSE build/innosetup/
wine '/root/.wine/drive_c/Program Files/Inno Setup 6/ISCC.exe' build/innosetup/phrase-client.iss
wine '/root/.wine/drive_c/Program Files/Inno Setup 6/ISCC.exe' build/innosetup/phrase-client-386.iss
mv build/innosetup/Output/phrase_setup_386.exe dist/phrase_windows_setup_386.exe
mv build/innosetup/Output/phrase_setup.exe dist/phrase_windows_setup.exe
rm build/innosetup/LICENSE
rm -rf build/innosetup/Output
