
# 第一种解法

1. 如果总数模3剩余1，那么就减去模1的最小数或两个最小的模2的数
1. 如果总数模3剩余2，那么就减去模2的最小数或两个最小的模1的数

在Rust实现中，我们取Vec中前面的两个数字并求和，最普通的做法是
```Rust
let r2: Option<i32> = match mod2.len() {
    2 => Some(mod2[0] + mod2[1]),
    _ => None,
};
```

高级一些的做法是:
```Rust

let r2 = mod2.chunks_exact(2).next().map(|x| x[0] + x[1]);
```
这里要注意 `chunks`, `chunks_exact` 两个方法的不同，`chunks_exact` 会保证生成的每一个chunk的长度都是满足chunk_size大小的，这个题目中我们要用 `chunks_exact`
```Rust

    let slice = ['l', 'o', 'r', 'e', 'm'];
    let mut iter = slice.chunks_exact(2);
    assert_eq!(iter.next().unwrap(), &['l', 'o']);
    assert_eq!(iter.next().unwrap(), &['r', 'e']);
    assert!(iter.next().is_none());
    assert_eq!(iter.remainder(), &['m']);
```

```Rust
    let slice = ['l', 'o', 'r', 'e', 'm'];
    let mut iter = slice.chunks(2);
    assert_eq!(iter.next().unwrap(), &['l', 'o']);
    assert_eq!(iter.next().unwrap(), &['r', 'e']);
    assert_eq!(iter.next().unwrap(), &['m']);
    assert!(iter.next().is_none());
```

更高级的用法,使用 first_chunk， 但是这个需要开启 `![feature(slice_first_last_chunk)]`
```Rust

let r2 = mod2.first_chunk::<2>().map(|x| x[0] + x[1]);
```