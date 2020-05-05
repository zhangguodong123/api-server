# Type Alias

**类型别名** 是 `Go 1.9` 版本添加的新功能。**主要应用于代码升级、工程重构、迁移中类型的兼容性问题。** 
C/C++ 语言中，代码的重构升级可以使用宏快速定义新的代码。Go 语言中并未选择通过宏，而是选择通过类型别名解决重构中最复杂的类型名变更问题。

可参考: `${GOPATH}/src/builtin/builtin.go`

- `type int8 int8`
- `type uint8 uint8`
- **type byte = uint8**, **type alias**
- **type rune = int32**, **type alias**
- `type IntegerType int`
- `type FloatType float32`
- `type ComplexType complex64`

## Glossary

- 类型别名规定：Type Alias 只是 Type 的别名，本质上 Type Alias 与 Type 是同一个类型，即基本数据类型是一致的。
- 类型定义：依据基本类型声明一个新的数据类型。

## Tips

- 类型别名只会在代码中存在，编译完成时，不会有类型别名对应的类型存在。
- 类型定义: 根据已存在类型定义新类型，新类型!

## [case](https://github.com/xwi88/go-study/tree/master/type_alias_define)

```go
package main

import (
	"fmt"
)

// Type define, according the exist type, builtin or your type.
type myInt int

// Type alias
type intAlias = int

func main() {
	var a myInt
	fmt.Printf("Type define, a Type: %T, value： %d\n", a, a)

	var b intAlias
	fmt.Printf("Type  alias, b Type: %T, value： %d\n", b, b)
}

// output:
// a Type: main.myInt, value： 0
// b Type: int, value： 0
```
