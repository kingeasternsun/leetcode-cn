fn main() {
    println!("Hello, world!");
}
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

pub struct Solution;
use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn largest_values(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = vec![];
        if root.is_none() {
            return res;
        }
        let mut nodes = vec![];
        nodes.push(root.unwrap());

        while nodes.len() > 0 {
            let mut new_nodes = vec![];
            let mut max_val = std::i32::MIN;

            // for v in nodes{

            //     if v.borrow().val>max_val{
            //         max_val = v.borrow().val;
            //     }

            //     if let Some(left) = v.borrow().left.clone(){
            //         new_nodes.push(left);
            //     }

            //     if let Some(right) = v.borrow().right.clone(){
            //         new_nodes.push(right);
            //     }

            // }

            nodes.into_iter().for_each(|v| {
                if v.borrow().val > max_val {
                    max_val = v.borrow().val;
                }

                if let Some(left) = v.borrow().left.clone() {
                    new_nodes.push(left);
                }

                if let Some(right) = v.borrow().right.clone() {
                    new_nodes.push(right);
                }
            });

            res.push(max_val);
            nodes = new_nodes;
        }

        res
    }
}
