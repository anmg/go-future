#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

pthread_mutex_t lock;

class Singleton{
    private:
        Singleton(){}
        static Singleton* instance;

    public:
        ~Singleton(){}
        static Singleton* getInstance(){
            if(instance==NULL){
                pthread_mutex_lock(&lock);
                if(instance==NULL){
                    instance=new Singleton;
                    printf("obj created.\n");
                }
                pthread_mutex_unlock(&lock);
            }
            return instance;
        }
};

Singleton* Singleton::instance = NULL;
//Singleton* Singleton::instance = new Singleton;

void* fun1(void*){
    Singleton::getInstance();
    printf("%d\n", 1);
}

void* fun2(void*){
    Singleton::getInstance();
    printf("%d\n", 2);
}

int main(void){
    pthread_mutex_init(&lock, NULL);

    pthread_t pid1, pid2;
    pthread_create(&pid1, NULL, fun1, NULL);
    pthread_create(&pid2, NULL, fun2, NULL);

    pthread_join(pid1, NULL);
    pthread_join(pid2, NULL);

    pthread_mutex_destroy(&lock);

    return 0;
}