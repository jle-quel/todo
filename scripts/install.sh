#!/usr/bin/env bash

C_RED="\033[31;01m"
C_GREEN="\033[32;01m"
C_YELLOW="\033[33;01m"
C_BLUE="\033[34;01m"
C_PINK="\033[35;01m"
C_CYAN="\033[36;01m"
C_NO="\033[0m"

os=$(uname)
cpu=$(uname -m)

################################################################################
###                                FUNCTIONS                                 ###
################################################################################

function	_err {
	if [ $1 != 0 ]; then
		echo $2
		exit 1
	fi
}

function	install {
	git clone "https://github.com/jle-quel/todo" /tmp/todo 1>/dev/null 2>/dev/null
	result=$?
	_err ${result} "todo: error while cloning the project"

	mv /tmp/todo/bin/$1 /usr/local/bin/todo 1>/dev/null 2>/dev/null
	result=$?
	_err ${result} "todo: error while adding binary to /usr/local/bin"

	echo
	echo "Successful installation 🚀 "
}

################################################################################
###                                   MAIN                                   ###
################################################################################

if [ ${cpu} != "x86_64" ]; then
	echo "64_bit CPU only supported 😬 "
	echo
	echo "If you want todo for your architecture, create an issue on \"https://github.com/jle-quel/todo\""
	exit 0
fi

case ${os} in
	"Linux"		)
		echo "Installing for Darwin 🍏 "
		install "todo_darwin_64bit" ;;
	"Darwin"	)
		echo "Installing for Linux 🐧 "
		install "todo_linux_64bit" ;;
	*			)
		"💻  → 🗑 " ;;
esac
