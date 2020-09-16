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
      right: None
    }
  }
}
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {

        if let Some(root) = root{

            let left = Self::invert_tree(root.borrow().left.clone());
            let right = Self::invert_tree(root.borrow().right.clone());
            
            root.borrow_mut().left = right;
            root.borrow_mut().right = left;

            Some(root)
            // return Some(Rc::new(RefCell::new(TreeNode::new())));

        }else{
            None
        }

    }
}