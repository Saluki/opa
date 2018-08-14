// Copyright 2018 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package profiler

import (
	"context"
	_ "encoding/json"
	"testing"
	"time"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/topdown"
	"github.com/open-policy-agent/opa/types"
)

func TestProfilerLargeArray(t *testing.T) {
	profiler := New()
	module := `package test

foo {
	bar
	not baz
	bee
}

bee {
	nums = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i", "b", "c", "d", "e", "f", "g", "h", "i"]
	num = nums[_]
	contains(num, "test")
}

bar {
	a := 1
	b := 2
	a != b
}

baz {
	true
	false
	true
}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	report := profiler.ReportByFile()

	fr, ok := report.Files["test.rego"]
	if !ok {
		t.Fatal("Expected file report for test.rego")
	}

	if len(fr.Result) != 11 {
		t.Fatalf("Expected file report length to be 11 instead got %v", len(fr.Result))
	}

	expectedNumEval := []int{1, 2, 1, 1, 1, 1633, 1, 1, 1, 1, 1}
	expectedNumRedo := []int{1, 0, 0, 1, 1633, 0, 1, 1, 1, 1, 0}
	expectedRow := []int{4, 5, 6, 10, 11, 12, 16, 17, 18, 22, 23}

	for idx, actualExprStat := range fr.Result {
		if actualExprStat.NumEval != expectedNumEval[idx] {
			t.Fatalf("Index %v: Expected number of evals %v but got %v", idx, expectedNumEval[idx], actualExprStat.NumEval)
		}

		if actualExprStat.NumRedo != expectedNumRedo[idx] {
			t.Fatalf("Index %v: Expected number of redos %v but got %v", idx, expectedNumRedo[idx], actualExprStat.NumRedo)
		}

		if actualExprStat.Location.Row != expectedRow[idx] {
			t.Fatalf("Index %v: Expected row %v but got %v", idx, expectedRow[idx], actualExprStat.Location.Row)
		}

	}
}

func TestProfileCheckExprDuration(t *testing.T) {
	profiler := New()

	ast.RegisterBuiltin(&ast.Builtin{
		Name: "test.sleep",
		Decl: types.NewFunction(
			types.Args(types.S),
			types.NewNull(),
		),
	})

	topdown.RegisterFunctionalBuiltin1("test.sleep", func(a ast.Value) (ast.Value, error) {
		d, _ := time.ParseDuration(string(a.(ast.String)))
		time.Sleep(d)
		return ast.Null{}, nil
	})

	module := `package test

	foo {
	  test.sleep("100ms")
	}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	report := profiler.ReportByFile()

	fr, ok := report.Files["test.rego"]
	if !ok {
		t.Fatal("Expected file report for test.rego")
	}

	if len(fr.Result) != 1 {
		t.Fatalf("Expected file report length to be 1 instead got %v", len(fr.Result))
	}

	if string(fr.Result[0].Location.Text) != "test.sleep(\"100ms\")" {
		t.Fatalf("Expected text is test.sleep(\"100ms\") but got %v", string(fr.Result[0].Location.Text))
	}

	if fr.Result[0].ExprTimeNs <= time.Duration(50*time.Millisecond).Nanoseconds() {
		t.Fatalf("Expected eval time is atleast 100 msec but got %v", fr.Result[0].ExprTimeNs)
	}

}

func TestProfilerReportTopNResultsNoCriteria(t *testing.T) {
	profiler := New()
	module := `package test

foo {
	bar
	not baz
	bee
}

bee {
	nums = ["a", "b", "c", "d"]
	num = nums[_]
	contains(num, "test")
}

bar {
	a := 1
	b := 2
	a != b
}

baz {
	true
	false
	true
}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	stats := profiler.ReportTopNResults(0, []string{})

	expectedResLen := 12
	if len(stats) != expectedResLen {
		t.Fatalf("Expected %v stats instead got %v", expectedResLen, len(stats))
	}
}

func TestProfilerReportTopNResultsOneCriteria(t *testing.T) {
	profiler := New()
	module := `package test

foo {
	bar
	not baz
	bee
}

bee {
	nums = ["a", "b", "c", "d"]
	num = nums[_]
	contains(num, "test")
}

bar {
	a := 1
	b := 2
	a != b
}

baz {
	true
	false
	true
}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	stats := profiler.ReportTopNResults(5, []string{"total_time_ns"})

	expectedResLen := 5
	if len(stats) != expectedResLen {
		t.Fatalf("Expected %v stats instead got %v", expectedResLen, len(stats))
	}

	var i int
	for i = 0; i < len(stats)-1; i++ {
		if stats[i].ExprTimeNs < stats[i+1].ExprTimeNs {
			t.Fatalf("Results not sorted in decreasing order of evaluation times")
		}
	}
}

func TestProfilerReportTopNResultsTwoCriteria(t *testing.T) {
	profiler := New()
	module := `package test

foo {
	bar
	not baz
	bee
}

bee {
	nums = ["a", "b", "c", "d"]
	num = nums[_]
	contains(num, "test")
}

bar {
	a := 1
	b := 2
	a != b
}

baz {
	true
	false
	true
}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	stats := profiler.ReportTopNResults(5, []string{"num_eval", "total_time_ns"})

	expectedResLen := 5
	if len(stats) != expectedResLen {
		t.Fatalf("Expected %v stats instead got %v", expectedResLen, len(stats))
	}

	var i int
	for i = 0; i < len(stats)-1; i++ {
		if stats[i].NumEval < stats[i+1].NumEval {
			t.Fatalf("Results not sorted in decreasing order of number of evaluations")
		}

		if stats[i].NumEval == stats[i+1].NumEval {
			if stats[i].ExprTimeNs < stats[i+1].ExprTimeNs {
				t.Fatalf("Results not sorted in decreasing order of evaluation times")
			}
		}
	}
}

func TestProfilerReportTopNResultsThreeCriteria(t *testing.T) {
	profiler := New()
	module := `package test

foo {
	bar
	not baz
	bee
}

bee {
	nums = ["a", "b", "c", "d"]
	num = nums[_]
	contains(num, "test")
}

bar {
	a := 1
	b := 2
	a != b
}

baz {
	true
	false
	true
}`

	_, err := ast.ParseModule("test.rego", module)
	if err != nil {
		t.Fatal(err)
	}

	eval := rego.New(
		rego.Module("test.rego", module),
		rego.Query("data.test.foo"),
		rego.Tracer(profiler),
	)

	ctx := context.Background()
	_, err = eval.Eval(ctx)

	if err != nil {
		t.Fatal(err)
	}

	stats := profiler.ReportTopNResults(10, []string{"num_eval", "num_redo", "total_time_ns"})

	expectedResLen := 10
	if len(stats) != expectedResLen {
		t.Fatalf("Expected %v stats instead got %v", expectedResLen, len(stats))
	}

	var i int
	for i = 0; i < len(stats)-1; i++ {
		if stats[i].NumEval < stats[i+1].NumEval {
			t.Fatalf("Results not sorted in decreasing order of number of evaluations")
		}

		if stats[i].NumEval == stats[i+1].NumEval {

			if stats[i].NumRedo < stats[i+1].NumRedo {
				t.Fatalf("Results not sorted in decreasing order of number of redos")
			}

			if stats[i].NumRedo == stats[i+1].NumRedo {
				if stats[i].ExprTimeNs < stats[i+1].ExprTimeNs {
					t.Fatalf("Results not sorted in decreasing order of evaluation times")
				}
			}
		}
	}
}