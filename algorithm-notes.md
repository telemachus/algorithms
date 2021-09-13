# Notes on Algorithms

## Loop Invariants

Loop invariants help us to see that an algorithm is correct. Three things should be true about loop invariants. (See *Introduction to Algorithms*[^1], pages 18-20.)

1. Initialization: the loop invariant is true before the first iteration of the loop.
1. Maintenance: the loop invariant is true after each iteration of the loop. (Note that the loop invariant may become temporarily false during the loop, but by the end of the loop, it must be true.)
1. Termination: when the loop ends, the invariant gives us a useful property that helps show that the algorithm is correct.

## Loop Invariants and Insertion Sort

As an example, let’s consider the loop invariants for insertion sort. The loop invariant for insertion sort are the following: first, for each iteration through the outer loop, where `j` is an index of the loop from `1` to `len(slice)-1`, the subslice `slice[0:j-1]` contains the elements from `slice[0:j-1]`, but in sorted order; second, the inner loop only adds items to `xs[i:j]` if those elements are less than or equal to `key`. (I think that there should be a better way to describe the second invariant.)

### Loop Invariant 1

1. Initialization: initially the subslice `xs[0:1]` contains one item; therefore, it is trivially in order.
1. Maintenance: informally, the point of the inner loop is to move `key` down as far as possible; therefore, at the end of the inner loop, everything in `xs[i:j]` is greater than or equal to key. More formally, we explicitly test whether `xs[i] > key` before doing anything. If `xs[i]` is not greater than `key`, then `i` remains unchanged and when we run `xs[i+1] = key`, we are placing `key` back where it started since `i` was defined as `j-1`. Therefore, `i+1` *is* j whenever i has not changed. On the other hand, if `xs[i]` is greater than `key`, then we do two things. First, we slide `xs[i]` rightward. Second, we decrease `i`. Thus, when we assign key to `xs[i+1]`, `xs[i+1]` has become `j-1`. Since the key was previously `xs[j]`, we have thus moved `key` down one and the larger number up one. Thus, `xs[0:j-1]` maintains the first invariant.
1. Termination: when the outer loop finishes, `j` is `len(xs)`. Thus, we can say that the entire slice (`xs[0:len(xs)]`) is in order when the outer loop finishes.

### Loop Invariant 2

1. Initialization: before the loop begins, nothing has been added to `xs[i:j]`, so the invariant is true in a trivial way.
1. Maintenance: the inner loop has a test that insures that anything larger than `key` is moved upwards. Thus, anything larger than key will end up somewhere in `xs[j:]`, and the invariant remains true.
1. Termination: the inner loop ends in one of two ways. First, the loop ends if reach a value of `i` where `xs[i]` is less than or equal to `key`. (Any such items are kept in `xs[i:j]`.) Thus, if the loop ends in this way, the invariant survives termination. Second, the loop explicitly tests for when `i` becomes less than zero. (This test protects us from attempting to index outside the bounds of a slice.) The lower bound of `xs[i:j]` is therefore zero, and based on initialization and maintenance, we know that `xs[0:j]` will have nothing added to it that isn’t smaller than or equal to `key`.

[^1]: *Introduction to Algorithms*, 3rd edition, Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein. The MIT Press (Cambridge, MA), 2009.
