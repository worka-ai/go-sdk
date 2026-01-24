package ui

import (
	"testing"
)

func renderSingle(kind WidgetKind) Widget {
	return Widget{ID: "node", Kind: kind}
}

func assertComponentKind(t *testing.T, widget Widget, expected string) {
	t.Helper()
	result := Render(widget)
	if len(result.Components) == 0 {
		t.Fatalf("no components rendered")
	}
	component := result.Components[0].Component
	switch expected {
	case "Text":
		if component.Text == nil {
			t.Fatalf("expected Text component")
		}
	case "Image":
		if component.Image == nil {
			t.Fatalf("expected Image component")
		}
	case "Icon":
		if component.Icon == nil {
			t.Fatalf("expected Icon component")
		}
	case "Divider":
		if component.Divider == nil {
			t.Fatalf("expected Divider component")
		}
	case "Row":
		if component.Row == nil {
			t.Fatalf("expected Row component")
		}
	case "Column":
		if component.Column == nil {
			t.Fatalf("expected Column component")
		}
	case "List":
		if component.List == nil {
			t.Fatalf("expected List component")
		}
	case "Button":
		if component.Button == nil {
			t.Fatalf("expected Button component")
		}
	case "TextField":
		if component.TextField == nil {
			t.Fatalf("expected TextField component")
		}
	case "CheckBox":
		if component.CheckBox == nil {
			t.Fatalf("expected CheckBox component")
		}
	case "Card":
		if component.Card == nil {
			t.Fatalf("expected Card component")
		}
	case "Modal":
		if component.Modal == nil {
			t.Fatalf("expected Modal component")
		}
	case "Tabs":
		if component.Tabs == nil {
			t.Fatalf("expected Tabs component")
		}
	case "MultipleChoice":
		if component.MultipleChoice == nil {
			t.Fatalf("expected MultipleChoice component")
		}
	case "Slider":
		if component.Slider == nil {
			t.Fatalf("expected Slider component")
		}
	case "DateTimeInput":
		if component.DateTimeInput == nil {
			t.Fatalf("expected DateTimeInput component")
		}
	case "AudioPlayer":
		if component.AudioPlayer == nil {
			t.Fatalf("expected AudioPlayer component")
		}
	case "Video":
		if component.Video == nil {
			t.Fatalf("expected Video component")
		}
	case "Timeline":
		if component.Timeline == nil {
			t.Fatalf("expected Timeline component")
		}
	case "TimelineItem":
		if component.TimelineItem == nil {
			t.Fatalf("expected TimelineItem component")
		}
	case "TimelineGroup":
		if component.TimelineGroup == nil {
			t.Fatalf("expected TimelineGroup component")
		}
	case "TimelineLane":
		if component.TimelineLane == nil {
			t.Fatalf("expected TimelineLane component")
		}
	default:
		t.Fatalf("unknown component %s", expected)
	}
}

func TestRenderWidgets(t *testing.T) {
	assertComponentKind(t, renderSingle(Text{Text: Str("Hello"), UsageHint: "h1"}), "Text")
	assertComponentKind(t, renderSingle(Image{URL: Str("https://example.com"), Fit: "cover"}), "Image")
	assertComponentKind(t, renderSingle(Icon{Name: Str("check")}), "Icon")
	assertComponentKind(t, renderSingle(Divider{Axis: "horizontal"}), "Divider")
	assertComponentKind(t, renderSingle(Row{Children: ChildrenItems()}), "Row")
	assertComponentKind(t, renderSingle(Column{Children: ChildrenItems()}), "Column")
	assertComponentKind(t, renderSingle(List{Children: ChildrenItems(), Direction: "vertical"}), "List")
	assertComponentKind(t, renderSingle(Button{
		Child:  Widget{Kind: Text{Text: Str("Click")}},
		Action: ActionCall("submit"),
	}), "Button")
	assertComponentKind(t, renderSingle(TextField{Text: ptrString(StrPath("/text"))}), "TextField")
	assertComponentKind(t, renderSingle(CheckBox{Label: Str("Agree"), Value: Bool(true)}), "CheckBox")
	assertComponentKind(t, renderSingle(Card{Child: Widget{Kind: Text{Text: Str("Card")}}}), "Card")
	assertComponentKind(t, renderSingle(Modal{
		EntryPoint: Widget{Kind: Text{Text: Str("Open")}},
		Content:    Widget{Kind: Text{Text: Str("Body")}},
	}), "Modal")
	assertComponentKind(t, renderSingle(Tabs{Items: []TabItem{{Title: Str("Tab"), Child: Widget{Kind: Text{Text: Str("Body")}}}}}), "Tabs")
	assertComponentKind(t, renderSingle(MultipleChoice{Selections: StringsPath("/choices")}), "MultipleChoice")
	assertComponentKind(t, renderSingle(Slider{Value: Num(1.0)}), "Slider")
	assertComponentKind(t, renderSingle(DateTimeInput{Value: Str("2024-01-01")}), "DateTimeInput")
	assertComponentKind(t, renderSingle(AudioPlayer{URL: Str("https://example.com/audio.mp3")}), "AudioPlayer")
	assertComponentKind(t, renderSingle(Video{URL: Str("https://example.com/video.mp4")}), "Video")
	assertComponentKind(t, renderSingle(Timeline{Children: ChildrenItems()}), "Timeline")
	assertComponentKind(t, renderSingle(TimelineItem{ItemID: "item-1"}), "TimelineItem")
	assertComponentKind(t, renderSingle(TimelineGroup{GroupID: "group-1", Children: ChildrenItems()}), "TimelineGroup")
	assertComponentKind(t, renderSingle(TimelineLane{LaneID: "lane-1", Children: ChildrenItems()}), "TimelineLane")
}

func ptrString(value StringValue) *StringValue { return &value }
