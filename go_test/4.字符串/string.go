package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Printf：
- 需要格式化字符串和占位符
- 可以精确控制输出格式
- 不会自动换行
Println：
- 直接输出参数值
- 自动在参数间添加空格
- 自动在末尾添加换行符
*/
//对于字符串操作的有4个包：
// bytes、strings
// strconv、unicode

/**
 * bytes 包操作 []byte。因为字符串是只读的，因此逐步构创建字符串会导致很多分配和复制，使用 bytes.Buffer 类型会更高
 * strings 包提供 切割、索引、前缀、查找、替换等功能
 * strconv 包提供 布尔型、整型数、浮点数 和对应字符串的相互转换，还提供了双引号转义相关的转换
 * nicode 包提供了 IsDigit、IsLetter、IsUpper、IsLower 等类似功能，用于给字符分类
 */

func main() {
	//utf-8 编码,一个汉字需要3个字节，通过len()获取的是字符串占据的字数
	str1 := "hello 世界"
	fmt.Println(len(str1)) //12

	//如果想要得到字符串本身的长度，可以将 string 转为 rune 数组再计算
	str2 := "hello 世界"
	fmt.Println(len([]rune(str2))) //8
	// byte 是 uint8 的别名
	// rune 是 int32 的别名，相当于 Go 里面的 char

	//以下遍历方式会乱码
	str := "你好 世界！"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	//解决方案1；转成rune切片遍历
	newstr := []rune(str)
	for i := 0; i < len(newstr); i++ {
		fmt.Printf("%c", newstr[i])
	}
	fmt.Println()
	//解决方案2：使用range遍历
	for _, value := range str {
		fmt.Printf("%c", value)
	}
	fmt.Println()

	//string包

	//字符串比较：使用strings.Compare() 比较两个字符串的字典序
	//第一个字符串在字典序上小于第二个字符串，返回-1
	//第一个字符串在字典序上大于第二个字符串，返回1
	//两个字符串完全相等，返回0
	fmt.Println(strings.Compare("aaa", "bbb")) // -1
	fmt.Println(strings.Compare("baa", "abb")) // 1
	fmt.Println(strings.Compare("aaa", "aaa")) // 0

	//使用strings.Index查找字符串字串出现的第一个位置，不存在返回-1
	//Go语言中字符串索引从0开始计数
	fmt.Println(strings.Index("hello world", "l")) // 2

	//使用strings.LastIndex查找字符串字串出现的最后一个位置，不存在返回-1
	//Go语言中字符串索引从0开始计数
	fmt.Println(strings.LastIndex("hello world", "l")) // 9

	//使用strings.Count统计字串在整体中出现的次数
	//使用strings.Repeat将字符串重复指定次数
	fmt.Println(strings.Count("abcabc cababababc", "abc")) // 3
	fmt.Println(strings.Repeat("abc", 3))                  //abcabcabc

	//使用strings.Replace实现字符串替换
	//使用strings.Split实现字符串切割
	//使用strings.join实现字符串拼接
	s1 := "acaacccc"
	//局部替换,替换次数<0 则替换所有
	fmt.Println(strings.Replace(s1, "a", "b", 2))  //将前两个a换成b，bcbacccc
	fmt.Println(strings.Replace(s1, "a", "b", -1)) //将所有a换成b，bcbbcccc
	//全部替换
	fmt.Println(strings.ReplaceAll(s1, "b", "a")) //acaacccc

	s2 := "abc,bbc,bbd"
	fmt.Printf("s2的类型：%T\n", s2)
	slice := strings.Split(s2, ",")
	fmt.Printf("slice的类型：%T\n", slice)
	// []string是Go中的一个切片类型，字符串切片
	fmt.Println(slice)

	s3 := strings.Join(slice, "-")
	fmt.Println(s3)

	//btyes包
	//创建缓冲器：bytes.NewBufferString、bytes.NewBuffer
	//高效的字符串构建：当需要频繁拼接字符串时，使用 bytes.Buffer 比直接字符串拼接更高效
	//直接传入字符串参数
	buf1 := bytes.NewBufferString("hello")
	//将字符串 "hello" 转换为字节切片 []byte("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	buf3 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Printf("%v,%v,%v\n", buf1, buf2, buf3)
	fmt.Printf("%v,%v,%v\n", buf1.Bytes(), buf2.Bytes(), buf3.Bytes())
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf4.Bytes(), buf5.Bytes())

	//写入缓冲器：Write、WriteString、WriteByte、WriteRune、WriteTo
	buf := bytes.NewBufferString("a")
	fmt.Printf("%v,%v\n", buf, buf.Bytes()) // a, [97]
	buf.Write([]byte("b"))                  //写入字节切片
	buf.WriteString("c")                    //写入字符串
	buf.WriteByte('d')                      //写入单个字节
	buf.WriteRune('e')                      //写入单个rune字符
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())

	/**
	* 关于缓冲区
	* Go 字节缓冲区底层以字节切片做存储，切片存在长度 len 与容量 cap
	* 缓冲区从长度 len 的位置开始写，当 len > cap 时，会自动扩容
	* 缓冲区从内置标记 off 位置开始读（off 始终记录读的起始位置）
	* 当 off == len 时，表明缓冲区已读完，读完就重置缓冲区 len = off = 0
	* len和off都从0开始技术，例如：字符串 "abc" 的 len = 3（索引 0, 1, 2）
	 */
	/**
	* ReadByte(): 读取当前位置字节，off++
	* Next(n): 跳过n个字节，off += n
	* Write(): 在末尾追加数据，len 增加
	* WriteByte(): 写入单字节，len++
	* Truncate(n): 设置未读字节数为n，调整len
	* Reset(): 清空缓冲区，len=0, off=0
	 */
	byteSlice := make([]byte, 20)                     //创建长度为20的字节切片，所有元素初始值为0
	byteSlice[0] = 1                                  // 将缓冲区第一个字节置1
	byteBuffer := bytes.NewBuffer(byteSlice)          // 创建20字节缓冲区 len = 20 off = 0
	c, _ := byteBuffer.ReadByte()                     // 读取当前位置的字节，byteslice0=1，读取操作会使off增加1
	fmt.Printf("len:%d, c=%d\n", byteBuffer.Len(), c) // len = 20 off =1  byteBuffer.Len() = len - off = 19  打印c=1
	byteBuffer.Reset()                                // 重置缓冲区 len = 0 off = 0
	fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=0
	byteBuffer.Write([]byte("hello byte buffer"))     // 写缓冲区  len+=17 17个字节
	fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=17 off = 0 len-off = 17
	byteBuffer.Next(4)                                // 跳过4个字节 off+=4 off = 4
	c, _ = byteBuffer.ReadByte()                      // 读第5个字节 off+=1 off = 5
	fmt.Printf("第5个字节:%d\n", c)                       // 打印:111(对应字母o)    len=17 off=5
	byteBuffer.Truncate(3)                            // 截断操作，将未字节数置为3        len=off+3=8   off=5
	fmt.Printf("len:%d\n", byteBuffer.Len())          // 打印len=3为未读字节数  上面len=8是底层切片长度 8-5 = 3
	byteBuffer.WriteByte(96)                          // 写入单字节，len+=1=9 将y改成A
	byteBuffer.Next(3)                                // 跳过三个字节，len=9 off+=3=8
	c, _ = byteBuffer.ReadByte()                      // off+=1=9    c=96
	fmt.Printf("第9个字节:%d\n", c)                       // 打印:96

	//缓冲区
	buf6 := &bytes.Buffer{} //创建一个新的bytes.Buffer类型指针 初始len = 0, off = 0
	// 写缓冲区
	buf6.WriteString("abc?def") // len = 7 off = 0
	// 从缓冲区读（分隔符为 ?）
	/*
	* 从 off=0 位置开始读取
	* 读取 'a', 'b', 'c', '?' 四个字符
	* 遇到分隔符 '?' 后停止读取
	* readstr 的值为 "abc?"
	* 读取后 off=4（跳过了已读取的4个字节）
	 */
	readstr, _ := buf6.ReadString('?') //读取的内容包含分隔符

	fmt.Println("readstr = ", readstr)
	fmt.Println("buff = ", buf6.String()) // buf6.String() 返回缓冲区中未读取的部分

	/** 缓冲区读数据
	* Read() 从缓冲区读取数据
	* ReadByte() 读取当前位置的字节
	* ReadBytes(delim byte) 读取当前位置到指定字符的子串
	* ReadString(delim byte) 读取当前位置到指定字符的子串
	* ReadRune() 读取当前位置的字符
	* ReadUnicodeRune() 读取当前位置的字符
	* ReadLine() 读取当前位置到行末的子串
	* ReadFrom(r io.Reader) 从 io.Reader 中读取数据并写入缓冲区
	 */
	log.SetFlags(log.Lshortfile)               //设置日志输出格式，log.Lshortfile 表示在日志中显示文件名和行号
	buff := bytes.NewBufferString("123456789") // 创建缓冲区，len = 9，off = 0
	log.Println("buff = ", buff.String())      // buff = 123456789

	// 从缓冲区读取4个字节
	s := make([]byte, 4)                  //创建长度为4的字节切片
	n, _ := buff.Read(s)                  // 从缓冲区读取最多4个字节到切片 s 中 off = 4
	log.Println("buff = ", buff.String()) // buff =  56789
	log.Println("s = ", string(s))        // s =  1234
	log.Println("n = ", n)                // n =  4 表示实际读取4个字节

	// 从缓冲区读取4个字节
	n, _ = buff.Read(s)                   //继续从off = 4的位置读取4个字节  off = 8
	log.Println("buff = ", buff.String()) // buff =  9
	log.Println("s = ", string(s))        // s =  5678
	log.Println("n = ", n)                // n =  4

	n, _ = buff.Read(s)
	log.Println("buff = ", buff.String()) // buff =  缓冲区已空：buff.String() 返回空字符串
	log.Println("s = ", string(s))        // s =  9678 读取了 "9" 到 s[0]，s[1:3] 保持之前的值 "678"
	log.Println("n = ", n)                // n = 1 表示实际读取1个字节

	buff.Reset()                          // 重置缓冲区 len = 0, off = 0
	buff.WriteString("abcdefg")           // len = 7, off = 0
	log.Println("buff = ", buff.String()) // buff =  abcdefg

	b, _ := buff.ReadByte()               // 读取一个字节，off = 1
	log.Println("b = ", string(b))        // b =  a
	log.Println("buff = ", buff.String()) // buff =  bcdefg

	b, _ = buff.ReadByte()                // 读取一个字节，off = 2
	log.Println("b = ", string(b))        // b =  b
	log.Println("buff = ", buff.String()) // buff =  cdefg

	bs, _ := buff.ReadBytes('e')          //从当前位置读取直到遇到字符 'e'（包含 'e'）
	log.Println("bs = ", string(bs))      // bs =  cde
	log.Println("buff = ", buff.String()) // buff =  fg

	buff.Reset()
	buff.WriteString("编译输出GO")
	/**
	ReadRune() 读取一个 UTF-8 字符（rune）
	返回值：
	r: rune 值（Unicode 码点）32534 对应汉字 "编"
	l: 该字符占用的字节数，汉字在 UTF-8 中占3个字节
	错误值（被忽略）
	*/
	r, l, _ := buff.ReadRune()
	log.Println("r = ", r, ", l = ", l, ", string(r) = ", string(r))
	// r =  32534 , l =  3 , string(r) =  编

	buff.Reset()
	buff.WriteString("qwer")              // len = 4, off = 0
	readstr1, _ := buff.ReadString('?')   // 当分隔符不存在时，ReadString() 会读取到缓冲区末尾，off = 4
	log.Println("readstr1 = ", readstr1)  // str =  qwer 读取到完整的字符串
	log.Println("buff = ", buff.String()) // buff =   空字符串，表示缓存区已经读完

	buff.WriteString("qwer")              // len = 8，off = 4
	readstr2, _ := buff.ReadString('w')   //off = 6
	log.Println("readstr2 = ", readstr2)  // str =  qw  读取到分隔符为止的内容
	log.Println("buff = ", buff.String()) // buff =  er   剩余未读的内容

	file, _ := os.Open("doc.go")
	buff.Reset()
	buff.ReadFrom(file)
	log.Println("doc.go = ", buff.String()) // doc.go =  123

	buff.Reset()
	buff.WriteString("中国人")        //写入中文字符串
	cbyte := buff.Bytes()          //返回缓冲区的底层字节切片
	log.Println("cbyte = ", cbyte) // cbyte =  [228 184 173 229 155 189 228 186 186]
	//"中" = 228, 184, 173 (3字节)

	//strconv包
	//字符串转[]byte
	sum := []byte("hello") // 将字符串 "hello" 转换为字节切片
	fmt.Println(sum)
	fmt.Printf("sum的类型：%T\n", sum) //byte是unit8的别名

	//字符串->整数  strconv.Atoi()或strconv.ParseInt()
	// 按照 10进制 转换，返回 int 类型
	i, _ := strconv.Atoi("33234") //将字符串 "33234" 转换为 int 类型
	fmt.Printf("%T\n", i)         // int
	fmt.Println(i)
	// param1：要转化的字符串
	// param2：转换的进制，如 2,8,16,32
	// param3：返回bit的大小（注意，字面量显示还是 int64） 0 根据平台自动选择，表示 int64
	i2, _ := strconv.ParseInt("33234", 10, 0)
	fmt.Printf("%T\n", i2) // int64
	fmt.Println(i2)

	//字符串->浮点数  strconv.ParseFloat()
	// 参数类似 ParseInt
	val, _ := strconv.ParseFloat("33.33", 32)
	fmt.Printf("type: %T\n", val) // type: float64

	val2, _ := strconv.ParseFloat("33.33", 64)
	fmt.Printf("type: %T\n", val2) // type: float64

	//整数->字符串  strconv.Itoa()或strconv.FormatInt()
	num2 := 180
	// 默认按照10进制转换
	f1 := strconv.Itoa(num2)
	// param1: 要转换的数字(必须是int64类型)
	// param2: 转换的进制
	f2 := strconv.FormatInt(int64(num2), 10)
	fmt.Printf("type: %T\n", f1) // type:string
	fmt.Printf("type: %T\n", f2) // type:string

	//浮点数 —> 整数：使用 strconv.FormatFloat
	// param1: 要转换的浮点数
	// param2: 格式 ('f', 'e', 'E', 'g', 'G')
	// param3: 精度（小数点后位数）
	// param4: 位大小（32 或 64）
	num3 := 23423134.323422
	fmt.Println(strconv.FormatFloat(float64(num3), 'f', -1, 64)) // 普通模式
	fmt.Println(strconv.FormatFloat(float64(num3), 'b', -1, 64)) // 二进制模式
	fmt.Println(strconv.FormatFloat(float64(num3), 'e', -1, 64)) // 科学记数法
	fmt.Println(strconv.FormatFloat(float64(num3), 'E', -1, 64)) // 同上，显示为E
	fmt.Println(strconv.FormatFloat(float64(num3), 'g', -1, 64)) // 指数大时用科学记数，否则普通模式
	fmt.Println(strconv.FormatFloat(float64(num3), 'G', -1, 64)) // 同上，显示为E

	//字符串 和 bool 类型转换
	// string --> bool
	flagBool, _ := strconv.ParseBool("true")
	fmt.Println(flagBool)
	fmt.Printf("type: %T\n", flagBool) // type:bool
	// bool --> string
	flagStr := strconv.FormatBool(true)
	fmt.Println(flagStr)
	fmt.Printf("type: %T\n", flagStr) // type:string

	//unicode
}
