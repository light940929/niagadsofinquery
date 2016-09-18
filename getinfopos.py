#!/usr/bin/env python
# -*- coding: UTF-8 -*-
"""NIAGADSOFINQUERY API application.
  exampleforjacky.py get  -m <api_id>
Usage:
  exampleforjacky.py get  -m <api_id>
  exampleforjacky.py (-h | --help)
  exampleforjacky.py (-v | --version)
Options:
  -m  --api_id <api_id> input api_id
  -h --help     show this screen
  -v --version  show version and exit
"""
import os
import re
import json
import sys
import getopt
import argparse
from docopt import docopt
from urllib2 import urlopen, Request
import urllib
import urllib2
import requests
from numpy import genfromtxt, savetxt


url_login = 'http://localhost:9000/login'
values = {"name":"testh","email":"testh@gmail.com","password":"testh"}
data = json.dumps(values)
headers = {'Content-Type': 'application/json'}
request_login = Request(url_login, data, headers)
response_login = urlopen(request_login)
data_login = json.loads(response_login.read())
key = data_login['token']


arguments = docopt(__doc__, version='0.0.1')
#print arguments

url_genotypebypositions = 'http://localhost:9000/api/genotypebypositions/' + arguments['<api_id>']
print url_genotypebypositions
token = 'Bearer ' + key
headers = {'Authorization': '%s' % token}
request_genotypebypositions = Request(url_genotypebypositions, headers=headers)
response_genotypebypositions = urlopen(request_genotypebypositions)
data_genotypebypositions = json.loads(response_genotypebypositions.read())
print data_genotypebypositions


def getGenotypebypositions(url_genotypebypositions, token, headers):
    jobid = data_genotypebypositions['success']['id']
    niagads_id = data_genotypebypositions['success']['niagads_id']
    individual_id = data_genotypebypositions['success']['individual_id']
    dataset_id = data_genotypebypositions['success']['dataset_id']
    snp_pos = data_genotypebypositions['success']['snp_pos']
    snp_pos_na = data_genotypebypositions['success']['snp_pos_na']
    genotypes = data_genotypebypositions['success']['genotypes']

    os.system("mkdir"+" "+"pos"+arguments['<api_id>'])
    ID=open('./'+"pos"+arguments['<api_id>']+'/ID_list_raw.f','w')
    ID.write('%s\n'%(niagads_id))
    ID.write('%s\n'%(individual_id))
    ID.write('%s\n'%(dataset_id))
    ID.close()

    SNP=open('./'+"pos"+arguments['<api_id>']+'/snp_ID_list_raw','w')
    SNP.write('%s\n'%(snp_pos))
    SNP.close()

getGenotypebypositions(url_genotypebypositions, token, headers)


#awk '{print $2}' ./example.merged.tped | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'
