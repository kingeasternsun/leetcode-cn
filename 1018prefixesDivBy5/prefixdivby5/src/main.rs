fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn prefixes_div_by5(a: Vec<i32>) -> Vec<bool> {
        let mut res = Vec::new();

        let mut num = 0;

        for v in a{
            num = (num<<1)+v;
            num = num%5;
            if num ==0{
                res.push(true);
            }else{
                res.push(false)
            }
        }

        res

    }
}

#[cfg(test)]
mod tests{
    use super::*;
    #[test]
    fn one(){

        assert_eq!(Solution::prefixes_div_by5(vec![0, 1, 1, 1, 1, 1]),vec![true, false, false, false, true, false]);
        assert_eq!(Solution::prefixes_div_by5(vec![1, 1, 1, 0, 1]),vec![false, false, false, false, false]);
    }
}