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

use std::cmp::Ordering;
impl Ord for ListNode {
    fn cmp(&self, other: &Self) -> Ordering {
        other.val.cmp(&self.val)
    }
}

impl PartialOrd for ListNode {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    // 4ms 2.96mb
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        if lists.len() == 0 {
            return None;
        }
        if lists.len() == 1 {
            return lists[0].clone();
        }

        let mut min_heap = std::collections::BinaryHeap::new();

        for node in lists {
            if let Some(n) = node {
                min_heap.push(n)
            }
        }

        let mut head = None;
        let mut pre_next = &mut head;

        while let Some(mut cur) = min_heap.pop() {
            // 断链 cur - X -> cur.next
            if let Some(cur_next) = cur.next.take() {
                min_heap.push(cur_next)
            }

            // cur 填入到 pre_next, 然后 pre_next 更新为 cur.next
            pre_next = &mut pre_next.insert(cur).next;
        }

        head
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        // assert_eq!(result, 4);
    }
}
