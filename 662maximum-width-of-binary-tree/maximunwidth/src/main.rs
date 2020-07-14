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
pub struct Solution {}
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn width_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

        if root ==None{
            return 0;
        }
    
        let mut stack = Vec::new();
        let mut ids = Vec::new();
        stack.push(root);
        ids.push(0 as u64);

        let mut maxWidth = 1;
        while !stack.is_empty(){


            if ids[ids.len()-1] - ids[0]+1 > maxWidth{
                maxWidth = ids[ids.len()-1] - ids[0]+1
            }


            let mut new_stack = Vec::new();
            let mut new_ids = Vec::new();

            while let Some(top) = stack.pop(){

                let id  = ids.pop().unwrap();
                let top = top.unwrap();
                if top.borrow().left.is_some(){
                    new_stack.push(top.borrow().left.clone());
                    new_ids.push(id<<1 as u64);
                }
                if top.borrow().right.is_some(){
                    new_stack.push(top.borrow().right.clone());
                    new_ids.push(id<<1 as u64 + 1);
                }
            }

            stack = new_stack;
            ids = new_ids;

        }
        
        maxWidth as i32
    }

}