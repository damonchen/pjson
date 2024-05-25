# pjson

using goparsec2 to parse json

## depend

`goparsec2`

    go get github.com/Dwarfartisan/goparsec2


## understand parsec

使用parsec之前，需要先理解monad，理解monad，可以这么考虑

1. 有一个算子，`parsec`处为 `P`，经过一系列操作后，返回的结果依旧是 `P`
2. 这个算子在运算过程中，会先将`P`解开来，得到里面的值，然后封装为`P`

