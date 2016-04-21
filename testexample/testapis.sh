//GET the TOKEN info
curl -u username:password http://localhost:9000/admin/secrets

//GET the list of genotypes
curl -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json'  http://localhost:9000/api/genotypes
//POST a genotype
curl -X POST  -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json' -d @test.json  http://localhost:9000/api/genotypes
//PUT a genotype
curl -X PUT  -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json' -d @test.json  http://localhost:9000/api/genotypes/:id

//Get the list of phenotypes
curl -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json'  http://localhost:9000/api/phenotypes
//POST a phenotype
curl -X POST -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json'  -d @test.json  http://localhost:9000/api/phenotypes
//PUT a phenotype
curl -X PUT  -H "Authorization: Bearer YOUR_TOKEN" -H 'Accept:application/json'  -d @test.json  http://localhost:9000/api/phenotypes/:id
