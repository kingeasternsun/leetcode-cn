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
    // 0ms 2.6mb
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut tmp = vec![];
        let pre_sum = 0;
        Self::dfs(root, pre_sum, target_sum, & mut tmp, & mut ret);
        ret
    }
  
    fn dfs(root:Option<Rc<RefCell<TreeNode>>>,pre_sum:i32, target_sum:i32,tmp:& mut Vec<i32>, ret:& mut Vec<Vec<i32>>){
        if let Some(root)=root{
            if root.borrow().left.is_none() && root.borrow().right.is_none(){
                if pre_sum + root.borrow().val == target_sum{
                    tmp.push(root.borrow().val);
                    ret.push(tmp.clone());
                    tmp.pop();
                }
                return 
            }

            tmp.push(root.borrow().val);
            Self::dfs(root.borrow().left.clone(), pre_sum+root.borrow().val, target_sum, tmp, ret);
            Self::dfs(root.borrow().right.clone(), pre_sum+root.borrow().val, target_sum, tmp, ret);
            tmp.pop();

        }
        
    }
}


