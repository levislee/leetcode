<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../maximum-sum-bst-in-binary-tree "Maximum Sum BST in Binary Tree")
　　　　　　　　　　　　　　　　
[Next >](../bulb-switcher-iii "Bulb Switcher III")

## [1374. Generate a String With Characters That Have Odd Counts (Easy)](https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts "生成每种字符都是奇数个的字符串")

<p>Given an&nbsp;integer <code>n</code>, <em>return a string with <code>n</code>&nbsp;characters such that each character in such string occurs <strong>an odd number of times</strong></em>.</p>

<p>The returned string must contain only lowercase English letters. If there are multiples valid strings, return <strong>any</strong> of them. &nbsp;</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 4
<strong>Output:</strong> &quot;pppz&quot;
<strong>Explanation:</strong> &quot;pppz&quot; is a valid string since the character &#39;p&#39; occurs three times and the character &#39;z&#39; occurs once. Note that there are many other valid strings such as &quot;ohhh&quot; and &quot;love&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 2
<strong>Output:</strong> &quot;xy&quot;
<strong>Explanation:</strong> &quot;xy&quot; is a valid string since the characters &#39;x&#39; and &#39;y&#39; occur once. Note that there are many other valid strings such as &quot;ag&quot; and &quot;ur&quot;.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> n = 7
<strong>Output:</strong> &quot;holasss&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 500</code></li>
</ul>

### Related Topics
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
If n is odd, return a string of size n formed only by 'a', else return string formed with n-1 'a' and 1 'b''.
</details>