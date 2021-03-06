# redisvars [![Build Status](https://travis-ci.org/saromanov/redisvars.svg?branch=master)](https://travis-ci.org/saromanov/redisvars) [![Coverage Status](https://coveralls.io/repos/saromanov/redisvars/badge.svg)](https://coveralls.io/r/saromanov/redisvars)
golang maps with redis backend

## Why?
For storing information which will be used later

## Usage
Dict
```go
rv := redisvars.New("localhost:6379")
dict := rv.NewDict()
dict.Set("A", "B")
dict.Set("V", "A")
dict.Commit()
dict.Set("V", "B")
fmt.Println(dict.Get("V")) //returns "A"
```

List
```go
rv := redisvars.New("localhost:6379")
lst := rv.NewList()
lst.SetList("DDDDD", []string{"A", "B", "C"})
lst.CommitList()
```
Struct
```go

import
(
	"encoding/json"
)

type Foobar struct {
	Item string
	Data int
}

elem := &Foobar{"first",1}
rv := redisvars.New("localhost:6379")
rstr := rv.NewStruct()
rstr.SetStruct("foobar1",elem)
rstr.CommitStruct()
res := rstr.GetStruct("foobar1")
var foobar Foobar
err := json.Unmarshal([]byte(res), &foobar)
fmt.Println(foobar.Item, foobar.Data)
```
## License
MIT