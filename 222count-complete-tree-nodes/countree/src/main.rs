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

impl Solution {
    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {

        let h = Self::get_depth(root.clone());
        if h <=1{
            return h
        }
        
        (1 << (h-1))+Self::get_node_id(root, h)
    }

    pub fn get_depth(root:Option<Rc<RefCell<TreeNode>>>)->i32{

        let mut d = 0;
        let mut cur = root;
        while let Some(node) = cur{
            d = d+1;
            cur = node.borrow().left.clone();
        }

        d
    }

    pub fn exist(root:Option<Rc<RefCell<TreeNode>>>,  id:  i32, mut h:i32)->bool{

        h = h-1;
        let mut cur = root;
        while let Some(node)=cur{
            
            // let Some(node) = cur;
            h = h -1;
            let bit_flag = 1 << h;
            if (id & bit_flag) == bit_flag{
                cur = node.borrow().right.clone();
            }else{
                cur = node.borrow().left.clone();
            }

            if h == 0 { 

                if cur.is_some(){
                    return true;
                }

                return false;
            }

        }

        false
    }

    pub fn get_node_id(root:Option<Rc<RefCell<TreeNode>>>, h:  i32)->i32{
        let mut start = 0;
        let mut end = (1 << (h-1))-1;

        while start>=0 && start< end{

            let mid = (end  - start+1)/2+start;

            if !Self::exist(root.clone(),mid, h){
                end = mid-1;
                continue;
            }


            start = mid;

        }

        start

    }
}