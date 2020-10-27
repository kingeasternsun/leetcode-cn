fn main() {
    println!("Hello, world!");
}

pub struct Solution;
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
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = Vec::new();
        Self::recur(root, &mut res);
        return res;
    }

    pub fn recur(root: Option<Rc<RefCell<TreeNode>>>, result: &mut Vec<i32>) {
        // if root.is_none(){
        //     return
        // }

        if let Some(cur) = root {
            result.push(cur.borrow().val);
            Self::recur(cur.borrow().left.clone(), result);
            Self::recur(cur.borrow().right.clone(), result);
        } else {
            return;
        }

        return;
    }

    //2.1 MB
    pub fn iter(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = Vec::new();
        let mut stack = Vec::new();
        stack.push(root);
        // let mut left = root.clone();

        while stack.len() > 0 {
            let mut left = stack.pop().unwrap();

            while let Some(cur) = left {
                res.push(cur.borrow().val);
                if cur.borrow().right.is_some() {
                    stack.push(cur.borrow().right.clone());
                }

                left = cur.borrow().left.clone();
            }
        }

        return res;
    }


    // 1.9MB 击败96.97
    pub fn iter2(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = Vec::new();
        if root.is_none(){
            return res
        }


        let mut stack = Vec::new();

        stack.push(root.unwrap());
        // let mut left = root.clone();

        while stack.len() > 0 {
            let mut left = stack.pop();

            while let Some(cur) = left {
                res.push(cur.borrow().val);
                if cur.borrow().right.is_some() {
                    stack.push(cur.borrow().right.clone().unwrap());
                }

                left = cur.borrow().left.clone();
            }
        }

        return res;
    }
}
