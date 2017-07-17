# logistic-map

> Logistic map algorithm and its improved algorithms.

[![Build Status](https://travis-ci.org/WindomZ/logistic-map.svg?branch=master)](https://travis-ci.org/WindomZ/logistic-map)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/logistic-map/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/logistic-map?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/logistic-map)](https://goreportcard.com/report/github.com/WindomZ/logistic-map)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

The [logistic map](https://en.wikipedia.org/wiki/Logistic_map) is a polynomial mapping (equivalently, recurrence relation) of degree 2, 
often cited as an archetypal example of how complex, chaotic behaviour can arise from very simple non-linear dynamical equations.

## Usage

[![GoDoc](https://godoc.org/github.com/WindomZ/logistic-map?status.svg)](https://godoc.org/github.com/WindomZ/logistic-map)

### Logistic map

```go
import "github.com/WindomZ/logistic-map"

// Parameter u=3.999, x=0.2 and 5000 logistic map iterates.
result := logistic.LogisticMap(3.999, 0.2, 5000)
// len(result) == 5001', include result[0] == x
...
```

### Improved logistic map

Automatically adjust parameter `u` in `(3.569945673, 4)`.

```go
import "github.com/WindomZ/logistic-map"

// Parameter x=0.2 and 5000 logistic map iterates.
// Regulation amplification factor is 100.
result := logistic.LogisticImprovedMap(0.2, 100, 5000)
// len(result) == 5001', include result[0] == x
...
```

## Reference

[Wiki](https://en.wikipedia.org/wiki/Logistic_map)

## License

[MIT](https://github.com/WindomZ/logistic-map/blob/master/LICENSE)
