fn main() {
    println!("Hello, world!");

    let word_list = vec![
        String::from("hot"),
        String::from("dot"),
        String::from("dog"),
        String::from("lot"),
        String::from("log"),
        String::from("cog"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hit"), String::from("cog"), word_list),
        5
    );
    let word_list2 = vec![
        String::from("hot"),
        String::from("dot"),
        String::from("dog"),
        String::from("lot"),
        String::from("log"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hit"), String::from("cog"), word_list2),
        0
    );

    let word_list3 = vec![
        String::from("hot"),
        String::from("dog"),
        String::from("cog"),
        String::from("pot"),
        String::from("dot"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hot"), String::from("dog"), word_list3),
        3
    );
}
pub struct Solution;


impl Solution {
    /*
    1. BFS
    2. 从 begin_word end_word分别轮流开始遍历 。
    */
    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        let mut exist = false;
        for v in &word_list {
            if *v == end_word {
                exist = true;

                break;
            }
        }

        if !exist {
            return 0;
        }

        let mut word = std::collections::HashSet::new();

        let mut queue = std::collections::VecDeque::new();
        queue.push_back(begin_word.clone());
        word.insert(begin_word.clone());

        let mut step = 1;

        while queue.len() > 0 {
            let num = queue.len();

            for i in 0..num {
                if let Some(node) = queue.pop_front() {
                    for x in &word_list {
                        //如果词典里的这个单词和当前节点的单词 差一个字符,并且之前没有访问过
                        if Self::one_char_diff(&node, x) && (!word.contains(x)) {

                            if *x == end_word{
                                return  step+1
                            }
                           
                            queue.push_back(x.clone());
                            word.insert(x.clone());

                        }
                    }
                }
            }

            step += 1;
        }

        0
    }

    //判断两个字符是否只差一个字符不一样
    fn one_char_diff(from: &String, to: &String) -> bool {
        if from.len() != to.len() {
            return false;
        }

        let mut cnt = 0;
        for i in 0..from.len() {
            if from.as_bytes()[i] != to.as_bytes()[i] {
                cnt += 1;
                if cnt == 2 {
                    return false;
                }
            }
        }

        return cnt == 1;
    }
}
