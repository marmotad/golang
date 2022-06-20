package main

import (
	"fmt"
	"strings"
)

func main() {
	//比较字符串是否相等，不相等返回-1，相等返回0
	fmt.Println(strings.Compare("aaa", "ccc"))
	fmt.Println(strings.Compare("aaa", "aaa"))
	//比较字符串是否包含子字符串，返回false/true
	fmt.Println(strings.Contains("aaa", "ac"))
	// 判断字符串 s 中是否包含 chars 中的任何一个字符
	fmt.Println(strings.ContainsAny("aaa", "adc"))
	//子字符串出现次数
	fmt.Println(strings.Count("abcdefabadat", "ab"))
	//按照空白字符分割（空格,\n,\t,\r）
	fmt.Printf("%q\n", strings.Fields("aa ddr\neee\tttt\rrrr"))
	// 判断字符串 s 中是否以 chars 中的字符串开头
	fmt.Println(strings.HasPrefix("abcd", "ab"))
	// 判断字符串 s 中是否以 chars 中的字符串结尾
	fmt.Println(strings.HasSuffix("abcd", "ab"))
	// chars 中的字符串在s字符串中第一次出现的位置
	fmt.Println(strings.Index("abcdefgabcdefg", "cd"))
	fmt.Println(strings.Index("abcdefgabcdefg", "de"))
	// chars 中的字符串在s字符串中最后一次出现的位置
	fmt.Println(strings.LastIndex("abcdefgabcdefg", "cd"))
	fmt.Println(strings.LastIndex("abcdefgabcdefg", "de"))
	// 以sep为分割符，将s的字符串分割为切片
	fmt.Println(strings.Split("abcdef;abc;abc", ";"))
	//以sep为分隔符，将s中的数组拼接起来
	fmt.Println(strings.Join([]string{"abc", "def", "eee"}, ";"))
	//将s字符串重复count次
	fmt.Println(strings.Repeat("abc", 3))
	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xxx", 1))
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xxx", -2))
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xxx", -1))
	//返回将s中前所有不重叠old子串都替换为new的新字符串
	fmt.Println(strings.ReplaceAll("abcabcabcab", "ab", "xxx"))
	//将大写字母转为小写
	fmt.Println(strings.ToLower("abcABC"))
	//将小写字母转为大写
	fmt.Println(strings.ToUpper("abcABC"))
	//将首字母转换为大写，在18版本中被弃用，参考：D:\code\golong\golang.cncf.online\src\day02\08stringsModule.go
	fmt.Println(strings.Title("abcABC"))
	//返回将 s 前后端所有 cutset 包含的 utf-8 码值都去掉的字符串。
	fmt.Println(strings.Trim("xyzabcyzx", "xz"))
	//TrimSpace 返回字符串 s 的切片，删除所有前导和尾随空格，如 Unicode 所定义。
	fmt.Println(strings.TrimSpace("abcd xxx \n \r"))
}
