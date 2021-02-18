# see: http://nsis.sourceforge.net/A_simple_installer_with_start_menu_shortcut_and_uninstaller
!define APPNAME "Frothy"
!define COMPANYNAME "Left Shift Logical, LLC."
!define DESCRIPTION "Cross-Platform TOTP client."
!addplugindir .\

# require admin rights
RequestExecutionLevel admin

SetCompressor /SOLID lzma

!include LogicLib.nsh

!macro VerifyUserIsAdmin
UserInfo::GetAccountType
pop $0
${If} $0 != "admin" # Require admin rights on NT4+
  messageBox mb_iconstop "Administrator rights required!"
  setErrorLevel 740 # ERROR_ELEVATION_REQUIRED
  quit
${EndIf}
!macroend

function .onInit
  # make global install
  setShellVarContext all

  # ensure admin
  !insertmacro VerifyUserIsAdmin

  StrCpy $InstDir "$PROGRAMFILES64\${APPNAME}"
functionEnd

function un.onInit
	SetShellVarContext all
 
	# Verify the uninstaller - last chance to back out
	MessageBox MB_OKCANCEL "Are you sure you want to uninstall ${APPNAME}?" IDOK next
		Abort
	next:
	!insertmacro VerifyUserIsAdmin
functionEnd

!include install.nsh
!include uninstall.nsh