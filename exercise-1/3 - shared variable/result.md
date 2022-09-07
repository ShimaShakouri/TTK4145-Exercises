In the first attempt of this exercise, we just incremented and decremented the value 'i' with for loops without using channels.
they are running independantly and not communicationg with each other and just communicating with the shared resource which is 'i'. In summary, the main problem is that we have two threads and swapping can occur at any time. Maybe the first one reads but not write back the result yet which will lead to an answer different from 0 while we are looping for thousands of time and overwriting occurs while we do not want this to happen.


Mutex vs Semaphore:
We chose mutex over semaphore since it allows the threads to access a single shared resource (which is our variable "i") but only one at a time. But, semaphore allows threads to access the finite instance of the resource until available.
