# niagadsofinquery
=======

##There are  *RESTful API* & *Authenticaiton* for the simple Phenotypes/Genotypes query.

** This project is done using [Gin Framework](https://github.com/gin-gonic/gin)+ [MysqlDB](https://github.com/go-sql-driver/mysql) + [JWT](https://github.com/dgrijalva/jwt-go)(JSON Web Token) + Golang. **

### Configuration
Pre-requisites: Go1.6, MysqlDB, Ginv1.0<br>


```
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
  <td>/signup</td>
  <td>GET</td>
  <td>name,password,email</td>
  <td>Signup Info</td>
 </tr>
 <tr>
 <td>/login</td>
  <td>POST</td>
  <td>name,password,email</td>
  <td>Login Info</td>
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
  <td>name, sex, birth, age_on_set, family_id, individual_id, paternal_id, maternal_id, affection_status</td>
  <td>To Create a New phenotype</td>
 </tr>
 <tr>
 <tr>
  <td>/api/phenotypes/:name</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single phenotypes</td>
 </tr>
 <tr>
  <td>/api/phenotypes/:name</td>
  <td>PUT</td>
  <td>name, sex, birth, age_on_set, family_id, individual_id, paternal_id, maternal_id, affection_status</td>
  <td>Updating a Single phenotype</td>
 </tr>
 <tr>
  <td>/api/phenotypes/:name</td>
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
  <td>name, chr, coordinate, variant_id, location, call</td>
  <td>To Create a New genotype</td>
 </tr>
 <tr>
 <tr>
  <td>/api/genotypes/:name</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single genotype</td>
 </tr>
 <tr>
  <td>/api/genotypes/:name</td>
  <td>PUT</td>
  <td>name, chr, coordinate, variant_id, location, call</td>
  <td>Updating a Single genotype</td>
 </tr>
 <tr>
  <td>/api/genotypes/:name</td>
  <td>DELETE</td>
  <td>id</td>
  <td>Delete a Single genotype</td>
 </tr>
 <tr>
  <td>/api/inqueries</td>
  <td>GET</td>
  <td></td>
  <td>Fetching All inqueries</td>
 </tr>
 <tr>
  <td>/api/inqueries</td>
  <td>POST</td>
  <td>name, individual_id, variant_id, did</td>
  <td>To Create a New inquery</td>
 </tr>
 <tr>
 <tr>
  <td>/api/inqueries/:name</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single inquery</td>
 </tr>
 <tr>
  <td>/api/inqueries/:name</td>
  <td>PUT</td>
  <td>name, individual_id, variant_id, did</td>
  <td>Updating a Single inquery</td>
 </tr>
 <tr>
  <td>/api/inqueries/:name</td>
  <td>DELETE</td>
  <td>id</td>
  <td>Delete a Single inquery</td>
 </tr>


 <tr>
  <td>/api/datasets</td>
  <td>GET</td>
  <td></td>
  <td>Fetching All datasets</td>
 </tr>
 <tr>
  <td>/api/datasets</td>
  <td>POST</td>
  <td>name, description, type, created_at</td>
  <td>To Create a New dataset</td>
 </tr>
 <tr>
 <tr>
  <td>/api/datasets/:name</td>
  <td>GET</td>
  <td></td>
  <td>Fetching A Single dataset</td>
 </tr>
 <tr>
  <td>/api/datasets/:name</td>
  <td>PUT</td>
  <td>name, description, type, created_at</td>
  <td>Updating a Single dataset</td>
 </tr>
 <tr>
  <td>/api/datasets/:name</td>
  <td>DELETE</td>
  <td>id</td>
  <td>Delete a Single datasets</td>
 </tr>
