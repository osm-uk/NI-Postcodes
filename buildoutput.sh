#!/usr/bin/env bash

cat data/*.csv | sort -ru > data/NI-postcodes.temp
mv data/NI-postcodes.temp data/NI-postcodes.csv
