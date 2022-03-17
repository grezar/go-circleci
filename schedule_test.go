package circleci

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_schedules_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	projectSlug := "gh/org1/prj1"

	mux.HandleFunc(fmt.Sprintf("/project/%s/schedule", projectSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Circle-Token", client.token)
		fmt.Fprint(w, `{"items": [{"id": "1","timetable":{"per-hour":0,"hours-of-day":[0]}}], "next_page_token": "1"}`)
	})

	ctx := context.Background()
	sl, err := client.Schedule.List(ctx, projectSlug, ScheduleListOptions{})
	if err != nil {
		t.Errorf("Schedules.List got error: %v", err)
	}

	want := &ScheduleList{
		Items: []*Schedule{
			{
				ID: "1",
				Timetable: Timetable{
					PerHour:    0,
					HoursOfDay: []int{0},
				},
			},
		},
		NextPageToken: "1",
	}

	if !cmp.Equal(sl, want) {
		t.Errorf("Schedules.List got %+v, want %+v", sl, want)
	}
}
