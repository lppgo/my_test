package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// strings 库对字符串操作具有很高的效率
func main_13() {
	// Compare()
	// EqualFold()

	// Contains()
	// Count()

	// Fields()

	// Join()

	// Index()

	// Map()

	// Repeat()

	// Replace()

	// Split()

	Trim()

	Builder()
	Reader()

	Seek()

}

// Compare
// 按照字典序比较两个字符串，通常情况下直接使用=，>，<会更快一些
func Compare() {
	s1 := "Hello"
	s2 := "hello"
	n := strings.Compare(s1, s2)
	fmt.Println(n) // -1
}

// Contains
func Contains() {
	// 字符串s中是否包含substr，返回true或者false
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true

	// ContainsAny用于判断子串中是否具有一个字符在源串s中。子串为空，返回false
	fmt.Println(strings.ContainsAny("team", "i"))     // false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     // true
	fmt.Println(strings.ContainsAny("failure", "ui")) // true
	fmt.Println(strings.ContainsAny("foo", ""))       // false
	fmt.Println(strings.ContainsAny("", ""))          // false

	// ContainsRune用于判断Ascall码代表的字符是否在源串s中
	fmt.Println(strings.ContainsRune("aardvark", 'a'))
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
}

// Count
// 判断子串在源串中的数量，如果子串为空，则长度为源串的长度+1
func Count() {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune 5=4+1

}

// EqualFold
// 在不区分大小写的情况下，判断两个字符串是否相同。
func EqualFold() {
	s1 := "Hello"
	s2 := "hello"
	b := strings.EqualFold(s1, s2)
	fmt.Println(b)
}

// Fields
func Fields() {
	// Fields：使用空白分割字符串。
	fmt.Printf("Fields are: %q", strings.Fields(" foo bar baz  ")) // ["foo" "bar" "baz"]

	// FieldsFunc：根据传入的函数分割字符串，如果当前参数c不是数字或者字母，返回true作为分割符号
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc(" foo1;bar2,baz3...", f)) // ["foo1" "bar2" "baz3"]
}

// HasPrefix
func HasPrefix() {
	// 判断字符串是否是以某个子串作为开头
}

// HasSuffix
func HasSuffix() {
	// 判断字符串是否是以某个子串作为者结尾
}

// Join
// 使用某个sep，连接字符串
func Join() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo,bar,baz

}

// Index
func Index() {
	// Index，IndexAny，IndexByte，IndexFunc，IndexRune
	// 都是返回满足条件的第一个位置，如果没有满足条件的数据，返回-1

	// LastIndex，LastIndexAny，LastIndexByte和LastIndexFunc

	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1

	// IndexAny 子串中的任意字符在源串出现的位置
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1

	// IndexByte，字符在字符串中出现的位置
	fmt.Println(strings.IndexByte("golang", 'g'))  // 0
	fmt.Println(strings.IndexByte("gophers", 'h')) // 3
	fmt.Println(strings.IndexByte("golang", 'x'))  // -1

	// IndexFunc 满足条件的作为筛选条件
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1

	// 某个字符在源串中的位置
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1

}

// Map
// 对字符串s中每一个字符执行map函数中的操作
func Map() {
	rot13 := func(r rune) rune { // r是遍历的每一个字符
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

}

// Repeat
// 重复一下s，count是重复的次数，不能传负数
func Repeat() {
	fmt.Println("ba: " + strings.Repeat("na", 2))
}

// Replace
func Replace() {
	// Replace和ReplaceAll
	// 使用new来替换old，替换的次数为n。如果n为负数，则替换所有的满足条件的子串
	// ReplaceAll使用new替换所有的old，相当于使用Replace时n<0
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinkky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
}

// Split
func Split() {
	/*
		func Split(s, sep string) []string
		func SplitAfter(s, sep string) []string
		func SplitAfterN(s, sep string, n int) []string
		func SplitN(s, sep string, n int) []string
	*/
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a","b","c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	// SplitN 定义返回之后的切片中包含的长度，最后一部分是未被处理的。
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a", "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)

	// 使用sep分割，分割出来的字符串中包含sep，可以限定分割之后返回的长度。
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a,", "b,c"]

	// 完全分割
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a,","b,", "c"]
}

// Trim
func Trim() {
	/*
		func Trim(s string, cutset string) string
		func TrimFunc(s string, f func(rune) bool) string
		func TrimLeft(s string, cutset string) string
		func TrimLeftFunc(s string, f func(rune) bool) string
		func TrimPrefix(s, prefix string) string
		func TrimSuffix(s, suffix string) string
		func TrimRight(s string, cutset string) string
		func TrimRightFunc(s string, f func(rune) bool) string
	*/
	// Trim 包含在cutset中的元素都会被去掉
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers

	// TrimFunc去掉满足条件的字符
	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// TrimLeft 去掉左边满足包含在cutset中的元素，直到遇到不在cutset中的元素为止
	fmt.Println(strings.TrimLeft("¡¡¡Hello,¡¡ Gophers!!!", "!¡")) // Hello, Gophers!!!

	// TrimLeftFunc 去掉左边属于函数返回值部分，直到遇到不在cutset中的元素为止
	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // Hello, Gophers!!!

	// TrimPrefix 去掉开头部分；TrimSuffix 去掉结尾部分
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Println("s:", s)
}

// Builder
// strings.Builder使用Write方法来高效的构建字符串。它最小化了内存拷贝，耗费零内存，不要拷贝非零的Builder
func Builder() {
	/*
		strings.Builder作为字符串拼接的利器，建议加大使用力度。
		func (b *Builder) Cap() int // 容量，涉及批量内存分配机制
		func (b *Builder) Grow(n int) // 手动分配内存数量
		func (b *Builder) Len() int // 当前builder中含有的所有字符长度
		func (b *Builder) Reset() // 清空builder
		func (b *Builder) String() string // 转化为字符串输出
		func (b *Builder) Write(p []byte) (int, error) // 往builder写入数据
		func (b *Builder) WriteByte(c byte) error // 往builder写入数据
		func (b *Builder) WriteRune(r rune) (int, error) // 往builder写入数据
		func (b *Builder) WriteString(s string) (int, error) // 往builder写入数据
	*/
	var builder strings.Builder
	for i := 3; i > 0; i-- {
		fmt.Fprintf(&builder, "%d...", i)
	}
	builder.WriteString("tail")
	fmt.Println(builder.String())
}

// Reader
func Reader() {
	// Reader通过读取字符串的方式，实现了接口io.Reader, io.ReaderAt, io.Seeker, io.WriterTo, io.ByteScanner和io.RuneScanner。
	// 零值Reader操作起来就像操作空字符串的io.Reader一样。

	/*
		func NewReader(s string) *Reader // 初始化reader实例
		func (r *Reader) Len() int // 未读字符长度
		func (r *Reader) Read(b []byte) (n int, err error)
		func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
		func (r *Reader) ReadByte() (byte, error)
		func (r *Reader) ReadRune() (ch rune, size int, err error)
		func (r *Reader) Reset(s string) // 重置以从s中读
		func (r *Reader) Seek(offset int64, whence int) (int64, error) // Seek implements the io.Seeker interface.
		func (r *Reader) Size() int64 // 字符串的原始长度
		func (r *Reader) UnreadByte() error
		func (r *Reader) UnreadRune() error
		func (r *Reader) WriteTo(w io.Writer) (n int64, err error) // WriteTo implements the io.WriterTo interface.
	*/

	// Len 返回未读字符串的长度
	// Size 返回字符串的原始长度大小
	// Read 读取字符串信息，读取之后会改变Len的返回值

	reader := strings.NewReader("abcdefghijklmn")
	fmt.Println(reader.Len()) // 输出14 初始时，未读长度等于字符串长度
	var buf []byte
	buf = make([]byte, 5)
	readLen, err := reader.Read(buf)
	fmt.Println("读取到的长度:", readLen) //读取到的长度5
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(buf)           //adcde
	fmt.Println(reader.Len())  //9  读取到了5个 剩余未读是14-5
	fmt.Println(reader.Size()) //14  字符串的长度

	// ReadAt 读取偏移off字节后的剩余信息到b中，ReadAt函数不会影响Len的数值
	r := strings.NewReader("abcdefghijklmn")
	var bufAt, buf2 []byte
	buf2 = make([]byte, 5)
	r.Read(buf2)
	fmt.Println("剩余未读的长度", r.Len())     //剩余未读的长度 9
	fmt.Println("已读取的内容", string(buf2)) //已读取的内容 abcde
	bufAt = make([]byte, 256)
	r.ReadAt(bufAt, 5)
	fmt.Println(string(bufAt)) //fghijklmn

	//测试下是否影响Len和Read方法
	fmt.Println("剩余未读的长度", r.Len())    //剩余未读的长度 9
	fmt.Println("已读取的内容", string(buf)) //已读取的内容 abcde

	// ReadByte 从当前已读取位置继续读取一个字节
	// UnreadByte 将当前已读取位置回退一位，当前位置的字节标记成未读取字节
	// ReadByte和UnreadByte会改变reader对象的长度
	r3 := strings.NewReader("abcdefghijklmn")
	//读取一个字节
	b, _ := r3.ReadByte()
	fmt.Println(string(b)) // a
	//int(r3.Size()) - r3.Len() 已读取字节数
	fmt.Println(int(r3.Size()) - r3.Len()) // 1

	//读取一个字节
	b, _ = r3.ReadByte()
	fmt.Println(string(b))                 // b
	fmt.Println(int(r3.Size()) - r3.Len()) // 2

	//回退一个字节
	r3.UnreadByte()
	fmt.Println(int(r3.Size()) - r3.Len()) // 1

	//读取一个字节
	b, _ = r3.ReadByte()
	fmt.Println(string(b))
}

// Seek
func Seek() {
	/*
		ReadAt方法并不会改变Len()的值，Seek的移位操作可以改变。
		offset是偏移的位置，whence是偏移起始位置，
		支持三种位置：io.SeekStart起始位，io.SeekCurrent当前位，io.SeekEnd末位。
		offset可以是负数，当时偏移起始位与offset相加得到的值不能小于0或者大于size()的长度
	*/
	r := strings.NewReader("abcdefghijklmn")

	var buf []byte
	buf = make([]byte, 5)
	r.Read(buf)
	fmt.Println(string(buf), r.Len()) //adcde 9

	buf = make([]byte, 5)
	r.Seek(-2, io.SeekCurrent) //从当前位置向前偏移两位 （5-2)
	r.Read(buf)
	fmt.Println(string(buf), r.Len()) //defgh 6

	buf = make([]byte, 5)
	r.Seek(-3, io.SeekEnd) //设置当前位置是末尾前移三位
	r.Read(buf)
	fmt.Println(string(buf), r.Len()) //lmn 0

	buf = make([]byte, 5)
	r.Seek(3, io.SeekStart) //设置当前位置是起始位后移三位
	r.Read(buf)
	fmt.Println(string(buf), r.Len()) //defgh 6
}
