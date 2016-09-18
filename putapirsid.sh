python ./getinforsid.py get -m $1
less ./rsid$1/ID_list_raw.f | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}' > ./rsid$1/ID_list.f
less ./rsid$1/snp_ID_list_raw | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}' >  ./rsid$1/snp_ID_list

##add back
#python ./extract_rsID.py -i ./rsid$1/ID_list.f -v ./rsid$1/snp_ID_list -j $1


# ./putapifile.sh  ./200/ID.output ./200/example.merged.tped  ./200/snp_na "1"
curl  -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"testh","email":"testh@gmail.com","password":"testh"}' http://localhost:9000/login   --stderr /Users/hannahlin/Desktop/curl_err.log | python -c "import json,sys;obj=json.load(sys.stdin);print obj['token'];" > ./token
token=$(cat ./token)

# ./ID.output
niagads_id=`awk '{print $3}' ./rsid$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
individual_id=`awk '{print $2}' ./rsid$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
dataset_id=`awk '{print $1}' ./rsid$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`

# ./example.merged.tped
snp_id=`awk '{print $2}' ./rsid$1/example.merged.tped  | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
genotypes=`awk '{for ( i=5; i <= NF; i++){ if(i==NF) print $i","; else print $i} }' ./rsid$1/example.merged.tped  | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str   a[i,j]}print str}}'`

# ./snp_na
snp_na_id=`awk '{print $1}' ./rsid$1/snp_na | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`


echo  $niagads_id
echo  $individual_id
echo  $snp_id
echo  $genotypes
echo  $snp_na_id

# url=http://localhost:9000/api/genotypebyrsids/$1
# contact={"id":"$1","niagads_id":"${niagads_id}","individual_id":"${individual_id}","dataset_id":"${dataset_id}","snp_id":"${snp_id}","snp_na_id":"${snp_na_id}","genotypes":"${genotypes}"}
# echo $contact
# echo $contact | python -mjson.tool > contact.json
# echo curl  -H "Accept: application/json" -H "Content-type: application/json" -H "Authorization: Bearer "$token  -X PUT -d ./contact.json $url

curl  -H "Accept: application/json" -H "Content-type: application/json" -H "Authorization: Bearer $token"  -X PUT -d "{\"id\": \"$1\", \"niagads_id\": \"${niagads_id}\",\"individual_id\": \"${individual_id}\",\"dataset_id\": \"${dataset_id}\",\"snp_id\": \"${snp_id}\",\"snp_na_id\": \"${snp_na_id}\", \"genotypes\": \"${genotypes}\"}" $url
