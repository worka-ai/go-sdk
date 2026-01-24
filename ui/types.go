package ui

type StringValue struct {
	literal *string
	path    *string
}

func Str(value string) StringValue { return StringValue{literal: &value} }
func StrPath(path string) StringValue {
	return StringValue{path: &path}
}

type NumberValue struct {
	literal *float64
	path    *string
}

func Num(value float64) NumberValue { return NumberValue{literal: &value} }
func NumPath(path string) NumberValue {
	return NumberValue{path: &path}
}

type BoolValue struct {
	literal *bool
	path    *string
}

func Bool(value bool) BoolValue { return BoolValue{literal: &value} }
func BoolPath(path string) BoolValue {
	return BoolValue{path: &path}
}

type StringArrayValue struct {
	literal []string
	path    *string
}

func Strings(values ...string) StringArrayValue { return StringArrayValue{literal: values} }
func StringsPath(path string) StringArrayValue {
	return StringArrayValue{path: &path}
}

type ValueKind int

const (
	ValuePath ValueKind = iota
	ValueString
	ValueNumber
	ValueBool
)

type Value struct {
	Kind   ValueKind
	Path   string
	String string
	Number float64
	Bool   bool
}

func PathValue(path string) Value {
	return Value{Kind: ValuePath, Path: path}
}
func StringValueLiteral(value string) Value {
	return Value{Kind: ValueString, String: value}
}
func NumberValueLiteral(value float64) Value {
	return Value{Kind: ValueNumber, Number: value}
}
func BoolValueLiteral(value bool) Value {
	return Value{Kind: ValueBool, Bool: value}
}

type Action struct {
	Name    string
	Context map[string]Value
}

func ActionCall(name string) Action { return Action{Name: name} }
func ActionWithContext(name string, ctx map[string]Value) Action {
	return Action{Name: name, Context: ctx}
}

type Children struct {
	Items    []Widget
	Template *TemplateChildren
}

type TemplateChildren struct {
	DataBinding string
	Template    Widget
}

func ChildrenItems(items ...Widget) Children { return Children{Items: items} }
func ChildrenTemplate(dataBinding string, template Widget) Children {
	return Children{Template: &TemplateChildren{DataBinding: dataBinding, Template: template}}
}

type Widget struct {
	ID     string
	Weight *float64
	Kind   WidgetKind
}

func (w Widget) WithID(id string) Widget {
	w.ID = id
	return w
}

func (w Widget) WithWeight(weight float64) Widget {
	w.Weight = &weight
	return w
}

type WidgetKind interface {
	kindName() string
}

type Text struct {
	Text      StringValue
	UsageHint string
}

func (Text) kindName() string { return "Text" }

type Image struct {
	URL       StringValue
	Fit       string
	UsageHint string
}

func (Image) kindName() string { return "Image" }

type Icon struct {
	Name StringValue
}

func (Icon) kindName() string { return "Icon" }

type Divider struct {
	Axis string
}

func (Divider) kindName() string { return "Divider" }

type Row struct {
	Children     Children
	Distribution string
	Alignment    string
}

func (Row) kindName() string { return "Row" }

type Column struct {
	Children     Children
	Distribution string
	Alignment    string
}

func (Column) kindName() string { return "Column" }

type List struct {
	Children  Children
	Direction string
	Alignment string
}

func (List) kindName() string { return "List" }

type Button struct {
	Child   Widget
	Action  Action
	Primary *bool
}

func (Button) kindName() string { return "Button" }

type TextField struct {
	Text              *StringValue
	Label             *StringValue
	TextFieldType     string
	ValidationRegexp  string
	OnSubmittedAction *Action
}

func (TextField) kindName() string { return "TextField" }

type CheckBox struct {
	Label StringValue
	Value BoolValue
}

func (CheckBox) kindName() string { return "CheckBox" }

type Card struct {
	Child Widget
}

func (Card) kindName() string { return "Card" }

type Modal struct {
	EntryPoint Widget
	Content    Widget
}

func (Modal) kindName() string { return "Modal" }

type TabItem struct {
	Title StringValue
	Child Widget
}

type Tabs struct {
	Items []TabItem
}

func (Tabs) kindName() string { return "Tabs" }

type ChoiceOption struct {
	Label StringValue
	Value string
}

type MultipleChoice struct {
	Selections           StringArrayValue
	Options              []ChoiceOption
	MaxAllowedSelections *int64
}

func (MultipleChoice) kindName() string { return "MultipleChoice" }

type Slider struct {
	Value    NumberValue
	MinValue *float64
	MaxValue *float64
}

func (Slider) kindName() string { return "Slider" }

type DateTimeInput struct {
	Value      StringValue
	EnableDate *bool
	EnableTime *bool
	FirstDate  string
	LastDate   string
}

func (DateTimeInput) kindName() string { return "DateTimeInput" }

type AudioPlayer struct {
	URL StringValue
}

func (AudioPlayer) kindName() string { return "AudioPlayer" }

type Video struct {
	URL StringValue
}

func (Video) kindName() string { return "Video" }

type Timeline struct {
	Children      Children
	Orientation   string
	Alignment     string
	AutoFollow    *BoolValue
	LaneMode      string
	CurrentItemID *StringValue
}

func (Timeline) kindName() string { return "Timeline" }

type TimelineItem struct {
	ItemID    string
	Title     *StringValue
	Subtitle  *StringValue
	Timestamp *StringValue
	Kind      string
	State     string
	Severity  string
	Icon      *StringValue
	Content   *Widget
	Action    *Action
}

func (TimelineItem) kindName() string { return "TimelineItem" }

type TimelineGroup struct {
	GroupID    string
	Title      *StringValue
	Summary    *StringValue
	Children   Children
	Collapsed  *BoolValue
	BadgeCount *NumberValue
	GroupState string
}

func (TimelineGroup) kindName() string { return "TimelineGroup" }

type TimelineLane struct {
	LaneID   string
	Title    *StringValue
	Children Children
}

func (TimelineLane) kindName() string { return "TimelineLane" }
