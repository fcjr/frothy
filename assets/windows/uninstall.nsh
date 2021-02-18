Section "Uninstall"
  # uninstall for all users
  setShellVarContext all

  # TODO close program

  # delete start menu items
  Delete "$SMPROGRAMS\${APPNAME}\${APPNAME}.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Uninstall.lnk"
  rmDir "$SMPROGRAMS\${APPNAME}"
  
  # delete install dir
	delete $INSTDIR\Frothy.exe
	delete $INSTDIR\logo.ico
  rmDir $INSTDIR

  # delete registry keys
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${COMPANYNAME} ${APPNAME}"

  # delete uninstaller
  Delete $INSTDIR\uninstall.exe
SectionEnd