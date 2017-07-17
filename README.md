# logistic-map

> Logistic map algorithm and its improved algorithms.

The [logistic map](https://en.wikipedia.org/wiki/Logistic_map) is a polynomial mapping (equivalently, recurrence relation) of degree 2, 
often cited as an archetypal example of how complex, chaotic behaviour can arise from very simple non-linear dynamical equations.

## Usage

### Logistic map

```go
import "github.com/WindomZ/logistic-map"

// Parameter u=3.999, x=0.2 and 5000 logistic map iterates.
result := logistic.LogisticMap(3.999, 0.2, 5000)
// len(result) == 5001', include result[0] == x
...
```

#### Improved

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
