## 快速入门

### Go语言环境搭建

#### Linux

* 根据你的 Linux 操作系统选择相应的压缩包，下载成功后，需要先进行解压

    ```shell
    sudo tar -C /usr/local -xzf go1.15.linux-amd64.tar.gz
    ```

* 把 /usr/local/go/bin 添加到 PATH 环境变量中，即添加到 /etc/profile 或者 $HOME/.bash_profile 文件中，保存后退出即可成功添加环境变量

    ```shell
    export PATH=$PATH:/usr/local/go/bin
    source ~/.bash_profile
    ```

#### MacOS

* 采用 PKG 安装包。下载 go1.15.darwin-amd64.pkg 后，双击按照提示安装即可
* 安装成功后，路径 /usr/local/go/bin 应该已经被添加到了 PATH 环境变量中（如果没有的话，可以手动添加，和上面 Linux 的方式一样）

#### 安装测试

* 安装成功后，可以打开终端或者命令提示符，输入 go version 来验证 Go 语言开发工具包是否安装成功

    ```shell
    go version
    # go version go1.15 darwin/amd64
    ```



### 环境变量设置

* Go 语言开发工具包安装好之后，还需要设置两个重要的环境变量，它们分别是 GOPATH 和 GOBIN。
    * GOPATH：代表 Go 语言项目的工作目录，在 Go Module 模式之前非常重要，现在基本上用来存放使用 go get 命令获取的项目
    * GOBIN：代表 Go 编译生成的程序的安装目录，比如通过 go install 命令，会把生成的 Go 程序安装到 GOBIN 目录下，以供在终端使用

* 在 Linux 和 macOS 下，把以下内容添加到 /etc/profile 或者 $HOME/.bash_profile 文件保存即可

    ```shell
    export GOPATH=/Users/sikey/WorkSpace/Go/GoPath
    export GOBIN=$GOPATH/bin
    source ~/.bash_profile
    ```



### 项目结构

* 采用 Go Module 的方式，可以在任何位置创建 Go 语言项目

* 打开终端，输入如下命令切换到项目位置

    ```shell
    cd /Users/sikey/WorkSpace/Go/GoLearn
    ```

* 执行如下命令创建一个 Go Module 项目，执行成功后，会生成一个 go.mod 文件。

    ```shell
    go mod init
    ```

* 然后在当前目录下创建一个 main.go 文件，这样整个项目目录结构是：

    ```shell
    GoLearn
    ├── go.mod
    ├── lib
    └── main.go
    ```

    * 其中 main.go 是整个项目的入口文件，里面有 main 函数
    * lib 目录是项目的子模块，根据项目需求可以新建很多个目录作为子模块，也可以继续嵌套为子模块的子模块



### 编译发布

* 在项目根目录输入以下命令，即可编译一个可执行文件

    ```shell
    go build ./main.go
    ```

* 测试下它是否可用

    ```shell
    ./main 
    # Hello, 世界
    ```

* 以上生成的可执行文件在当前目录，也可以把它安装到 $GOBIN 目录

    ```shell
    go install ./main.go
    ```



