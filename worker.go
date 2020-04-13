package celeriac

import (
	"fmt"
	"sync"
)

type WorkerOptions struct {
	Concurrency int
}

type Worker struct {
	broker Broker
	opts   *WorkerOptions
	tasks  []Task
	queues []Queue
	wg     *sync.WaitGroup
}

func (w *Worker) Run(stopCh <-chan struct{}) error {
	for i, t := range w.tasks {
		q, err := w.broker.GetQueue(QueueNameForTask(t))
		if err != nil {
			return err
		}

		w.queues[i] = q
	}

	for i, q := range w.queues {
		messages, err := q.Consume()
		if err != nil {
			return err
		}

		w.wg.Add(1)

		go w.consumer(fmt.Sprintf("%d", i), messages, stopCh)
	}

	w.wg.Wait()

	return nil
}

func (w *Worker) consumer(id string, messages <-chan Message, stopCh <-chan struct{}) {
	defer w.wg.Done()

	for {
		select {
		case msg := <-messages:
			body, err := ParseMessageBody(msg.Body)
			if err != nil {
				fmt.Printf("[ERROR] unable to parse message body: %s\n", err)
				continue
			}

			fmt.Printf("Worker %s received Task: %s\n", id, msg.MessageID)
			task := GetTask(body.Task)
			if task == nil {
				fmt.Printf("[ERROR] task %s does not exist: %s\n", body.Task)
				continue
			}

			if err := task.Init(); err != nil {
				fmt.Printf("[ERROR] task %s(%s).Init: %s\n", err)
				continue
			}

			if err := task.Run(body.Args...); err != nil {
				fmt.Printf("[ERROR] task %s(%s).Run: %s\n", err)
			}

			if err := task.Exit(); err != nil {
				fmt.Printf("[ERROR] task %s(%s).Exit: %s\n", err)
				continue
			}

		case <-stopCh:
			fmt.Printf("Worker %s shutting down\n", id)
			return
		}
	}
}

func NewWorker(b Broker, opts *WorkerOptions, t ...Task) *Worker {
	if opts == nil {
		opts = &WorkerOptions{
			Concurrency: len(t),
		}
	}

	if opts.Concurrency == 0 {
		opts.Concurrency = len(t)
	}

	return &Worker{
		broker: b,
		opts:   opts,
		tasks:  t,
		queues: make([]Queue, len(t)),
		wg:     new(sync.WaitGroup),
	}
}
