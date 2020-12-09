## GO Error 学习笔记

### 关于错误处理的几点理解

1. 错误只应该被处理一次，打日志、降级处理、返回错误等都算作对错误的处理。打了日志、做了降级就不应该继续返回错误；如果返回错误，则不做别的处理。
2. 错误应该包含错误信息和堆栈信息；以便于后续的处理。
3. 项目中应该尽量简洁明了的进行错误的判断处理。

### GO 1.13 错误处理

```
//errors package

//创建一个错误信息
func New(text string) error {}

//判断是否实现 interface { Unwrap() error } 接口，实现则调用返回根错误
func Unwrap(err error) error {}

//判断 err 错误是否 等于 target 错误，或者包含 target 错误
//如果 err 和 target 是同一个，那么返回true
//如果 err 是一个打包(wrap) 错误;并且 err 包含 target 错误 返回true
//不等于也不包含则返回false
func Is(err, target error) bool {}

//获取错误中包含的 具体错误，方法内部进行断言
func As(err error, target interface{}) bool {}

```

```
//fmt package

//打包根错误、添加错误信息；实现 interface { Unwrap() error } 接口来返回根错误
func Errorf(format string, a ...interface{}) error {}

```

GO 1.13 中提供了对错误的打包（wrap）功能；同时提供了比较好用的 Is、As 方法，但是却没有包含错误的堆栈信息

### github.com/pkg/errors

```
//github.com/pkg/errors

//创建包含 错误信息 和 堆栈信息 的错误
func New(message string) error {}
func Errorf(format string, args ...interface{}) error {}

//打包（wrap）错误;添加 堆栈信息；
//实现了interface { Cause() error } 接口
//实现了interface { Unwrap() error } 接口兼容 GO 1.13
func WithStack(err error) error {}

//打包（wrap）错误;添加 错误信息；
//实现了interface { Cause() error } 接口
//实现了interface { Unwrap() error } 接口兼容 GO 1.13
func WithMessage(err error, message string) error {}
func WithMessagef(err error, format string, args ...interface{}) error {}

//打包（wrap）错误;添加 错误信息 和 堆栈信息
//实现了interface { Cause() error } 接口
//实现了interface { Unwrap() error } 接口兼容 GO 1.13
func Wrap(err error, message string) error {}
func Wrapf(err error, format string, args ...interface{}) error {}

//判断是否实现 interface { Cause() error } 接口，实现则调用返回根错误；等同于 GO 1.13 的 Unwrap()
func Cause(err error) error {}

```

### 使用注意

1. 综合使用 GO 1.13 和 pgk/errors；pkg/errors 打包处理的错误 GO 1.13 提供的 Unwrap、Is、As 同样可以使用
2. 自己应用层的代码，出现错误；可以使用 pkg/errors 的包，创建一个包含错误信息 和 堆栈信息的错误进行返回
3. 对于别人返回的错误，如果是底层的未包含堆栈信息的错误可以使用 wrap 打包错误，添加错误信息 和 堆栈信息
4. 对于别人返回的错误，如果是已经打包，包含 堆栈信息 的错误则直接返回。

