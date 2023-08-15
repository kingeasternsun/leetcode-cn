struct Solution;

#[derive(Debug, PartialEq, Eq)]
struct Value {
    id: usize,
    source: String,
    target: String,
}

impl Solution {
    pub fn find_replace_string(
        s: String,
        indices: Vec<i32>,
        sources: Vec<String>,
        targets: Vec<String>,
    ) -> String {
        let bytes = s.as_bytes();
        let mut values = Vec::new();
        for (i, &id) in indices.iter().enumerate() {
            if id as usize + sources[i].len() <= bytes.len()
                && bytes[id as usize..id as usize + sources[i].len()] == sources[i].as_bytes()[..]
            {
                values.push(Value {
                    id: indices[i] as usize,
                    source: sources[i].clone(),
                    target: targets[i].clone(),
                })
            }
        }

        values.sort_unstable_by_key(|k| k.id);

        if values
            .windows(2)
            .any(|x| x[0].id + x[0].source.len() - 1 >= x[1].id)
        {
            return s;
        }

        let mut ret = Vec::new();
        ret.extend_from_slice(&bytes[0..values[0].id]);

        values.windows(2).for_each(|x| {
            ret.extend_from_slice(x[0].target.as_bytes());
            ret.extend_from_slice(&bytes[x[0].id + x[0].source.len()..x[1].id]);
        });

        ret.extend_from_slice(values[values.len() - 1].target.as_bytes());
        ret.extend_from_slice(
            &bytes
                [values[values.len() - 1].id + values[values.len() - 1].source.len()..bytes.len()],
        );

        unsafe { String::from_utf8_unchecked(ret) }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_replace_string(
                "abcd".to_string(),
                vec![0, 2],
                vec!["a".to_string(), "cd".to_string()],
                vec!["eee".to_string(), "ffff".to_string()]
            ),
            "eeebffff".to_string()
        );
    }
}
