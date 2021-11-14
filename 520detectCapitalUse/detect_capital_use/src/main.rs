fn main() {
    println!("Hello, world!");
}

pub fn detect_capital_use(word: String) -> bool {

    if word.len() ==0{
        return true
    }
    if word.len() ==1{
        return true
    }

    if word.chars().all(|x|x.is_lowercase()){
        return true 
    }

    if word.chars().all(|x|x.is_uppercase()){
        return true 
    }

    if word.chars().enumerate().all(|(id,x)|if id == 0{x.is_uppercase()} else{x.is_lowercase()}){
        return true 
    }
    

    false

}
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_name() {
        assert!(detect_capital_use("USA".to_string()));
        assert!(detect_capital_use("abc".to_string()));
        assert!(detect_capital_use("Abc".to_string()));
        assert!(!detect_capital_use("bAc".to_string()));
        assert!(!detect_capital_use("FlaG".to_string()));
    }
}