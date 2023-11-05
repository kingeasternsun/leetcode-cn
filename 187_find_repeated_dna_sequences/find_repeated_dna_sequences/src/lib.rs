pub struct Solution;
impl Solution {
    pub fn find_repeated_dna_sequences_a(s: String) -> Vec<String> {
        let mut sequence_map = std::collections::HashMap::new();
        s.as_bytes().windows(10).for_each(|x| {
            let v = sequence_map.entry(x).and_modify(|e| *e += 1).or_insert(1);
        });

        let mut ret = vec![];
        for (k, v) in sequence_map {
            if v >= 2 {
                ret.push(String::from_utf8(k.to_vec()).unwrap());
            }
        }
        ret
    }

    // 8ms 6.92MB
    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        let mut sequence_map = std::collections::HashMap::new();
        s.as_bytes().windows(10).for_each(|x| {
            let v = sequence_map.entry(x).and_modify(|e| *e += 1).or_insert(1);
        });

        sequence_map
            .into_iter()
            .filter(|x| x.1 >= 2)
            .map(|k| String::from_utf8(k.0.to_vec()).unwrap())
            .collect()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
