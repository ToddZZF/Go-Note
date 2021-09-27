# map

哈希表是一种巧妙并且实用的数据结构。它是一个无序的 key/value 对的集合，其中所有的 key 都是不同的，然后通过给定的 key 可以在常数时间复杂度内检索、更新或删除对应的 value。

在 Go 语言中，一个 map 就是一个哈希表的引用。

## map定义

`map` 的定义语法如下：

```go
map[KeyType]ValueType
```

其中，

* KeyType:表示键的类型。
* ValueType:表示键对应的值的类型。

map 类型的变量默认初始值为 `nil`，需要使用 `make()` 函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

示例：

```go
ages := make(map[string]int) // mapping from strings to ints
```

我们也可以用 map 字面值的语法创建 map，同时还可以指定一些最初的 key/value：

```go
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```

这相当于：

```go
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

因此，另一种创建空的 map 的表达式是 `map[string]int{}`。

## map的使用

Map 中的元素通过 key 对应的下标语法访问：

```go
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"
```

使用内置的 delete 函数可以删除元素：

```go
delete(ages, "alice") // remove element ages["alice"]
```

所有这些操作是安全的，即使这些元素不在 map 中也没有关系；如果一个查找失败将返回 value 类型对应的零值，例如，即使 map 中不存在“bob”下面的代码也可以正常工作，因为 `ages["bob"]` 失败时将返回 0。

```go
ages["bob"] = ages["bob"] + 1
```

而且 `x += y` 和 `x++` 等简短赋值语法也可以用在map上，所以上面的代码可以改写成：

```go
ages["bob"] += 1
ages["bob"]++
```

但是map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：

```go
_ = &ages["bob"] // compile error: cannot take address of map element
```

禁止对 map 元素取址的原因是 map 可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

## map的遍历

要想遍历map中全部的 key/value 对的话，可以使用 range 风格的 for 循环实现，和之前的 slice 遍历语法类似。下面的迭代语句将在每次迭代时设置 name 和 age 变量，它们对应下一个键/值对：

```go
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
```

## 判断某个键是否存在

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。

如果要按顺序遍历 key/value 对，我们必须显式地对key进行排序，可以使用 sort 包的 Strings 函数对字符串 slice 进行排序。下面是常见的处理方式：

```go
import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```

因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效。下面的代码创建了一个空的slice，但是slice的容量刚好可以放下map中全部的key：

```go
names := make([]string, 0, len(ages))
```

在上面的第一个range循环中，我们只关心map中的key，所以我们忽略了第二个循环变量。在第二个循环中，我们只关心names中的名字，所以我们使用“_”空白标识符来忽略第一个循环变量，也就是迭代slice时的索引。

## 使用delete()函数删除键值对

使用 `delete()` 内建函数从map中删除一组键值对，`delete()` 函数的格式如下：

```go
delete(map, key)
```

其中，

* map:表示要删除键值对的map
* key:表示要删除的键值对的键

## 判断map是否为空/未引用

map类型的零值是nil，也就是没有引用任何哈希表。

```go
var ages map[string]int
fmt.Println(ages == nil)    // "true"
fmt.Println(len(ages) == 0) // "true"
```

map 上的大部分操作，包括查找、删除、len 和 range 循环都可以安全工作在 nil 值的 map上，它们的行为和一个空的 map 类似。但是向一个 nil 值的map存入元素将导致一个 panic 异常：

```go
ages["carol"] = 21 // panic: assignment to entry in nil map
```

在向 map 存数据前必须先创建 map。

通过 key 作为索引下标来访问 map 将产生一个 value。如果 key 在 map 中是存在的，那么将得到与 key 对应的 value；如果 key 不存在，那么将得到 value 对应类型的零值，正如我们前面看到的 `ages["bob"]` 那样。这个规则很实用，但是有时候可能需要知道对应的元素是否真的是在 map 之中。例如，如果元素类型是一个数字，你可能需要区分一个已经存在的 0，和不存在而返回零值的 0，可以像下面这样测试：

```go
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
```

## 比较两个map是否相等

和 slice 一样，map 之间也不能进行相等比较；唯一的例外是和 nil 进行比较。要判断两个 map 是否包含相同的 key 和 value，我们必须通过一个循环实现：

```go
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```

## 元素为map类型的切片

下面的代码演示了切片中的元素为map类型时的操作：

```go
func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
```

## 值为切片类型的map

下面的代码演示了map中值为切片类型的操作：

```go
func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
```