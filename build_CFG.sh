#!/usr/bin/env bash
cd sys-configs/
CFG_FILES=$(ls -d *.cfg)
FILES="$CFG_FILES"
if test -z $1; then
	echo "Param is NULL!"
else
	for F in $FILES; do
		Num=$(grep -n --line-number "ignores" $F | awk -F: '{print $1}')
		sed -i "${Num}a \\\t\\t\"$1\"," $F	
	done
	cd ..
	echo "Successful!"
fi
