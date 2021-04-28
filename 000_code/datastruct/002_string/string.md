[toc]

# 1: [串String](https://github.com/lppgo/over-algorithm/blob/master/05-%E4%B8%B2%E4%B8%8E%E5%B9%BF%E4%B9%89%E8%A1%A8/01-%E4%B8%B2%E7%AE%80%E4%BB%8B.md)
串是由零到多个字符组成的有限序列，常称为字符串。基本所有语言都有对字符串的实现.

零个字符的串称为空串（nullstring），长度为 0，即"";

**场景**
```go
1: 翻转String
2: 查找子串
3: KMP 模式匹配算法

```
1: 朴素匹配算法的缺陷
经典模式匹配算法缺陷是，每次匹配失败后都要回溯到主串中i的位置，例如：主串为“000000000000000000001”，模式串为“0001”时，匹配正确之前的每次匹配都要到模式串的最后一个字符，才会回溯，这其中浪费的时间不容小觑。

KMP算法（克努特-莫里斯-普拉特算法）解决了上述问题，告别了回溯，KMP算法可以使事件复杂度从O(n*m)转变为O(n+m)。

2: [理解KMP模式匹配算法](https://github.com/lppgo/over-algorithm/blob/master/05-%E4%B8%B2%E4%B8%8E%E5%B9%BF%E4%B9%89%E8%A1%A8/03-KMP%E6%A8%A1%E5%BC%8F%E5%8C%B9%E9%85%8D%E7%AE%97%E6%B3%95.md)




# 2: 广义表