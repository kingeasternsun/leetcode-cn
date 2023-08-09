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
        Self::dp(root.clone(), true)
    }

    // 由于二叉树无法像数组那样直接跳到孙子节点，所以只能一级一级跳，通过一个can_rob 标识当前节点是否要跳过
    // 从当前节点开始，rob 当前节点可以得到的最大值 和 不rob 当前节点可以得到的最大值
    // 如果父节点被rob了， 当前节点就不能rob
    fn dp(root: Option<Rc<RefCell<TreeNode>>>, skip: bool) -> i32 {
        if root.is_none() {
            return 0;
        }

        let cur = root.unwrap();

        if !skip {
            // rob cur
            let ret1 = cur.borrow().val
                + Self::dp(cur.borrow().left.clone(), true)
                + Self::dp(cur.borrow().right.clone(), true);

            // not rob cur
            let ret0 = Self::dp(cur.borrow().left.clone(), false)
                + Self::dp(cur.borrow().right.clone(), false);
            return ret1.max(ret0);
        }

        // can not rob cur
        return Self::dp(cur.borrow().left.clone(), false)
            + Self::dp(cur.borrow().right.clone(), false);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {}
}
