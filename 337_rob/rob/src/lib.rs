use std::{cell::RefCell, rc::Rc};

struct Solution;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl Solution {

    pub fn rob(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        
        Self::dp(root.clone(),true)
    }

    // 从当前节点开始，rob 当前节点可以得到的最大值 和 不rob 当前节点可以得到的最大值
    // 如果父节点被rob了， 当前节点就不能rob
    fn dp(root: Option<Rc<RefCell<TreeNode>>>, can_rob: bool) -> i32 {
        if root.is_none() {
            return 0;
        }

        let cur = root.unwrap();

        if can_rob {
            // rob cur
            let ret1= cur.as_ref().borrow().val + Self::dp(cur.as_ref().borrow().left.clone(), false) + Self::dp(cur.as_ref().borrow().right.clone(),false);

            // not rob cur
            let ret0 =  Self::dp(cur.as_ref().borrow().left.clone(), true) + Self::dp(cur.as_ref().borrow().right.clone(),true);
            return ret1.max(ret0);
        }

        // can not rob cur
        return Self::dp(cur.as_ref().borrow().left.clone(), true) + Self::dp(cur.as_ref().borrow().right.clone(),true);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {}
}
