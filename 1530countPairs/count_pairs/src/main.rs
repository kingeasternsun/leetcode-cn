fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn count_pairs(root: Option<Rc<RefCell<TreeNode>>>, distance: i32) -> i32 {
        let mut sum = 0;
        if distance < 2 {
            return 0;
        }

        Self::dfs2(root, distance, &mut sum);
        sum
    }

    pub fn dfs(root: Option<Rc<RefCell<TreeNode>>>, distance: i32, sum: &mut i32) -> Vec<i32> {
        let mut res = Vec::new();

        if root.is_none() {
            return res;
        }

        let cur = root.unwrap();
        if cur.borrow().left.is_none() && cur.borrow().right.is_none() {
            res.push(0);
            return res;
        }

        let left = Self::dfs(cur.borrow().left.clone(), distance, sum);
        let right = Self::dfs(cur.borrow().right.clone(), distance, sum);

        for dis1 in &left {
            for dis2 in &right {
                if dis1 + dis2 + 2 <= distance {
                    *sum = *sum + 1;
                }
            }
        }

        for dis1 in left {
            if dis1 + 1 < distance {
                res.push(dis1 + 1);
            }
        }

        for dis2 in right {
            if dis2 + 1 < distance {
                res.push(dis2 + 1);
            }
        }
        res
    }

    //因为题目中distance小于等于10,所以从根节点到叶子节点的距离最大为9； 返回从当前节点到各个叶子节点距离为0到9的个数
    pub fn dfs2(root: Option<Rc<RefCell<TreeNode>>>, distance: i32, sum: &mut i32) -> [i32; 10] {
        // let mut res = Vec::new();

        if let Some(cur) = root {
            let mut res = [0; 10];
            if cur.borrow().left.is_none() && cur.borrow().right.is_none() {
                res[0] = 1; //到叶子节点为0的个数为1
                return res;
            }

            let left = Self::dfs2(cur.borrow().left.clone(), distance, sum);
            let right = Self::dfs2(cur.borrow().right.clone(), distance, sum);

            for i in 0..10 {
                for j in 0..10 {
                    if i + j + 2 <= distance {
                        *sum = *sum + left[i as usize] * right[j as usize]; //左子树中距离为i的个数 程序 右子树中距离为j的个数
                    }
                }
            }

            for i in 0..9 {
                res[i + 1] += left[i] //从当前节点到叶子节点距离为i+1， 就等于从左子节点到叶子节点距离为i
            }
            for i in 0..9 {
                res[i + 1] += right[i]
            }

            res
        } else {
            [0; 10]
        }
    }
}
