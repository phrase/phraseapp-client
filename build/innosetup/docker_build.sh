#!/bin/bash

set -e

cd /client
cp LICENSE build/innosetup/
wine '/root/.wine/drive_c/Program Files/Inno Setup 5/ISCC.exe' build/innosetup/phraseapp-client.iss
wine '/root/.wine/drive_c/Program Files/Inno Setup 5/ISCC.exe' build/innosetup/phraseapp-client-386.iss
mv build/innosetup/Output/phraseapp_setup_386.exe dist/phraseapp_windows_setup_386.exe
mv build/innosetup/Output/phraseapp_setup.exe dist/phraseapp_windows_setup.exe
rm build/innosetup/LICENSE
rm -rf build/innosetup/Output
