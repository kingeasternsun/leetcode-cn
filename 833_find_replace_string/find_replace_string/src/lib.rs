struct Solution;

#[derive(Debug, PartialEq, Eq)]
struct Value {
    id: usize,
    source: String,
    target: String,
}

#[derive(Debug, PartialEq, Eq)]
struct Item<'a> {
    id: usize,
    source: &'a str,
    target: &'a str,
}

impl Solution {
    // 0ms 2.3mb
    pub fn find_replace_string(
        s: String,
        indices: Vec<i32>,
        sources: Vec<String>,
        targets: Vec<String>,
    ) -> String {
        let bytes = s.as_bytes();
        let mut values = Vec::new();
        // 选取source能够匹配的替换
        for (i, &id) in indices.iter().enumerate() {
            let start_id = id as usize;
            if start_id + sources[i].len() <= bytes.len()
                && bytes[start_id..start_id + sources[i].len()] == sources[i].as_bytes()[..]
            {
                values.push(Value {
                    id: start_id,
                    source: sources[i].clone(),
                    target: targets[i].clone(),
                })
            }
        }

        if values.len() == 0 {
            return s;
        }

        values.sort_unstable_by_key(|k| k.id);

        // 如果存在两个替换的索引冲突了
        if values
            .windows(2)
            .any(|x| x[0].id + x[0].source.len() - 1 >= x[1].id)
        {
            return s;
        }

        // 根据替换规则构造新的字符串
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

    pub fn find_replace_string1(
        s: String,
        indices: Vec<i32>,
        sources: Vec<String>,
        targets: Vec<String>,
    ) -> String {
        let bytes = s.as_bytes();
        let mut values = Vec::new();
        // 选取source能够匹配的替换
        for (i, &id) in indices.iter().enumerate() {
            let start_id = id as usize;
            if start_id + sources[i].len() <= bytes.len()
                && bytes[start_id..start_id + sources[i].len()] == sources[i].as_bytes()[..]
            {
                values.push(Item {
                    id: start_id,
                    source: &sources[i],
                    target: &targets[i],
                });
            }
        }

        if values.len() == 0 {
            return s;
        }

        values.sort_unstable_by_key(|k| k.id);

        // 如果存在两个替换的索引冲突了
        if values
            .windows(2)
            .any(|x| x[0].id + x[0].source.len() - 1 >= x[1].id)
        {
            return s;
        }

        // 根据替换规则构造新的字符串
        let mut ret = Vec::new();
        ret.extend_from_slice(&bytes[0..values[0].id]);

        values.windows(2).for_each(|x| {
            ret.extend_from_slice(x[0].target.as_bytes());
            ret.extend_from_slice(&bytes[x[0].id + x[0].source.len()..x[1].id]);
        });

        let last_val = values.last().unwrap();
        ret.extend_from_slice(last_val.target.as_bytes());
        ret.extend_from_slice(&bytes[last_val.id + last_val.source.len()..bytes.len()]);

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

        assert_eq!(
            Solution::find_replace_string(
                "ynqoknliec".to_string(),
                vec![1, 8, 3],
                vec!["nq".to_string(), "ec".to_string(), "ok".to_string()],
                vec!["dh".to_string(), "tc".to_string(), "pc".to_string()]
            ),
            "ydhpcnlitc".to_string()
        );
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::find_replace_string1(
                "abcd".to_string(),
                vec![0, 2],
                vec!["a".to_string(), "cd".to_string()],
                vec!["eee".to_string(), "ffff".to_string()]
            ),
            "eeebffff".to_string()
        );
        assert_eq!(
            Solution::find_replace_string1(
                "ynqoknliec".to_string(),
                vec![1, 8, 3],
                vec!["nq".to_string(), "ec".to_string(), "ok".to_string()],
                vec!["dh".to_string(), "tc".to_string(), "pc".to_string()]
            ),
            "ydhpcnlitc".to_string()
        );
    }
}
