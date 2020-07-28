fn main() {
    println!("Hello, world!");
}


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
      right: None
    }
  }
}


pub struct Solution ;

use std::rc::Rc;
use std::cell::RefCell;
use std::cmp;
impl Solution {
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

        if root == None {
            return 0
        }

        Self::depth_helper(root.as_ref())

    }

    fn depth_helper(root:Option<&Rc<RefCell<TreeNode>>>)-> i32 {
        if let Some(node)=root {
            let left = Self::depth_helper(node.borrow().left.as_ref());
            let right = Self::depth_helper(node.borrow().right.as_ref());

            return cmp::max(left,right)+1
        }
        0
    }
}