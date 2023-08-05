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
    // 递归1：0ms 2.1MB
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

    // https://github.com/kingeasternsun/leetcode-cn/tree/master/21_merge_two_lists/merge_two_lists
    // 递归2：4ms 2.1MB
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

    // 迭代1： 0ms 2.0MB
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

    // 迭代2： 0ms 1.9MB
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = None;
        let mut cur = &mut head;
        let mut cur1 = list1;
        let mut cur2 = list2;

        loop {
            if cur1.is_some() && cur2.is_some() {
                if cur1.as_ref().unwrap().val < cur2.as_ref().unwrap().val {
                    let mut node1 = cur1.unwrap();
                    cur1 = node1.next.take();
                    cur = &mut cur.insert(node1).next;
                } else {
                    let mut node2 = cur2.unwrap();
                    cur2 = node2.next.take();
                    cur = &mut cur.insert(node2).next;
                }
            } else {
                *cur = cur1.or(cur2); // 这个 or 的用法牛逼
                return head;
            }
        }
    }
}
