#!/bin/bash

set -e

cd /client
cp LICENSE innosetup/
wine '/root/.wine/drive_c/Program Files/Inno Setup 5/ISCC.exe' innosetup/phraseapp-client.iss
mv innosetup/Output/phraseapp_setup.exe dist/phraseapp_windows_setup.exe
rm innosetup/LICENSE
rm -rf innosetup/Output
