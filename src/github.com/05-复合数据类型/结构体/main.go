package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型
type myInt int     // 自定义类型
type yourInt = int // 类型别名

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// 结构体是值类型
type boy struct {
	name, gender string
}

func f(x person) {
	x.gender = "女" // 修改的是副本，go语言中函数参数传参永远是拷贝
}

func f2(x *person) {
	//(*x).gender = "女" // 根据内存地址找到变量，修改的是原来的变量
	x.gender = "女" // 语法糖，自动根据指针找到对应的变量
}

// 结构体占用一块连续的内存空间
type x struct {
	a int8
	b int8
	c int8
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name:   name,
		age:    age,
		gender: "boy",
		hobby:  nil,
	}
}

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{name: name}
}

// 方法是作用于特定类型的函数
// 接收者表示是调用具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

func main() {
	d1 := newDog("pidan")
	d1.wang()

	var zbb person
	zbb.name = "zbb"
	zbb.age = 18
	zbb.gender = "boy"
	zbb.hobby = []string{"唱", "跳", "rapper"}
	fmt.Println(zbb)
	fmt.Printf("%T \n", zbb)
	f(zbb)
	fmt.Println(zbb)
	f2(&zbb)
	fmt.Println(zbb)

	var p = new(person)
	p.name = "jzz"
	fmt.Printf("%T\n %v\n", p, p)

	// 结构体指针
	// 1. key-value初始化
	var p1 = person{name: "xxx", gender: "boy"}
	fmt.Printf("%#v\n", p1)
	// 2. 使用值列表形式初始化，值的顺序要喝结构体定义字段顺序一致
	p2 := person{"yyy", 18, "boy", []string{"rapper"}}
	fmt.Printf("%#v\n", p2)

	// 匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "xxx"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	p3 := newPerson("xx", 1)
	fmt.Println(p3, *p3)

}
