fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn unique_occurrences(mut arr: Vec<i32>) -> bool {
        let mut counter = vec![0;2001];

        for v in arr{
            counter[v as usize + 1000] = counter[v as usize+1000] +1;
        }

        let mut set = std::collections::HashSet::new();
        for v in counter{

            if v >0{

                if set.contains(&v){
                    return false;
                }else{
                    set.insert(v);
                }
            }

        }

        true
    }

}