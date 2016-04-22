# niagadsofinquery
=======

##There are  *RESTful API* & *Authenticaiton* for the simple Phenotypes/Genotypes query.

** This project is done using [Gin Framework](https://github.com/gin-gonic/gin)+ [MongoDB](https://github.com/mongodb/mongo) + [JWT](https://github.com/dgrijalva/jwt-go)(JSON Web Token) + Golang. **

### Configuration
Pre-requisites: Go1.6, MongoDBv3.2.1, Ginv1.0<br>


```

mongod --dbpath ~/PATH/DB

mongo
> use admin
switched to db admin
> db.createUser(
...   {
...     user: "name",
...     pwd: "password",
...     roles: [ { role: "readWrite", db: "test" } ]
...   }
... )

mongod --auth --dbpath ~/PATH/DB

go get && go install && PORT=7000 DEBUG=* gin -p 9000 -a 7000 -i run

```

###Request
--------
<table>
 <tr>
   <th>URL</th>
   <th>Method</th>
   <th>Parameters</th>
   <th>Description</th>
 </tr>
 <tr>
  <td>/admin/secrets</td>
  <td>GET</td>
  <td>name,password</td>
  <td>Admin Login</td>
 </tr>
 <tr>
 <td>/api/oauth2/token</td>
  <td>GET</td>
  <td></td>
  <td>Token Info</td>
 </tr>
 <tr>
  <td>/api/phenotypes</td>
  <td>GET</td>
  <td></td>
  <td>Fetching All phenotypes</td>
 </tr>
 <tr>
  <td>/api/phenotypes</td>
  <td>POST</td>
  <td>title, sex, birth, ageonset, famliyID, individualID, paternalID, maternalID</td>
  <td>To Create a New phenotype</td>
 </tr>
 <tr>
 <tr>
  <td>/api/phenotypes/:id</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single phenotypes</td>
 </tr>
 <tr>
  <td>/api/phenotypes/:id</td>
  <td>PUT</td>
  <td>title, sex, birth, ageonset, famliyID, individualID, paternalID, maternalID</td>
  <td>Updating a Single phenotype</td>
 </tr>
 <tr>
  <td>/api/phenotypes/:id</td>
  <td>DELETE</td>
  <td>id</td>
  <td>Delete a Single phenotype</td>
 </tr>
 <tr>
  <td>/api/genotypes</td>
  <td>GET</td>
  <td></td>
  <td>Fetching All genotypes</td>
 </tr>
 <tr>
  <td>/api/genotypes</td>
  <td>POST</td>
  <td>title, variantID, location, call</td>
  <td>To Create a New genotype</td>
 </tr>
 <tr>
 <tr>
  <td>/api/genotypes/:id</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single genotype</td>
 </tr>
 <tr>
  <td>/api/genotypes/:id</td>
  <td>PUT</td>
  <td>title, variantID, location, call</td>
  <td>Updating a Single genotype</td>
 </tr>
 <tr>
  <td>/api/genotypes/:id</td>
  <td>DELETE</td>
  <td>id</td>
  <td>Delete a Single genotype</td>
 </tr>
