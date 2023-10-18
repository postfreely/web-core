package activitystreams

import (
	"reflect"
	"testing"
)

func TestNewImageAttachment(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Attachment
	}{
		{name: "good svg", args: args{"https://writefreely.org/img/writefreely.svg"}, want: Attachment{
			Type:      "Image",
			URL:       "https://writefreely.org/img/writefreely.svg",
			MediaType: "image/svg+xml",
		}},
		{name: "good png", args: args{"https://i.snap.as/12345678.png"}, want: Attachment{
			Type:      "Image",
			URL:       "https://i.snap.as/12345678.png",
			MediaType: "image/png",
		}},
		{name: "no extension", args: args{"https://i.snap.as/12345678"}, want: Attachment{
			Type:      "Image",
			URL:       "https://i.snap.as/12345678",
			MediaType: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageAttachment(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVideoAttachment(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Attachment
	}{
		{name: "good avi", args: args{"https://videos.test.example/path/to/something.avi"}, want: Attachment{
			Type:      "Video",
			URL:       "https://videos.test.example/path/to/something.avi",
			MediaType: "video/x-msvideo",
		}},
		{name: "good flv", args: args{"https://files.example.com/path/to/something.flv"}, want: Attachment{
			Type:      "Video",
			URL:       "https://files.example.com/path/to/something.flv",
			MediaType: "video/x-flv",
		}},
		{name: "good mkv", args: args{"https://orange.example/v/VIEW.mkv"}, want: Attachment{
			Type:      "Video",
			URL:       "https://orange.example/v/VIEW.mkv",
			MediaType: "video/x-matroska",
		}},
		{name: "good mov", args: args{"https://example.com/final.mov"}, want: Attachment{
			Type:      "Video",
			URL:       "https://example.com/final.mov",
			MediaType: "video/quicktime",
		}},
		{name: "good mp4", args: args{"https://www.videos.example/a/bb/ccc/untitled.mp4"}, want: Attachment{
			Type:      "Video",
			URL:       "https://www.videos.example/a/bb/ccc/untitled.mp4",
			MediaType: "video/mp4",
		}},
		{name: "good ogv", args: args{"https://files.example.com/path/to/file.ogv"}, want: Attachment{
			Type:      "Video",
			URL:       "https://files.example.com/path/to/file.ogv",
			MediaType: "video/ogg",
		}},
		{name: "good webm", args: args{"https://something.example/videos/something.webm"}, want: Attachment{
			Type:      "Video",
			URL:       "https://something.example/videos/something.webm",
			MediaType: "video/webm",
		}},
		{name: "good wmv", args: args{"https://v.example.com/f/something.wmv"}, want: Attachment{
			Type:      "Video",
			URL:       "https://v.example.com/f/something.wmv",
			MediaType: "video/x-ms-wmv",
		}},
		{name: "no extension", args: args{"https://some.example/watch/123456789"}, want: Attachment{
			Type:      "Video",
			URL:       "https://some.example/watch/123456789",
			MediaType: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoAttachment(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDocumentAttachment(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Attachment
	}{
		{name: "mp3", args: args{"https://listen.as/matt/abc.mp3"}, want: Attachment{
			Type:      "Document",
			URL:       "https://listen.as/matt/abc.mp3",
			MediaType: "audio/mpeg",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDocumentAttachment(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocumentAttachment() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
