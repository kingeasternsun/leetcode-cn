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
    //一路向左，最后出栈，取右节点，再一路向左
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = vec![];

        let mut node_stack = vec![];
        let mut cur_node = root;


        loop{


            loop{
                if let Some(cur) = cur_node{

                    // let cur = cur.borrow();
                    node_stack.push(cur.clone());
                    cur_node = cur.borrow().left.clone();

                }else{
                    break
                }

            }

            if node_stack.is_empty(){
                break
            }


            let left = node_stack.pop().unwrap();
            result.push(left.borrow().val);
            cur_node = left.borrow().right.clone();

        }
        
        result
    }
}