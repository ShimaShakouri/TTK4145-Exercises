Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> → They are both ways to handle data flow, but they do it in
different ways.
→ Concurrency is about running and managing different
computations at the same time with only using one CPU (they
share a cashier).
→ Parallelism is about running multiple computations
simultaneously with different CPU s so that it works in
parallel (they have their own cashier).

What is the difference between a *race condition* and a *data race*? 
> 
> A race condition is a situation, in which the result of an operation depends on the interleaving of certain individual operations while a data race is a situation, in which at least two threads access a shared variable at the same time. At least one thread tries to modify the variable.
A data race is undefined behaviour and can cause all reasoning about the program to go wrong. A race condition can be the reason for a data race but it is not necesserily bad.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> → A scheduler assigns resources to perform tasks.
It does this by deciding witch process runs at a given point in
time. It often has the ability to pause, and/or replace a process.
It can also change the priority of a process by moving it
backward in the stack.
→ The scheduler is the activity of the process manager that
handles the removal of the running process from the CPU, and
the selection of another process on the basis of a particular
strategy.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> → This could be to make a program faster and is simpler to
program.
→ Multiple threads (parallelism) helps solving problems with
application performance. Threads require fewer resources and
generate less overhead.
Multiple threads leads to the computer becoming faster and
being able to calculate and/or do many more things or complete
many more processes over the same time period. But as a result
the computer needs many more CPU s for it to manage all the
→ Multiple threads (parallelism) helps solving problems with
application performance. Threads require fewer resources and
generate less overhead.
Multiple threads leads to the computer becoming faster and
being able to calculate and/or do many more things or complete
many more processes over the same time period. But as a result
the computer needs many more CPU s for it to manage all the
threads.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> → "Green threads" is virtual threads that are scheduled by a
runtime library or virtual machine, in contrast to threads that
are controlled by the underlying operating systems (OS)
→ "Fibers" is a particularly lightweight thread of execution, and
uses cooperative multitasking.
→ "Coroutines" are computer program components that
generalize subroutines for cooperative-multitasking.
→ Threads are pre-emptively scheduled while coroutines/ fibers
are not. Threads is determined in advance what it is supposed
to do, but coroutines/ fibers can change there objective along
the way.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> → I would say it makes it harder for the programmer. Because the
programmer must have more control over the hole program, so
he or she does not make a change that has effect on something
else that was not intended.

What do you think is best - *shared variables* or *message passing*?
> → Message passing model allows multiple processes to read and
write data to the message queue without being connected to
each other.
→ Shared memory model is the memory that can be
simultaneously accessed by multiple processes.
→ Shared memory has the advantage that it is faster then the
message passing model when on the same computer.
I/we think shared memory is the best because it is the fastest
and have the smallest change of problems since it can only be
addressed.


