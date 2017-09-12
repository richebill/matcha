// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/ios/tabview.proto

package ios

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha "gomatcha.io/matcha/pb"
import matcha_text "gomatcha.io/matcha/pb/text"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TabChildView struct {
	Title        string                  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Icon         *matcha.ImageOrResource `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	SelectedIcon *matcha.ImageOrResource `protobuf:"bytes,4,opt,name=selectedIcon" json:"selectedIcon,omitempty"`
	Badge        string                  `protobuf:"bytes,5,opt,name=badge" json:"badge,omitempty"`
}

func (m *TabChildView) Reset()                    { *m = TabChildView{} }
func (m *TabChildView) String() string            { return proto.CompactTextString(m) }
func (*TabChildView) ProtoMessage()               {}
func (*TabChildView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TabChildView) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TabChildView) GetIcon() *matcha.ImageOrResource {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *TabChildView) GetSelectedIcon() *matcha.ImageOrResource {
	if m != nil {
		return m.SelectedIcon
	}
	return nil
}

func (m *TabChildView) GetBadge() string {
	if m != nil {
		return m.Badge
	}
	return ""
}

type TabView struct {
	Screens             []*TabChildView        `protobuf:"bytes,1,rep,name=screens" json:"screens,omitempty"`
	SelectedIndex       int64                  `protobuf:"varint,2,opt,name=selectedIndex" json:"selectedIndex,omitempty"`
	BarColor            *matcha.Color          `protobuf:"bytes,3,opt,name=barColor" json:"barColor,omitempty"`
	SelectedColor       *matcha.Color          `protobuf:"bytes,6,opt,name=selectedColor" json:"selectedColor,omitempty"`
	UnselectedColor     *matcha.Color          `protobuf:"bytes,7,opt,name=unselectedColor" json:"unselectedColor,omitempty"`
	SelectedTextStyle   *matcha_text.TextStyle `protobuf:"bytes,8,opt,name=selectedTextStyle" json:"selectedTextStyle,omitempty"`
	UnselectedTextStyle *matcha_text.TextStyle `protobuf:"bytes,9,opt,name=unselectedTextStyle" json:"unselectedTextStyle,omitempty"`
}

func (m *TabView) Reset()                    { *m = TabView{} }
func (m *TabView) String() string            { return proto.CompactTextString(m) }
func (*TabView) ProtoMessage()               {}
func (*TabView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TabView) GetScreens() []*TabChildView {
	if m != nil {
		return m.Screens
	}
	return nil
}

func (m *TabView) GetSelectedIndex() int64 {
	if m != nil {
		return m.SelectedIndex
	}
	return 0
}

func (m *TabView) GetBarColor() *matcha.Color {
	if m != nil {
		return m.BarColor
	}
	return nil
}

func (m *TabView) GetSelectedColor() *matcha.Color {
	if m != nil {
		return m.SelectedColor
	}
	return nil
}

func (m *TabView) GetUnselectedColor() *matcha.Color {
	if m != nil {
		return m.UnselectedColor
	}
	return nil
}

func (m *TabView) GetSelectedTextStyle() *matcha_text.TextStyle {
	if m != nil {
		return m.SelectedTextStyle
	}
	return nil
}

func (m *TabView) GetUnselectedTextStyle() *matcha_text.TextStyle {
	if m != nil {
		return m.UnselectedTextStyle
	}
	return nil
}

type TabEvent struct {
	SelectedIndex int64 `protobuf:"varint,1,opt,name=selectedIndex" json:"selectedIndex,omitempty"`
}

func (m *TabEvent) Reset()                    { *m = TabEvent{} }
func (m *TabEvent) String() string            { return proto.CompactTextString(m) }
func (*TabEvent) ProtoMessage()               {}
func (*TabEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TabEvent) GetSelectedIndex() int64 {
	if m != nil {
		return m.SelectedIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*TabChildView)(nil), "matcha.view.ios.TabChildView")
	proto.RegisterType((*TabView)(nil), "matcha.view.ios.TabView")
	proto.RegisterType((*TabEvent)(nil), "matcha.view.ios.TabEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/ios/tabview.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdd, 0x6a, 0xe2, 0x40,
	0x14, 0xc7, 0x89, 0xf1, 0x73, 0x54, 0x64, 0x67, 0x97, 0xdd, 0x20, 0xec, 0xe2, 0x4a, 0x0b, 0x96,
	0x96, 0xa4, 0xe8, 0x85, 0x17, 0x85, 0x5e, 0x68, 0x0b, 0xf5, 0xa2, 0x28, 0x31, 0xf4, 0xa2, 0x77,
	0x33, 0xc9, 0x41, 0x07, 0x62, 0x46, 0x92, 0xf1, 0xa3, 0x8f, 0xd1, 0x57, 0x28, 0x7d, 0xd0, 0x92,
	0xc9, 0x87, 0x8d, 0x06, 0x6f, 0x92, 0x99, 0x33, 0xbf, 0xff, 0xff, 0xcc, 0x39, 0x67, 0xd0, 0xcd,
	0x82, 0xaf, 0x88, 0xb0, 0x97, 0x44, 0x67, 0xdc, 0x88, 0x56, 0xc6, 0x9a, 0x1a, 0x5b, 0x06, 0x3b,
	0x83, 0xf1, 0xc0, 0x10, 0x84, 0x86, 0x6b, 0x7d, 0xed, 0x73, 0xc1, 0x71, 0x2b, 0x66, 0x65, 0x88,
	0xf1, 0xa0, 0xfd, 0x3f, 0x5f, 0xce, 0x56, 0x64, 0x01, 0x91, 0xa6, 0x7d, 0x99, 0x8f, 0x08, 0xd8,
	0x0b, 0xf9, 0x89, 0xb0, 0xee, 0xa7, 0x82, 0x1a, 0x16, 0xa1, 0xe3, 0x25, 0x73, 0x9d, 0x17, 0x06,
	0x3b, 0xfc, 0x0b, 0x95, 0x04, 0x13, 0x2e, 0x68, 0x85, 0x8e, 0xd2, 0xab, 0x99, 0xd1, 0x06, 0x5f,
	0xa3, 0x22, 0xb3, 0xb9, 0xa7, 0xa9, 0x1d, 0xa5, 0x57, 0xef, 0xff, 0xd1, 0x63, 0xeb, 0x49, 0x98,
	0x70, 0xea, 0x9b, 0x10, 0xf0, 0x8d, 0x6f, 0x83, 0x29, 0x21, 0x7c, 0x87, 0x1a, 0x01, 0xb8, 0x60,
	0x0b, 0x70, 0x26, 0xa1, 0xa8, 0x78, 0x5e, 0x94, 0x81, 0xc3, 0xfc, 0x94, 0x38, 0x0b, 0xd0, 0x4a,
	0x51, 0x7e, 0xb9, 0xe9, 0xbe, 0xab, 0xa8, 0x62, 0x11, 0x2a, 0x6f, 0x38, 0x44, 0x95, 0xc0, 0xf6,
	0x01, 0xbc, 0x40, 0x53, 0x3a, 0x6a, 0xaf, 0xde, 0xff, 0xab, 0x1f, 0xf5, 0x47, 0xff, 0x5e, 0x91,
	0x99, 0xd0, 0xf8, 0x02, 0x35, 0xd3, 0x54, 0x9e, 0x03, 0x7b, 0x59, 0xa2, 0x6a, 0x66, 0x83, 0xf8,
	0x0a, 0x55, 0x29, 0xf1, 0xc7, 0xdc, 0xe5, 0x7e, 0x5c, 0x6e, 0x33, 0xf1, 0x97, 0x41, 0x33, 0x3d,
	0xc6, 0x83, 0x83, 0x61, 0xc4, 0x97, 0xf3, 0xf8, 0x2c, 0x83, 0x87, 0xa8, 0xb5, 0xf1, 0xb2, 0xb2,
	0x4a, 0x9e, 0xec, 0x98, 0xc2, 0x0f, 0xe8, 0x47, 0x12, 0xb0, 0x60, 0x2f, 0xe6, 0xe2, 0xcd, 0x05,
	0xad, 0x2a, 0xa5, 0xbf, 0x13, 0xa9, 0x9c, 0x6c, 0x7a, 0x6a, 0x9e, 0x0a, 0xf0, 0x13, 0xfa, 0x79,
	0x30, 0x3e, 0xf8, 0xd4, 0xce, 0xfa, 0xe4, 0x49, 0xba, 0xb7, 0xa8, 0x6a, 0x11, 0xfa, 0xb8, 0x05,
	0x4f, 0x9c, 0xb6, 0x56, 0xc9, 0x69, 0xed, 0xe8, 0x1e, 0xfd, 0x63, 0x5c, 0x4f, 0x1f, 0x66, 0xfc,
	0x5b, 0xd3, 0x74, 0x70, 0xa3, 0xda, 0x8c, 0xc6, 0x63, 0x7e, 0x55, 0x19, 0x0f, 0x3e, 0x0a, 0xf5,
	0x67, 0x09, 0xb1, 0xe9, 0x7c, 0x36, 0xa2, 0x65, 0xf9, 0x66, 0x07, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x87, 0x09, 0xa7, 0x4a, 0x3e, 0x03, 0x00, 0x00,
}
