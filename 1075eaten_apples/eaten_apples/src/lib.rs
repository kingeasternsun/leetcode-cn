mod lt1705 {
    struct Solution;
    // 36ms 2.8MB
    impl Solution {
        // 贪心法，1. 有苹果就吃 2.吃保质期最近的那个苹果
        pub fn eaten_apples(apples: Vec<i32>, days: Vec<i32>) -> i32 {
            use std::cmp::Reverse;
            use std::collections;
            // 记录在第key天腐烂的苹果有多少个
            let mut day2number = collections::HashMap::new();
            // 根据腐烂的日期排序
            let mut day_heap = collections::BinaryHeap::new();

            let mut eat_num = 0;

            let n = apples.len();
            for i in 0..n {
                if apples[i] > 0 {
                    let day = i as i32 + days[i];
                    if !day2number.contains_key(&day) {
                        day_heap.push(Reverse(day));
                    }

                    let entry = day2number.entry(day).or_insert(0);
                    *entry += apples[i];
                }

                while let Some(Reverse(day)) = day_heap.pop() {
                    if day <= i as i32 {
                        // 已经过期了
                        day2number.remove(&day);
                        continue;
                    }

                    let mut number = day2number.get_mut(&day).unwrap();
                    eat_num += 1;
                    *number -= 1;

                    // 这个腐烂日期的苹果吃完了
                    if *number == 0 {
                        day2number.remove(&day);
                    } else {
                        // 记得重新放回去
                        day_heap.push(Reverse(day));
                    }

                    // 一天只能吃一个
                    break;
                }
            }

            // 继续吃
            let mut i = n;
            loop {
                while let Some(Reverse(day)) = day_heap.pop() {
                    if day <= i as i32 {
                        // 已经过期了
                        day2number.remove(&day);
                        continue;
                    }

                    let number = day2number.get_mut(&day).unwrap();
                    eat_num += 1;
                    *number -= 1;

                    // 这个腐烂日期的苹果吃完了
                    if *number == 0 {
                        day2number.remove(&day);
                    } else {
                        // 记得重新放回去
                        day_heap.push(Reverse(day));
                    }

                    // 一天只能吃一个
                    break;
                }

                if day_heap.is_empty() {
                    break;
                }

                i += 1;
            }

            eat_num
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_name() {
            assert_eq!(
                Solution::eaten_apples(vec![1, 2, 3, 5, 2], vec![3, 2, 1, 4, 2]),
                7
            );

            assert_eq!(
                Solution::eaten_apples(vec![3, 0, 0, 0, 0, 2], vec![3, 0, 0, 0, 0, 2]),
                5
            );

            assert_eq!(Solution::eaten_apples(vec![2, 1, 10], vec![2, 10, 1]), 4);
        }
    }
}

mod lt1705b {
    struct Solution;

    // 20ms 2.5MB
    impl Solution {
        // 贪心法，1. 有苹果就吃 2.吃保质期最近的那个苹果
        // 相比第一种的实现方法，
        //    1. 不会将相同截止日期的苹果放在一起，而是单独处理
        //    2. 直接把截止日期和对应的苹果数量都放入到最小堆中，
        //  这样的好处在于代码实现更加简单，每次取截止日期最小的吃，如果吃完还有剩余，再放回去
        //    3. 过去n天后，后面直接根据截止日期和截止日期的苹果树快速计算可以吃的苹果数目

        pub fn eaten_apples(apples: Vec<i32>, days: Vec<i32>) -> i32 {
            use std::cmp::Ordering;
            use std::collections;
            #[derive(Eq)]
            struct Item(i32, i32);
            impl PartialEq for Item {
                fn eq(&self, other: &Self) -> bool {
                    self.0 == other.0
                }
            }

            impl Ord for Item {
                fn cmp(&self, other: &Self) -> Ordering {
                    other.0.cmp(&self.0)
                }
            }

            impl PartialOrd for Item {
                fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
                    Some(self.cmp(other))
                }
            }

            // 根据腐烂的日期排序
            let mut day_heap = collections::BinaryHeap::new();

            let mut eat_num = 0;

            let n = apples.len();
            for i in 0..n {
                // 根据腐烂的日期加入堆中
                if apples[i] > 0 {
                    day_heap.push(Item(i as i32 + days[i], apples[i]))
                }

                while let Some(mut item) = day_heap.pop() {
                    if item.0 <= i as i32 {
                        continue;
                    }
                    // 吃一个苹果
                    eat_num += 1;
                    item.1 -= 1;
                    // 重新放回去
                    if item.1 > 0 {
                        day_heap.push(item);
                    }
                    break;
                }
            }

            // 继续吃
            let mut cur_day = n as i32;
            while let Some(item) = day_heap.pop() {
                if item.0 <= cur_day {
                    continue;
                }
                // 计算当前距离截止日期还有多少天  item.0 - cur_day
                let off_day = item.0 - cur_day;
                // 天数比剩余苹果还多 ，这堆苹果一天一个都可以吃完
                if off_day > item.1 {
                    eat_num += item.1;
                    cur_day += item.1; // ⚠️ 注意，当前算法中相同截止日期可能会有多项，所以这里加 item.1 而不是直接跳到截止日期
                } else {
                    eat_num += off_day;
                    cur_day += off_day;
                }
            }

            eat_num
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_name() {
            assert_eq!(
                Solution::eaten_apples(vec![1, 2, 3, 5, 2], vec![3, 2, 1, 4, 2]),
                7
            );

            assert_eq!(
                Solution::eaten_apples(vec![3, 0, 0, 0, 0, 2], vec![3, 0, 0, 0, 0, 2]),
                5
            );

            // assert_eq!(Solution::eaten_apples(vec![2, 1, 10], vec![2, 10, 1]), 4);
            assert_eq!(Solution::eaten_apples(vec![1], vec![2]), 1);
        }
    }
}

mod lt1705c {
    struct Solution;

    // 8ms 2.5MB
    impl Solution {
        // 贪心法，1. 有苹果就吃 2.吃保质期最近的那个苹果
        // 相比第二种的实现方法，第二种方式是取出来吃，如果吃完还有剩余再放回去
        //  这个方法是直接在最小堆中吃，如果吃完没有了，就移除
        pub fn eaten_apples(apples: Vec<i32>, days: Vec<i32>) -> i32 {
            use std::cmp::Ordering;
            use std::collections;
            #[derive(Eq)]
            struct Item(i32, i32);
            impl PartialEq for Item {
                fn eq(&self, other: &Self) -> bool {
                    self.0 == other.0
                }
            }

            impl Ord for Item {
                fn cmp(&self, other: &Self) -> Ordering {
                    other.0.cmp(&self.0)
                }
            }

            impl PartialOrd for Item {
                fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
                    Some(self.cmp(other))
                }
            }

            // 根据腐烂的日期排序
            let mut day_heap = collections::BinaryHeap::new();

            let mut eat_num = 0;

            let n = apples.len();
            for i in 0..n {
                // 根据腐烂的日期加入堆中
                if apples[i] > 0 {
                    day_heap.push(Item(i as i32 + days[i], apples[i]))
                }

                while let Some(item) = day_heap.peek() {
                    if item.0 <= i as i32 {
                        day_heap.pop();
                        continue;
                    }
                    // 吃一个苹果
                    eat_num += 1;
                    if item.1 == 1 {
                        day_heap.pop();
                    } else {
                        day_heap.peek_mut().unwrap().1 -= 1;
                    }

                    break;
                }
            }

            // 继续吃
            let mut cur_day = n as i32;
            while let Some(item) = day_heap.pop() {
                if item.0 <= cur_day {
                    continue;
                }
                // 计算当前距离截止日期还有多少天  item.0 - cur_day
                let off_day = item.0 - cur_day;
                // 天数比剩余苹果还多 ，这堆苹果一天一个都可以吃完
                if off_day > item.1 {
                    eat_num += item.1;
                    cur_day += item.1; // ⚠️ 注意，当前算法中相同截止日期可能会有多项，所以这里加 item.1 而不是直接跳到截止日期
                } else {
                    eat_num += off_day;
                    cur_day += off_day;
                }
            }

            eat_num
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_name() {
            assert_eq!(
                Solution::eaten_apples(vec![1, 2, 3, 5, 2], vec![3, 2, 1, 4, 2]),
                7
            );

            assert_eq!(
                Solution::eaten_apples(vec![3, 0, 0, 0, 0, 2], vec![3, 0, 0, 0, 0, 2]),
                5
            );

            // assert_eq!(Solution::eaten_apples(vec![2, 1, 10], vec![2, 10, 1]), 4);
            assert_eq!(Solution::eaten_apples(vec![1], vec![2]), 1);
        }
    }
}
