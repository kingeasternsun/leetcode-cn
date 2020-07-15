fn main() {
    println!("Hello, world!");
}


pub struct Solution {}

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
use std::cmp::max;
impl Solution {
    pub fn rob(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        

        let (tmp0,tmp1) = Self::rob_help(root);
        return max(tmp0,tmp1)
    }

    pub fn rob_help(root: Option<Rc<RefCell<TreeNode>>>)->(i32,i32){

        if root.is_none(){
            return (0,0);
        }

        let cur = root.unwrap();
        let (left0,left1) = Self::rob_help(cur.borrow().left.clone());
        let (right0,right1) = Self::rob_help(cur.borrow().right.clone());

        //当前不rob, 左节点可以抢劫 也可以不抢劫
        let tmp0 = max(left0,left1)+max(right0,right1);

        //当前抢劫 左右节点不能rob
        let tmp1 = left0+right0+cur.borrow().val;

        (tmp0,tmp1)
    }
}