// Compile with `gcc foo.c -std=c99 -lpthread`, or use the makefile

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t Mutex;

// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 0 ; j<1000000 ; j++){
      pthread_mutex_lock(&Mutex);
      i++;
      pthread_mutex_unlock(&Mutex);
    }

    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0 ; j<1000000 ; j++){
      pthread_mutex_lock(&Mutex);
      i--;
      pthread_mutex_unlock(&Mutex);
    }
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    pthread_mutex_init(&Mutex,NULL);
  
    pthread_t First_Thread;
    pthread_create(&First_Thread, NULL, incrementingThreadFunction, NULL);
    
    
    pthread_t Second_Thread;
    pthread_create(&Second_Thread, NULL, decrementingThreadFunction, NULL);

    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    
    pthread_join(First_Thread, NULL);
    pthread_join(Second_Thread, NULL);
    
    printf("The magic number is: %d\n", i);
  
    pthread_mutex_destroy(&Mutex);
    return 0;
}
