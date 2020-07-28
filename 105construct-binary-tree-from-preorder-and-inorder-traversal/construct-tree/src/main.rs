fn main() {
    println!("Hello, world!");
}
// Definition for a binary tree node.

pub struct Solution;
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
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::build_helper(&preorder[..], &inorder[..])
    }


    pub fn build_helper(preorder:&[i32],inorder:&[i32])->Option<Rc<RefCell<TreeNode>>> {
        if preorder.len()==0{
            return None;
        }


        let root_value = preorder[0];
        let mut mid = inorder.len();

        for i in 0..inorder.len(){
            if inorder[i]==root_value{
                mid = i;
            }
        }

        if mid == inorder.len(){
            return None;
        }

        let left = Self::build_helper(&preorder[1..mid+1], &inorder[0..mid]);
        let right = Self::build_helper(&preorder[mid+1..], &inorder[mid+1..]);
        let root = Some(Rc::new(RefCell::new(TreeNode{val:root_value,left:left,right:right})));
        root

    }
}