package main

import "github.com/s8sg/goflow/flow"

// define provide definition of the workflow
func DefineWorkFlow() {

}

// workload
func doSomething(ctx *flow.Context,f *flow.Workflow) error{
	f.SyncNode(options ...flow.BranchOption)
}

func main() {

}
