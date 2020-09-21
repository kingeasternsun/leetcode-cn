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
      right: None
    }
  }
}
pub struct Solution;
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn convert_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut cur_sum = 0;

        Self::sum(root.clone(),& mut cur_sum);
        root
    }

    fn sum(root: Option<Rc<RefCell<TreeNode>>>,cur_sum :& mut i32){
        if let Some(cur_root) = root{

            Self::sum(cur_root.borrow().right.clone(), cur_sum);
            *cur_sum = *cur_sum + cur_root.borrow().val;
            cur_root.borrow_mut().val = *cur_sum;
            Self::sum(cur_root.borrow().left.clone(), cur_sum);

        }else{
            return 
        }
    }
}