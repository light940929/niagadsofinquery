#!/bin/bash

awk '{print $1 " , "$2}' ./CBD.ped >  ./individualist.txt
awk '{print $2}' ./CBD.map >  ./snpID.txt
awk '{print $1 " , "$2 " , " $4}' ./CBD.map >  ./snplist.txt
