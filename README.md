# MoneyFormat
格式化金额
### 安装方式
> go get github.com/houzhongjian/money

### 示例
```
package main
import(
    "github.com/houzhongjian/money"
    "log"
)
func main(){
    m := money.Format(1234567.89)
    log.Println(m)
}
```
输出
```
$1,234,567.89
```
