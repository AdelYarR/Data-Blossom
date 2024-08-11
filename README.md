Data Blossom implements a simple CRUD system with standard Go project layout. This is the result of my learning the Go language.

The main functional options are:
  · curl -X POST -H "Content-Type:application/json" -d "{"name":"Newname", "type_id":NewtypeID}" localhost:8000/api/lang — CREATE
  · cult localhost:8000/api/lang — READ
  · curl -X POST -H "Content-Type:application/json" -d "{"name":"Newname", "type_id":NewtypeID}" localhost:8000/api/lang?id=ID  — UPDATE
  · curl -X DELETE localhost:8000/api/lang?id=ID — DELETE
