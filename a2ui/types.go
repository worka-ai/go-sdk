// Package a2ui provides typed builders for the default A2UI catalog.
//
// It mirrors the default A2UI widgets supported by GenUI:
// AudioPlayer, Button, Card, CheckBox, Column, DateTimeInput, Divider, Icon,
// Image, List, Modal, MultipleChoice, Row, Slider, Tabs, TextField, Text, Video.
package a2ui

import "encoding/json"

// StringRef represents a bound or literal string value.
type StringRef struct {
	Path          *string `json:"path,omitempty"`
	LiteralString *string `json:"literalString,omitempty"`
}

func StringLiteral(value string) StringRef { return StringRef{LiteralString: &value} }
func StringPath(path string) StringRef     { return StringRef{Path: &path} }

// NumberRef represents a bound or literal numeric value.
type NumberRef struct {
	Path          *string  `json:"path,omitempty"`
	LiteralNumber *float64 `json:"literalNumber,omitempty"`
}

func NumberLiteral(value float64) NumberRef { return NumberRef{LiteralNumber: &value} }
func NumberPath(path string) NumberRef      { return NumberRef{Path: &path} }

// BoolRef represents a bound or literal boolean value.
type BoolRef struct {
	Path           *string `json:"path,omitempty"`
	LiteralBoolean *bool   `json:"literalBoolean,omitempty"`
}

func BoolLiteral(value bool) BoolRef { return BoolRef{LiteralBoolean: &value} }
func BoolPath(path string) BoolRef   { return BoolRef{Path: &path} }

// StringArrayRef represents a bound or literal string array.
type StringArrayRef struct {
	Path         *string  `json:"path,omitempty"`
	LiteralArray []string `json:"literalArray,omitempty"`
}

func StringArrayLiteral(values ...string) StringArrayRef { return StringArrayRef{LiteralArray: values} }
func StringArrayPath(path string) StringArrayRef         { return StringArrayRef{Path: &path} }

// ActionValue represents a bound or literal action context value.
type ActionValue struct {
	Path           *string  `json:"path,omitempty"`
	LiteralString  *string  `json:"literalString,omitempty"`
	LiteralNumber  *float64 `json:"literalNumber,omitempty"`
	LiteralBoolean *bool    `json:"literalBoolean,omitempty"`
}

func ActionValuePath(path string) ActionValue     { return ActionValue{Path: &path} }
func ActionValueString(value string) ActionValue  { return ActionValue{LiteralString: &value} }
func ActionValueNumber(value float64) ActionValue { return ActionValue{LiteralNumber: &value} }
func ActionValueBoolean(value bool) ActionValue   { return ActionValue{LiteralBoolean: &value} }

// ActionContextEntry represents a key/value entry for an action context.
type ActionContextEntry struct {
	Key   string      `json:"key"`
	Value ActionValue `json:"value"`
}

// Action represents a UI action binding.
type Action struct {
	Name    string               `json:"name"`
	Context []ActionContextEntry `json:"context,omitempty"`
}

func NewAction(name string) Action { return Action{Name: name} }
func ActionWithContext(name string, ctx ...ActionContextEntry) Action {
	return Action{Name: name, Context: ctx}
}

// ComponentTemplate defines a templated child list binding.
type ComponentTemplate struct {
	ComponentID string `json:"componentId"`
	DataBinding string `json:"dataBinding"`
}

// ComponentChildren specifies explicit children or a template.
type ComponentChildren struct {
	ExplicitList []string           `json:"explicitList,omitempty"`
	Template     *ComponentTemplate `json:"template,omitempty"`
}

func ChildrenExplicit(ids ...string) ComponentChildren { return ComponentChildren{ExplicitList: ids} }
func ChildrenTemplate(componentID, dataBinding string) ComponentChildren {
	return ComponentChildren{Template: &ComponentTemplate{ComponentID: componentID, DataBinding: dataBinding}}
}

// ComponentEntry is a component instance entry.
type ComponentEntry struct {
	ID        string    `json:"id"`
	Weight    *float64  `json:"weight,omitempty"`
	Component Component `json:"component"`
}

func Entry(id string, component Component) ComponentEntry {
	return ComponentEntry{ID: id, Component: component}
}
func (e ComponentEntry) WithWeight(weight float64) ComponentEntry {
	e.Weight = &weight
	return e
}

// Component holds one concrete widget payload.
type Component struct {
	AudioPlayer    *AudioPlayerProps    `json:"AudioPlayer,omitempty"`
	Button         *ButtonProps         `json:"Button,omitempty"`
	Card           *CardProps           `json:"Card,omitempty"`
	CheckBox       *CheckBoxProps       `json:"CheckBox,omitempty"`
	Column         *ColumnProps         `json:"Column,omitempty"`
	DateTimeInput  *DateTimeInputProps  `json:"DateTimeInput,omitempty"`
	Divider        *DividerProps        `json:"Divider,omitempty"`
	Icon           *IconProps           `json:"Icon,omitempty"`
	Image          *ImageProps          `json:"Image,omitempty"`
	List           *ListProps           `json:"List,omitempty"`
	Modal          *ModalProps          `json:"Modal,omitempty"`
	MultipleChoice *MultipleChoiceProps `json:"MultipleChoice,omitempty"`
	Row            *RowProps            `json:"Row,omitempty"`
	Slider         *SliderProps         `json:"Slider,omitempty"`
	Tabs           *TabsProps           `json:"Tabs,omitempty"`
	TextField      *TextFieldProps      `json:"TextField,omitempty"`
	Text           *TextProps           `json:"Text,omitempty"`
	Video          *VideoProps          `json:"Video,omitempty"`
}

func AudioPlayerComponent(props AudioPlayerProps) Component { return Component{AudioPlayer: &props} }
func ButtonComponent(props ButtonProps) Component           { return Component{Button: &props} }
func CardComponent(props CardProps) Component               { return Component{Card: &props} }
func CheckBoxComponent(props CheckBoxProps) Component       { return Component{CheckBox: &props} }
func ColumnComponent(props ColumnProps) Component           { return Component{Column: &props} }
func DateTimeInputComponent(props DateTimeInputProps) Component {
	return Component{DateTimeInput: &props}
}
func DividerComponent(props DividerProps) Component { return Component{Divider: &props} }
func IconComponent(props IconProps) Component       { return Component{Icon: &props} }
func ImageComponent(props ImageProps) Component     { return Component{Image: &props} }
func ListComponent(props ListProps) Component       { return Component{List: &props} }
func ModalComponent(props ModalProps) Component     { return Component{Modal: &props} }
func MultipleChoiceComponent(props MultipleChoiceProps) Component {
	return Component{MultipleChoice: &props}
}
func RowComponent(props RowProps) Component       { return Component{Row: &props} }
func SliderComponent(props SliderProps) Component { return Component{Slider: &props} }
func TabsComponent(props TabsProps) Component     { return Component{Tabs: &props} }
func TextFieldComponent(props TextFieldProps) Component {
	return Component{TextField: &props}
}
func TextComponent(props TextProps) Component   { return Component{Text: &props} }
func VideoComponent(props VideoProps) Component { return Component{Video: &props} }

// Widget property definitions.
type RowProps struct {
	Children     ComponentChildren `json:"children"`
	Distribution *string           `json:"distribution,omitempty"`
	Alignment    *string           `json:"alignment,omitempty"`
}

type ColumnProps struct {
	Children     ComponentChildren `json:"children"`
	Distribution *string           `json:"distribution,omitempty"`
	Alignment    *string           `json:"alignment,omitempty"`
}

type ListProps struct {
	Children  ComponentChildren `json:"children"`
	Direction *string           `json:"direction,omitempty"`
	Alignment *string           `json:"alignment,omitempty"`
}

type TextProps struct {
	Text      StringRef `json:"text"`
	UsageHint *string   `json:"usageHint,omitempty"`
}

type ImageProps struct {
	URL       StringRef `json:"url"`
	Fit       *string   `json:"fit,omitempty"`
	UsageHint *string   `json:"usageHint,omitempty"`
}

type IconProps struct {
	Name StringRef `json:"name"`
}

type DividerProps struct {
	Axis *string `json:"axis,omitempty"`
}

type ButtonProps struct {
	Child   string `json:"child"`
	Action  Action `json:"action"`
	Primary *bool  `json:"primary,omitempty"`
}

type TextFieldProps struct {
	Text              *StringRef `json:"text,omitempty"`
	Label             *StringRef `json:"label,omitempty"`
	TextFieldType     *string    `json:"textFieldType,omitempty"`
	ValidationRegexp  *string    `json:"validationRegexp,omitempty"`
	OnSubmittedAction *Action    `json:"onSubmittedAction,omitempty"`
}

type CheckBoxProps struct {
	Label StringRef `json:"label"`
	Value BoolRef   `json:"value"`
}

type CardProps struct {
	Child string `json:"child"`
}

type ModalProps struct {
	EntryPointChild string `json:"entryPointChild"`
	ContentChild    string `json:"contentChild"`
}

type TabItem struct {
	Title StringRef `json:"title"`
	Child string    `json:"child"`
}

type TabsProps struct {
	TabItems []TabItem `json:"tabItems"`
}

type ChoiceOption struct {
	Label StringRef `json:"label"`
	Value string    `json:"value"`
}

type MultipleChoiceProps struct {
	Selections           StringArrayRef `json:"selections"`
	Options              []ChoiceOption `json:"options"`
	MaxAllowedSelections *int64         `json:"maxAllowedSelections,omitempty"`
}

type SliderProps struct {
	Value    NumberRef `json:"value"`
	MinValue *float64  `json:"minValue,omitempty"`
	MaxValue *float64  `json:"maxValue,omitempty"`
}

type DateTimeInputProps struct {
	Value      StringRef `json:"value"`
	EnableDate *bool     `json:"enableDate,omitempty"`
	EnableTime *bool     `json:"enableTime,omitempty"`
	FirstDate  *string   `json:"firstDate,omitempty"`
	LastDate   *string   `json:"lastDate,omitempty"`
}

type AudioPlayerProps struct {
	URL StringRef `json:"url"`
}

type VideoProps struct {
	URL StringRef `json:"url"`
}

// SurfaceUpdate updates the surface with components.
type SurfaceUpdate struct {
	SurfaceID  string           `json:"surfaceId"`
	Components []ComponentEntry `json:"components"`
}

// SurfaceStyles controls per-surface styling.
type SurfaceStyles struct {
	Font         *string `json:"font,omitempty"`
	PrimaryColor *string `json:"primaryColor,omitempty"`
}

// BeginRendering starts rendering a surface.
type BeginRendering struct {
	SurfaceID string         `json:"surfaceId"`
	Root      string         `json:"root"`
	CatalogID *string        `json:"catalogId,omitempty"`
	Styles    *SurfaceStyles `json:"styles,omitempty"`
}

// DataModelUpdate updates the data model for a surface.
type DataModelUpdate struct {
	SurfaceID string          `json:"surfaceId"`
	Path      *string         `json:"path,omitempty"`
	Contents  json.RawMessage `json:"contents"`
}

// DeleteSurface removes a surface.
type DeleteSurface struct {
	SurfaceID string `json:"surfaceId"`
}

// A2uiMessage encodes a surface message.
type A2uiMessage struct {
	SurfaceUpdate   *SurfaceUpdate   `json:"surfaceUpdate,omitempty"`
	BeginRendering  *BeginRendering  `json:"beginRendering,omitempty"`
	DataModelUpdate *DataModelUpdate `json:"dataModelUpdate,omitempty"`
	DeleteSurface   *DeleteSurface   `json:"deleteSurface,omitempty"`
}
