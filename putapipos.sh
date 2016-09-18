python ./getinfopos.py get -m $1
less ./pos$1/ID_list_raw.f | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str " "  a[i,j]}print str}}' > ./pos$1/ID_list.f
less ./pos$1/snp_ID_list_raw | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str " "  a[i,j]}print str}}' >  ./pos$1/snp_ID_list

##add back
#python ./extract_pos.py -i ./pos$1/ID_list.f -v ./pos$1/snp_ID_list -j $1

curl  -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"testh","email":"testh@gmail.com","password":"testh"}' http://localhost:9000/login   --stderr /Users/hannahlin/Desktop/curl_err.log | python -c "import json,sys;obj=json.load(sys.stdin);print obj['token'];" > ./token
token=$(cat ./token)

echo $token


# ./ID.output
niagads_id=`awk '{print $3}' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
individual_id=`awk '{print $2}' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
dataset_id=`awk '{print $1}' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/ID.output | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`

# ./example.merged.tped
snp_id=`awk '{print $2}' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/example.merged.tped  | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`
genotypes=`awk '{for ( i=5; i <= NF; i++){ if(i==NF) print $i","; else print $i} }' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/example.merged.tped | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str   a[i,j]}print str}}'`

# ./snp_na
snp_na_id=`awk '{print $1}' /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/pos$1/snp_na | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ","  a[i,j]}print str}}'`

echo  $niagads_id
echo  $individual_id
echo  $snp_id
echo  $genotypes
echo  $snp_na_id


url=http://localhost:9000/api/genotypebypositions/$1
echo {  >> ./contact.json
echo "  \"id\"": \"$1\", >> ./contact.json
echo "  \"niagads_id\"": \"${niagads_id}\", >> ./contact.json
echo "  \"individual_id\"": \"${individual_id}\",  >> ./contact.json
echo "  \"dataset_id\"": \"${dataset_id}\", >> ./contact.json
echo "  \"snp_pos\"": \"${snp_pos}\", >> ./contact.json
echo "  \"snp_pos_na\"": \"${snp_pos_na}\", >> ./contact.json
echo "  \"genotypes\"": \"${genotypes}\" >> ./contact.json
echo } >> ./contact.json

#contact={"id":"$1","niagads_id":"${niagads_id}","individual_id":"${individual_id}","dataset_id":"${dataset_id}","snp_pos":"${snp_pos}","snp_pos_na":"${snp_pos_na}","genotypes":"${genotypes}"}
cat ./contact.json
#echo $contact | python -mjson.tool > contact.json
#echo curl  -H "Accept: application/json" -H "Content-type: application/json" -H "Authorization: Bearer "$token  -X PUT -d ./contact.json $url

#curl  -H "Accept: application/json" -H "Content-type: application/json" -H "Authorization: Bearer $token"  -X PUT -d /Users/hannahlin/Documents/workspace/Go/src/github.com/user/niagadsofinquery/contact.json $url

curl  -H "Accept: application/json" -H "Content-type: application/json" -H "Authorization: Bearer $token"  -X PUT -d "{\"id\": \"$1\", \"niagads_id\": \"${niagads_id}\",\"individual_id\": \"${individual_id}\",\"dataset_id\": \"${dataset_id}\",\"snp_pos\": \"${snp_pos}\",\"snp_pos_na\": \"${snp_pos_na}\", \"genotypes\": \"${genotypes}\"}" $url
