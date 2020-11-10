fn main() {
    println!("Hello, world!");
}
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub struct Solution;


impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut cur = head;
        let mut pre = ListNode::new(0);
        let mut tail = &mut pre; //记录每一次翻转后的的最后一个节点

        //一开始 tail  first->second
        while let Some(mut first) = cur {
            cur = first.next.take(); // first和seconde 分开了。  tail->first second 
            if let Some(mut second) = cur {
                cur = second.next.take(); // first second   断开变为 first second cur

                //一开始的顺序  tail  first->second
                second.next = Some(first); //second->first
                // cur = second.next.replace(first);
                tail.next = Some(second);
                tail = tail.next.as_mut().unwrap().next.as_mut().unwrap();
            } else {
                //一开始的顺序  tail  first
                tail.next = Some(first);
                tail = tail.next.as_mut().unwrap()
            }
        }

        return pre.next;
    }
}
