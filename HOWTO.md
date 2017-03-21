# Update OUI

## Generating ouis.json
Scripts to generate the database are hosted in the [ouiner](https://github.com/snowhigh/ouiner) project.

## Generating code
Place your ouis.json file in data/ouis.json.

```
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata data/
```
Replace package name from main to libmacouflage
```
$ vim bindata.go
```
Done!
