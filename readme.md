Option monad for upcoming go 1.18

`
Experimental
`

`
Can be used now (go 1.17.3) with [gotip](https://pkg.go.dev/golang.org/dl/gotip)
`

install
```sh 
gotip get github.com/jh9aea/go-option
```


```go 
s := option.Some(1)
n := option.None[int]()

s.Get() == 1
n.Get() // panic

s.isSome() == true
n.isSome() == false

option.GetOr(func () int { return 2 }, s) == 1
option.GetOr(func () int { return 2 }, n) == 2

Map(func(v int) string { return strconv.Itoa(v + 1) }, Some(1)).Get()   // "2"
```