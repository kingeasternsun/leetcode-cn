use std::cell::RefCell;
use std::rc::Rc;
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

struct Solution;
impl Solution {
    pub fn longest_univalue_path(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let (_, _, max_len) = Self::dfs(root);
        return max_len;
    }

    // ret.0 表示节点的value, ret.1 表示以节点为起点的最大路径长度， ret.2 表示节点中的最大路径
    // 20ms 2.84mb
    fn dfs(root: Option<Rc<RefCell<TreeNode>>>) -> (Option<i32>, i32, i32) {
        match root {
            None => (None, 0, 0),
            Some(root) => {
                let (left_val, left_len, left_max) = Self::dfs(root.borrow().left.clone());
                let (right_val, right_len, right_max) = Self::dfs(root.borrow().right.clone());

                let mut cur_total_len = 0;
                let mut cur_len = 0;
                if left_val.map_or(false, |v| v == root.borrow().val) {
                    cur_total_len += left_len + 1;
                    cur_len = cur_len.max(left_len + 1);
                }

                if right_val.map_or(false, |v| v == root.borrow().val) {
                    cur_total_len += right_len + 1;
                    cur_len = cur_len.max(right_len + 1);
                }

                return (
                    Some(root.borrow().val),
                    cur_len,
                    cur_total_len.max(left_max).max(right_max),
                );
            }
        }
    }

    // 以当前节点作为起点的最长路径 和 当前二叉树中的最长路径
    // fn dfs1(root: Rc<RefCell<TreeNode>>) -> (i32, i32) {
    //     match (root.borrow().left.clone(), root.borrow().right.clone()) {
    //         (None, None) => (0, 0),
    //         (None, Some(right)) => {
    //             let tmp = Self::dfs1(right.clone());
    //             if root.borrow().val == right.borrow().val {
    //                 return (tmp.0 + 1, (tmp.1).max(tmp.0 + 1));
    //             }
    //             return (0, tmp.1);
    //         }
    //         (Some(left), None) => {
    //             let tmp = Self::dfs1(left.clone());
    //             if root.borrow().val == left.borrow().val {
    //                 return (tmp.0 + 1, (tmp.1).max(tmp.0 + 1));
    //             }
    //             return (0, tmp.1);
    //         }

    //         (Some(left), Some(right)) => {
    //             let mut tmp = (0, 0);
    //             let left_tmp = Self::dfs1(left.clone());
    //             let right_tmp = Self::dfs1(right.clone());
    //             if root.borrow().val == left.borrow().val {
    //                 tmp.0 += left_tmp.0 + 1;
    //             }
    //             if root.borrow().val == right.borrow().val {
    //                 tmp.0 += right_tmp.0 + 1;
    //             }
    //             (tmp.0, tmp.0.max(left_tmp.1).max(right_tmp.1))
    //         }
    //     }
    // }

    // 求以当前点为终点的最长路径
    // pre 表示父亲节点的值
    // pre_len 表示以父亲节点为终点的最长路径
    /*
    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, pre: Option<i32>, pre_len: i32, max_len: &mut i32) {
        if root.is_none() {
            return;
        }

        let root = root.unwrap();
        if let Some(pre_value) = pre {
            if root.borrow().val == pre_value {
                let new_len = pre_len + 1;
                *max_len = new_len.max(*max_len);
                Self::dfs(
                    root.borrow().left.clone(),
                    Some(root.borrow().val),
                    new_len,
                    max_len,
                );
                Self::dfs(
                    root.borrow().right.clone(),
                    Some(root.borrow().val),
                    new_len,
                    max_len,
                );
            } else {
                Self::dfs(
                    root.borrow().left.clone(),
                    Some(root.borrow().val),
                    0,
                    max_len,
                );
                Self::dfs(
                    root.borrow().right.clone(),
                    Some(root.borrow().val),
                    0,
                    max_len,
                );
            }
        } else {
            Self::dfs(
                root.borrow().left.clone(),
                Some(root.borrow().val),
                0,
                max_len,
            );
            Self::dfs(
                root.borrow().right.clone(),
                Some(root.borrow().val),
                0,
                max_len,
            );
        }
    }

    */
}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
