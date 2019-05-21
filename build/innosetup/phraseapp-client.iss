; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{DC598D8E-8A9B-4CAD-AFD8-0324FDF4E0F1}
AppName=PhraseApp Client
AppVersion=1.14.0
AppPublisher=PhraseApp GmbH
AppPublisherURL=https://phraseapp.com/cli
AppSupportURL=https://phraseapp.com/cli
AppUpdatesURL=https://phraseapp.com/cli
ArchitecturesAllowed=x64
DefaultDirName={pf}\PhraseApp
DefaultGroupName=PhraseApp-Client
DisableProgramGroupPage=yes
LicenseFile=LICENSE
InfoAfterFile=postinstall.rtf
OutputBaseFilename=phraseapp_setup
SetupIconFile=parrot.ico
Compression=lzma
SolidCompression=yes

[Files]
Source: "../../dist/phraseapp_windows_amd64.exe"; DestDir: "{app}"; DestName: "phraseapp.exe"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; \

[Setup]
AlwaysRestart = yes

[Icons]
Name: "{group}\PhraseApp Client"; Filename: "{app}"
Name: "{group}\{cm:UninstallProgram,PhraseApp Client}"; Filename: "{uninstallexe}"
