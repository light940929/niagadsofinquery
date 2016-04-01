//GET the list of genotypes
curl -s http://localhost:9000/genotypes -H 'Accept:application/json'
//POST a genotype
curl -X POST  -H 'content-type: application/json'  -d @test.json  http://localhost:7000/genotypes
//PUT a genotype
curl -X PUT  -H 'content-type: application/json'  -d @test.json  http://localhost:7000/genotypes/:id



//Get the list of phenotypes
curl -s http://localhost:9000/phenotypes -H 'Accept:application/json'
//POST a phenotype
curl -X POST  -H 'content-type: application/json'  -d @test.json  http://localhost:7000/phenotypes
//PUT a phenotype
curl -X PUT  -H 'content-type: application/json'  -d @test.json  http://localhost:7000/phenotypes/:id
