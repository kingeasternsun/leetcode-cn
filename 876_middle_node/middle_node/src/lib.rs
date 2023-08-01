#[derive(PartialEq, Eq, Debug, Clone)]
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
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut one = head.as_ref();
        let mut two = head.as_ref();

        while let Some(two_cur) = two {
            if let Some(two_next) = two_cur.next.as_ref() {
                one = one.unwrap().next.as_ref();
                two = two_next.next.as_ref();
            } else {
                break;
            }
        }

        one.cloned()
    }

    pub fn middle_node2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut one = head.clone();
        let mut two = head;

        while let Some(two_cur) = two {
            if let Some(two_next) = two_cur.next {
                one = one.unwrap().next;
                two = two_next.next;
            } else {
                break;
            }
        }

        one
    }
}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
