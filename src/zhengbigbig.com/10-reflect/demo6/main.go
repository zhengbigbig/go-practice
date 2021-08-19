package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MySQLConfig ini配置文件解析器
type MySQLConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MySQLConfig `int:"mysql"`
	RedisConfig `int:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数校验，传入data必须是指针类型，因为需要在函数中对其赋值，必须是结构体类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data must be a pointer") // 格式化输出之后返回一个error类型
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data must be a struct pointer") // 格式化输出之后返回一个error类型
		return
	}
	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)
	var structName string
	// 2. 一行一行读数据 如果注释就跳过，如果[]开头就表示节(section)，如果不是[开头就是=分割的键值对
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		} else if strings.HasPrefix(line, "[") {
			if line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 提取掉[]，取到中间的内容把首尾空格去除
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("int") {
					// 说明找到了对应的嵌套结构体，将字段名记录
					structName = field.Name
					fmt.Printf("找到%s对应嵌套结构体%s\n", sectionName, structName)
					break
				}
			}
		} else if strings.Contains(line, "=") {
			// 1. 以等号分割，左key右value
			s := strings.Split(line, "=")
			key := strings.TrimSpace(s[0])
			value := strings.TrimSpace(s[1])
			// 2. 根据structName 去data里面把对应的嵌套结构体提取
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 嵌套结构体的值信息
			sType := sValue.Type()                     // 嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("%s该字段应该为结构体", structName)
				return
			}
			// 3. 遍历嵌套结构体的每个字段，判断tag是不是等于key
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存储在类型信息中
				fieldType = field
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			// 找不到对应的字段
			if len(fieldName) == 0 {
				continue
			}
			// 4. 如果key = tag，赋值
			// 4.1 根据fieldName取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对其赋值
			fmt.Println(fileName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var parseInt int64
				parseInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error", idx+1)
					return
				}
				fileObj.SetInt(parseInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var parseFloat float64
				parseFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error", idx+1)
					return
				}
				fileObj.SetFloat(parseFloat)
			}
		}
	}
	// 3.
	return nil
}

func main() {
	var c Config
	err := loadIni("./conf.ini", &c)
	if err != nil {
		fmt.Printf("load int failed,err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", c)

}
