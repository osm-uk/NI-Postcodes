#!/usr/bin/env bash

cat data/*.csv | sort -ru > data/NI-postcodes.temp

cat excludes.csv | while read line
do
#echo $line
	sed -i "/$line/d" data/NI-postcodes.temp
done

mv data/NI-postcodes.temp data/NI-postcodes.csv
