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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut resut = vec![];
        if let Some(root ) = &root{
            Self::dfs(root,& mut resut);
        }
        resut 
    }

    pub fn dfs(root:&Rc<RefCell<TreeNode>>, resut:& mut Vec<i32>)  {
        
        if let Some(left ) = &root.borrow().left{
            Self::dfs(left,resut);
        }

        resut.push(root.borrow().val);

        if let Some(right ) = &root.borrow().right{
            Self::dfs(right,resut);
        }

    }
}