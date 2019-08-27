// 2019-08-27 13:24
pub struct List<T> {
    pub head: Link<T>,
}

type Link<T> = Option<Box<Node<T>>>;

pub struct Node<T> {
    pub elem: T,
    pub next: Link<T>,
}

impl<T> List<T> {
    pub fn new() -> Self {
        List { head: None }
    }

    pub fn push(&mut self, elem: T) {
        let new_node = Box::new(Node {
            elem,
            next: self.head.take(),
        });
        self.head = Some(new_node);
    }

    pub fn pop(&mut self) -> Option<T> {
        match self.head.take() {
            None => None,
            Some(node) => {
                self.head = node.next;
                Some(node.elem)
            }
        }
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_ref().map(|node| &node.elem)
    }

    pub fn peek_mut(&mut self) -> Option<&mut T> {
        self.head.as_mut().map(|node| &mut node.elem)
    }
}

#[test]
fn test_push_and_pop() {
    let mut list = List::new();
    list.push(2);
    assert_eq!(list.pop().unwrap(), 2);
    assert_eq!(list.pop(), None);
}

#[test]
fn test_peek() {
    let mut list = List::new();
    list.push(2);
    list.peek_mut().map(|value| {
        *value = 3;
    });
    assert_eq!(list.peek().unwrap(), &3);
    list.pop();
    assert!(list.peek().is_none());
}
