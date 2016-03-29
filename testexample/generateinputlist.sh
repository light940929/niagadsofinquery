#!/bin/bash

awk '{print $1 " , "$2}' ./CBD.ped >  ./individualist.txt
awk '{print $2}' ./CBD.map >  ./snplist.txt
