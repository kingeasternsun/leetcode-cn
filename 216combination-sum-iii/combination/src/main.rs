fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    pub fn combination_sum3(k: i32, n: i32) -> Vec<Vec<i32>> {

        let mut res = Vec::new();
        let mut vec = Vec::new();
        Self::combination(k, n, 0, & mut vec, & mut res);
        res

    }

    // k 个 大于 max_v 且小于等于9 的数组成 n；
    pub fn combination(k:i32,n:i32,max_v:i32, vec: & mut Vec<i32>,res: & mut Vec<Vec<i32>> ){


        if k == 1{

            if n > max_v && n<= 9{
                vec.push(n);
                res.push(vec.clone());
                vec.pop();
            }

            return 

        }

        // 新添加的数字肯定比 maxv 要大，如果 k*max_v >=n ，那么后续得到的和 肯定 大于n 
        if k*max_v >= n {
            return 
        }


        // 如果 n/k >=9 ，那么肯定至少有一个 是 大于 9 
        if n/k >=9{
            return 
        }

        for i in max_v+1..n/k+1{
            // 选择 i 
            vec.push(i);
            Self::combination(k-1, n-i, i, vec, res);
            vec.pop();
        }

    }
}