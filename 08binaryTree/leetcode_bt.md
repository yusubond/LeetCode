### 二叉树的相关题目

[TOC]

#### 95 不同的二叉搜索树 II【中等】

题目要求：给定一个整数n，生成所有由1...n为节点所组成的二叉搜索树。

思路分析

题目看着很难，实际利用递归思想没有那么难。

从序列`1..n`取出数字`1` 作为当前树的根节点，那么就有i-1个元素用来构造左子树，二另外的i+1..n个元素用来构造右子树。最后我们将得到G(i-1)棵不同的左子树，G(n-i)棵不同的右子树，其中G为卡特兰数。

这样就可以以i为根节点，和两个可能的左右子树序列，遍历一遍左右子树序列，将左右子树和根节点链接起来。

```go
func generateTrees(n int) []*TreeNode {
  if n == 0 { return nil }
  return generateTreesWithNode(1, n)
}

func generateTreesWithNode(start end int) []*TreeNode {
    res := make([]*TreeNode, 0)
    if start > end {
      // 表示空树
      res = append(res, nil)
      return res
    }
    for i := start; i <= end; i++ {
      left := generateTreesWithNode(start, i-1)
      right := generateTreesWithNode(i+1, end)
      for _, l := range left {
        for _, r := range right {
          res = append(res, &TreeNode{
            Val: i,
            Left: l,
            Right: r,
          })
        }
      }
    }
    return res
  }
```

#### 98 验证二叉搜索树【中等】

题目要求：https://leetcode-cn.com/problems/validate-binary-search-tree/

算法分析：

```go
// date 2020/03/21
// 算法1：遍历其中序遍历序列，并判断其序列是否为递增
func isValidBST(root *TreeNode) bool {
    nums := inOrder(root)
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] >= nums[i+1] { return false }
    }
    return true
}

func inOrder(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := make([]int, 0)
    if root.Left != nil { res = append(res, inOrder(root.Left)...) }
    res = append(res, root.Val)
    if root.Right != nil { res = append(res, inOrder(root.Right)...) }
    return res
}
// 算法2：递归
// 递归判断当前结点的值是否位于上下边界之中
// 时间复杂度O(N)
func isValidBST(root *TreeNode) bool {
  if root == nil { return true }
  var isValid func(root *TreeNode, min, max float64) bool
  isValid = func(root *TreeNode, min, max float64) bool {
    if root == nil { return true }
    if float64(root.Val) <= min || float64(root.Val) >= max { return false }
    return isValid(root.Left, min, float64(root.Val)) && isValid(root.Right, float64(root.Val), max)
  }
  return isValid(root, math.Inf(-1), math.Inf(0))
}

// 算法3
// 其思想是迭代版的中序遍历，因为栈中已经保留所有的结点，只需要出栈判断即可
// 时间复杂度O(N)，空间复杂度O(1),最优解
func isValidBST(root *TreeNode) bool {
    if root == nil { return true }
    stack := make([]*TreeNode, 0)
    preValue := math.Inf(-1)

    for len(stack) != 0 || root != nil {
        // 将左子树全部放入栈
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 出栈，即取出左子树的最后一个节点
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 判断其是否大于前继结点
        if float64(root.Val) <= preValue { return false }
        // 更新前继节点和前继节点的值
        preValue = float64(root.Val)
        root = root.Right
    }
    return true
}
```

#### 101 对称二叉树【简单】

题目链接：https://leetcode.com/problems/symmetric-tree/

思路分析

算法1：

如果一个树的左子树和右子树对称，那么这个树就是对称的。那么两个树在什么情况下互为镜像呢？

> 如果两个树同时满足两个条件，则互为镜像。
>
> 1.它们的根节点具有相同的值；
>
> 2.每个树的右子树和另一个树的左子树对称。

```go
// 算法一:递归
func isSymmetrics(root *TreeNode) bool {
  var isMirror func(r1, r2 *TreeNode) bool
  isMirror = func(r1, r2 *TreeNode) bool {
    if r1 == nil && r2 == nil { return true }
    if r1 == nil || r2 == nil { return false }
    return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
  }
  return isMirror(root, root)
}
```

算法2：类似层序遍历，将子树的左右结点依次放入队列，通过判断队列中的连续的两个值是否相同来确认是否对称。

```go
// 算法二:迭代
func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root, root)
    var n int
    var t1, t2 *TreeNode
    for len(queue) != 0 {
        n = len(queue)
        for i := 0; i < n; i += 2 {
            t1, t2 = queue[i], queue[i+1]
            if t1 == nil && t2 == nil { continue }
            if t1 == nil || t2 == nil { return false }
            if t1.Val != t2.Val { return false }
            queue = append(queue, t1.Left, t2.Right, t1.Right, t2.Left)
        }
        queue = queue[n:]
    }
    return true
}
```

#### 104 二叉树的最大深度【简单】

题目链接：https://leetcode.com/problems/maximum-depth-of-binary-tree/

题目要求：给定一棵二叉树，找出其最大深度。

思路分析

算法：递归，递归调用分别得到左子树和右子树的深度，然后比较取最大值，并+1返回结果

```go
// 算法一: 递归，采用自底向上的递归思想
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自底向上的递归
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    l, r := maxDepth(root.Left), maxDepth(root.Right)
    if l > r { return l+1 }
    return r + 1
}

// 算法二
// bfs广度优先搜索
// 时间复杂度O(N)，空间复杂度O(N)
func maxDepth(root *TreeNode) int {
  if root == nil {return 0}
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  depth, n := 0, 0
  for len(queue) != 0 {
    n = len(queue)
    for i := 0; i < n; i++ {
      if queue[i].Left != nil {
        queue = append(queue, queue[i].Left)
      }
      if queue[i].Right != nil {
        queue = append(queue, queue[i].Right)
      }
    }
    queue = queue[n:]
    depth++
  }
  return depth
}

// 算法三 dfs深度优先搜索
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自顶向下的递归
func maxDepth(root *TreeNode) int {
    // dfs
    var dfs func(root *TreeNode, level int) int
    dfs = func(root *TreeNode, level int) int {
        if root == nil { return level }
        if root.Left == nil && root.Right == nil { return level+1 }
        if root.Left == nil {
            return dfs(root.Right, level+1)
        }
        if root.Right == nil {
            return dfs(root.Left, level+1)
        }
        l, r := dfs(root.Left, level+1), dfs(root.Right, level+1)
        if l > r { return l }
        return r
    }
    return dfs(root, 0)
}
```

#### 105 从前序与中序遍历序列构造二叉树【中等】

算法分析：1）从中序序列中找到根节点，递归左右子树。

```go
// date 2020/02/26
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 { return nil }
  root := &TreeNode{Val: preorder[0]}
  index := 0
  for i, v := range inorder {
    if v == preorder[0] {
      index = i
      break
    }
  }
  root.Left = buildTree(preorder[1:index+1], inorder[:index])
  root.Right = buildTree(preorder[index+1:], inorder[index+1:])
  return root
}
```

####  106 从中序和后序遍历序列构造二叉树【中等】

思路分析

利用根节点在中序遍历的位置

```go
// 递归, 中序和后序
func buildTree(inorder, postorder []int) *TreeNode {
  if len(postorder) == 0 {return nil}
  n := len(postorder)
  root = &TreeNode{
    Val: postorder[n-1],
  }
  index := -1
  for i := 0; i < len(inorder); i++ {
    if inorder[i] == root.Val {
      index = i
      break
    }
  }
  if index >= 0 && index < n {
    root.Left = buildTree(inorder[:index], postorder[:index])
    root.Right = buildTree(inorder[index+1:], postorder[index:n-1]))
  }
  return root
}
// 递归，前序和中序
func buildTree(preorder, inorder []int) *TreeNode {
  if 0 == len(preorder) {return nil}
  n := len(preorder)
  root := &TreeNode{
    Val: preorder[0],
  }
  index := -1
  for i := 0; i < len(inorder); i++ {
    if inorder[i] == root.Val{
      index = i
      break
    }
  }
  if index >= 0 && index < n {
    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
  }
  return root
}
```

#### 110 平衡二叉树【简单】

题目要求：给定一个二叉树，判断其是否为高度平衡二叉树。【一个二叉树的每个节点的左右两个子树的高度差的绝对值不超过1，视为高度平衡二叉树】

算法一：递归

```go
// date 2020/02/18
func isBalanced(root *TreeNode) bool {
  if root == nil { return true }
  if !isBalanced(root.Left) || !isBalanced(root.Right) { return false }
  if Math.Abs(float64(getHeight(root.Left)-getHeight(root.Right))) > float64(1) { return false }
  return true
}
func getHeight(root *TreeNode) int {
  if root == nil { return 0 }
  if root.Left == nil && root.Right == nil { return 1 }
  l, r := getHeight(root.Left), getHeight(root.Right)
  if l > r { return l+1 }
  return r+1
}
```

#### 111 二叉树的最小深度【简单】

题}要求：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

算法分析

```go
// date 2020/03/21
// 递归算法
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil { return 1 + minDepth(root.Right) }
    if root.Right == nil { return 1 + minDepth(root.Left) }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if l > r { return r+1 }
    return l+1
}
// bfs广度优先搜索
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    var n, depth int
    for len(queue) != 0 {
        n = len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left == nil && queue[i].Right == nil { return depth+1 }
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        depth++
        queue = queue[n:]
    }
    return depth
}
// dfs深度优先搜索
func minDepth(root *TreeNode) int {
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int {
        // 空结点直接返回
        if root == nil { return depth }
        // 叶子结点
        if root.Left == nil && root.Right == nil { return depth+1 }
        if root.Left == nil {
            return dfs(root.Right, depth+1)
        }
        if root.Right == nil {
            return dfs(root.Left, depth+1)
        }
        l, r := dfs(root.Left, depth+1), dfs(root.Right, depth+1)
        if l < r { return l }
        return r
    }
    return dfs(root, 0)
}
```

#### 112 路径总和Path Sum【简单】

思路分析

算法1：自顶向下的递归思想

```go
// 自顶向下递归
func hasPathSum(root *TreeNode, sum int) bool {
  if root == nil {return false}
  sum -= root.Val
  if root.Left == nil && root.Right == nil {return sum == 0}
  return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```

算法2：迭代，利用queue保存当前结果，类型层序遍历。

```go
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {return false}
    queue, csum := make([]*TreeNode, 0), make([]int, 0)
    queue = append(queue, root)
    csum = append(csum, sum - root.Val)
    for 0 != len(queue) {
        n := len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left == nil && queue[i].Right == nil && csum[i] == 0 {return true}
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
                csum = append(csum, csum[i] - queue[i].Left.Val)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
                csum = append(csum, csum[i] - queue[i].Right.Val)
            }
        }
        queue = queue[n:]
        csum = csum[n:]
    }
    return false
}
```

#### 156 上下翻转二叉树【中等】

题目要求：给定一个二叉树，其中所有的右节点要么是具有兄弟节点（拥有相同父节点的左节点）的叶节点，要么为空，将此二叉树上下翻转并将它变成一棵树， 原来的右节点将转换成左叶节点。返回新的根。

算法分析：

```go
// date 2020/02/25
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var parent, parent_right *TreeNode
  for root != nil {
    // 0.保存当前left, right
    root_left := root.Left
    root.Left = parent_right  // 重新构建root,其left为上一论的right
    parent_right := root.Right // 保存right
    root.Right = parent // 重新构建root,其right为上一论的root
    parent = root
    root = root_left
  }
  return parent
}
```

#### 199 二叉树的右视图【中等】

算法分析：层序遍历的最后一个节点值。

```go
// date 2020/02/26
func rightSideView(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  for len(stack) != 0 {
    n := len(stack)
    res = append(res, stack[0].Val)
    for i := 0; i < n; i++ {
      if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
      if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
    }
    stack = stack[n:]
  }
  return res
}
```

#### 236 二叉树的最近公共祖先【中等】

题目要求：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

思路分析：

```go
// date 2020/03/29
// 递归
// 在左右子树中分别查找p,q结点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil || root == p || root == q { return root }
     if p == q { return p }
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)
     // 左右均不为空，表示p,q分别位于左右子树中
     if left != nil && right != nil { return root }
     // 左为空，表示p,q位于右子树,否则位于左子树
     if left == nil { return right }
     return left
}
```

#### 257 二叉树的所有路径【简单】

题目要求：给定一个二叉树，返回所有从根节点到叶子节点的路径。

算法分析：

```go
// date 2020/02/23
func binaryTreePaths(root *TreeNode) []string {
  res := make([]string, 0)
  if root == nil { return res }
  if root.Left == nil && root.Right == nil {
    res = append(res, fmt.Sprintf("%d", root.Val))
    return res
  } else if root.Left != nil {
    for _, v := range binaryTreePaths(root.Left) {
      res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
    }
  } else if root.Right != nil {
    for _, v := range binaryTreePaths(root.Right) {
      res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
    }
  }
  return res
}
```

```go
class MinStack {
public:
    /** initialize your data structure here. */
    stack<int> _stack;
    int _min = INT_MAX;
    MinStack() {
        
    }
    
    void push(int x) {
        if(_min >= x){
            if(!_stack.empty()){
                _stack.push(_min);
            }
            _min = x;
        }
        _stack.push(x);
    }
    
    void pop() {
        if(_stack.empty())
            return;
        if(_stack.size() == 1)
            _min = INT_MAX;
        else if(_min == _stack.top()){//下一个元素是下一个最小值
            _stack.pop();
            _min = _stack.top();
        }
        _stack.pop();
    }
    
    int top() {
        return _stack.top();
    }
    
    int getMin() {
        return _min;
    }
};
```

#### 270 最接近的二叉搜索树值【简单】

题目要求：给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

算法分析：使用层序遍历，广度优先搜索。

```go
// date 2020/02/23
func closestValue(root *TreeNode, target float64) int {
    if root == nil { return 0 }
    var res int
    cur := math.Inf(1)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        n := len(stack)
        for i := 0; i < n; i++ {
            if math.Abs(float64(stack[i].Val) - target) < cur {
                cur = math.Abs(float64(stack[i].Val) - target)
                res = stack[i].Val
            }
            if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
            if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
        }
        stack = stack[n:]
    }
    return res
}
```

#### 538 把二叉树转换成累加树【简单】

算法：逆序的中序遍历，查找。

```go
// date 2020/02/26
func convertBST(root *TreeNode) *TreeNode {
  var sum int
  decOrder(root, &sum)
  return root
}

func decOrder(root *TreeNode, sum *int) {
  if root == nil { return }
  decOrder(root.Right, sum)
  *sum += root.Val
  root.Val = *sum
  decOrder(root.Left, sum)
}
```

#### 543 二叉树的直径【简单】

题目要求：给定一棵二叉树，返回其直径。

算法分析：左右子树深度和的最大值。

```go
// date 2020/02/23
func diameterOfBinaryTree(root *TreeNode) int {
    v1, _ := findDepth(root)
    return v1
}

func findDepth(root *TreeNode) (int, int) {
  if root == nil { return 0, 0 }
  var v int
  v1, l := findDepth(root.Left)
  v2, r := findDepth(root.Right)
  if v1 > v2 {
    v = v1
  } else {
    v = v2
  }
  if l+r > v { v = l+r }
  if l > r { return v, l+1}
  return v, r+1
}
```

#### 545 二叉树的边界【中等】

题目要求：给定一个二叉树，按逆时针返回其边界。

思路分析：分别求其左边界，右边界，和所有的叶子节点，用visited辅助去重，算法如下：

```go
// date 2020/02/23
func boundaryOfBinaryTree(root *TreeNode) []int {
  if root == nil { return []int{} }
  res := make([]int, 0)
  left, right := make([]*TreeNode, 0), make([]*TreeNode, 0)
  visited := make(map[*TreeNode]int, 0)
  res = append(res, root.Val)
  visited[root] = 1
  // find left node
  pre := root.Left
  for pre != nil {
    left = append(left, pre)
    visited[pre] = 1
    if pre.Left != nil {
      pre = pre.Left
    } else {
      pre = pre.Right
    }
  }
  // find right node
  pre = root.Right
  for pre != nil {
    right = append(right, pre)
    visited[pre] = 1
    if pre.Right != nil {
      pre = pre.Right
    } else {
      pre = pre.Left
    }
  }
  leafs := findLeafs(root)
  // make res
  for i := 0; i < len(left); i++ {
    res = append(res, left[i].Val)
  }
  for i := 0; i < len(leafs); i++ {
    if _, ok := visited[leafs[i]]; ok { continue }
    res = append(res, ,leafs[i].Val)
  }
  for i := len(right) - 1; i >= 0; i-- {
    res = append(res, right[i].Val)
  }
  return res
}

func findLeafs(root *TreeNode) []*TreeNode {
  res := make([]*TreeNode, 0)
  if root == nil { return res }
  if root.Left == nil && root.Right == nil {
    res = append(res, root)
    return res
  }
  res = append(res, findLeafs(root.Left)...)
  res = append(res, findLeafs(root.Right)...)
  return res
}
```

#### 563 二叉树的坡度【简单】

题目要求：给定一棵二叉树，返回其坡度。

算法一：根据定义，递归计算。

```go
func findTilt(root *TreeNode) int {
    if root == nil || root.Left == nil && root.Right == nil { return 0 }
    left := sum(root.Left)
    right := sum(root.Right)
    var res int
    if left > right {
        res = left - right
    } else {
        res = right - left
    }
    return res + findTilt(root.Left) + findTilt(root.Right)
}

func sum(root *TreeNode) int {
    if root == nil { return 0 }
    return sum(root.Left) + sum(root.Right) + root.Val
}
```

#### 617 合并二叉树【简单】

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

```go
// date 2020/02/23
// 递归
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
  if t1 == nil { return t2 }
  if t2 == nil { return t1 }
  t1.Val += t2.Val
  t1.Left = mergeTrees(t1.Left, t2.Left)
  t1.Right = mergeTrees(t1.Right, t2.Right)
  return t1
}
```

#### 654 最大二叉树【中等】

题目要求：

给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
			左子树是通过数组中最大值左边部分构造出的最大二叉树。
			右子树是通过数组中最大值右边部分构造出的最大二叉树。
			通过给定的数组构建最大二叉树，并且输出这个树的根节点。

算法分析：

找到数组中的最大值，构建根节点，然后递归调用。

```go
// date 2020/02/25
// 递归版
func constructMaximumBinaryTree(nums []int) *TreeNode {
  if len(nums) == 0 { return nil }
  if len(nums) == 1 { return &TreeNode{Val: nums[0]} }
  p := 0
  for i, v := range nums {
    if v > nums[p] { p = i }
  }
  root := &TreeNode{Val: nums[p]}
  root.Left = constructMaximumBinaryTree(nums[:p])
  root.Right = constructMaximumBinaryTree(nums[p+1:])
  return root
}
```

#### 655 输出二叉树【中等】

题目要求：将二叉树输出到m*n的二维字符串数组中。

算法思路：先构建好res，然后逐层填充。

```go
// date 2020/02/25
func printTree(root *TreeNode) [][]string {
  depth := findDepth(root)
  length := 1 << depth - 1
  res := make([][]string, depth)
  for i := 0; i < depth; i++ {
    res[i] = make([]string, length)
  }
  fill(res, root, 0, 0, length)
}

func fill(res [][]string, t *TreeNode, i, l, r int) {
  if t == nil { return }
  res[i][(l+r)/2] = fmt.Sprintf("%d", root.Val)
  fill(res, t.Left, i+1, l, (l+r)/2)
  fill(res, t.Right, i+1, (l+r+1)/2, r)
}

func findDepth(root *TreeNode) int {
  if root == nil { return 0 }
  l, r := findDepth(root.Left), findDepth(root.Right)
  if l > r { return l+1 }
  return r+1
}
```

#### 662 二叉树的最大宽度【中等】

题目要求：给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
		链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree

算法分析：使用栈stack逐层遍历，计算每一层的宽度，注意全空节点的时候要返回。时间复杂度O(n)，空间复杂度O(logn)。

```go
// date 2020/02/26
func widthOfBinaryTree(root *TreeNode) int {
  res := 0
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  var n, i, j int
  var isAllNil bool
  for 0 != len(stack) {
    i, j, n = 0, len(stack)-1, len(stack)
    // 去掉每一层两头的空节点
    for i < j && stack[i] == nil { i++ }
    for i < j && stack[j] == nil { j-- }
    // 计算结果
    if j-i+1 > res { res = j-i+1 }
    // 添加下一层
    isAllNil = true
    for ; i <= j; i++ {
      if stack[i] != nil {
        stack = append(stack, stack[i].Left, stack[i].Right)
        isAllNil = false
      } else {
        stack = append(stack, nil, nil)
      }
    }
    if isAllNil { break }
    stack = stack[n:]
  }
  return res
}
```

#### 669 修剪二叉搜索树【简单】

题目要求：给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

算法分析：

```go
// date 2020/02/25
func trimBST(root *TreeNode, L, R int) *TreeNode {
  if root == nil { return nil }
  if root.Val < L { return trimBST(root.Right, L, R) }
  if root.Val > R { return trimBST(root.Left, L, R) }
  root.Left = trimBST(root.Left, L, R)
  root.Right = trimBST(root.Right, L, R)
  return root
}
```

#### 703 数据流中第K大元素【简单】

解题思路：

0。使用数据构建一个容量为k的小顶堆。

1。如果小顶堆堆满后，向其增加元素，若大于堆顶元素，则重新建堆，否则掉弃。

2。堆顶元素即为第k大元素。

```go
// date 2020/02/19
type KthLargest struct {
  minHeap []int
  size
}

func Constructor(k int, nums []int) KthLargest {
  res := KthLargest{
    minHeap: make([]int, k),
  }
  for _, v := range nums {
    res.Add(v)
  }
}

func (this *KthLargest) Add(val int) int {
  if this.size < len(this.minHeap) {
    this.size++
    this.minHeap[this.size-1] = val
    if this.size == len(this.minHeap) {
      this.makeMinHeap()
    }
  } else if this.minHeap[0] < val {
    this.minHeap[0] = val
    this.minHeapify(0)
  }
  return this.minHeap[0]
}

func (this *KthLargest) makeMinHeap() {
  // 为什么最后一个根节点是n >> 1 - 1
  // 长度为n，最后一个叶子节点的索引是n-1; left = i >> 1 + 1; right = i >> 1 + 2
  // 父节点是 n >> 1 - 1
  for i := len(this.minHeap) >> 1 - 1; i >= 0; i-- {
    this.minHeapify(i)
  }
}

func (this *KthLargest) minHeapify(i int) {
  if i > len(this.minHeap) {return}
  temp, n := this.minHeap[i], len(this.minHeap)
  left := i << 1 + 1
  right := i << 1 + 2
  for left < n {
    right = left + 1
    // 找到left, right中较小的那个
    if right < n && this.minHeap[left] > this.minHeap[right] { left++ }
    // left, right 均小于root否和小顶堆，跳出
    if this.minHeap[left] >= this.minHeap[i] { break }
    this.minHeap[i] = this.minHeap[left]
    this.minHeap[left] = temp
		// left发生变化，继续调整
    i = left
    left = i << 1 + 1
  }
}
```

#### 814 二叉树剪枝【中等】

题目要求：给定一个二叉树，每个节点的值要么是0，要么是1。返回移除了所有不包括1的子树的原二叉树。

算法分析：

使用containOne()函数判断节点以及节点的左右子树是否包含1。

```go
// date 2020/02/25
func pruneTree(root *TreeNode) *TreeNode {
  if containsOne(root) { return root }
  return nil
}

func containsOne(root *TreeNode) bool {
  if root == nil { return false }
  l, r := containsOne(root.Left), containsOne(root.Right)
  if !l { root.Left = nil }
  if !r { root.Right = nil }
  return root.Val == 1 || l || r
}
```

#### 993 二叉树的堂兄弟节点【简单】

题目链接：https://leetcode-cn.com/problems/cousins-in-binary-tree/

题目要求：给定一个二叉树和两个节点元素值，判断两个节点值是否是堂兄弟节点。

堂兄弟节点的定义：两个节点的深度一样，但父节点不一样。

算法分析：

使用findFatherAndDepth函数递归找到每个节点的父节点和深度，并进行比较。

```go
// date 2020/02/25
func isCousins(root *TreeNode, x, y int) bool {
  xf, xd := findFatherAndDepth(root, x)
  yf, yd := findFatherAndDepth(root, y)
  return xd == yd && xf != yf
}
func findFatherAndDepth(root *TreeNode, x int) (*TreeNode, int) {
  if root == nil || root.Val == x { return nil, 0 }
  if root.Left != nil && root.Left.Val == x { return root, 1 }
  if root.Right != nil && root.Right.Val == y { return root, 1 }
  l, lv := findFatherAndDepth(root.Left, x)
  r, rv := findFatherAndDepth(root.Right, x)
  if l != nil { return l, lv+1 }
  return r, rv+1
}
```

#### 1104 二叉树寻路【中等】

题目要求：https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/

算法分析：

```go
// date 2020/02/25
func pathInZigZagTree(label int) []int {
  res := make([]int, 0)
  for label != 1 {
    res = append(res, label)
    label >>= 1
    y := highBit(label)
    lable = ^(label-y)+y
  }
  res = append(res, 1)
  i, j := 0, len(res)-1
  for i < j {
    t := res[i]
    res[i] = res[j]
    res[j] = t
    i++
    j--
  }
  return res
}label = label ^(1 << (label.bit_length() - 1)) - 1

func bitLen(x int) int {
  res := 0
  for x > 0 {
    x >>= 1
    res++
  }
  return res
}
```

#### 面试题27 二叉树的镜像【简单】

题目链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

题目要求：给定一棵二叉树，返回其镜像二叉树。

```go
// date 2020/02/23
// 递归算法
func mirrorTree(root *TreeNode) *TreeNode {
  if root == nil { return root }
  left := root.Left
  root.Left = root.Right
  root.Right = left
  root.Left = mirrorTree(root.Left)
  root.Right = mirrorTree(root.Right)
  return root
}
```

