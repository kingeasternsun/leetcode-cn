/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-26 16:51:27
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-26 17:24:11
 * @FilePath: \387firstUniqChar\first_uniq_char\src\lib.rs
 */

pub struct Solution;

impl Solution {
    /*
    通过记录每个字母出现的开始和结束位置，就可以快速查到 只出现一次的第一个位置
    */
    pub fn first_uniq_char1(s: String) -> i32 {
        let mut char_pos = vec![vec![-1, -1]; 26]; //每个字母第一次和最后一次出现位置

        let bs = s.into_bytes();

        for i in 0..bs.len(){
            let c = (bs[i] - 'a' as u8) as usize;
            if char_pos[c][0] == - 1{ // 第一次出现
                char_pos[c][0] = i as i32;
            }else{
                char_pos[c][1] = i as i32;
            }
        }

        let mut res = bs.len() as i32;
        for v in &char_pos{
            
            //只出现一次
            if  v[0]!=-1 && v[1]==-1 && v[0]<res {
                    res = v[0]
            }
        }


        if res == bs.len() as i32{
            return -1
        }

        return res
    }

    pub fn first_uniq_char(s: String) -> i32 {
        let mut char_pos = vec![vec![None, None]; 26]; //每个字母第一次和最后一次出现位置

        let bs = s.into_bytes();

        for i in 0..bs.len(){
            let c = (bs[i] - 'a' as u8) as usize;
            if char_pos[c][0].is_none(){ // 第一次出现
                char_pos[c][0] = Some(i);
            }else{
                char_pos[c][1] = Some(i);
            }
        }

        let mut res = bs.len() as i32;
        for v in &char_pos{
            
            //只出现一次
            if let Some(pos) = v[0]{
                if v[1].is_none() && (pos as i32) < res{
                    res = pos as i32;
                }
                
            }
        }


        if res == bs.len() as i32{
            return -1
        }

        return res
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::first_uniq_char(String::from("leetcode")),0);
        assert_eq!(Solution::first_uniq_char(String::from("loveleetcode")),2);
        assert_eq!(Solution::first_uniq_char(String::from("aabb")),-1);

        assert_eq!(Solution::first_uniq_char1(String::from("leetcode")),0);
        assert_eq!(Solution::first_uniq_char1(String::from("loveleetcode")),2);
        assert_eq!(Solution::first_uniq_char1(String::from("aabb")),-1);
    }
}
