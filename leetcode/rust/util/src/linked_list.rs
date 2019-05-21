// 2018-12-29 22:05
#[derive(PartialEq, Eq, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub struct LinkedList {}
impl LinkedList {
    pub fn new(nums: Vec<i32>) -> (Option<Box<ListNode>>) {
        let mut fake_root = Box::new(ListNode::new(0));
        let mut head = &mut fake_root;
        for i in nums {
            head.next = Some(Box::new(ListNode::new(i)));
            head = (head.next).as_mut().unwrap();
        }
        return fake_root.next;
    }
}

#[macro_export]
macro_rules! list {
    ($($v:expr),*) => {
        LinkedList::new(vec![$($v), *])
    }
}
