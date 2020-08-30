package todo

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type instrumentingService struct {
	createCounter  prometheus.Counter
	createLatency  prometheus.Histogram
	findOneCounter prometheus.Counter
	findOneLatency prometheus.Histogram
	findAllCounter prometheus.Counter
	findAllLatency prometheus.Histogram
	deleteCounter  prometheus.Counter
	deleteLatency  prometheus.Histogram
	updateCounter  prometheus.Counter
	updateLatency  prometheus.Histogram
	doneCounter    prometheus.Counter
	doneLatency    prometheus.Histogram

	next Service
}

func NewInstrumentingService(createCounter prometheus.Counter, createLatency prometheus.Histogram, findOneCounter prometheus.Counter, findOneLatency prometheus.Histogram, findAllCounter prometheus.Counter, findAllLatency prometheus.Histogram, deleteCounter prometheus.Counter, deleteLatency prometheus.Histogram, updateCounter prometheus.Counter, updateLatency prometheus.Histogram, doneCounter prometheus.Counter, doneLatency prometheus.Histogram, next Service) *instrumentingService {
	return &instrumentingService{
		createCounter:  createCounter,
		createLatency:  createLatency,
		findOneCounter: findOneCounter,
		findOneLatency: findOneLatency,
		findAllCounter: findAllCounter,
		findAllLatency: findAllLatency,
		deleteCounter:  deleteCounter,
		deleteLatency:  deleteLatency,
		updateCounter:  updateCounter,
		updateLatency:  updateLatency,
		doneCounter:    doneCounter,
		doneLatency:    doneLatency,
		next:           next,
	}
}

func (i *instrumentingService) Create(title, description string) (Todo, error) {
	defer func(begin time.Time) {
		i.createCounter.Inc()
		i.createLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Create(title, description)
}

func (i *instrumentingService) FindById(id int) (Todo, error) {
	defer func(begin time.Time) {
		i.findOneCounter.Inc()
		i.findOneLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.FindById(id)
}

func (i *instrumentingService) FindAll(limit, offset int) ([]Todo, error) {
	defer func(begin time.Time) {
		i.findAllCounter.Inc()
		i.findAllLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.FindAll(limit, offset)
}

func (i *instrumentingService) Delete(id int) (Todo, error) {
	defer func(begin time.Time) {
		i.deleteCounter.Inc()
		i.deleteLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Delete(id)
}

func (i *instrumentingService) Update(id int, title, description string, status bool) (Todo, error) {
	defer func(begin time.Time) {
		i.updateCounter.Inc()
		i.updateLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Update(id, title, description, status)
}

func (i *instrumentingService) Done(id int, status bool) (Todo, error) {
	defer func(begin time.Time) {
		i.doneCounter.Inc()
		i.doneLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Done(id, status)
}
