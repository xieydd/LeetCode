package list

/*
Doubly Linked List(Recurrent, can start from every node of the list, not only the first one)
*/

type Element struct {
	next, prev *Element
	list       *List // The element belong to which list
	Value      interface{}
}

type List struct {
	length int
	// root as the sentinel node, avoid when delete or insert node for the list need consider
	// the boundiary conditions
	root   Element
}

// return the prev list element or nil
func (e *Element) Prev() *Element {
	if p := e.prev; p != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Return the next list element or nil
func (e *Element) Next() *Element {
	if n := e.next; n != nil && n != &e.list.root { // because when init the list, set the list.root.next and prev = &list.root, just to avoid the Init list
		return n
	}

	return nil
}

// Create a list instance
func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.length = 0
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

// lazyInit

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// Return the list length
func (l *List) Len() int { return l.length }

// Return the List next first element or nil if list is empty
func (l *List) Head() *Element {
	if l.length == 0 {
		return nil
	}

	return l.root.next
}

// Return the List last element or nil if list is empty
func (l *List) Back() *Element {
	if l.length == 0 {
		return nil
	}

	return l.root.prev
}

// Insert a Element in element at, focus element belong to a list, so the list len will be added
func (l *List) insert(e, at *Element) *Element {
	temp := at.next
	at.next = e
	e.prev = at
	e.next = temp
	temp.prev = e
	l.length++
	e.list = l
	return e
}

// Notice insert a value, will bind to a element, so the element will bind to the list
func (l *List) insertValue(v interface{}, at *Element) *Element{
	return l.insert(&Element{Value: v}, at)
}

func (l *List) remove(e *Element) *Element {
	prev := e.prev
	next := e.next

	prev.next = next
	next.prev = prev
	e.prev = nil // avoid memory leak
	e.next = nil
	e.list = nil
	l.length--
	return e
}

// move e to next of at(first to move e). move meanings they belong to one list
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	temp := at.next
	at.next = e
	e.prev = at
	e.next = temp
	temp.prev = e //just to make sure, don`t know
	return e
}

// remove a element , if element`s list is the list of l, will only return the value of element, don`t move
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}

	return e.Value
}

// PushFront insert a new element with value in front of list l(notice the lazy init, make sure the list is not null)
func (l *List) PushFront(value interface{}) *Element {
	l.lazyInit()
	return l.insert(&Element{Value: value}, &l.root)
}

// PushPrev insert a element with value in prev of  list l
func (l *List) PushPrev(value interface{}) *Element {
	l.lazyInit()
	return l.insert(&Element{Value: value}, l.root.prev)
}

// InsertBefore insert a element with value before of at element, if at is not element of list l, will insert nothing, return e with value
func (l *List) InsertBefore(value interface{}, at *Element) *Element {
	if at.list != l {
		return at
	}
	return l.insertValue(value, at.prev)
}

// InsertAfter insert a element with value after of at element, if at is not element of list l, will insert nothing, return e with value
func (l *List) InsertAfter(value interface{}, at *Element) *Element {
	if at.list != l {
		return at
	}
	return l.insertValue(value, at)
}

// MoveToFront move the element to the front of the list of element e, when e is not the element of list, will do nothing
func (l *List) MoveToFront(e *Element) {
	if e.list != l || e.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack move the element to the back of list of element e , when e is not the element of list, will do nothing
func (l *List) MoveToBack(e *Element) {
	if e.list != l || e.next == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore move element e to after of element mark, if the list of e or the list of at is not l, or e == at will not modify the list
func (l *List) MoveBefore(e, at *Element) {
	if e.list != l || at.list != l || e == at {
		return
	}
	l.move(e, at.prev)
}

// MoveAfter
func (l *List) MoveAfter(e, at *Element) {
	if e.list != l || at.list != l || e == at {
		return
	}
	l.move(e, at)
}

//PushBackList insert a copy of other list after l, other list and l maybe same , but not nil
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Head(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList insert a copy of other list front l ...
func (l *List) PushFrontList(other *List) {

	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
