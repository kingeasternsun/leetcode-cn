/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 14:03:33
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-05 16:07:41
 * @FilePath: \1208equalSubstring\equal_substring\src\lib.rs
 */
pub struct Solution;
impl Solution {

/*
执行用时：
0 ms, 在所有 Rust 提交中击败了100.00%的用户
内存消耗：
2.2 MB, 在所有 Rust 提交中击败了100.00%的用户
*/
    pub fn equal_substring(s: String, t: String, max_cost: i32) -> i32 {
        let s: Vec<u8> = s
            .chars()
            .zip(t.chars())
            .map(|(x, y)| {
                if x > y {
                    x as u8 - y as u8
                } else {
                    y as u8 - x as u8
                }
            })
            .collect();

        let mut start = 0;
        let mut end = 0;
        let mut tmp_cost = 0;
        let mut res = 0;

        while start <= end && end < s.len() {
            if tmp_cost <= max_cost {
                if (end - start) as i32 > res {
                    res = (end - start) as i32;
                }

                tmp_cost += s[end] as i32;
                end += 1;
                continue;
            }

            //消耗超过了阈值，就需要把start去掉
            tmp_cost -= s[start] as i32;
            start += 1;
        }

        if tmp_cost <= max_cost {
            if (end - start) as i32 > res {
                res = (end - start) as i32;
            }
        }

        res
    }

/*
执行用时：
0 ms, 在所有 Rust 提交中击败了100.00%的用户
内存消耗：
2.5 MB, 在所有 Rust 提交中击败了100.00%的用户
*/
    pub fn equal_substring2(s: String, t: String, max_cost: i32) -> i32 {
        let s: Vec<i32> = s
            .chars()
            .zip(t.chars())
            .map(|(x, y)| (x as i32 - y as i32).abs())
            .collect();

        let mut start = 0;
        let mut end = 0;
        let mut tmp_cost = 0;
        let mut res = 0;

        while start <= end && end < s.len() {
            if tmp_cost <= max_cost {
                if (end - start) > res {
                    res = end - start;
                }

                tmp_cost += s[end];
                end += 1;
                continue;
            }

            //消耗超过了阈值，就需要把start去掉
            tmp_cost -= s[start];
            start += 1;
        }

        if tmp_cost <= max_cost {
            if (end - start) > res {
                res = end - start;
            }
        }

        res as i32
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::equal_substring(String::from("abcd"), String::from("bcdf"), 3),
            3
        );
        assert_eq!(
            Solution::equal_substring(String::from("abcd"), String::from("cdef"), 3),
            1
        );
        assert_eq!(
            Solution::equal_substring(String::from("abcd"), String::from("acde"), 0),
            1
        );
        assert_eq!(
            Solution::equal_substring(
                String::from("krpgjbjjznpzdfy"),
                String::from("nxargkbydxmsgby"),
                14
            ),
            4
        );
    }
}
