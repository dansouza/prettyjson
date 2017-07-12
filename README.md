# prettyjson
command-line utility to pretty-print JSON content from files/stdin

# install
```
$ go install github.com/dansouza/prettyjson
```
# usage
```
$ curl 'https://maps.googleapis.com/maps/api/geocode/json?&address=statue%20of%20liberty' 2> /dev/null | prettyjson
```
or
```
$ wget 'https://maps.googleapis.com/maps/api/geocode/json?&address=statue%20of%20liberty' -O liberty.json && prettyjson liberty.json
```

# example
[![asciicast](https://asciinema.org/a/8o0kNockI6UJx2oJQhIroSdkR.png)](https://asciinema.org/a/8o0kNockI6UJx2oJQhIroSdkR)
