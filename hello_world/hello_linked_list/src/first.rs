// 2019-08-27 13:24
pub struct List {
    pub head: Link,
}

type Link = Option<Box<Node>>;

pub struct Node {
    pub elem: i32,
    pub next: Link,
}

impl List {
    pub fn new() -> Self {
        List { head: None }
    }

    pub fn push(&mut self, elem: i32) {
        let new_node = Box::new(Node {
            elem,
            next: self.head.take(),
        });
        self.head = Some(new_node);
    }

    pub fn pop(&mut self) -> Option<i32> {
        match self.head.take() {
            None => None,
            Some(node) => {
                self.head = node.next;
                Some(node.elem)
            }
        }
    }
}

#[test]
fn test_push_and_pop() {
    let mut list = List::new();
    list.push(2);
    assert_eq!(list.pop().unwrap(), 2);
    assert_eq!(list.pop(), None);
}
