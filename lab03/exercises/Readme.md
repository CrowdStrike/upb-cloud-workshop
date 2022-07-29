1. Given an array of INTs, calculate the sum of even numbers in the array using 4 go routines in two different ways(using channels, using WaitGroup).
2. In the crawler directory, make the Crawl function parse the urls in a parallel manner(go routines)
Hint: you can use a map for checking which urls are visited

3. Implement a parralel binary tree construction inside the tree directory.

The binary tree should have a method insert which should add a value to the tree, with the signature below. <br>

    insert(value int) *BinaryTree

Adding a value should respect the generic insertion rules, if the value is smaller than the current one we add it to the left and if the value is greater we add it to the right. Do not add the value if we already have it in the tree.

Syncronisation should not be done over the entire tree, but only for the actual node where you need to add the value. Hint: mutex

Create n threads that build the tree with randomly generated values, you can start with n=2, then wait for the tree to be fully constructed using a barrier/ wait group. Afterwards send all the values on a channel making sure they are sent in ascending order.

Use another thread to read from the channel to check if the values are really sent in ascending order.
