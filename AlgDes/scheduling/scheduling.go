package main 

import (
	"fmt"
	"sort"
	"UJS/file"
	"strings"
	"strconv"
)

// type job is a prioritized task which needs to be scheduled
type job struct {
	length int
	weight int
	priority float32
}

type jobList []*job

func (jL jobList) Len() int {
	return len(jL)
}

func (jL jobList) Less(i, j int) bool {
	retVal := false
	if jL[i].priority > jL[j].priority {
		retVal = true
	} else if jL[i].priority == jL[j].priority && jL[i].weight > jL[j].weight{
		retVal = true
	}
	return retVal
}

func (jL jobList) Swap(i, j int) {
	jL[i], jL[j] = jL[j], jL[i]
	return
}

const fileName string = "D:\\James\\Temp\\jobs.txt"

func main() {
	jobSlice := loadJobs()

	prioritize(jobSlice)

	totalTime := schedule(jobSlice)

	fmt.Println(totalTime)
}

func schedule(jobSlice *jobList) int{
	weightedTime, runTime := 0, 0
	sort.Sort(jobSlice)
	for _, v := range *jobSlice {
		runTime += v.length
		weightedTime += runTime * v.weight
	}
	return weightedTime
}

func prioritize(jobSlice *jobList) {
	for _, v := range *jobSlice {
		v.priority = float32(v.weight) / float32(v.length)
	}
	return
}

func loadJobs() *jobList{
	a := file.FileLines(fileName)[1:]
	jobSlice := make(jobList, 10000)

	for i, v := range a {
		jobSlice[i] = addJob(v)
	}
	return &jobSlice
}

func addJob(a string) *job{
	parts := strings.Split(a, "\n")
	parts = strings.Split(parts[0], " ")
	jobWeight, _ := strconv.Atoi(parts[0])
	jobLength, _ := strconv.Atoi(parts[1])
	newJob := &job{jobLength, jobWeight, 0}
	return newJob
}