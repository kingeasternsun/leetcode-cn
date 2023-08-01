struct Solution;

// https://www.youtube.com/watch?v=6c7pZYP_iIE
#[derive(Debug,Default)]
pub struct Widget(Option<Box<i32>>);
impl Widget {
    pub fn data_b(&self)->Option<&i32>{
        self.0.as_deref()
    }

}

fn i_need_ownership(data: Option<&i32>){
    let _:Option<i32> = data.map(ToOwned::to_owned);
    let _:Option<i32> = data.cloned();
    let _:Option<i32> = data.copied();
}

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

impl Solution {
    pub fn reorder_list(head: &mut Option<Box<ListNode>>) {
        let len = Self::get_list_len(head.as_ref());
        let (mut first, mut second) = Self::split_list(head.take(), len);
        let second_revert = Self::revert(second);
        *head = Self::merge(first, second_revert);
    }

    fn get_list_len(mut head: Option<&Box<ListNode>>) -> usize {
        let mut len = 0;
        while let Some(node) = head {
            head = node.next.as_ref();
            len += 1;
        }
        len
    }
    // 当List只有3个节点，那么需要从 head 跳 0 次，才能达到中间节点的上一个节点
    // 当List只有4个节点，那么需要从 head 跳 1 次，才能达到中间节点的上一个节点
    // 当List只有5个节点，那么需要从 head 跳 1 次，才能达到中间节点的上一个节点
    fn split_list(
        mut head: Option<Box<ListNode>>,
        len: usize,
    ) -> (Option<Box<ListNode>>, Option<Box<ListNode>>) {
        let skip_cnt: usize = (len - 2) / 2;
        let mut cur = head.as_mut().unwrap();
        for _ in 0..skip_cnt {
            cur = cur.next.as_mut().unwrap();
        }

        let next = cur.next.take();
        (head, next)
    }

    // 利用快慢指针找到中间点
    // find the mid node by using quick and slow pointer
    // fn find_mid(head: Option<Box<ListNode>>) -> & Option<Box<ListNode>>{
    //     let mut a = &head;
    //     let mut b = &head;

    //     while let Some(node) = b.as_ref() {
    //         if let Some(nx) = node.next.as_ref() {
    //             a = &(a.as_ref().unwrap().next);
    //             b = &(nx.next);
    //         } else {
    //             break;
    //         }
    //     }
    //    a
    // }

    fn revert(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut pre = None;
        while let Some(mut cur) = head {
            head = cur.next.take();
            cur.next = pre.take(); // 这里是关键技术点
            pre.replace(cur);
        }
        pre
    }

    fn merge(
        mut one: Option<Box<ListNode>>,
        mut two: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut one_cur = one.as_mut().unwrap();
        while let Some(mut two_cur) = two {
            two = two_cur.next.take();
            two_cur.next = one_cur.next.take();
            one_cur.next = Some(two_cur);

            if let Some(new_one) = one_cur.next.as_mut().unwrap().next.as_mut() {
                one_cur = new_one;
            } else {
                break;
            }
        }

        one
    }

    // fn merge(mut one: Option<Box<ListNode>>, mut two: Option<Box<ListNode>>) {
    //     while one.is_some() && two.is_some() {
    //         let mut one_cur = one.as_mut().unwrap();
    //         let mut two_cur = two.as_mut().unwrap();

    //         // one_cur.next may be None, but two_cur.next may be Some
    //         if one_cur.next.is_some(){
    //             two = two_cur.next.replace(one_cur.next.take().unwrap());
    //         }

    //         one = one_cur.next.replace(two_cur);
    //     }
    // }

    // fn revert(
    //     head: &mut Option<Box<ListNode>>,
    // ) -> (Option< & Box<ListNode>>, Option<&mut Box<ListNode>>) {
    //     if let Some(h) = head {
    //         if h.next.is_none() {
    //             return (head.as_ref(), head.as_mut());
    //         }
    //         // 断链
    //         // break list
    //         let mut nx = h.next.take();
    //         let (nx_head, mut nx_tail) = Self::revert(&mut nx);
    //         nx_tail.as_mut().unwrap().next = *head;
    //     }
    //     (None, None)
    // }
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
