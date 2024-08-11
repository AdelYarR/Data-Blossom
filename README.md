<h2>Data Blossom implements a simple CRUD system with standard Go project layout. This is the result of my learning the Go language.</h2>

<h3>The main functional options are:</h3>
<ul>
	<li>curl -X POST -H "Content-Type:application/json" -d "{"name":"Newname", "type_id":NewtypeID}" localhost:8000/api/lang — CREATE</li>
  <li>cult localhost:8000/api/lang — READ</li>
  <li>curl -X POST -H "Content-Type:application/json" -d "{"name":"Newname", "type_id":NewtypeID}" localhost:8000/api/lang?id=ID  — UPDATE</li>
  <li>curl -X DELETE localhost:8000/api/lang?id=ID — DELETE</li>
</ul>
