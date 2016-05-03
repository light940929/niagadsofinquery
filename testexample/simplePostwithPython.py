"""NIAGADSOFINQUERY API application.
   simplePostwithPython.py get -n <titlename> -i <individualnum> -s <snpnum> -f <tfampath> -p <tpedpath> -a <apitoken>
Usage:
  simplePostwithPython.py get -n <titlename> -i <individualnum> -s <snpnum> -f <tfampath> -p <tpedpath> -a <apitoken>
  simplePostwithPython.py (-h | --help)
  simplePostwithPython.py (-v | --version)
Options:
  -n --titlename <titlename> input title
  -i --individualnum <individualnum> input individual num
  -s --snpnum <snpnum> input snp num
  -f --tfampath <tfampath> input tfam path
  -p --tpedpath <tpedpath> input tped path
  -a --apitoken <apitoken> input api token
  -h --help     show this screen
  -v --version     show version and exit
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
url_phenotypes = 'http://localhost:9000/api/phenotypes'
url_genotypes = 'http://localhost:9000/api/genotypes'
token = 'Bearer ' + arguments['--apitoken']
headers = {'Authorization': '%s' % token}
request_phenotypes = Request(url_phenotypes, headers=headers)
request_genotypes = Request(url_genotypes, headers=headers)

response_phenotypes = urlopen(request_phenotypes)
response_genotypes = urlopen(request_genotypes)

data_phenotypes = json.loads(response_phenotypes.read())
data_genotypes = json.loads(response_genotypes.read())


def  postPhenotypes(url_phenotypes, token, headers):

    #tfam included familyID individualID dadID momID sex 1or2 or other phenotype 1or2 means case or control
    list = []
    lines = [line.strip() for line in open(arguments['--tfampath'])]
    for line in lines:
        ids=line.split()
        #print ids
        print "{title:"+arguments['--titlename']+",family_id:"+ids[0]+",individual_id:"+ids[1]+",paternal_id:"+ids[2]+",maternal_id:"+ids[3]+",sex:"+ids[4]+",affection_status:"+ids[5]+"}"
        values = {"title": arguments['--titlename'], "family_id": ids[0], "individual_id": ids[1], "paternal_id": ids[2], "maternal_id": ids[3], "sex": ids[4], "affection_status": ids[5]}
        data = json.dumps(values)
        print data
        req = requests.post(url_phenotypes, data, headers=headers)
        print  req.status_code



def  postGenotypes(url_genotypes, token, headers):

    list = []
    lines = [line.strip() for line in open(arguments['--tpedpath'])]
    for line in lines:
        ids=line.split()
        indnum=int(arguments['--individualnum'])
        snpnum=int(arguments['--snpnum'])
        num = indnum*snpnum
        #print ids
        strina = ''.join(ids[4:num+4])
        call = strina.strip(',')
        print "{title:"+arguments['--titlename']+",chr:"+ids[0]+",variant_id:"+ids[1]+",location:"+ids[2]+",coordinate:"+ids[3]+",call:"+call+"}"
        values = {"title": arguments['--titlename'], "chr": ids[0], "variant_id": ids[1], "location": ids[2], "coordinate": ids[3], "call": call}
        data = json.dumps(values)
        print data
        req = requests.post(url_genotypes, data, headers=headers)
        print  req.status_code



postPhenotypes(url_phenotypes, token, headers)
postGenotypes(url_genotypes, token, headers)
