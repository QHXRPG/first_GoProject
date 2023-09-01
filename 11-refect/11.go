package main

import (
	"fmt"
	"reflect"
)

/*
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
os.OpenFile 是os包中的一个函数，用于打开一个文件并返回一个文件对象。
/dev/tty 是要打开的文件路径，它表示当前终端设备。
os.O_RDWR 是打开文件的模式，表示以读写的方式打开文件。
0  是打开文件时的权限标志，表示不设置额外的权限。

- `tty`是一个`*os.File`类型的变量，它是通过打开`/dev/tty`文件后返回的文件对象。
	通过这个文件对象，可以进行与终端的交互，如读取终端输入或向终端输出内容。

- `err`是一个`error`类型的变量，它表示打开文件时可能出现的错误。
	如果打开文件成功，则`err`的值为`nil`，表示没有发生错误。如果打开文件失败，
	则`err`将包含一个描述错误的具体信息。
*/

/*
	var r io.Reader
	r = tty

	_, ok := r.(io.Reader)
	fmt.Println(ok)
	_, ok = r.(*os.File)
	fmt.Println(ok)
	_, ok = r.(io.Writer)
	fmt.Println(ok)

这段代码执行了一系列类型断言操作，用于检查变量 `r` 的实际类型。下面是对每个断言操作的解释：

1. `_, ok := r.(io.Reader)`
   这个断言操作检查 `r` 是否实现了 `io.Reader` 接口。如果 `r` 实现了该接口，则 `ok` 的值为 `true`，否则为 `false`。

2. `_, ok = r.(*os.File)`
   这个断言操作检查 `r` 是否是 `*os.File` 类型的实例。如果 `r` 是 `*os.File` 类型的实例，则 `ok` 的值为 `true`，否则为 `false`。

3. `_, ok = r.(io.Writer)`
   这个断言操作检查 `r` 是否实现了 `io.Writer` 接口。如果 `r` 实现了该接口，则 `ok` 的值为 `true`，否则为 `false`。

最后，通过调用 `fmt.Println` 打印每个断言操作的结果，即 `ok` 的值，用于判断 `r` 的实际类型是否符合断言的条件。
*/

func main() {

	var num float64 = 1.2345

	fmt.Println("--------------------")
	numType, numValue := reflect.TypeOf(num), reflect.ValueOf(num)
	fmt.Println(numType)
	fmt.Println(numValue)

	fmt.Println("1--------------------")
	pointer, value := reflect.ValueOf(&num), reflect.ValueOf(num)
	fmt.Println(pointer)                 //0x14000110018
	fmt.Println(value)                   //1.2345
	fmt.Println(reflect.TypeOf(pointer)) //reflect.Value
	fmt.Println(reflect.TypeOf(value))   //reflect.Value

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	fmt.Println("3--------------------")
	convertPointer, convertValue := pointer.Interface(), value.Interface()
	fmt.Println(convertPointer)                 //0x14000110018
	fmt.Println(convertValue)                   //1.2345
	fmt.Println(reflect.TypeOf(convertPointer)) //*float64
	fmt.Println(reflect.TypeOf(convertValue))   //float64

	fmt.Println("4--------------------")
	convertPointer, convertValue = pointer.Interface().(*float64), value.Interface().(float64)
	fmt.Println(convertPointer)                 //0x14000110018
	fmt.Println(convertValue)                   //1.2345
	fmt.Println(reflect.TypeOf(convertPointer)) //*float64
	fmt.Println(reflect.TypeOf(convertValue))   //float64
}

/*
这段代码使用了反射（reflect）包来进行变量类型的转换和获取。下面是对每个操作的解释：

1. `pointer := reflect.ValueOf(&num)`
   这行代码使用 `reflect.ValueOf` 函数将 `&num` 的地址（指针）转换为 `reflect.Value` 类型的值 `pointer`。
	`pointer` 是一个 `reflect.Value` 类型的变量，表示 `num` 的指针。

2. `value := reflect.ValueOf(num)`
   这行代码使用 `reflect.ValueOf` 函数将 `num` 的值转换为 `reflect.Value` 类型的值 `value`。
	`value` 是一个 `reflect.Value` 类型的变量，表示 `num` 的值。

3. `fmt.Println(reflect.TypeOf(pointer))`
   这行代码使用 `reflect.TypeOf` 函数获取 `pointer` 的类型，并通过 `fmt.Println` 打印出来。
	它打印的是 `*reflect.Value`，表示 `pointer` 是一个指向 `reflect.Value` 类型的指针。

4. `fmt.Println(reflect.TypeOf(value))`
   这行代码使用 `reflect.TypeOf` 函数获取 `value` 的类型，并通过 `fmt.Println` 打印出来。
	它打印的是 `reflect.Value`，表示 `value` 是一个 `reflect.Value` 类型的值。

5. `convertPointer := pointer.Interface().(*float64)`
   这行代码使用 `Interface` 方法将 `pointer` 转换为 `interface{}` 类型，
	然后使用类型断言 `.(*float64)` 将其强制转换为 `*float64` 类型的指针。如果类型不匹配，会导致 panic。

6. `convertValue := value.Interface().(float64)`
   这行代码使用 `Interface` 方法将 `value` 转换为 `interface{}` 类型，
	然后使用类型断言 `.(float64)` 将其强制转换为 `float64` 类型的值。如果类型不匹配，会导致 panic。

7. `fmt.Println(convertPointer)`
   这行代码通过 `fmt.Println` 打印 `convertPointer` 的值，即 `num` 的指针。

8. `fmt.Println(convertValue)`
   这行代码通过 `fmt.Println` 打印 `convertValue` 的值，即 `num` 的值。

9. `fmt.Println(reflect.TypeOf(convertPointer))`
   这行代码使用 `reflect.TypeOf` 函数获取 `convertPointer` 的类型，并通过 `fmt.Println` 打印出来。
	它打印的是 `*float64`，表示 `convertPointer` 是一个指向 `float64` 类型的指针。

10. `fmt.Println(reflect.TypeOf(convertValue))`
    这行代码使用 `reflect.TypeOf` 函数获取 `convertValue` 的类型，并通过 `fmt.Println` 打印出来。
	它打印的是 `float64`，表示 `convertValue` 是一个 `float64` 类型的值。

*/
