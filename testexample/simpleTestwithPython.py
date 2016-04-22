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

url_phenotypes = 'http://localhost:9000/api/phenotypes'
url_genotypes = 'http://localhost:9000/api/genotypes'
token = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Ik5JQUdBRFMiLCJleHAiOjE0NjEzNjI0NTV9.-Roix0YvuPy9VHaWm9wE83yB7NiSunyVXsVlR74lu2Y'
headers = {'Authorization': '%s' % token}
request_phenotypes = Request(url_phenotypes, headers=headers)
request_genotypes = Request(url_genotypes, headers=headers)

response_phenotypes = urlopen(request_phenotypes)
response_genotypes = urlopen(request_genotypes)

data_phenotypes = json.loads(response_phenotypes.read())
data_genotypes = json.loads(response_genotypes.read())


def  loadPhenotypes(data_phenotypes):
    phenotypes_list = data_phenotypes['phenotypes']
    for phenotype in phenotypes_list:
       print(phenotype['title'])
       print(phenotype['family_id'])
       print(phenotype['individual_id'])
       print(phenotype['paternal_id'])
       print(phenotype['maternal_id'])


def  loadGenotypes(data_genotypes):
    genotypes_list = data_genotypes['genotypes']
    for genotype in genotypes_list:
       print(genotype['title'])
       print(genotype['chr'])
       print(genotype['coordinate'])
       print(genotype['variant_id'])

def  postGenotypes(url_genotypes, token, headers):
    values = {"title":"test","chr":"2","variant_id":"snp4","location":"0","coordinate":"1111830","call":"G T G T G G T T G T T T"}
    data = json.dumps(values)
    req = requests.post(url_genotypes, data, headers=headers)
    print  req.status_code


loadPhenotypes(data_phenotypes)
loadGenotypes(data_genotypes)
postGenotypes(url_genotypes, token, headers)
