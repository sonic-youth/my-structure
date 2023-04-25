## Skiplist Implement In Go
Skiplist is a simple but high-performance structure, which is designed for replacing complex self balancing tree.  
As we know, AVL tree will do many times of rotating operation(no more than O(logn)), which will impact the efficiency of insert and delete operation, so we created the Red-Black-Tree.
But there's something we need to face to
1. In the concurrency scenarios, when we need to update value, the Red-Black-Tree will also need to do some rotating/staining operation, which will involving many nodes.
2. Red-Black-Tree cannot quickly support **Range Search** operation.
3. Red-Black-Tree is too complex to implement.

So skiplist emerge as the times require, to help us implement less, and create more.

