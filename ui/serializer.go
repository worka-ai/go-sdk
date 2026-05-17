package ui

import "github.com/modelcontextprotocol/go-sdk/uiwire"

type RenderResult struct {
	Root       string
	Components []uiwire.ComponentEntry
}

func Render(root Widget) RenderResult {
	s := &serializer{counter: 0}
	rootID := s.renderWidget(root)
	return RenderResult{Root: rootID, Components: s.components}
}

type serializer struct {
	counter    int
	components []uiwire.ComponentEntry
}

func (s *serializer) nextID() string {
	s.counter++
	return "ui-" + itoa(s.counter)
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	buf := make([]byte, 0, 20)
	for value > 0 {
		buf = append(buf, byte('0'+value%10))
		value /= 10
	}
	for i := 0; i < len(buf)/2; i++ {
		j := len(buf) - 1 - i
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func (s *serializer) renderChildren(children Children) uiwire.ComponentChildren {
	if children.Template != nil {
		templateID := s.renderWidget(children.Template.Template)
		return uiwire.ChildrenTemplate(templateID, children.Template.DataBinding)
	}
	ids := make([]string, 0, len(children.Items))
	for _, child := range children.Items {
		ids = append(ids, s.renderWidget(child))
	}
	return uiwire.ChildrenExplicit(ids...)
}

func (s *serializer) renderWidget(widget Widget) string {
	id := widget.ID
	if id == "" {
		id = s.nextID()
	}
	component := s.renderKind(widget.Kind)
	entry := uiwire.Entry(id, component)
	if widget.Weight != nil {
		entry = entry.WithWeight(*widget.Weight)
	}
	s.components = append(s.components, entry)
	return id
}

func (s *serializer) renderKind(kind WidgetKind) uiwire.Component {
	switch v := kind.(type) {
	case Text:
		props := uiwire.TextProps{Text: toStringRef(v.Text)}
		if v.UsageHint != "" {
			props.UsageHint = v.UsageHint
		}
		return uiwire.TextComponent(props)
	case Image:
		props := uiwire.ImageProps{URL: toStringRef(v.URL)}
		if v.Fit != "" {
			props.Fit = v.Fit
		}
		if v.UsageHint != "" {
			props.UsageHint = v.UsageHint
		}
		return uiwire.ImageComponent(props)
	case Icon:
		return uiwire.IconComponent(uiwire.IconProps{Name: toStringRef(v.Name)})
	case Divider:
		props := uiwire.DividerProps{}
		if v.Axis != "" {
			props.Axis = v.Axis
		}
		return uiwire.DividerComponent(props)
	case Row:
		props := uiwire.RowProps{Children: s.renderChildren(v.Children)}
		if v.Distribution != "" {
			props.Distribution = v.Distribution
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return uiwire.RowComponent(props)
	case Column:
		props := uiwire.ColumnProps{Children: s.renderChildren(v.Children)}
		if v.Distribution != "" {
			props.Distribution = v.Distribution
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return uiwire.ColumnComponent(props)
	case List:
		props := uiwire.ListProps{Children: s.renderChildren(v.Children)}
		if v.Direction != "" {
			props.Direction = v.Direction
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return uiwire.ListComponent(props)
	case Button:
		childID := s.renderWidget(v.Child)
		props := uiwire.ButtonProps{Child: childID, Action: toAction(v.Action)}
		if v.Primary != nil {
			props.Primary = v.Primary
		}
		return uiwire.ButtonComponent(props)
	case TextField:
		props := uiwire.TextFieldProps{}
		if v.Text != nil {
			props.Text = ptrStringRef(toStringRef(*v.Text))
		}
		if v.Label != nil {
			props.Label = ptrStringRef(toStringRef(*v.Label))
		}
		if v.TextFieldType != "" {
			props.TextFieldType = v.TextFieldType
		}
		if v.ValidationRegexp != "" {
			props.ValidationRegexp = v.ValidationRegexp
		}
		if v.OnSubmittedAction != nil {
			props.OnSubmittedAction = ptrAction(toAction(*v.OnSubmittedAction))
		}
		return uiwire.TextFieldComponent(props)
	case CheckBox:
		return uiwire.CheckBoxComponent(uiwire.CheckBoxProps{
			Label: toStringRef(v.Label),
			Value: toBoolRef(v.Value),
		})
	case Card:
		childID := s.renderWidget(v.Child)
		return uiwire.CardComponent(uiwire.CardProps{Child: childID})
	case Modal:
		entryID := s.renderWidget(v.EntryPoint)
		contentID := s.renderWidget(v.Content)
		return uiwire.ModalComponent(uiwire.ModalProps{
			EntryPointChild: entryID,
			ContentChild:    contentID,
		})
	case Tabs:
		items := make([]uiwire.TabItemProps, 0, len(v.Items))
		for _, item := range v.Items {
			childID := s.renderWidget(item.Child)
			items = append(items, uiwire.TabItemProps{Title: toStringRef(item.Title), Child: childID})
		}
		return uiwire.TabsComponent(uiwire.TabsProps{TabItems: items})
	case MultipleChoice:
		opts := make([]uiwire.ChoiceOptionProps, 0, len(v.Options))
		for _, option := range v.Options {
			opts = append(opts, uiwire.ChoiceOptionProps{Label: toStringRef(option.Label), Value: option.Value})
		}
		props := uiwire.MultipleChoiceProps{Selections: toStringArrayRef(v.Selections), Options: opts}
		if v.MaxAllowedSelections != nil {
			props.MaxAllowedSelections = v.MaxAllowedSelections
		}
		return uiwire.MultipleChoiceComponent(props)
	case Slider:
		props := uiwire.SliderProps{Value: toNumberRef(v.Value)}
		props.MinValue = v.MinValue
		props.MaxValue = v.MaxValue
		return uiwire.SliderComponent(props)
	case DateTimeInput:
		props := uiwire.DateTimeInputProps{Value: toStringRef(v.Value)}
		if v.EnableDate != nil {
			props.EnableDate = v.EnableDate
		}
		if v.EnableTime != nil {
			props.EnableTime = v.EnableTime
		}
		if v.FirstDate != "" {
			props.FirstDate = v.FirstDate
		}
		if v.LastDate != "" {
			props.LastDate = v.LastDate
		}
		return uiwire.DateTimeInputComponent(props)
	case AudioPlayer:
		return uiwire.AudioPlayerComponent(uiwire.AudioPlayerProps{URL: toStringRef(v.URL)})
	case Video:
		return uiwire.VideoComponent(uiwire.VideoProps{URL: toStringRef(v.URL)})
	case Timeline:
		props := uiwire.TimelineProps{Children: s.renderChildren(v.Children)}
		if v.Orientation != "" {
			props.Orientation = v.Orientation
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		if v.AutoFollow != nil {
			props.AutoFollow = ptrBoolRef(toBoolRef(*v.AutoFollow))
		}
		if v.LaneMode != "" {
			props.LaneMode = v.LaneMode
		}
		if v.CurrentItemID != nil {
			props.CurrentItemID = ptrStringRef(toStringRef(*v.CurrentItemID))
		}
		return uiwire.TimelineComponent(props)
	case TimelineItem:
		props := uiwire.TimelineItemProps{}
		if v.ItemID != "" {
			props.ItemID = v.ItemID
		}
		if v.Title != nil {
			props.Title = ptrStringRef(toStringRef(*v.Title))
		}
		if v.Subtitle != nil {
			props.Subtitle = ptrStringRef(toStringRef(*v.Subtitle))
		}
		if v.Timestamp != nil {
			props.Timestamp = ptrStringRef(toStringRef(*v.Timestamp))
		}
		if v.Kind != "" {
			props.Kind = v.Kind
		}
		if v.State != "" {
			props.State = v.State
		}
		if v.Severity != "" {
			props.Severity = v.Severity
		}
		if v.Icon != nil {
			props.Icon = ptrStringRef(toStringRef(*v.Icon))
		}
		if v.Content != nil {
			props.ContentChild = s.renderWidget(*v.Content)
		}
		if v.Action != nil {
			props.Action = ptrAction(toAction(*v.Action))
		}
		return uiwire.TimelineItemComponent(props)
	case TimelineGroup:
		props := uiwire.TimelineGroupProps{
			GroupID:  v.GroupID,
			Children: s.renderChildren(v.Children),
		}
		if v.Title != nil {
			props.Title = ptrStringRef(toStringRef(*v.Title))
		}
		if v.Summary != nil {
			props.Summary = ptrStringRef(toStringRef(*v.Summary))
		}
		if v.Collapsed != nil {
			props.Collapsed = ptrBoolRef(toBoolRef(*v.Collapsed))
		}
		if v.BadgeCount != nil {
			props.BadgeCount = ptrNumberRef(toNumberRef(*v.BadgeCount))
		}
		if v.GroupState != "" {
			props.GroupState = v.GroupState
		}
		return uiwire.TimelineGroupComponent(props)
	case TimelineLane:
		props := uiwire.TimelineLaneProps{LaneID: v.LaneID, Children: s.renderChildren(v.Children)}
		if v.Title != nil {
			props.Title = ptrStringRef(toStringRef(*v.Title))
		}
		return uiwire.TimelineLaneComponent(props)
	default:
		return uiwire.Component{}
	}
}

func toStringRef(value StringValue) uiwire.StringRef {
	if value.path != nil {
		return uiwire.StringRef{Path: *value.path}
	}
	return uiwire.StringRef{LiteralString: value.literal}
}

func toNumberRef(value NumberValue) uiwire.NumberRef {
	if value.path != nil {
		return uiwire.NumberRef{Path: *value.path}
	}
	return uiwire.NumberRef{LiteralNumber: value.literal}
}

func toBoolRef(value BoolValue) uiwire.BoolRef {
	if value.path != nil {
		return uiwire.BoolRef{Path: *value.path}
	}
	return uiwire.BoolRef{LiteralBoolean: value.literal}
}

func toStringArrayRef(value StringArrayValue) uiwire.StringArrayRef {
	if value.path != nil {
		return uiwire.StringArrayRef{Path: *value.path}
	}
	return uiwire.StringArrayRef{LiteralArray: value.literal}
}

func toAction(action Action) uiwire.Action {
	if len(action.Context) == 0 {
		return uiwire.NewAction(action.Name)
	}
	entries := make([]uiwire.ActionContextEntry, 0, len(action.Context))
	for key, value := range action.Context {
		entries = append(entries, uiwire.ActionContextEntry{
			Key:   key,
			Value: toActionValue(value),
		})
	}
	return uiwire.ActionWithContext(action.Name, entries...)
}

func toActionValue(value Value) uiwire.ActionValue {
	switch value.Kind {
	case ValuePath:
		return uiwire.ActionValue{Path: value.Path}
	case ValueString:
		return uiwire.ActionValue{LiteralString: &value.String}
	case ValueNumber:
		return uiwire.ActionValue{LiteralNumber: &value.Number}
	case ValueBool:
		return uiwire.ActionValue{LiteralBoolean: &value.Bool}
	default:
		return uiwire.ActionValue{}
	}
}

func ptrStringRef(value uiwire.StringRef) *uiwire.StringRef { return &value }
func ptrNumberRef(value uiwire.NumberRef) *uiwire.NumberRef { return &value }
func ptrBoolRef(value uiwire.BoolRef) *uiwire.BoolRef       { return &value }
func ptrAction(value uiwire.Action) *uiwire.Action          { return &value }
