package worker

type IWorkerProcessor interface {
	WorkerStart() error
}
