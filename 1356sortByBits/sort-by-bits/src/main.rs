fn main() {
    println!("Hello, world!");
}


pub struct Solution;
impl Solution {
    pub fn sort_by_bits(mut arr:  Vec<i32>) -> Vec<i32> {

        arr.sort_by_cached_key(|x|(x.count_ones(),*x));
        arr
    }
}