# fixedgrep

## how to install
```
go get -u github.com/umaumax/fixedgrep
```

## how to use
```
cat README.md | fixedgrep -prefix='#' how to
```

### help
```
Usage of fixedgrep:
  -max int
    	max hit number (default -1)
  -prefix string
    	fixed prefix string
  -suffix string
    	fixed suffix string
```
