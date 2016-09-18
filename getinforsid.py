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


arguments = docopt(__doc__, version='0.0.1')
#print arguments

url_login = 'http://localhost:9000/login'
values = {"name":"testh","email":"testh@gmail.com","password":"testh"}
data = json.dumps(values)
headers = {'Content-Type': 'application/json'}
request_login = Request(url_login, data, headers)
response_login = urlopen(request_login)
data_login = json.loads(response_login.read())
key = data_login['token']



url_genotypebyrsids = 'http://localhost:9000/api/genotypebyrsids/' + arguments['<api_id>']
print url_genotypebyrsids
token = 'Bearer ' + key
headers = {'Authorization': '%s' % token}
request_genotypebyrsids = Request(url_genotypebyrsids, headers=headers)
response_genotypebyrsids = urlopen(request_genotypebyrsids)
data_genotypebyrsids = json.loads(response_genotypebyrsids.read())
print data_genotypebyrsids

def getGenotypebyrsids(url_genotypebyrsids, token, headers):
    jobid = data_genotypebyrsids['success']['id']
    niagads_id = data_genotypebyrsids['success']['niagads_id']
    individual_id = data_genotypebyrsids['success']['individual_id']
    dataset_id = data_genotypebyrsids['success']['dataset_id']
    snp_id = data_genotypebyrsids['success']['snp_id']
    snp_na_id = data_genotypebyrsids['success']['snp_na_id']
    genotypes = data_genotypebyrsids['success']['genotypes']

    os.system("mkdir"+" "+"rsid"+arguments['<api_id>'])
    ID=open('./'+"rsid"+arguments['<api_id>']+'/ID_list_raw.f','w')
    ID.write('%s\n'%(niagads_id))
    ID.write('%s\n'%(individual_id))
    ID.write('%s\n'%(dataset_id))
    ID.close()

    SNP=open('./'+"rsid"+arguments['<api_id>']+'/snp_ID_list_raw','w')
    SNP.write('%s\n'%(snp_id))
    SNP.close()

getGenotypebyrsids(url_genotypebyrsids, token, headers)
