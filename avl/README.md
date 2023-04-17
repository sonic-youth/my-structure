## AVL Tree implement in go

AVL tree have the following feature
1. It must be a binary-search tree
2. For each node, the absolute value (**balance factor**) of the height difference between the left subtree and right subtree must **at most 1**;


So for AVL tree, we need to implement the following function to support the two feature

### Calculate height and balance factor
For one node, we need to maintain its height and balance factor, so we can see whether the node is balanced, and what we will do when we need to make it balanced.

### Verify the properties and balance of tree
The tree must be a binary-search tree, and we can use the in-order iterator to just the properties;
And for the balance, we can the balance factor for every node.

### Maintain balance
When we found one node is unbalance, we need to do some operator to make it balance, here's some function we need to implement to cover all the situations.

#### 1. Rotate Right
When a new node insert to the left-left children node, the node will be unbalanced(with balance factor=2),
So we need to do the following things:
1. Set y as x's right children(y > x)
2. 
```

        y                              x
       / \                           /   \
      x   T4    rotate right        z     y
     / \       - - - - - - - ->    / \   / \
    z   T3                       T1  T2 T3 T4
   / \
 T1   T2
```



