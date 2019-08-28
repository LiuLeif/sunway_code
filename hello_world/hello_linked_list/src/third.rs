// 2019-08-27 13:24
#[derive(Default)]
pub struct List<T> {
    pub head: Link<T>,
}

type Link<T> = Option<Box<Node<T>>>;

pub struct Node<T> {
    pub elem: T,
    pub next: Link<T>,
}

impl<T> List<T> {
    pub fn push_front(&mut self, elem: T) {
        let new_node = Box::new(Node {
            elem,
            next: self.head.take(),
        });
        self.head = Some(new_node);
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn push_back(&mut self, elem: T) {
        if self.is_empty() {
            self.push_front(elem);
            return;
        }
        let new_node = Box::new(Node { elem, next: None });
        let mut curr = self.head.as_mut();
        while let Some(node) = curr.take() {
            if node.next.as_ref().is_none() {
                curr = Some(node);
                break;
            }
            curr = node.next.as_mut();
        }
        curr.unwrap().next = Some(new_node);
    }

    pub fn pop_front(&mut self) -> Option<T> {
        match self.head.take() {
            None => None,
            Some(node) => {
                self.head = node.next;
                Some(node.elem)
            }
        }
    }

    pub fn peek_front(&self) -> Option<&T> {
        self.head.as_ref().map(|node| &node.elem)
    }
}

#[test]
fn test_push_and_pop_front() {
    let mut list = List::default();
    list.push_front(2);
    assert_eq!(list.pop_front().unwrap(), 2);
    assert_eq!(list.pop_front(), None);
}

#[test]
fn test_push_back() {
    let mut list = List::default();
    list.push_back(1);
    list.push_back(2);
    list.push_back(3);
    assert_eq!(list.pop_front(), Some(1));
    assert_eq!(list.pop_front(), Some(2));
    assert_eq!(list.pop_front(), Some(3));
}
