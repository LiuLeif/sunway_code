// Accepted
// @lc id=208 lang=rust
// problem: implement
type Link = Option<Box<Trie>>;

struct Trie {
    is_leaf: bool,
    children: [Link; 26],
}

impl Trie {
    fn new() -> Self {
        Trie {
            is_leaf: false,
            children: Default::default(),
        }
    }

    fn insert(&mut self, word: String) {
        let mut root = self;
        for b in word.as_bytes() {
            let index = (b - b'a') as usize;
            if root.children[index].is_none() {
                root.children[index] = Some(Box::new(Trie::new()));
            };
            root = root.children[index].as_mut().unwrap();
        }
        root.is_leaf = true;
    }

    fn search(&self, word: String) -> bool {
        let mut root = self;
        for b in word.as_bytes() {
            let index = (b - b'a') as usize;
            if root.children[index].is_none() {
                return false;
            }
            root = root.children[index].as_ref().unwrap();
        }
        root.is_leaf
    }

    fn starts_with(&self, prefix: String) -> bool {
        let mut root = self;
        for b in prefix.as_bytes() {
            let index = (b - b'a') as usize;
            if root.children[index].is_none() {
                return false;
            }
            root = root.children[index].as_ref().unwrap();
        }
        return true;
    }
}

fn main() {
    let mut trie = Trie::new();
    trie.insert("hello".to_owned());
    println!("{:?}", trie.search("hello".to_owned()));
    println!("{:?}", trie.search("hell".to_owned()));
    println!("{:?}", trie.starts_with("hello".to_owned()));
}
