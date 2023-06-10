use std::{iter::Iterator, result::Result::Ok, vec::Vec};

pub struct Solution;

pub fn get_num_frequency(word: &String) -> i32 {
    let mut ret = 0;
    let mut min = 'z';
    for c in word.chars() {
        if c == min {
            ret = ret + 1;
        } else if c < min {
            min = c;
            ret = 1
        } else {
        }
    }
    ret
}

use std::cmp::{Ord, Ordering};
impl Solution {
    pub fn num_smaller_by_frequency(queries: Vec<String>, words: Vec<String>) -> Vec<i32> {
        let mut word_ret: Vec<i32> = words.iter().map(|word| get_num_frequency(word)).collect();
        word_ret.sort_unstable();

        queries
            .iter()
            .map(|query| {
                let query_frequency = get_num_frequency(query);
                match word_ret.binary_search_by(|&y| match y.cmp(&query_frequency) {
                    Ordering::Greater => Ordering::Greater,
                    Ordering::Less | Ordering::Equal => Ordering::Less,
                }) {
                    Ok(id) => word_ret.len() as i32 - id as i32 - 1,
                    Err(id) => word_ret.len() as i32 - id as i32,
                }
            })
            .collect()
    }
    // #[stable(feature = "partition_point", since = "1.52.0")]
    // pub fn num_smaller_by_frequency(queries: Vec<String>, words: Vec<String>) -> Vec<i32> {
    //     let mut word_ret: Vec<i32> = words.iter().map(|word| get_num_frequency(word)).collect();
    //     word_ret.sort_unstable();

    //     queries
    //         .iter()
    //         .map(|query| {
    //             let query_frequency = get_num_frequency(query);
    //             let id = word_ret.partition_point(|&x|x<=query_frequency);
    //             word_ret.len() as i32 - id as i32
                
    //         })
    //         .collect()
    // }
}

#[cfg(test)]
#[test]
fn it_works() {
    assert_eq!(
        Solution::num_smaller_by_frequency(vec!["cbd".to_string()], vec!["zaaaz".to_string()]),
        vec![1]
    );
    assert_eq!(
        Solution::num_smaller_by_frequency(
            vec!["bbb".to_string(), "cc".to_string()],
            vec![
                "a".to_string(),
                "aa".to_string(),
                "aaa".to_string(),
                "aaaa".to_string()
            ]
        ),
        vec![1, 2]
    );

    assert_eq!(
        Solution::num_smaller_by_frequency(
            vec!["a".to_string(), "cccc".to_string(), "ccccc".to_string()],
            vec![
                "a".to_string(),
                "aa".to_string(),
                "aaa".to_string(),
                "aaaa".to_string()
            ]
        ),
        vec![3, 0, 0]
    );
}
