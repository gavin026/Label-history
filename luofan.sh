touch luoTemp.txt

cat struct.h | while read line
do
	echo $line > luotemp.txt
	line2=$(sed 's/\n//g' luotemp.txt)
	result=$(echo "${line2}" | grep "{")
	result2=$(echo "${line2}" | grep "}")
	if [[ "$result" != "" ]];then
		echo "$line2"  >> $1
	elif [[ "$result2" != "" ]];then
		echo "$line2" >> $1
	else
		echo "$line2" | awk -F " " '{print "\t"$2"\t\t"$1}' >> $1 
	fi

done
rm luoTemp.txt
echo "Good Luck!!!"



#!/usr/bin/env bash
cat data.dat | awk -F '/' '{print $3}' >TempN.txt  
echo "Start..."
for line in $(cat TempN.txt)
do 
	echo "Hello:$line"
	grep -r --include="*.c" name[[:space:]]=[[:space:]]\"$line\"  /home/mount/Ares  >> FilePath.txt
	grep -r --include="*.h" \"{$line}\"  /home/mount/Ares >> FilePath.txt
done
echo "End..."
rm TempN.txt

#!/usr/bin/env bash
#sort -n data.dat | uniq > temp.txt
#rm Ares/*
cat data.dat | awk -F '/' '{print $3}' >temp.txt
cat temp.txt | awk -F ':' '{print $1}' >temp0.txt
for line in $(cat temp0.txt)
do 
	cp xxx.txt Ares/cbg_$line.txt
	sed -i "s/XXX/${line}/g" Ares/cbg_$line.txt
done
rm temp0.txt
