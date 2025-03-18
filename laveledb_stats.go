package main

import (
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/some/ledgerdb")

func CollectMetrics() {
	store, err := NewETCDStore("ledgerdb_data_metrics")
	if err != nil {
		log.Fatalf("Failed to initialize ETCD store: %v", err)
	}
	defer store.Close()
	insertionTimes := []int64{}
	searchTimes := []int64{}
	deletionTimes := []int64{}
	var memStats runtime.MemStats
	for i := 1; i <= 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		start := time.Now()
		err := store.Insert(key, value)
		if err == nil {
			insertionTimes = append(insertionTimes, time.Since(start).Microseconds())
		}
		start = time.Now()
		_, err = store.Search(key)
		if err == nil {
			searchTimes = append(searchTimes, time.Since(start).Microseconds())
		}
		start = time.Now()
		err = store.Delete(key)
		if err == nil {
			deletionTimes = append(deletionTimes, time.Since(start).Microseconds())
		}
	}
	runtime.ReadMemStats(&memStats)
	cpuUsage := runtime.NumCPU()
	fmt.Printf("Average Insertion Time: %v µs\n", avg(insertionTimes))
	fmt.Printf("Average Search Time: %v µs\n", avg(searchTimes))
	fmt.Printf("Average Deletion Time: %v µs\n", avg(deletionTimes))
	fmt.Printf("CPU Usage: %d cores\n", cpuUsage)
	fmt.Printf("Memory Usage: %d bytes\n", memStats.Alloc)
}
func avg(times []int64) int64 {
	var total int64
	for _, t := range times {
		total += t
	}
	return total / int64(len(times))
}
