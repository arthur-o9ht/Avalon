---
title: "Leetcode"
date: 2018-08-05T23:50:24+08:00
---

# LeetCode Record

>记录所有目前所有已经解出题目的解法及具体代码，以下代码若无具体说明均为**Go**<!--more-->

-----------
[TOC]

## 初级题目
### 数组

####   1.从排序数组中删除重复项
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 $O(1)$ 额外空间的条件下完成。

**示例：**
>给定数组 nums = [1,1,2], 
>函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2
> 你不需要考虑数组中超出新长度后面的元素。

**解答：**
```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    var l int = 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[l] {
            l++
            nums[l] = nums[i]
        }
    }
    return l+1
}
```
**知识点：**

因为这个题目不需要构建出新的数组，所以只需使用**原地算法**进行替换元素即可。使用两个下标，一个进行逐个元素循环，另一个只有当元素不同时才前进一。
>**原地算法**：
>在计算机科学中，一个原地算法（in-place algorithm）是一种使用小的，固定数量的额外之空间来转换资料的算法。当算法执行时，输入的资料通常会被要输出的部份覆盖掉。不是原地算法有时候称为非原地（not-in-place）或不得其所（out-of-place）。


#### 2.旋转数组
将包含 n 个元素的数组向右旋转 k 步。空间复杂度$O(1)$，且有最少三种办法
**示例：**
>如果  n = 7 ,  k = 3，给定数组  [1,2,3,4,5,6,7]  ，向右旋转后的结果为 [5,6,7,1,2,3,4]

**解答：**
```go
func reverse(nums []int, start int, end int) []int {
	var temp int
	for ; start < end; start, end = start+1, end-1 {
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
	}
	return nums
}

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	k = k % len(nums)
	nums = reverse(nums, 0, l-k-1)
	nums = reverse(nums, l-k, l-1)
	nums = reverse(nums, 0, l-1)
}
```

**知识点：**
采用**分治**的方法，先通过k算出真正需要移动的步数k，然后对不需要越过数组长度的元素(最终下标为：l-k-1)长度内的元素调换顺序，然后在对越过数组长度的元素进行调换顺序，然后整体调换顺序。


#### 3.存在重复
给定一个整数数组，判断是否存在重复元素。
如果任何值在数组中出现至少两次，函数应该返回 true。如果每个元素都不相同，则返回 false

**示例：无**

**解答：**
```go
//方法一，先排序，后使用类似原地算法的检测方式来检测是否有重复
func containsDuplicateSort(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Ints(nums)
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l] {
			l++
		} else {
			return true
		}
	}
	return false
}

//方法二，使用额外空间来进行检测
func containsDuplicateMemory(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	var numsMap = make(map[int]bool, l)
	for i := 0; i < l; i++ {
		if _, ok := numsMap[nums[i]]; !ok {
			numsMap[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}
```
**知识点：**

#### 4.只出现一次的数字
给定一个整数数组，除了某个元素外其余元素均出现两次。请找出这个只出现一次的元素。你的算法应该是一个线性时间复杂度。 你可以不用额外空间来实现它吗？

**示例：无**
**解答：**
```go
func singleNumber(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result 
}
```
**知识点：**
>异或运算
>如果a、b两个值不相同，则异或结果为1。如果a、b两个值相同，异或结果为0。

题目中说明除开只有单个的元素，其余元素均为出现两次，则通过
A^B^C^A^B = A^A^B^B^C = 0 ^ 0 ^ C = C 可得出单次出现的元素

#### 5.两个数组的交集 II
给定两个数组，写一个方法来计算它们的交集。

**示例：**
>给定 nums1 = [1, 2, 2, 1], nums2 = [2, 2], 返回 [2, 2].

注意：
	输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
	我们可以不考虑输出结果的顺序。
	
跟进:
如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果nums2的元素存储在磁盘上，内存是有限的，你不能一次加载所有的元素到内存中，你该怎么办？

**解答：**
```go
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	var countMap = make(map[int]int)
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	for i := 0; i < len(nums1); i++ {
		if v, ok := countMap[nums1[i]]; !ok {
			countMap[nums1[i]] = 1
		} else {
			countMap[nums1[i]] = v + 1
		}
	}

	for i := 0; i < len(nums2); i++ {
		if v, ok := countMap[nums2[i]]; ok && v >= 1 {
			res = append(res, nums2[i])
			countMap[nums2[i]]--
		}
	}
	return res
}
```

**知识点：**
先循环第一个数组，记录出现的元素及出现次数，记录的map A 用元素做key，value为出现次数。再循环第二个数组，当出现元素存在于第一个数组中，**并且出现次数>=1**，则推入最终的结果记录数组，并且A中该元素为k的value自减。

**TODO：**
后续思考跟进内容并给出答案
这里记录下别人的思考
>1. 如果不排序，$O(mn)$。  
>2. 如果m和n都在合理范围内，先排序，再一个一个对比，时间复杂度$O(nlgn + mlgm + m+n)$。  
>3. 如果m远小于n, 对n排序，m也排序$O(nlgn+mlgm+m+n)$，或者m不排序$O(nlgn + mn)$。 这两种都差不多。也可以对m不排序，在n里做binary search，这样复杂度降低为nlgn+mlgn, 降很低。 
>4. 如果n很大，n只能存在disk上。只能把m load到内存，然后n一个一个的读进来，和m对比，这时m可以用hash存，这样复杂度就为$O(n)$了。

[引用CSDN地址](https://blog.csdn.net/whl_program/article/details/71244305)

#### 6.旋转图像
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。
你必须在**原地（不可使用新的空间）**旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

**示例：**
>给定 **matrix** = 
>[	
>	[1,2,3],
>	[4,5,6],
>	[7,8,9]
>],
>原地旋转输入矩阵，使其变为:
>[
>	[7,4,1],
>	[8,5,2],
>	[9,6,3]
>]
>
>给定 **matrix** =
>[
> 	[ 5, 1, 9,11],
> 	[ 2, 4, 8,10],
> 	[13, 3, 6, 7],
> 	[15,14,12,16]
>], 
>
>原地旋转输入矩阵，使其变为:
>[
>	[15,13, 2, 5],
>	[14, 3, 4, 1],
>	[12, 6, 8, 9],
>	[16, 7,10,11]
>]
**解答：**
```go
func rotate(matrix [][]int)  {
    var temp int

	n := len(matrix)

	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < j {
				temp = matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = temp
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n/2 {
				temp = matrix[i][j]
				matrix[i][j] = matrix[i][len(matrix)-j-1]
				matrix[i][len(matrix)-j-1] = temp
			}
		}
	}
}
```
**知识点：**
>	[1,2,3],
>	[4,5,6],
>	[7,8,9]
>   先通过下标，xy的值 调换为 yx的值
>  [1,4,7]
>  [2,5,8]
>  [,3,6,9]
>  然后再行数据首尾相调换即可为
>	[7,4,1],
>	[8,5,2],
>	[9,6,3]

#### 7.加一

给定一个非负整数组成的非空数组，给整数加一。
可以假设整数不包含任何前导零，除了数字0本身。
最高位数字存放在列表的首位。

**示例：无**
**解答：**
```go
func plusOne(digits []int) []int {
    if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	l := len(digits)
	for i := 0; i < l; i++ {
		if digits[i] != 0 {
			return digits
		}
	}
	var res = []int{1}
	res = append(res, digits...)
	return res
}
```
**知识点：**
如果数组当前值为9，改为0并继续循环，如果不为9则直接返回数据。但是最终要检查如遇到最坏情况的话，数组中所有数据均为9时，则需要在数组头添加1后再输出

#### 8.移动0
给定一个数组 nums, 编写一个函数将所有 0 移动到它的末尾，同时保持非零元素的相对顺序。
**注意事项：**

1. 必须在原数组上操作，不要为一个新数组分配额外空间。
2. 尽量减少操作总数。


**示例：**
>定义 nums = [0, 1, 0, 3, 12]，调用函数之后， nums 应为 [1, 3, 12, 0, 0]。

**解答：**
```go
func moveZeroes(nums []int)  {
    if len(nums) == 0 {
		return
	}
	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			j++
		}
	}
}
```
**知识点：**
两个下表，一个顺序移动，一个停留在0处，只要还有非零的数值，就交换，这样后面只会剩下0
#### 9.两数之和

**示例：**
>给定 nums = [2, 7, 11, 15], target = 9
>因为 nums[0] + nums[1] = 2 + 7 = 9
>所以返回 [0, 1]

**解答：**
```go

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

```
**知识点：**
1、双循环暴力解决，时间复杂度 $O({n^2})$ 
2、可以通过map，但golang不知道怎么解决下标不同的同一个值的问题，时间复杂度$O(n)$，可以后续想想办法
3、结合1、2、方法，将1中的$O(n^2)$  降至$O(nlgn)$

#### 10.买卖股票的最佳时机 II

**示例：无**
**解答：**
```go

```
**知识点：**

#### 11.有效的数独

**示例：无**
**解答：**
```go

```
**知识点：**



### 字符串

#### 1.反转字符串
请编写一个函数，其功能是将输入的字符串反转过来。

**示例：**
>输入：s = "hello"
>返回："olleh"

**解答：**
```go
func reverseString(s string) string {
    var res = []uint8{}
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}
```
**知识点：**
字符串作为数组来理解的话，字符串其实就是一个[]uint8的数组，通过空间换时间，空间复杂度$O(n)$，时间复杂度$O(n)$


#### 2.颠倒整数
给定一个 32 位有符号整数，将整数中的数字进行反转。

**示例：**
>输入: 123
>输出: 321
>
>输入: -123
>输出: -321
>
>输入: 120
>输出: 21

假设我们的环境只能存储 32 位有符号整数，其数值范围是$ [−2^{31},  2^{31} − 1]$。根据这个假设，如果反转后的整数溢出，则返回 0。


**解答：**
```go
func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}
```
**知识点：**
一个整数可以通过不停的%10 /10 来获取每一位的数据，在通过每次结果的乘10累加获取最终的翻转值

#### 3.字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

**示例：**
>s = "leetcode"
>返回 0.

>s = "loveleetcode",
>返回 2.

**解答：**
```go
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	var res map[uint8]int = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := res[s[i]]; !ok {
			res[s[i]] = 1
		} else {
			res[s[i]]++
		}
	}

	for i := 0; i < len(s); i++ {
		if v, ok := res[s[i]]; v == 1 && ok {
			return i
		}
	}
	return -1
}
```
**知识点：**
1、这类查找存在不存在的均可暴力多层循环解决，但无技术含量，不进行实现了
2、本来想采用原地算法，但由于未规定特殊元素个数且需要查找的是特殊元素，而非重复元素，所以当特殊元素处于尾部或根本没特殊元素时，会检查失败，本方案抛弃
3、使用空间换时间，新建立一个数组放入字符串中每个元素且标记元素个数，循环完毕后在一次循环检查就能得出结论，方案可行且复杂度较低
#### 4.有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

**示例：**
>s = "anagram"，t = "nagaram"，返回 true
>s = "rat"，t = "car"，返回 false
**解答：**
```go
func isAnagram(s string, t string) bool {
    sMap := make([]int, 26)
    tMap := make([]int, 26)
    
    for i := 0; i < len(s); i++ {
        sMap[s[i]-'a']++
    }
    for i := 0; i < len(t); i++ {
        tMap[t[i]-'a']++
    }
    
    for i := 0; i < 26; i++ {
        if sMap[i] != tMap[i] {
            return false
        }
    }
    return true
}
```
**知识点：**
1. 两个字符串，首先想到用两个数组来标记各个字符串中字母出现的字符及字符的个数，然后对数组进行比对，又因为字母一共只有26个，所以数组设置长度只需要26即可，最终检测循环次数为26


#### 5.验证回文字符串
给定一个字符串，确定它是否是回文，只考虑字母数字字符和忽略大小写。

**示例：**
>"A man, a plan, a canal: Panama" 是回文字符串。
>"race a car" 不是回文字符串。

**解答：**
```go
func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	s = strings.ToLower(s)
	l := len(s) - 1
	for i := 0; i < len(s); {
		if s[i] < 'a' || s[i] > 'z' {
			i++
			continue
		}
		if s[l] < 'a' || s[l] > 'z' {
			l--
			continue
		}
		if s[l] != s[i] && l >= i {
			return false
		}
		i++
		l--
	}
	return true
}
```
**知识点：**
1、比较疑惑的是，自测通过检测，但是提交代码时，" " 无法通过测试…在本机实验时可以通过……比较疑惑
2、思路很简单，一个数组两个指针，分别从头部和尾部开始验证，当遇到非字母字符时跳过并指针移动即可，网上大部分思路也是如此，bu
#### 6.字符串转整数（atoi）
实现 atoi，将字符串转为整数。

在找到第一个非空字符之前，需要移除掉字符串中的空格字符。如果第一个非空字符是正号或负号，选取该符号，并将其与后面尽可能多的连续的数字组合起来，这部分字符即为整数的值。如果第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

字符串可以在形成整数的字符后面包括多余的字符，这些字符可以被忽略，它们对于函数没有影响。

当字符串中的第一个非空字符序列不是个有效的整数；或字符串为空；或字符串仅包含空白字符时，则不进行转换。

若函数不能执行有效的转换，返回 0。

**示例：**
>输入: "42"
>输出: 42

>输入: "words and 987"
>输出: 0
>解释: 第一个非空字符是 'w', 但它不是数字或正、负号。因此无法执行有效的转换。

**解答：**
```go
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	b := 1
	numIndex := 0
	for ; numIndex < len(str); numIndex++ {
		if str[numIndex] == ' ' {
			continue
		}
		break
	}
	if numIndex+1 > len(str) {
		return 0
	}
	if str[numIndex] == '-' || str[numIndex] == '+' {
		if str[numIndex] == '-' {
			b = -1
		}
		numIndex++
	}
	if numIndex+1 > len(str) {
		return 0
	}
	res := 0
	for i := numIndex; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		res = 10*res + int(str[i]-'0')
		if b*res >= math.MaxInt32 {
			return math.MaxInt32
		}

		if b*res <= math.MinInt32 {
			return math.MinInt32
		}
	}
	res = res * b
	return res
}
```
**知识点：**
思考步骤如下：
1、先去掉空格，切割字符串
2、找到符号位置，再次切割字符串
3、循环累加*10并累加当前，并且在每次累加后判断是否超出范围，如超出的话直接返回范围边界
4、不可在全部累加完成后再判断范围，因为可能会出现溢出导致数据异常，如累加数*1 变为负数

#### 7.实现strStr()

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

**示例：**
>输入: haystack = "hello", needle = "ll"
>输出: 2

**解答：**
```go
func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	if nLen == 0 {
		return 0
	}
	if nLen > hLen {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		j := 0
		for ; j < nLen; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == nLen {
			return i
		}
	}
	return -1
}

```
**知识点：**
1、不用将haystack全部循环，循环至(hLen-nlen)处即可，因为再往后循环的话，剩余字符串长度已经小于needle长度

#### 8.最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

**示例：**
>输入: ["flower","flow","flight"]
>输出: "fl"

>输入: ["dog","racecar","car"]
>输出: ""
>解释: 输入不存在公共前缀。

**解答：**
```go

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}

	}
	return strs[0]
}
```
**知识点：**
思路就是用第一个字符串的每一个字母和后面的字符串的同样位置进行对比，如果一直就继续，当出现不一致就跳出循环给出第一个字符串，从第一个位置开始到当前位置-1的字符串


### 链表

#### 1.删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

**示例：**
>给定一个链表: 1->2->3->4->5, 和 n = 2.
>当删除了倒数第二个节点后，链表变为 1->2->3->5.

**解答：**
```go
 // Definition for singly-linked list.
 // type ListNode struct {
 //     Val int
 //     Next *ListNode
 // }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var len int
	var header = &ListNode{
		Val: -1,
	}
	node := head
	for {
		len++
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	header.Next = head
	res := header
	for i := 0; i <= len; i++ {
		if i == len-n {
			res.Next = res.Next.Next
		} else {
			res = res.Next
		}
	}
	return header.Next
}

//网上耗时较少的一个写法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 {
        return head
    }
    if head == nil {
        return head
    }
    // if n == 1 && head.Next == nil {
    //     return nil
    // }
    ahead := head
    current := head
    
    for n > 0 && ahead != nil {
        ahead = ahead.Next
        n--
    }
    
    for ahead != nil && ahead.Next != nil {
        ahead = ahead.Next
        current = current.Next
    }
    if ahead == nil {
        return head.Next
    } else if current.Next == nil {
        return nil
    } else {
        current.Next = current.Next.Next
    }
    
    return head
}
```
**知识点：**
思路：先计算出链表长度，然后从头部开始循环，到指定长度时删除链表节点。小技巧是给一个默认头部，这样哪怕是要删除给出的链表的原本头部都是可以的，不会造成报错。

同时贴一个网友的解法，另一个思路，耗时差别应该不特别大

#### 2. 反转链表
反转一个单链表

**示例：**
>输入: 1->2->3->4->5->NULL
>输出: 5->4->3->2->1->NULL

**进阶：**
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
**解答：**
```go
 // Definition for singly-linked list.
 // type ListNode struct {
 //     Val int
 //     Next *ListNode
 // }
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodeList := &ListNode{Val: head.Val, Next: nil}
	if head.Next != nil {
		head = head.Next
		for {
			tmp := &ListNode{}
			tmp.Val = head.Val
			tmp.Next = nodeList
			nodeList = tmp
			if head.Next == nil {
				break
			}
			head = head.Next
		}
	}
	return nodeList
}

//耗时更少的写法
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    pre := head
    current := pre.Next
    pre.Next = nil
    next := head
    for current != nil {
        next = current.Next
        current.Next = pre
        pre = current
        current = next
    }
    return pre
}
```
**知识点：**

#### 3.合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

**示例：**
>输入：1->2->4, 1->3->4
>输出：1->1->2->3->4->4

**解答：**
```go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{}
	f := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	f = f.Next
	return f
}
```
**知识点：**
技巧是设定一个头结点，后续运算后可直接返回头结点。
#### 4.

**示例：**
>

**解答：**
```go

```
**知识点：**

#### 5.

**示例：**
>

**解答：**
```go

```
**知识点：**