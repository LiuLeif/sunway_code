// Accepted
// @lc id=705 lang=rust
// problem: design_hash_set
use std::cell::RefCell;
struct Solution {}
struct MyHashSet {
    bucket: Vec<RefCell<Vec<i32>>>,
}

impl MyHashSet {
    fn hash(x: i32) -> usize {
        ((x as u32 * 2654435761) >> (32 - 5)) as usize
    }

    fn new() -> Self {
        MyHashSet {
            bucket: (vec![RefCell::new(vec![]); 1 << 5]),
        }
    }

    fn add(&self, key: i32) {
        if self.contains(key) {
            return;
        }
        let index = MyHashSet::hash(key);
        self.bucket[index].borrow_mut().push(key);
    }

    fn remove(&self, key: i32) {
        let index = MyHashSet::hash(key);
        let pos = self.bucket[index].borrow().iter().position(|&x| x == key);
        if pos.is_none() {
            return;
        }
        self.bucket[index].borrow_mut().swap_remove(pos.unwrap());
    }

    fn contains(&self, key: i32) -> bool {
        let index = MyHashSet::hash(key);
        self.bucket[index].borrow().iter().any(|&x| x == key)
    }
}

fn main() {
    let set = MyHashSet::new();
    set.add(1);
    set.remove(1);
    println!("{:?}", set.contains(1));
}
