type SinglyLinkedList struct {
    front *Node
  
    length int
  }
  
  // Init initializes an empty list
  func (s *SinglyLinkedList) Init() *SinglyLinkedList {
    s.length = 0
    return s
  }
  
  // New returns an initialized list
  func New() *SinglyLinkedList {
    return new(SinglyLinkedList).Init()
  }
  
  // Front returns the first node in list s
  func (s *SinglyLinkedList) Front() *Node {
    return s.front
  }
  
  // Back returns the last node in list s
  func (s *SinglyLinkedList) Back() *Node {
    currentNode := s.front
    for currentNode != nil && currentNode.next != nil {
      currentNode = currentNode.next
    }
    return currentNode
  }
  
  // Append appends node n to list s
  func (s *SinglyLinkedList) Append(n *Node) {
    if s.front == nil {
      s.front = n
    } else {
      currentNode := s.front
  
      for currentNode.next != nil {
        currentNode = currentNode.next
      }
  
      currentNode.next = n
    }
  
    s.length++
  }
  func (s *SinglyLinkedList) InsertAfter(insert *Node, after *Node) {
    currentNode := s.front
    for currentNode != nil && currentNode != after && currentNode.next != nil {
      currentNode = currentNode.next
    }
  
    if currentNode == after {
      insert.next = after.next
      after.next = insert
      s.length++
    }
  }
  
  // Remove removes node n from list s
  func (s *SinglyLinkedList) Remove(n *Node) {
    if s.front == n {
      s.front = n.next
      s.length--
    } else {
      currentNode := s.front
  
      // search for node n
      for currentNode != nil && currentNode.next != nil && currentNode.next != n {
        currentNode = currentNode.next
      }
  
      // see if current's next node is n
      // if it's not n, then node n wasn't found in list s
      if currentNode.next == n {
        currentNode.next = currentNode.next.next
        s.length--
      }
    }
  }