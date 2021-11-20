package main

import (
	"database/sql"
	"fmt"
	"go-di/demo"
	"go-di/di"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*

测试
单例对象在整个容器中只有一个实例，所以不管在何处注入，获取到的指针一定是一样的。
实例对象是通过同一个工厂方法创建的，所以每个实例的指针不可以相同。
下面是测试入口代码，完整代码在github仓库，有兴趣的可以翻阅：
*/

func main() {
	container := di.NewContainer()
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/mysql")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	container.SetSingleton("db", db)
	container.SetPrototype("b", func() (interface{}, error) {
		return demo.NewB(), nil
	})

	a := demo.NewA()
	if err := container.Ensure(a); err != nil {
		fmt.Println(err)
		return
	}
	// 打印指针，确保单例和实例的指针地址
	fmt.Printf("db: %p\n", a.Db)
	fmt.Printf("db1: %p\n", a.Db1)
	fmt.Printf("b: %p\n", &a.B)
	fmt.Printf("b1: %p\n", &a.B1)
}

/*

$ go run main.go
db: 0xc0000a6d00
db1: 0xc0000a6d00
b: 0xc0000be150
b1: 0xc0000be158

*/
