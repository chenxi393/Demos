#include <assert.h>
#include <iostream>
#include <malloc.h>
#include <pthread.h>
#include <queue>
#include <vector>
using namespace std;

static void* worker(void* arg);

struct Work {
    void (*func)(void*);
    void* arg;
};

struct ThreadPool {
    // 互斥锁 锁住全局队列
    pthread_mutex_t mutex;
    pthread_cond_t not_empty;

    deque<Work> queue;
    vector<pthread_t> threads;
};

void thread_pool_init(ThreadPool* p, size_t num_threads)
{
    // 防御性编程
    assert(num_threads > 0);

    // 这些api应该都是
    int rv = pthread_mutex_init(&p->mutex, NULL);
    assert(rv == 0);

    rv = pthread_cond_init(&p->not_empty, NULL);
    assert(rv == 0);

    p->threads.resize(num_threads);
    for (size_t i = 0; i < num_threads; i++) {
        // worker 即消费者代码
        rv = pthread_create(&p->threads[i], NULL, &worker, p);
        assert(rv == 0);
    }
}

void thread_pool_destroy(ThreadPool* p)
{
    //  回收线程
    for (size_t i = 0; i < p->threads.size(); i++) {
        // worker 即消费者代码
        pthread_join(p->threads[i], NULL);
    }
}

// 消费者代码
static void* worker(void* arg)
{
    ThreadPool* tp = (ThreadPool*)arg;
    while (true) {
        // 先拿锁 看看工作队列是否有任务
        pthread_mutex_lock(&tp->mutex);
        // pthread_cond_signal不能保证每次只唤醒1个等待线程，可能会唤醒多个，所以会导致惊群现象，
        // 为了解决惊群现象，我们可以在pthread_cond_wait前面 通过while 判断
        // 条件变量标识的队列或变量的状态是否可用，只有可用才去消费，如果不可用，则继续返回等待
        while (tp->queue.empty()) {
            pthread_cond_wait(&tp->not_empty, &tp->mutex);
        }

        // got the job
        Work w = tp->queue.front();
        tp->queue.pop_front();
        pthread_mutex_unlock(&tp->mutex);

        // 执行任务
        w.func(w.arg);
    }
    return NULL;
}

// 生产者代码
static void producer(ThreadPool* tp, void (*f)(void*), void* arg)
{
    Work w;
    w.arg = arg;
    w.func = f;
    pthread_mutex_lock(&tp->mutex);
    tp->queue.push_back(w);
    pthread_cond_signal(&tp->not_empty);
    pthread_mutex_unlock(&tp->mutex);
}

void func_test(void* i)
{
    int* num = (int*)(i);
    printf("I am Thread to print %d \n", *num);
}

int main()
{
    ThreadPool pool;
    thread_pool_init(&pool, 10);

    for (int i = 0; i < 200000; i++) {
        int* ppp = (int*)malloc(sizeof(int));
        *ppp = i;
        producer(&pool, &func_test, ppp);
    }
    thread_pool_destroy(&pool);
    return 0;
}
