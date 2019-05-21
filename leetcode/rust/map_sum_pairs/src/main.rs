// Accepted
// @lc id=677 lang=rust
// problem: map_sum_pairs

use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;

type Link = Rc<RefCell<TrieNode>>;

struct TrieNode {
    count: i32,
    children: HashMap<char, Link>,
}

impl TrieNode {
    fn new() -> TrieNode {
        TrieNode {
            count: 0,
            children: HashMap::new(),
        }
    }
}
struct MapSum {
    trie: Link,
}

impl MapSum {
    fn new() -> Self {
        MapSum {
            trie: Rc::new(RefCell::new(TrieNode::new())),
        }
    }

    fn insert(&self, key: String, val: i32) {
        let mut root = self.trie.clone();
        for c in key.chars() {
            root.borrow_mut()
                .children
                .entry(c)
                .or_insert_with(|| Rc::new(RefCell::new(TrieNode::new())));
            root = root.clone().borrow().children.get(&c).unwrap().clone();
        }
        root.borrow_mut().count = val;
    }

    fn get_count(root: Link) -> i32 {
        let mut ret = root.borrow().count;
        for (_, v) in root.borrow().children.clone() {
            ret += MapSum::get_count(v)
        }
        ret
    }

    fn sum(&self, prefix: String) -> i32 {
        let mut root = self.trie.clone();
        for c in prefix.chars() {
            if root.borrow().children.contains_key(&c) {
                root = root.clone().borrow().children.get(&c).cloned().unwrap();
            } else {
                return 0;
            }
        }

        MapSum::get_count(root)
    }
}

fn main() {
    let m = MapSum::new();
    m.insert("apple".to_owned(), 3);
    m.insert("app".to_owned(), 2);
    m.insert("appa".to_owned(), 1);
    println!("{:?}", m.sum("ap".to_owned()));
}
