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

// into iterator
pub struct ListIntoIterator<T>(List<T>);

impl<T> std::iter::IntoIterator for List<T> {
    type Item = T;
    type IntoIter = ListIntoIterator<T>;

    fn into_iter(self) -> Self::IntoIter {
        ListIntoIterator(self)
    }
}

impl<T> Iterator for ListIntoIterator<T> {
    type Item = T;
    fn next(&mut self) -> Option<T> {
        self.0.pop()
    }
}

// iter
impl<T> List<T> {
    pub fn iter(&mut self) -> ListIter<T> {
        ListIter {
            next: self.head.as_ref()
        }
    }
}
pub struct ListIter<'a, T> {
    next: Option<&'a Box<Node<T>>>,
}

impl<'a, T> std::iter::Iterator for ListIter<'a, T> {
    type Item = &'a T;
    fn next(&mut self) -> Option<&'a T> {
        self.next.map( |node| {

        }})
    }
}
#[test]
fn test_push_and_pop() {
    let mut list = List::default();
    list.push(2);
    assert_eq!(list.pop().unwrap(), 2);
    assert_eq!(list.pop(), None);
}

#[test]
fn test_peek() {
    let mut list = List::default();
    list.push(2);
    list.peek_mut().map(|value| {
        *value = 3;
    });
    assert_eq!(list.peek().unwrap(), &3);
    list.pop();
    assert!(list.peek().is_none());

    let mut list2 = List::default();
    list2.push(1);
    list2.push(2);
    assert_eq!(list2.into_iter().collect::<Vec<_>>(), vec![2, 1]);
}

#[test]
fn test_into_iterator() {
    let mut list2 = List::default();
    list2.push(1);
    list2.push(2);
    assert_eq!(list2.into_iter().collect::<Vec<_>>(), vec![2, 1]);
}
