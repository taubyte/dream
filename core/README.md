# README #

Dreamland starts taubyte services in one or multiple universes. nodes are automatically fully meshed

You can use this from command line. See example below.

Or as a library in tests.

## API
### Status
```
GET http://127.0.0.1:1421/status
```

## CLI
### Simple run
```shell
go run . start single universe
```

### Run with fixtures
```shell
go run . start single universe --fixtures fixture1,fixture2,fixture3...
``` 

### Start simple nodes
```shell
go run . start single universe --simples one,two,three
``` 

### Disable services
```shell
go run . start single universe --disable seer,tns
``` 


### Bind services
```shell
go run . start single universe --bind seer@11009,seer@11888/http
``` 

### Bind simple nodes
```shell
go run . start single universe --simple one@11444,two,three@10102
``` 
