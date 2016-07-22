[Setup]
AppName=Kiros Agent
AppVersion=0.1
DefaultDirName={pf}\kirosagent
UninstallDisplayIcon={app}\kirosagent.exe
Compression=lzma2
SolidCompression=yes
OutputDir=userdocs:Inno Setup Examples Output
PrivilegesRequired=admin
UseSetupLdr=yes

[Files]
Source: "E:\gowork\bin\kirosagent.exe"; DestDir: "{app}"

[Run]
Filename: "{app}\kirosagent.exe"; Parameters: "install"; Flags: runascurrentuser runhidden
Filename: "{app}\kirosagent.exe"; Parameters: "start"; Flags: runascurrentuser runhidden

[UninstallRun]
Filename: "{app}\kirosagent.exe"; Parameters: "stop"; Flags: runascurrentuser runhidden
Filename: "{app}\kirosagent.exe"; Parameters: "remove"; Flags: runascurrentuser runhidden
