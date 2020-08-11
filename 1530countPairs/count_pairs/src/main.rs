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

        Self::dfs(root, distance, &mut sum);
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
}
