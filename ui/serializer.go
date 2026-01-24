package ui

import "github.com/modelcontextprotocol/go-sdk/a2ui"

type RenderResult struct {
	Root       string
	Components []a2ui.ComponentEntry
}

func Render(root Widget) RenderResult {
	s := &serializer{counter: 0}
	rootID := s.renderWidget(root)
	return RenderResult{Root: rootID, Components: s.components}
}

type serializer struct {
	counter    int
	components []a2ui.ComponentEntry
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

func (s *serializer) renderChildren(children Children) a2ui.ComponentChildren {
	if children.Template != nil {
		templateID := s.renderWidget(children.Template.Template)
		return a2ui.ChildrenTemplate(templateID, children.Template.DataBinding)
	}
	ids := make([]string, 0, len(children.Items))
	for _, child := range children.Items {
		ids = append(ids, s.renderWidget(child))
	}
	return a2ui.ChildrenExplicit(ids...)
}

func (s *serializer) renderWidget(widget Widget) string {
	id := widget.ID
	if id == "" {
		id = s.nextID()
	}
	component := s.renderKind(widget.Kind)
	entry := a2ui.Entry(id, component)
	if widget.Weight != nil {
		entry = entry.WithWeight(*widget.Weight)
	}
	s.components = append(s.components, entry)
	return id
}

func (s *serializer) renderKind(kind WidgetKind) a2ui.Component {
	switch v := kind.(type) {
	case Text:
		props := a2ui.TextProps{Text: toStringRef(v.Text)}
		if v.UsageHint != "" {
			props.UsageHint = v.UsageHint
		}
		return a2ui.TextComponent(props)
	case Image:
		props := a2ui.ImageProps{URL: toStringRef(v.URL)}
		if v.Fit != "" {
			props.Fit = v.Fit
		}
		if v.UsageHint != "" {
			props.UsageHint = v.UsageHint
		}
		return a2ui.ImageComponent(props)
	case Icon:
		return a2ui.IconComponent(a2ui.IconProps{Name: toStringRef(v.Name)})
	case Divider:
		props := a2ui.DividerProps{}
		if v.Axis != "" {
			props.Axis = v.Axis
		}
		return a2ui.DividerComponent(props)
	case Row:
		props := a2ui.RowProps{Children: s.renderChildren(v.Children)}
		if v.Distribution != "" {
			props.Distribution = v.Distribution
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return a2ui.RowComponent(props)
	case Column:
		props := a2ui.ColumnProps{Children: s.renderChildren(v.Children)}
		if v.Distribution != "" {
			props.Distribution = v.Distribution
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return a2ui.ColumnComponent(props)
	case List:
		props := a2ui.ListProps{Children: s.renderChildren(v.Children)}
		if v.Direction != "" {
			props.Direction = v.Direction
		}
		if v.Alignment != "" {
			props.Alignment = v.Alignment
		}
		return a2ui.ListComponent(props)
	case Button:
		childID := s.renderWidget(v.Child)
		props := a2ui.ButtonProps{Child: childID, Action: toAction(v.Action)}
		if v.Primary != nil {
			props.Primary = v.Primary
		}
		return a2ui.ButtonComponent(props)
	case TextField:
		props := a2ui.TextFieldProps{}
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
		return a2ui.TextFieldComponent(props)
	case CheckBox:
		return a2ui.CheckBoxComponent(a2ui.CheckBoxProps{
			Label: toStringRef(v.Label),
			Value: toBoolRef(v.Value),
		})
	case Card:
		childID := s.renderWidget(v.Child)
		return a2ui.CardComponent(a2ui.CardProps{Child: childID})
	case Modal:
		entryID := s.renderWidget(v.EntryPoint)
		contentID := s.renderWidget(v.Content)
		return a2ui.ModalComponent(a2ui.ModalProps{
			EntryPointChild: entryID,
			ContentChild:    contentID,
		})
	case Tabs:
		items := make([]a2ui.TabItemProps, 0, len(v.Items))
		for _, item := range v.Items {
			childID := s.renderWidget(item.Child)
			items = append(items, a2ui.TabItemProps{Title: toStringRef(item.Title), Child: childID})
		}
		return a2ui.TabsComponent(a2ui.TabsProps{TabItems: items})
	case MultipleChoice:
		opts := make([]a2ui.ChoiceOptionProps, 0, len(v.Options))
		for _, option := range v.Options {
			opts = append(opts, a2ui.ChoiceOptionProps{Label: toStringRef(option.Label), Value: option.Value})
		}
		props := a2ui.MultipleChoiceProps{Selections: toStringArrayRef(v.Selections), Options: opts}
		if v.MaxAllowedSelections != nil {
			props.MaxAllowedSelections = v.MaxAllowedSelections
		}
		return a2ui.MultipleChoiceComponent(props)
	case Slider:
		props := a2ui.SliderProps{Value: toNumberRef(v.Value)}
		props.MinValue = v.MinValue
		props.MaxValue = v.MaxValue
		return a2ui.SliderComponent(props)
	case DateTimeInput:
		props := a2ui.DateTimeInputProps{Value: toStringRef(v.Value)}
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
		return a2ui.DateTimeInputComponent(props)
	case AudioPlayer:
		return a2ui.AudioPlayerComponent(a2ui.AudioPlayerProps{URL: toStringRef(v.URL)})
	case Video:
		return a2ui.VideoComponent(a2ui.VideoProps{URL: toStringRef(v.URL)})
	case Timeline:
		props := a2ui.TimelineProps{Children: s.renderChildren(v.Children)}
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
		return a2ui.TimelineComponent(props)
	case TimelineItem:
		props := a2ui.TimelineItemProps{}
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
		return a2ui.TimelineItemComponent(props)
	case TimelineGroup:
		props := a2ui.TimelineGroupProps{
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
		return a2ui.TimelineGroupComponent(props)
	case TimelineLane:
		props := a2ui.TimelineLaneProps{LaneID: v.LaneID, Children: s.renderChildren(v.Children)}
		if v.Title != nil {
			props.Title = ptrStringRef(toStringRef(*v.Title))
		}
		return a2ui.TimelineLaneComponent(props)
	default:
		return a2ui.Component{}
	}
}

func toStringRef(value StringValue) a2ui.StringRef {
	if value.path != nil {
		return a2ui.StringRef{Path: *value.path}
	}
	return a2ui.StringRef{LiteralString: value.literal}
}

func toNumberRef(value NumberValue) a2ui.NumberRef {
	if value.path != nil {
		return a2ui.NumberRef{Path: *value.path}
	}
	return a2ui.NumberRef{LiteralNumber: value.literal}
}

func toBoolRef(value BoolValue) a2ui.BoolRef {
	if value.path != nil {
		return a2ui.BoolRef{Path: *value.path}
	}
	return a2ui.BoolRef{LiteralBoolean: value.literal}
}

func toStringArrayRef(value StringArrayValue) a2ui.StringArrayRef {
	if value.path != nil {
		return a2ui.StringArrayRef{Path: *value.path}
	}
	return a2ui.StringArrayRef{LiteralArray: value.literal}
}

func toAction(action Action) a2ui.Action {
	if len(action.Context) == 0 {
		return a2ui.NewAction(action.Name)
	}
	entries := make([]a2ui.ActionContextEntry, 0, len(action.Context))
	for key, value := range action.Context {
		entries = append(entries, a2ui.ActionContextEntry{
			Key:   key,
			Value: toActionValue(value),
		})
	}
	return a2ui.ActionWithContext(action.Name, entries...)
}

func toActionValue(value Value) a2ui.ActionValue {
	switch value.Kind {
	case ValuePath:
		return a2ui.ActionValue{Path: value.Path}
	case ValueString:
		return a2ui.ActionValue{LiteralString: &value.String}
	case ValueNumber:
		return a2ui.ActionValue{LiteralNumber: &value.Number}
	case ValueBool:
		return a2ui.ActionValue{LiteralBoolean: &value.Bool}
	default:
		return a2ui.ActionValue{}
	}
}

func ptrStringRef(value a2ui.StringRef) *a2ui.StringRef { return &value }
func ptrNumberRef(value a2ui.NumberRef) *a2ui.NumberRef { return &value }
func ptrBoolRef(value a2ui.BoolRef) *a2ui.BoolRef       { return &value }
func ptrAction(value a2ui.Action) *a2ui.Action          { return &value }
