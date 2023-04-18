## AVL Tree Implement In Go

AVL tree have the following feature
1. It must be a binary-search tree
2. For each node, the absolute value (**balance factor**) of the height difference between the left subtree and right subtree must **at most 1**;


So for AVL tree, we need to implement the following function to support the two feature

### Calculate Height And Balance Factor
For one node, we need to maintain its height and balance factor, so we can see whether the node is balanced, and what we will do when we need to make it balanced.

### Verify The Properties And Balance Of Tree
The tree must be a binary-search tree, and we can use the in-order iterator to just the properties;
And for the balance, we can the balance factor for every node.

### Maintain Balance
When we found one node is unbalance, we need to do some operator to make it balance, here's some function we need to implement to cover all the situations.

#### 1. Rotate Right
When a new node insert to the left-left children node(node z), the node y will be unbalanced(with balance factor=2).
So we need to do the following operations:
1. Set y as x's right children(y > x)
2. Set T3 as y's left children(y > T3)
```

        y                              x
       / \                           /   \
      x   T4    rotate right        z     y
     / \       - - - - - - - ->    / \   / \
    z   T3                       T1  T2 T3 T4
   / \
 T1   T2
```

#### 2. Rotate Left
When a new node insert to the right-right children node(node z), the node y will be unbalanced(with balance factor=2).
So we need to do the same thing as **Rotate Right** do, just reverse the operations.

```
    y                                   x
   / \                                /   \
  T4  x          rorate left         y     z
     / \       - - - - - - - ->     / \   / \
    T3  z                          T4 T3 T2 T1
       / \
      T2 T1
```
#### 3. Rotate Left-Right
When a new node insert to the left-right children node(node z), the node y will be unbalanced(with balance factor=2).
So we need to do the following operations:
1. Rotate left for the node x.
2. Rotate right for the node y.

```
     y                      y                        z
    / \                    / \                     /   \ 
   x   T1  rotate left    z  T1   rotate right    x      y
  / \      - - - - - >   / \      - - - - - - >  / \    / \
T2   z                  x  T4                   T2 T3  T4 T1
    / \                / \
   T3  T4             T2 T3 
```

#### 4. Rotate Right-Left
When a new node insert to the right-left children(node z), the node y will be unbalanced(with balance factor=2).
So we need to the same things as **Rotate Left-Right** do, just reverse the operations.
1. Rotate right the node x.
2. Rotate left for the node y.
```
     y                      y                           z
    / \                    / \                        /   \
   T1  x   rotate right   T1  z    rotate left       y     x
      / \  - - - - - - >     / \   - - - - - - - >  / \   / \
     z  T2                  T3  x                  T1 T3 T4 T2
    / \                        / \
   T3  T4                     T4 T2
```
### Insert Node
When insert a new node, you need to do the following operations
1. Update the height of the node
2. Judge if the new node will make unbalance occured.
3. If unbalance occured, what rotate operation we need do.

### Delete Node
When delete a new node, there are some situations will happen.
1. When delete a leaf node, just delete it.
2. When delete a non-leaf node, you need to find the min-value in the node's right subtree, to reorganize a new root node.
3. Judge if the new node is balanced, if not, what rotate operation we need to do.
