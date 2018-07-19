## 1、包，Go语言标准库内建提供了regexp包
## 2、规则
    . 	匹配除换行符以外的任意字符
    \w 	匹配字母或数字或下划线或汉字
    \s 	匹配任意的空白符
    \d 	匹配数字
    \b 	匹配单词的开始或结束
    ^ 	匹配字符串的开始
    $ 	匹配字符串的结束
    * 	重复零次或更多次
    + 	重复一次或更多次
    ? 	重复零次或一次
    {n} 	重复n次
    {n,} 	重复n次或更多次
    {n,m} 	重复n到m次

    捕获 (exp) 	匹配exp,并捕获文本到自动命名的组里

    (?<name>exp) 	匹配exp,并捕获文本到名称为name的组里，也可以写成(?'name'exp)

    (?:exp) 		匹配exp,不捕获匹配的文本，也不给此分组分配组号
## 3、常用方法
func Match(pattern string, b []byte) (matched bool, err error)
func MatchString(pattern string, s string) (matched bool, err error)
func MustCompile(str string) *Regexp
func (re *Regexp) FindAllString(s string, n int) []string



