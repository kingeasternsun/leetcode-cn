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

pub struct Solution {}

use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {

        let mut stack = Vec::new();
        let mut k = k;

        let mut cur = root;

        loop {

            while(cur.is_some()){
                stack.push(cur.clone());
                cur = cur.unwrap().borrow().left.clone();
            }

            if let Some(head) = stack.pop(){
                k = k -1;
                if k==0{
                    return head.unwrap().borrow().val;
                }

                cur = head.unwrap().borrow().right.clone();

            }else{
                break
            }

            
        }


        0
    }
}