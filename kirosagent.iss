[Setup]
AppName=Kiros Agent
AppVersion=0.1
DefaultDirName=kirosagent
UninstallDisplayIcon={app}\kirosagent.exe
Compression=lzma2
SolidCompression=yes
OutputDir=userdocs:Inno Setup Examples Output
PrivilegesRequired=admin

[Files]
Source: "kirosagent.exe"; DestDir: "{app}"

[Run]
Filename: "{app}\kirosagent.exe"; Parameters: "install"; Flags: runascurrentuser runhidden