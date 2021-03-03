/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-03 16:38:41
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 17:47:27
 * @FilePath: \1178findNumOfValidWords\find_num_of_valid_words\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn find_num_of_valid_words(words: Vec<String>, puzzles: Vec<String>) -> Vec<i32> {
        let mut word_map = std::collections::HashMap::new();
        let mut res = Vec::new();

        //生成 word_map
        for word in words {
            let mut word_status: u32 = 0;
            word.into_bytes().iter().for_each(|x| {
                word_status |= 1 << (x - 'a' as u8);
            });

            let entry = word_map.entry(word_status).or_insert(0);
            *entry += 1;
        }

        for puzzle in puzzles {
            let mut num = 0;
            let p = puzzle.into_bytes();
            let bit_status =  1 << (p[0]-'a' as u8); //必须包含puzzle的第一个字母
            Self::find_match(&p[1..], bit_status, &word_map, &mut num);
            res.push(num);
        }

        res
    }

    pub fn find_match(
        s: &[u8],
       mut  bit_status: u32,
        word_map: &std::collections::HashMap<u32, i32>,
        res: &mut i32,
    ) {
        if s.len() == 0 {
            if let Some(v) = word_map.get(&bit_status) {
                *res = *res + *v;
            }
            return;
        }

        //不取第一个字母
        Self::find_match(&s[1..], bit_status, word_map, res);

        //取第一个字母
        // bit_status = bit_status | (1 << (s[0] - 'a' as u8));
        Self::find_match(&s[1..], bit_status | (1 << (s[0] - 'a' as u8)), word_map, res);
        //恢复bit_status
        // bit_status = bit_status ^ (1 << (s[0] - 'a' as u8));
    }
}


/*
10110

10101 -> 10100

10011 ->



*/
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_num_of_valid_words(
                vec![
                    "aaaa".to_string(),
                    "asas".to_string(),
                    "able".to_string(),
                    "ability".to_string(),
                    "actt".to_string(),
                    "actor".to_string(),
                    "access".to_string()
                ],
                vec![
                    "aboveyz".to_string(),
                    "abrodyz".to_string(),
                    "abslute".to_string(),
                    "absoryz".to_string(),
                    "actresz".to_string(),
                    "gaswxyz".to_string()
                ]
            ),
            vec![1, 1, 3, 2, 4, 0]
        );
    }
}
