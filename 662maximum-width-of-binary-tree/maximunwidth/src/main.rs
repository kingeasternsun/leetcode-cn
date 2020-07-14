fn main() {
    println!("Hello, world!");

    let mut left2 = Some(Rc::new(RefCell::new(TreeNode::new(4))));
    let mut left1 = TreeNode::new(2);
    left1.left = left2;

    let mut left = Some(Rc::new(RefCell::new(left1)));
    let mut right = Some(Rc::new(RefCell::new(TreeNode::new(3))));

    // let mut root = Some(Rc::new(RefCell::new(TreeNode::new(1))));
    let mut root = TreeNode::new(1);
    root.left = left;
    root.right = right;

    Solution::width_of_binary_tree(Some(Rc::new(RefCell::new(root))));

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

        if root == None{
            return 0;
        }
    
        let mut stack = Vec::new();
        let mut ids = Vec::new();
        stack.push(root);
        ids.push(1 as u64);

        let mut max_width = 1;
        while !stack.is_empty(){

            // println!("{:?}",ids);

            if ids[ids.len()-1] - ids[0]+1 > max_width{
                max_width = ids[ids.len()-1] - ids[0]+1
            }


            let mut new_stack = Vec::new();
            let mut new_ids = Vec::new();

            for i in 0..stack.len(){

                // let id  = ids[i];
                let top = stack[i].as_ref().unwrap();
                if top.borrow().left.is_some(){
                    new_stack.push(top.borrow().left.clone());
                    new_ids.push(ids[i]<<1 as u64);
                }
                if top.borrow().right.is_some(){
                    new_stack.push(top.borrow().right.clone());
                    new_ids.push((ids[i]<<1) as u64 + 1);
                }
            }

            stack = new_stack;
            ids = new_ids;

        }
        
        max_width as i32
    }

}