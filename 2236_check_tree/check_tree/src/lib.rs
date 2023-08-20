
use std::rc::Rc;
use std::cell::RefCell;
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

// use std::rc::Rc;
// use std::cell::RefCell;
struct Solution;
impl Solution {
    pub fn check_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(root_node) = root{
           return  match (root_node.borrow().left.clone(), root_node.borrow().right.clone()){
                (Some(left), Some(right))=>{
                     root_node.borrow().val == left.borrow().val + right.borrow().val
                },
                _ => false,
            }
        }

        false

    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
    }
}
