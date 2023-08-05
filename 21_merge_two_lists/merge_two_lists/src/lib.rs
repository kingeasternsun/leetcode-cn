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
    // 0ms 2.1MB
    pub fn merge_two_lists0(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        return match (list1, list2) {
            (None, None) => None,
            (None, Some(cur2)) => Some(cur2),
            (Some(cur1), None) => Some(cur1),
            (Some(mut cur1), Some(mut cur2)) => {
                if cur1.val <= cur2.val {
                    let next = cur1.next.take();
                    cur1.next = Self::merge_two_lists(next, Some(cur2));
                    return Some(cur1);
                } else {
                    let next = cur2.next.take();
                    cur2.next = Self::merge_two_lists(Some(cur1), next);
                    return Some(cur2);
                }
            }
        };
    }

    // 4ms 2.1MB
    pub fn merge_two_lists2(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        return match (list1, list2) {
            (Some(mut cur1), Some(mut cur2)) => {
                if cur1.val <= cur2.val {
                    let next = cur1.next.take();
                    cur1.next = Self::merge_two_lists(next, Some(cur2));
                    return Some(cur1);
                } else {
                    let next = cur2.next.take();
                    cur2.next = Self::merge_two_lists(Some(cur1), next);
                    return Some(cur2);
                }
            }
            (x, y) => x.or(y),
        };
    }

    // 0ms 2.0MB
    pub fn merge_two_lists3(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = None;
        let mut cur = &mut head;
        let mut cur1 = list1;
        let mut cur2 = list2;

        loop {
            match (cur1, cur2) {
                (Some(mut node1), Some(mut node2)) => {
                    if node1.val <= node2.val {
                        cur1 = node1.next.take();
                        cur2 = Some(node2);
                        cur = &mut cur.insert(node1).next; // 这个神来之笔
                    } else {
                        cur1 = Some(node1);
                        cur2 = node2.next.take();
                        cur = &mut cur.insert(node2).next; // 这个神来之笔
                    }
                }
                (x, y) => {
                    *cur = x.or(y); // 这个 or 的用法牛逼
                    return head;
                }
            }
        }
    }

     // 0ms 2.0MB
     pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = None;
        let mut cur = &mut head;
        let mut cur1 = list1;
        let mut cur2 = list2;

        loop {
            match (cur1, cur2) {
                (Some(mut node1), Some(mut node2)) => {
                    if node1.val <= node2.val {
                        cur1 = node1.next.take();
                        cur2 = Some(node2);
                        cur = &mut cur.insert(node1).next; // 这个神来之笔
                    } else {
                        cur1 = Some(node1);
                        cur2 = node2.next.take();
                        cur = &mut cur.insert(node2).next; // 这个神来之笔
                    }
                }
                (x, y) => {
                    *cur = x.or(y); // 这个 or 的用法牛逼
                    return head;
                }
            }
        }
    }
}
