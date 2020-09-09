fn main() {
    println!("Hello, world!");
    let mut tmp = Vec::new();
    let mut result = Vec::new();
    tmp.push(3i32);
    result.push(tmp.clone());
    tmp.push(4i32);
    result.push(tmp);
}

pub struct Solution {}
impl Solution {
    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        let mut result = Vec::new();

        if k > n {
            return result;
        }

        let mut tmp = Vec::new();
        Self::dfs(&mut tmp, &mut result, n, k);
        result
    }

    fn dfs(tmp: &mut Vec<i32>, result: &mut Vec<Vec<i32>>, n: i32, k: i32) {
        if tmp.len() == k as usize {
            result.push(tmp.clone());
            return;
        }

        let mut beg = 1;
        //提前返回
        if let Some(last) = tmp.last() {
            if tmp.len() as i32 + n - last < k {
                return;
            }

            beg = *last + 1;
        }

        // tmp.len() + n+1 - i >=k
        // tmp.len() + n + 1 -k >= i

        for i in beg..(tmp.len() as i32 + n + 2 -k){
        // for i in beg..n + 1 {
            tmp.push(i);
            Self::dfs(tmp, result, n, k);
            tmp.pop();
        }
    }
}
