# leetcode
leetcode的代码
- 每一题至少会提供一种解法，如果有更优的会补上
- 分别用golang，rust实现。code like gopher， code like rustor
- 提供测试代码
每一题的结构,以1528题目为例 如下
```shell
├── 1528.go   #go实现方法
├── 1528_test.go #go 测试代码
└── restore-string #cargo new --lib restore-string创建的rust项目
    └── src
        └── lib.rs  #包含rust实现和test代码
```