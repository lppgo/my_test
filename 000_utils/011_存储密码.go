package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

func readme() {
	// 1:<普通方案> 单向hash加密

	// 明文密码做单向哈希后存储，单向哈希算法有一个特征：无法通过哈希后的摘要(digest)恢复原始数据，这也是“单向”二字的来源。
	// 常用的单向哈希算法包括SHA-256, SHA-1, MD5等。

	// 单向哈希有两个特性：
	// 1）同一个密码进行单向哈希，得到的总是唯一确定的摘要。
	// 2）计算速度快。随着技术进步，一秒钟能够完成数十亿次单向哈希计算
	// 结合上面两个特点，考虑到多数人所使用的密码为常见的组合，攻击者可以将所有密码的常见组合进行单向哈希，得到一个摘要组合,
	// 然后与数据库中的摘要进行比对即可获得对应的密码。这个摘要组合也被称为rainbow table。
	// 因此通过单向加密之后存储的数据，和明文存储没有多大区别。因此，一旦网站的数据库泄露，所有用户的密码本身就大白于天下。

	// 2: <进阶方案> “加盐”salt:
	// 通常的做法是，先将用户输入的密码进行一次MD5（或其它哈希算法）加密；将得到的 MD5 值前后加上一些只有管理员自己知道的随机串，再进行一次MD5加密。
	// 这个随机串中可以包括某些固定的串，也可以包括用户名（用来保证每个用户加密使用的密钥都不一样）

	// 3: <专家方案>
	// 上面的进阶方案在几年前也许是足够安全的方案，因为攻击者没有足够的资源建立这么多的rainbow table。
	// 但是，时至今日，因为并行计算能力的提升，这种攻击已经完全可行。
	// 怎么解决这个问题呢？只要时间与资源允许，没有破译不了的密码，
	// 所以方案是:故意增加密码计算所需耗费的资源和时间，使得任何人都不可获得足够的资源建立所需的rainbow table。

	// 这类方案有一个特点，算法中都有个因子，用于指明计算密码摘要所需要的资源和时间，也就是计算强度。
	// 计算强度越大，攻击者建立rainbow table越困难，以至于不可继续。
	// 这里推荐scrypt方案，scrypt是由著名的FreeBSD黑客Colin Percival为他的备份服务Tarsnap开发的
	// 目前Go语言里面支持的库 https://github.com/golang/crypto/tree/master/scrypt

}

// Go语言对这三种加密算法的实现如下所示：
func hash() {
	//import "crypto/sha256"
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))

	//import "crypto/sha1"
	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))

	//import "crypto/md5"
	h = md5.New()
	io.WriteString(h, "需要加密的密码")
	fmt.Printf("%x", h.Sum(nil))
}

func hashSalt() string {
	//import "crypto/md5"
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
	return last
}

func crypt() {
	salt := hashSalt()
	dk, err := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
	fmt.Println(dk, err)
}
