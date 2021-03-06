# go-tutorial
go-tutorial

## Links
The Go Playground: https://play.golang.org/
Go: https://golang.org

## GO
Help Documentation Examples
```
  go doc json.Decoder.decode
  go doc json.Decoder
  go doc json
```

##Editors
Visual Studio Code: https://code.visualstudio.com
Goland: https://www.jetbrains.com/go/

## GO

`go run main.go`
`go run github.com/thampton/go-tutorial`

## Variables
3 ways to initialize - (keyword var)

Implicit initialization `:=` is most common

Local Variables must be used or results in compiler error

int, float32, float64, boolen, complex (real imag)

## Pointers
Location in memory that holds variable for usi

nil indicates a pointer not pointing to anything

`*` dereferences pointers and indicates type is a pointer

No pointer arithmetic 

& is the address of operator

## Constants
keyword const, can use implict type but must assign value

const blocks and iota allow for advanced constant patterns

## Collections
Arrays - brackets go before size and Arrays are constant size

Slices
*  : operator - complete array or slice, or range
*  Slice constructed from fixed size array points to that array
*  Slice constructed without fixed size, is dynamically sized and managed by go

Maps - Key Value Pairs - Can specify key and value types.

Struct - Only Data Type that can contain multiple data types

## Program Flow
Loops - For loop only - 4 types
*  loop till condition
*  loop till condition with post clause
*  infinite loops - `for ; ; {` or `for {`
*  loop over collections -  `for i, v range := slice` or `for _, v range := slice`

  break exists
  continue exists
  `_` readonly value

Panic - Throws Exception

Switch
   Cases break by default
   Default case 
   keyword break and fallthrough
  
  
