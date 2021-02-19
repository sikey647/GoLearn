## 数据类型

### 变量声明
* 变量声明
    * 通过 var 声明语句来定义一个变量，定义的时候需要指定这个变量的名称与类型，并且设置好变量的初始值
        ```
        var 变量名 类型 = 表达式
        ```
        ```go
        var i int = 10
        fmt.Println("i = ", i)
        ```
* 类型推导
    * Go 语言具有类型推导功能，所以也可以不去刻意地指定变量的类型，而是让 Go 语言自己推导
        ```go
        var j = 10
        fmt.Println("j = ", j)
        ```
    *  int、float64、bool、string 等基础类型都可以被自动推导
* 一次声明多个变量
    * 一次声明多个变量，把要声明的多个变量放到一个括号中
        ```go
        var (
            m int = 100
            n = 1000
        )
        fmt.Println("m =", m, ", n =", n)
        ```
* 变量简短声明
    * 借助类型推导，Go 语言提供了变量的简短声明 :=
    ```
    变量名 := 表达式
    ```
    ```go
    si := 10
    sb := false
    ss := "Hello"
    fmt.Println("si =", si, ", sb =", sb, ", ss = ", ss)
    ```

### 基础类型

#### 整型
* 整型分为：
    * 有符号整型：如 int、int8、int16、int32 和 int64
    * 无符号整型：如 uint、uint8、uint16、uint32 和 uint64
* 有符号整型表示的数值可以为负数、零和正数，而无符号整型只能为零和正数
* int 和 uint 这两个没有具体 bit 大小的整型，它们的大小可能是 32bit，也可能是 64bit，和硬件设备 CPU 有关
    * 在整型中，如果能确定 int 的 bit 就选择比较明确的 int 类型，因为这会让程序具备很好的移植性
* 还有一种字节类型 byte，它其实等价于 uint8 类型，可以理解为 uint8 类型的别名，用于定义一个字节，所以字节 byte 类型也属于整型

#### 浮点数
* Go 语言提供了两种精度的浮点数，分别是 float32 和 float64
    * 项目中最常用的是 float64，因为它的精度高，浮点计算的结果相比 float32 误差会更小
  ```go
  var f32 float32 = 1.23
  var f64 float64 = 12.3456
  fmt.Println("f32 =", f32, ", f64 =", f64)
  ```
* 对于数字类型之间，可以强制转换
    * 采用强制转换的方式转换数字类型，可能会丢失一些精度，比如浮点型转为整型时，小数点部分会全部丢失
    ```go
    i2f := float64(i)
    fmt.Printf("i = %v, i2f type = %T, value = %v\n", i, i2f, i2f)
    // i = 1, i2f type = float64, value = 1
    f2i := int(f64)
    fmt.Printf("f64 = %v, f2i type = %T, value = %v\n", f64, f2i, f2i)
    // f64 = 12.3456, f2i type = int, value = 12
    ```

#### 布尔型
* 一个布尔型的值只有两种：true 和 false，Go 语言中的布尔型使用关键字 bool 定义
    ```go
    var bf bool = false
    var bt bool = true
    fmt.Println("bf =", bf, ", bt =", bt)
    ```
* 布尔值可以用于一元操作符 !，表示逻辑非的意思，也可以用于二元操作符 &&、||，它们分别表示逻辑和、逻辑或

#### 字符串
* 字符串通过类型 string 声明
    ```go
    var s1 string = "Hello"
    var s2 string = "世界"
    fmt.Println("s1 =", s1, ", s2 =", s2)
    ```
* 可以通过操作符 + 把字符串连接起来，得到一个新的字符串
    ```go
    fmt.Println("s1 + s2 = ", s1 + s2)
    ```
* 字符串和数字互转
    * 通过包 strconv 的 Itoa 函数可以把一个 int 类型转为 string，Atoi 函数则用来把 string 转为 int
        ```go
        i2s := strconv.Itoa(i)
        fmt.Printf("i2s type = %T, value = %v\n", i2s, i2s)
        // i2s type = string, value = 1
        if s2i, err := strconv.Atoi(i2s); err == nil {
            fmt.Printf("s2i type = %T, value = %v\n", s2i, s2i)
        }
        // s2i type = int, value = 1
        ```
    * 对于浮点数、布尔型，Go 语言提供了 strconv.ParseFloat、strconv.ParseBool、strconv.FormatFloat 和 strconv.FormatBool 进行互转
        ```go
        vs := "3.1415926535"
        if s2f32, err := strconv.ParseFloat(vs, 32); err == nil {
            fmt.Printf("s2f32 type = %T, value = %v\n", s2f32, s2f32)
        }
        // s2f32 type = float64, value = 3.1415927410125732
        if s2f64, err := strconv.ParseFloat(vs, 64); err == nil {
            fmt.Printf("s2f64 type = %T, value = %v\n", s2f64, s2f64)
        }
        // s2f64 type = float64, value = 3.1415926535
    
        vf := 3.1415926535
        s32 := strconv.FormatFloat(vf, 'f', -1, 32)
        fmt.Printf("s32 type = %T, value = %v\n", s32, s32)
        // s32 type = string, value = 3.1415927
        s64 := strconv.FormatFloat(vf, 'f', -1, 64)
        fmt.Printf("s64 type = %T, value = %v\n", s64, s64)
        // s64 type = string, value = 3.1415926535
        ```
        ```go
        vbs := "true"
        if s2b, err := strconv.ParseBool(vbs); err == nil {
            fmt.Printf("s2b type = %T, value = %v\n", s2b, s2b)
        }
        // s2b type = bool, value = true
        vb := true
        b2s := strconv.FormatBool(vb)
        fmt.Printf("b2s type = %T, value = %v\n", b2s, b2s)
        // b2s type = string, value = true
        ```

#### 零值
* 零值其实就是一个变量的默认值，如果声明了一个变量，但是没有对其进行初始化，那么 Go 语言会自动初始化其值为对应类型的零值
    * 数字类的零值是 0，布尔型的零值是 false，字符串的零值是 "" 空字符串等 
    ```go
    var (
            zi int
            zf float64
            zb bool
            zs string
    )
    fmt.Println("zi =", zi, ", zf =", zf, ", zb =", zb, ", zs =", zs)
    ```
#### 指针
* 指针对应的是变量在内存中的存储位置，也就说指针的值就是变量的内存地址
* 通过 & 可以获取一个变量的地址，也就是指针
    ```go
    pi := &i
    fmt.Println("pi =", pi, ", *pi =", *pi, ", i =", i)
    ```

### 常量

#### 常量的定义
* 在程序中，常量的值是指在编译期就确定好的，一旦确定好之后就不能被修改，这样就可以防止在运行期被恶意篡改
* 常量的定义的关键字是 const，Go 语言可以类型推导，所以在常量声明时也可以省略类型
    ```go
    const cs1 string = "Hello"
    const cs2 = "世界"
    const (
        cs3 string = "你好"
        cs4 = "World"
    )
    fmt.Println("cs1 =", cs1, ", cs2 =", cs2, ", cs3 =", cs3, ", cs4 =", cs4)
    ```

#### iota
* iota 是一个常量生成器，它可以用来初始化相似规则的常量，避免重复的初始化
    * 不使用 iota，每个常量都要初始化，会比较繁琐
    ```go
    const (
        one = 1
        two = 2
        three = 3
        four = 4
    )
    fmt.Println("one =", one, ", two =", two, ", three =", three, ", four =", four)
    ```
    * 这些常量是有规律的（连续的数字），所以可以使用 iota 进行声明
    ```go
    const (
        five = iota + 5
        six
        seven
        eight
    )
    fmt.Println("five =", five, ", six =", six, ", seven =", seven, ", eight =", eight)
    ```
* iota 的初始值是 0，它的能力就是在每一个有常量声明的行后面 +1