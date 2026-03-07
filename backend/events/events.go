//go:generate go run ../cmd/generate_events.go

package events

import (
	"context"
	"media-organizer/backend/stores"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Event string

type Events struct {
	FileLoadingBegin   Event `json:"fileLoadingBegin"`
	FileLoadingEnd     Event `json:"fileLoadingEnd"`
	DialogSelectFailed Event `json:"dialogSelectFailed"`
}

var EventNames = Events{
	FileLoadingBegin:   "file-loading:begin",
	FileLoadingEnd:     "file-loading:end",
	DialogSelectFailed: "dialog-select:failed",
}

type BackToFront struct {
	ctx context.Context
}

func CreateBackToFrontEvents(ctx context.Context) *BackToFront {
	return &BackToFront{ctx: ctx}
}

func (evt *BackToFront) FileLoadingBegin() {
	runtime.EventsEmit(evt.ctx, string(EventNames.FileLoadingBegin))
}

func (evt *BackToFront) FileLoadingEnd(files []*stores.FileEntry) {
	runtime.EventsEmit(evt.ctx, string(EventNames.FileLoadingEnd), files)
}

func (evt *BackToFront) DialogSelectFailed(err error) {
	runtime.EventsEmit(evt.ctx, string(EventNames.DialogSelectFailed), err)
}
