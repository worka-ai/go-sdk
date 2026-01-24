package a2ui

import (
	"encoding/json"
	"reflect"
	"testing"
)

func assertJSONEqual(t *testing.T, got any, want string) {
	t.Helper()
	gotBytes, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("marshal got: %v", err)
	}
	var gotObj any
	var wantObj any
	if err := json.Unmarshal(gotBytes, &gotObj); err != nil {
		t.Fatalf("unmarshal got: %v", err)
	}
	if err := json.Unmarshal([]byte(want), &wantObj); err != nil {
		t.Fatalf("unmarshal want: %v", err)
	}
	if !reflect.DeepEqual(gotObj, wantObj) {
		t.Fatalf("json mismatch\n got: %s\nwant: %s", string(gotBytes), want)
	}
}

func TestA2UIComponentSerialization(t *testing.T) {
	text := Entry("text", TextComponent(TextProps{
		Text:      StringLiteral("Hello"),
		UsageHint: strPtr("h1"),
	}))
	assertJSONEqual(t, text, `{"id":"text","component":{"Text":{"text":{"literalString":"Hello"},"usageHint":"h1"}}}`)

	image := Entry("img", ImageComponent(ImageProps{
		URL:       StringLiteral("https://example.com/logo.png"),
		Fit:       strPtr("cover"),
		UsageHint: strPtr("mediumFeature"),
	}))
	assertJSONEqual(t, image, `{"id":"img","component":{"Image":{"url":{"literalString":"https://example.com/logo.png"},"fit":"cover","usageHint":"mediumFeature"}}}`)

	icon := Entry("icon", IconComponent(IconProps{
		Name: StringLiteral("check"),
	}))
	assertJSONEqual(t, icon, `{"id":"icon","component":{"Icon":{"name":{"literalString":"check"}}}}`)

	divider := Entry("divider", DividerComponent(DividerProps{
		Axis: strPtr("horizontal"),
	}))
	assertJSONEqual(t, divider, `{"id":"divider","component":{"Divider":{"axis":"horizontal"}}}`)

	row := Entry("row", RowComponent(RowProps{
		Children:     ChildrenExplicit("a", "b"),
		Distribution: strPtr("spaceBetween"),
		Alignment:    strPtr("center"),
	}))
	assertJSONEqual(t, row, `{"id":"row","component":{"Row":{"children":{"explicitList":["a","b"]},"distribution":"spaceBetween","alignment":"center"}}}`)

	column := Entry("column", ColumnComponent(ColumnProps{
		Children:     ChildrenTemplate("item", "/items"),
		Distribution: strPtr("start"),
		Alignment:    strPtr("stretch"),
	}))
	assertJSONEqual(t, column, `{"id":"column","component":{"Column":{"children":{"template":{"componentId":"item","dataBinding":"/items"}},"distribution":"start","alignment":"stretch"}}}`)

	list := Entry("list", ListComponent(ListProps{
		Children:  ChildrenExplicit("one", "two"),
		Direction: strPtr("vertical"),
		Alignment: strPtr("start"),
	}))
	assertJSONEqual(t, list, `{"id":"list","component":{"List":{"children":{"explicitList":["one","two"]},"direction":"vertical","alignment":"start"}}}`)

	button := Entry("button", ButtonComponent(ButtonProps{
		Child:   "button-text",
		Action:  NewAction("submit"),
		Primary: boolPtr(true),
	}))
	assertJSONEqual(t, button, `{"id":"button","component":{"Button":{"child":"button-text","action":{"name":"submit"},"primary":true}}}`)

	textField := Entry("text-field", TextFieldComponent(TextFieldProps{
		Text:              refPtr(StringPath("/email")),
		Label:             refPtr(StringLiteral("Email")),
		TextFieldType:     strPtr("shortText"),
		ValidationRegexp:  strPtr("^.+@.+$"),
		OnSubmittedAction: actionPtr(NewAction("submit_form")),
	}))
	assertJSONEqual(t, textField, `{"id":"text-field","component":{"TextField":{"text":{"path":"/email"},"label":{"literalString":"Email"},"textFieldType":"shortText","validationRegexp":"^.+@.+$","onSubmittedAction":{"name":"submit_form"}}}}`)

	checkBox := Entry("check", CheckBoxComponent(CheckBoxProps{
		Label: StringLiteral("Agree"),
		Value: BoolPath("/agree"),
	}))
	assertJSONEqual(t, checkBox, `{"id":"check","component":{"CheckBox":{"label":{"literalString":"Agree"},"value":{"path":"/agree"}}}}`)

	card := Entry("card", CardComponent(CardProps{
		Child: "card-body",
	}))
	assertJSONEqual(t, card, `{"id":"card","component":{"Card":{"child":"card-body"}}}`)

	modal := Entry("modal", ModalComponent(ModalProps{
		EntryPointChild: "open-modal",
		ContentChild:    "modal-content",
	}))
	assertJSONEqual(t, modal, `{"id":"modal","component":{"Modal":{"entryPointChild":"open-modal","contentChild":"modal-content"}}}`)

	tabs := Entry("tabs", TabsComponent(TabsProps{
		TabItems: []TabItem{
			{Title: StringLiteral("One"), Child: "tab-one"},
			{Title: StringLiteral("Two"), Child: "tab-two"},
		},
	}))
	assertJSONEqual(t, tabs, `{"id":"tabs","component":{"Tabs":{"tabItems":[{"title":{"literalString":"One"},"child":"tab-one"},{"title":{"literalString":"Two"},"child":"tab-two"}]}}}`)

	maxSelections := int64(2)
	multipleChoice := Entry("choices", MultipleChoiceComponent(MultipleChoiceProps{
		Selections: StringArrayPath("/selected"),
		Options: []ChoiceOption{
			{Label: StringLiteral("A"), Value: "a"},
			{Label: StringLiteral("B"), Value: "b"},
		},
		MaxAllowedSelections: &maxSelections,
	}))
	assertJSONEqual(t, multipleChoice, `{"id":"choices","component":{"MultipleChoice":{"selections":{"path":"/selected"},"options":[{"label":{"literalString":"A"},"value":"a"},{"label":{"literalString":"B"},"value":"b"}],"maxAllowedSelections":2}}}`)

	slider := Entry("slider", SliderComponent(SliderProps{
		Value:    NumberPath("/volume"),
		MinValue: floatPtr(0),
		MaxValue: floatPtr(100),
	}))
	assertJSONEqual(t, slider, `{"id":"slider","component":{"Slider":{"value":{"path":"/volume"},"minValue":0,"maxValue":100}}}`)

	dateInput := Entry("date", DateTimeInputComponent(DateTimeInputProps{
		Value:      StringPath("/date"),
		EnableDate: boolPtr(true),
		EnableTime: boolPtr(false),
		FirstDate:  strPtr("2026-01-01"),
		LastDate:   strPtr("2026-12-31"),
	}))
	assertJSONEqual(t, dateInput, `{"id":"date","component":{"DateTimeInput":{"value":{"path":"/date"},"enableDate":true,"enableTime":false,"firstDate":"2026-01-01","lastDate":"2026-12-31"}}}`)

	audio := Entry("audio", AudioPlayerComponent(AudioPlayerProps{
		URL: StringLiteral("https://example.com/audio.mp3"),
	}))
	assertJSONEqual(t, audio, `{"id":"audio","component":{"AudioPlayer":{"url":{"literalString":"https://example.com/audio.mp3"}}}}`)

	video := Entry("video", VideoComponent(VideoProps{
		URL: StringLiteral("https://example.com/video.mp4"),
	}))
	assertJSONEqual(t, video, `{"id":"video","component":{"Video":{"url":{"literalString":"https://example.com/video.mp4"}}}}`)

	timeline := Entry("timeline", TimelineComponent(TimelineProps{
		Children:      ChildrenExplicit("item-1", "item-2"),
		Orientation:   strPtr("vertical"),
		Alignment:     strPtr("alternate"),
		AutoFollow:    boolRefPtr(BoolLiteral(true)),
		LaneMode:      strPtr("single"),
		CurrentItemID: refPtr(StringLiteral("item-2")),
	}))
	assertJSONEqual(t, timeline, `{"id":"timeline","component":{"Timeline":{"children":{"explicitList":["item-1","item-2"]},"orientation":"vertical","alignment":"alternate","autoFollow":{"literalBoolean":true},"laneMode":"single","currentItemId":{"literalString":"item-2"}}}}`)

	timelineItem := Entry("item-1", TimelineItemComponent(TimelineItemProps{
		ItemID:       "item-1",
		Title:        refPtr(StringLiteral("Step Started")),
		Subtitle:     refPtr(StringLiteral("Preparing data")),
		Timestamp:    refPtr(StringLiteral("2026-01-24T00:00:00Z")),
		Kind:         strPtr("step"),
		State:        strPtr("running"),
		Severity:     strPtr("info"),
		Icon:         refPtr(StringLiteral("bolt")),
		ContentChild: strPtr("detail"),
		Action:       actionPtr(NewAction("timeline.focus_item")),
	}))
	assertJSONEqual(t, timelineItem, `{"id":"item-1","component":{"TimelineItem":{"itemId":"item-1","title":{"literalString":"Step Started"},"subtitle":{"literalString":"Preparing data"},"timestamp":{"literalString":"2026-01-24T00:00:00Z"},"kind":"step","state":"running","severity":"info","icon":{"literalString":"bolt"},"contentChild":"detail","action":{"name":"timeline.focus_item"}}}}`)

	timelineGroup := Entry("group-1", TimelineGroupComponent(TimelineGroupProps{
		GroupID:    "group-1",
		Title:      refPtr(StringLiteral("Task Group")),
		Summary:    refPtr(StringLiteral("2 tasks")),
		Children:   ChildrenExplicit("item-1"),
		Collapsed:  boolRefPtr(BoolLiteral(false)),
		BadgeCount: numberRefPtr(NumberLiteral(2)),
		GroupState: strPtr("running"),
	}))
	assertJSONEqual(t, timelineGroup, `{"id":"group-1","component":{"TimelineGroup":{"groupId":"group-1","title":{"literalString":"Task Group"},"summary":{"literalString":"2 tasks"},"children":{"explicitList":["item-1"]},"collapsed":{"literalBoolean":false},"badgeCount":{"literalNumber":2},"groupState":"running"}}}}`)

	timelineLane := Entry("lane-1", TimelineLaneComponent(TimelineLaneProps{
		LaneID:   "lane-1",
		Title:    refPtr(StringLiteral("Lane A")),
		Children: ChildrenExplicit("item-1"),
	}))
	assertJSONEqual(t, timelineLane, `{"id":"lane-1","component":{"TimelineLane":{"laneId":"lane-1","title":{"literalString":"Lane A"},"children":{"explicitList":["item-1"]}}}}`)
}

func TestA2uiMessages(t *testing.T) {
	surface := SurfaceUpdate{SurfaceID: "surface-1", Components: []ComponentEntry{
		Entry("text", TextComponent(TextProps{Text: StringLiteral("Hello")})),
	}}
	msg := A2uiMessage{SurfaceUpdate: &surface}
	assertJSONEqual(t, msg, `{"surfaceUpdate":{"surfaceId":"surface-1","components":[{"id":"text","component":{"Text":{"text":{"literalString":"Hello"}}}}]}}`)

	data := json.RawMessage(`{"hello":"world"}`)
	update := A2uiMessage{DataModelUpdate: &DataModelUpdate{SurfaceID: "surface-2", Contents: data}}
	assertJSONEqual(t, update, `{"dataModelUpdate":{"surfaceId":"surface-2","contents":{"hello":"world"}}}`)
}

func strPtr(v string) *string     { return &v }
func boolPtr(v bool) *bool        { return &v }
func floatPtr(v float64) *float64 { return &v }

func refPtr(ref StringRef) *StringRef       { return &ref }
func actionPtr(action Action) *Action       { return &action }
func boolRefPtr(ref BoolRef) *BoolRef       { return &ref }
func numberRefPtr(ref NumberRef) *NumberRef { return &ref }
