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

struct Solution;

impl Solution {
    // 0ms 1.88mb
    pub fn swap_pairs0(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut ret = None;
        let mut pre_next = &mut ret; // the & mut of the next field belongs to pre node

        let mut cur = head;
        while let Some(mut node1) = cur {
            match node1.next.take() {
                None => {
                    pre_next.replace(node1);
                    return ret;
                }
                Some(mut node2) => {
                    cur = node2.next.take();
                    pre_next = &mut pre_next.insert(node2).next;
                    pre_next = &mut pre_next.insert(node1).next;
                }
            }
        }
        ret
    }

    // 0ms 2.21mb
    pub fn swap_pairs1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut cur = head;
        let mut pre = ListNode::new(0);
        let mut tail = &mut pre; //记录每一次翻转后的的最后一个节点

        //一开始 tail  first->second
        while let Some(mut first) = cur {
            // 当前节点 和 当前节点的next节点 分开
            let cur_next = first.next.take();
            if let Some(mut second) = cur_next {
                cur = second.next.replace(first);
                tail.next = Some(second);
                tail = tail.next.as_mut().unwrap().next.as_mut().unwrap();
            } else {
                tail.next = Some(first);
                return pre.next;
            }
        }

        return pre.next;
    }

    // 0ms 2.01mb
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if let Some(mut node1) = head {
            match node1.next.take() {
                None => return Some(node1),
                Some(mut node2) => {
                    node1.next = Self::swap_pairs(node2.next.take());
                    node2.next = Some(node1);
                    return Some(node2);
                }
            }
        }

        return None;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {}
}
