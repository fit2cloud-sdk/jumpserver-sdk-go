package main

import (
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func testOps() {
	section("Ops Jobs")

	list, _, err := client.Self.ListAssets(ctx, nil, nil)
	if err != nil || len(list) == 0 {
		skip("Ops.CreateJob", "no self assets available")
		return
	}

	assetID := list[0].ID
	job, _, err := client.Ops.CreateJob(ctx, &model.OpsJobRequest{
		Assets:      []string{assetID},
		Module:      "shell",
		Args:        "date",
		RunAs:       "root",
		RunAsPolicy: "skip",
		Instant:     true,
		IsPeriodic:  false,
		Timeout:     -1,
	})
	if err != nil {
		fail("Ops.CreateJob", err)
		return
	}
	ok("Ops.CreateJob (task_id=" + job.TaskID + ")")

	result, _, err := client.Ops.GetJobResult(ctx, job.TaskID)
	if err != nil {
		fail("Ops.GetJobResult", err)
	} else {
		ok("Ops.GetJobResult (finished=" + boolStr(result.IsFinished) + ")")
	}
}
