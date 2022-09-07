Task 1:

  1. Because some tasks are more important or urgent than others. Tasks are given priority so that it lets urgent/important ones happen as fast as possible without wasting time on waiting for less important tasks to complete.

  2. In order for a scehdulaer to be used in real-time systems, it must be predictable so we can analyze it and bu analyzing, it is meant that it needs to verify that all tasks meet their deadlines.


Task 2:

  1.With FIFO scheduling algorithm:

Wait time for each task:

a: 0-0=0      |       b: 5-1=4

c: 7-4=3       |      d: 11-6=5

Average waiting time: (0+4+3+5)/4=3

The overall gant chart is as follows:

![image](https://user-images.githubusercontent.com/97433657/165918446-7af171cf-a690-4c90-b1d0-907b6a2a12c9.png)


  2.With round-robin scheduling algorithm and q=1:
  
  
Turn around time = completion time - arrival time

a: 14 - 0 = 14

b: 6 - 2 = 4

c: 13 -4 = 9

d: 11 - 6 = 5

Waiting time = turn around time - execution time

a: 15 - 5 = 9

b: 4 - 2 = 2

c: 9 - 4 = 5

d: 5 - 3 = 2

Average waiting time = (9+2+5+2)/4=4.5

![image](https://user-images.githubusercontent.com/89384876/165950228-fda4d7d3-db11-40bc-a4f9-4d616c99123e.png)




Task 3:
  1. Without priority inheritance:
  
|Task|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|
|----|-|-|-|-|-|-|-|-|-|-|--|--|--|--|--|
| a  | | | | |E| | | | | |  |Q |V |E |  |
| b  | | |E|V| |V|E|E|E| |  |  |  |  |  |
| c  |E|Q| | | | | | | |Q|Q |  |  |  |E |

  2. With priority inheritance:
  
|Task|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|
|----|-|-|-|-|-|-|-|-|-|-|--|--|--|--|--|
| a  | | | | |E| | |Q| |V|E |  |  |  |  |
| b  | | |E|V| | | | |V| |  |E |E |E |  |
| c  |E|Q| | | |Q|Q| | | |  |  |  |  |E |

  Explain:
  1.
  Priority inversion: It occurs when a high-priority task is forced to wait for the release of a shared resource owned by a lower-priority task. This occur when two tasks attempt to access a single shared resource.
  
  Unbounded priority inversion: It occurs when there are one or more medium-priority tasks that prevent the low-priority task from running - and thereby releasing the   shared resource.  It occurs when an intervening task extends a bounded priority inversion, possibly forever.
  
  2.
  No, it does not. 
  Because if thread1 is interupted by thread2 in priority while using a resourceX, then thread2 allocateY and
  tries to allocateX, but is blocked by thread1. Thread1 now inherites the other ones priority and tries to complete the work it needs X for. However, since this work   requires allocation of the resourceY, it can't be completed, since thread2 allready has allocated it.
  
Task 4:

1. Assumptions:

  . Tasks are periodic with known time. Realistic.

  . There are a fixed set of tasks. May require workarounds.

  . Tasks are independent. Realistic.

  . There is no overhead. Depends on the system.

  . The deadline equal the period. Pretty realistic.
  
  . The task switching times need to be negligble/fixed. Not realistic in all cases.

2. Utilization test of task 3:

  U = 15/50 + 10/30 + 5/20 = 0.8833  which is not less than or equal to   3(2^(1/3)-1) = 0.7798 ==> The task is not schedulable.

3. Response-time analysis of task 3:

  Task c: w0 = 5 => RT = 5 <= 20

  Task b: w0 = 10 and w1 = 10 + ceil(10/20)*5 = 15 and  w2 = 10 + ceil(15/20)*5 = 15 => RT = 15 <= 30

  Task a: w0 = 15

  w1 = 15 + ceil(15/30)*10 + ceil(15/20)*5 = 15 + 10 + 5 = 30

  w2 = 15 + ceil(30/30)*10 + ceil(30/20)*5 = 15 + 10 + 10 = 35

  w3 = 15 + ceil(35/30)*10 + ceil(35/20)*5 = 15 + 20 + 10 = 45

  w4 = 15 + ceil(45/30)*10 + ceil(45/20)*5 = 15 + 20 + 15 = 50

  w5 = 15 + ceil(50/30)*10 + ceil(50/20)*5 = 15 + 20 + 15 = 50

  => RT = 50 <= 50

Therefore, the task is schedulable.


  
  
  
