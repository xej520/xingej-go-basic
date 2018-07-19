## 1、处理HTTP请求
使用 net/http 包提供的 http.ListenAndServe() 方法，可以在指定的地址进行监听， 开启一个HTTP，服务端该方法的原型如下：
func ListenAndServe(addr string, handler Handler) error
该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连 接请求。
该方法有两个参数：
第一个参数 addr 即监听地址；
第二个参数表示服务端处理程序， 通常为空，这意味着服务端调用 http.DefaultServeMux 进行处理，
而服务端编写的业务逻 辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中
## 2、处理HTTPS请求
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler)      error

## 3、路由
http.HandleFunc()方法接受两个参数

第一个参数是HTTP请求的 目标路径"/hello"，该参数值可以是字符串，也可以是字符串形式的正则表达式
第二个参数指定具体的回调方法，比如helloHandler。

当我们的程序运行起来后，访问http://localhost:8080/hello  ， 程序就会去调用helloHandler()方法中的业务逻辑程序。


