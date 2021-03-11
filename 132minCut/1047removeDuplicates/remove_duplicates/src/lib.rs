/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-09 10:03:59
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-09 10:20:25
 * @FilePath: \1047removeDuplicates\remove_duplicates\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn remove_duplicates(s: String) -> String {
        // let s = s.into_bytes();
        let s = s.as_bytes(); //相比使用into_bytes(),内存从2.3M降低到2.2M
        let mut res = Vec::new();
        s.iter().for_each(|x|{
            if let Some(y) = res.last(){
                if y == x{
                    res.pop();
                }else{
                    res.push(*x);
                }
            }else{
                res.push(*x);
            }
        });

    //    unsafe { String::from_utf8_unchecked(res)}

       String::from_utf8_lossy(&res).to_string() //相比上面 从2.2M降低到2.1MB

    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::remove_duplicates(String::from("abbaca")),String::from("ca"));
        assert_eq!(Solution::remove_duplicates(String::from("aaa")),String::from("a"));
    }
}
