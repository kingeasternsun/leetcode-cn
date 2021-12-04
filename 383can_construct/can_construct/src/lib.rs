struct Solution;
impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut m = vec![0u32; 26];
        magazine
            .as_bytes()
            .iter()
            .for_each(|x| m[(x - b'a') as usize] = m[(x - b'a') as usize] + 1);

        for x in ransom_note.as_bytes().iter() {
            if m[(x - b'a') as usize] == 0 {
                return false;
            }
            m[(x - b'a') as usize] = m[(x - b'a') as usize] - 1;
        }

        true
    }
}

mod leet383 {
    struct Solution;
    impl Solution {
        pub fn can_construct(ransom_note: String, magazine: String) -> bool {
            let mut m = vec![0u32; 26];
            magazine
                .as_bytes()
                .iter()
                .for_each(|x| m[(x - b'a') as usize] = m[(x - b'a') as usize] + 1);

            ransom_note
                .as_bytes()
                .iter()
                .try_for_each(|x| {
                    if m[(x - b'a') as usize] == 0 {
                        return None;
                    }
                    m[(x - b'a') as usize] = m[(x - b'a') as usize] - 1;
                    Some(())
                })
                .is_some()
        }
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::can_construct("a".to_string(), "b".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aa".to_string(), "ab".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aa".to_string(), "aab".to_string()),
            true
        );
        assert_eq!(
            Solution::can_construct("aab".to_string(), "aab".to_string()),
            true
        );
        assert_eq!(
            Solution::can_construct("aab".to_string(), "aa".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aab".to_string(), "ab".to_string()),
            false
        );
    }
}
