/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 10:43:47
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 11:34:39
 * @FilePath: \13romanToInt\roman-to-int\src\lib.rs
 */

pub struct Solution;
impl Solution {
    pub fn roman_to_int1(s: String) -> i32 {

        let map = |x :char|{
            match x{
                'I' => 1,
                'V' => 5,
                'X' => 10,
                'L' => 50,
                'C' => 100,
                'D' => 500,
                'M' => 1000,
                _ => 0,

            }
        };

        let mut pre = 0;
        let mut res = 0;

        s.chars().rev().for_each(|x|{
            let cur = map(x);
            if cur<pre{
                res -= cur;
            }else{
                res += cur;
            }

            pre = cur;

        });
    
        res

    }

    pub fn roman_to_int(s: String) -> i32 {

        let map :std::collections::HashMap<_,_> = String::from("IVXLCDM").chars().zip(vec![1,5,10,50,100,500,1000].into_iter()).collect();

        let mut pre = &0;
        let mut res = 0;

        s.chars().rev().for_each(|x|{
            let cur = match map.get(&x){
                Some(a)=>a,
                None=>&0,
            };
            if pre>cur{
                res -= *cur;
            }else{
                res += *cur;
            }

            pre = cur;

        });
    
        res

    }
}


#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works1() {
        assert_eq!(Solution::roman_to_int1(String::from("III")),3);
        assert_eq!(Solution::roman_to_int1(String::from("IV")),4);
        assert_eq!(Solution::roman_to_int1(String::from("IX")),9);
        assert_eq!(Solution::roman_to_int1(String::from("LVIII")),58);
        assert_eq!(Solution::roman_to_int1(String::from("MCMXCIV")),1994);
    }
    fn it_works() {
        assert_eq!(Solution::roman_to_int(String::from("III")),3);
        assert_eq!(Solution::roman_to_int(String::from("IV")),4);
        assert_eq!(Solution::roman_to_int(String::from("IX")),9);
        assert_eq!(Solution::roman_to_int(String::from("LVIII")),58);
        assert_eq!(Solution::roman_to_int(String::from("MCMXCIV")),1994);
    }
}
