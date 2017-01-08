// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 08 Jan 2017 04:57:05 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package nk

/*
#include "nuklear.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Char type as declared in nk/nuklear.h:390
type Char byte

// Uchar type as declared in nk/nuklear.h:391
type Uchar byte

// Byte type as declared in nk/nuklear.h:392
type Byte byte

// Short type as declared in nk/nuklear.h:393
type Short int16

// Ushort type as declared in nk/nuklear.h:394
type Ushort uint16

// Int type as declared in nk/nuklear.h:395
type Int int32

// Uint type as declared in nk/nuklear.h:396
type Uint uint32

// Size type as declared in nk/nuklear.h:397
type Size uint

// Ptr type as declared in nk/nuklear.h:398
type Ptr uint

// Hash type as declared in nk/nuklear.h:400
type Hash uint32

// Flags type as declared in nk/nuklear.h:401
type Flags uint32

// Rune type as declared in nk/nuklear.h:402
type Rune uint32

// Glyph type as declared in nk/nuklear.h:455
type Glyph [4]byte

// Handle as declared in nk/nuklear.h:456
const sizeofHandle = unsafe.Sizeof(C.nk_handle{})

type Handle [sizeofHandle]byte

// PluginAlloc type as declared in nk/nuklear.h:475
type PluginAlloc func(arg0 Handle, old unsafe.Pointer, arg2 Size) unsafe.Pointer

// PluginFree type as declared in nk/nuklear.h:476
type PluginFree func(arg0 Handle, old unsafe.Pointer)

// PluginFilter type as declared in nk/nuklear.h:477
type PluginFilter func(arg0 *TextEdit, unicode Rune) int32

// PluginPaste type as declared in nk/nuklear.h:478
type PluginPaste func(arg0 Handle, arg1 *TextEdit)

// PluginCopy type as declared in nk/nuklear.h:479
type PluginCopy func(arg0 Handle, arg1 string, len int32)

// TextWidthF type as declared in nk/nuklear.h:1286
type TextWidthF func(arg0 Handle, h float32, arg2 string, len int32) float32

// QueryFontGlyphF type as declared in nk/nuklear.h:1287
type QueryFontGlyphF func(handle Handle, fontHeight float32, glyph *UserFontGlyph, codepoint Rune, nextCodepoint Rune)

// DrawIndex type as declared in nk/nuklear.h:2042
type DrawIndex uint16

// StyleItem as declared in nk/nuklear.h:427
type StyleItem C.struct_nk_style_item

// Command as declared in nk/nuklear.h:1783
type Command C.struct_nk_command

// MenuState as declared in nk/nuklear.h:2672
type MenuState C.struct_nk_menu_state

// StyleChart as declared in nk/nuklear.h:442
type StyleChart C.struct_nk_style_chart

// ConfigStackFlags as declared in nk/nuklear.h:2860
type ConfigStackFlags C.struct_nk_config_stack_flags

// ConvertConfig as declared in nk/nuklear.h:426
type ConvertConfig struct {
	GlobalAlpha        float32
	LineAa             AntiAliasing
	ShapeAa            AntiAliasing
	CircleSegmentCount uint32
	ArcSegmentCount    uint32
	CurveSegmentCount  uint32
	Null               DrawNullTexture
	VertexLayout       []DrawVertexLayoutElement
	VertexSize         Size
	VertexAlignment    Size
	ref82bf4c25        *C.struct_nk_convert_config
	allocs82bf4c25     interface{}
}

// CommandBuffer as declared in nk/nuklear.h:424
type CommandBuffer C.struct_nk_command_buffer

// BakedFont as declared in nk/nuklear.h:1325
type BakedFont C.struct_nk_baked_font

// FontConfig as declared in nk/nuklear.h:1338
type FontConfig C.struct_nk_font_config

// ConfigStackColor as declared in nk/nuklear.h:2861
type ConfigStackColor C.struct_nk_config_stack_color

// CommandCircle as declared in nk/nuklear.h:1858
type CommandCircle C.struct_nk_command_circle

// ConfigStackColorElement as declared in nk/nuklear.h:2853
type ConfigStackColorElement C.struct_nk_config_stack_color_element

// CommandPolygon as declared in nk/nuklear.h:1890
type CommandPolygon C.struct_nk_command_polygon

// DrawList as declared in nk/nuklear.h:429
type DrawList C.struct_nk_draw_list

// PropertyState as declared in nk/nuklear.h:2745
type PropertyState C.struct_nk_property_state

// CommandTriangle as declared in nk/nuklear.h:1841
type CommandTriangle C.struct_nk_command_triangle

// StyleSelectable as declared in nk/nuklear.h:436
type StyleSelectable C.struct_nk_style_selectable

// Image as declared in nk/nuklear.h:457
type Image C.struct_nk_image

// ConfigStackFlagsElement as declared in nk/nuklear.h:2852
type ConfigStackFlagsElement C.struct_nk_config_stack_flags_element

// Table as declared in nk/nuklear.h:2704
type Table C.struct_nk_table

// Context as declared in nk/nuklear.h:432
type Context C.struct_nk_context

// Vec2 as declared in nk/nuklear.h:451
type Vec2 C.struct_nk_vec2

// ChartSlot as declared in nk/nuklear.h:2634
type ChartSlot C.struct_nk_chart_slot

// CommandArc as declared in nk/nuklear.h:1873
type CommandArc C.struct_nk_command_arc

// CommandCircleFilled as declared in nk/nuklear.h:1866
type CommandCircleFilled C.struct_nk_command_circle_filled

// Style as declared in nk/nuklear.h:2582
type Style C.struct_nk_style

// StyleScrollbar as declared in nk/nuklear.h:439
type StyleScrollbar C.struct_nk_style_scrollbar

// Input as declared in nk/nuklear.h:2002
type Input C.struct_nk_input

// CommandArcFilled as declared in nk/nuklear.h:1882
type CommandArcFilled C.struct_nk_command_arc_filled

// Window as declared in nk/nuklear.h:2756
type Window C.struct_nk_window

// DrawCommand as declared in nk/nuklear.h:425
type DrawCommand C.struct_nk_draw_command

// CommandRect as declared in nk/nuklear.h:1814
type CommandRect C.struct_nk_command_rect

// Font as declared in nk/nuklear.h:1379
type Font C.struct_nk_font

// ConfigStackUserFont as declared in nk/nuklear.h:2862
type ConfigStackUserFont struct {
	Head           int32
	Elements       [8]ConfigStackUserFontElement
	refa664861d    *C.struct_nk_config_stack_user_font
	allocsa664861d interface{}
}

// TextUndoState as declared in nk/nuklear.h:1647
type TextUndoState C.struct_nk_text_undo_state

// DrawVertexLayoutElement as declared in nk/nuklear.h:433
type DrawVertexLayoutElement struct {
	Attribute      DrawVertexLayoutAttribute
	Format         DrawVertexLayoutFormat
	Offset         Size
	refeb0614d6    *C.struct_nk_draw_vertex_layout_element
	allocseb0614d6 interface{}
}

// CommandScissor as declared in nk/nuklear.h:1791
type CommandScissor C.struct_nk_command_scissor

// PageData as declared in nk/nuklear.h:2888
const sizeofPageData = unsafe.Sizeof(C.union_nk_page_data{})

type PageData [sizeofPageData]byte

// MemoryStatus as declared in nk/nuklear.h:1479
type MemoryStatus C.struct_nk_memory_status

// MouseButton as declared in nk/nuklear.h:1976
type MouseButton C.struct_nk_mouse_button

// StyleText as declared in nk/nuklear.h:2191
type StyleText C.struct_nk_style_text

// ConfigStackStyleItem as declared in nk/nuklear.h:2857
type ConfigStackStyleItem C.struct_nk_config_stack_style_item

// TextUndoRecord as declared in nk/nuklear.h:1640
type TextUndoRecord C.struct_nk_text_undo_record

// Panel as declared in nk/nuklear.h:431
type Panel C.struct_nk_panel

// Str as declared in nk/nuklear.h:1551
type Str C.struct_nk_str

// CommandImage as declared in nk/nuklear.h:1913
type CommandImage C.struct_nk_command_image

// Chart as declared in nk/nuklear.h:2644
type Chart C.struct_nk_chart

// StyleSlide as declared in nk/nuklear.h:437
type StyleSlide C.struct_nk_style_slide

// Key as declared in nk/nuklear.h:1992
type Key C.struct_nk_key

// Page as declared in nk/nuklear.h:2900
type Page C.struct_nk_page

// Buffer as declared in nk/nuklear.h:422
type Buffer C.struct_nk_buffer

// StyleWindowHeader as declared in nk/nuklear.h:445
type StyleWindowHeader C.struct_nk_style_window_header

// CommandLine as declared in nk/nuklear.h:1797
type CommandLine C.struct_nk_command_line

// StyleProgress as declared in nk/nuklear.h:438
type StyleProgress C.struct_nk_style_progress

// BufferMarker as declared in nk/nuklear.h:1499
type BufferMarker C.struct_nk_buffer_marker

// Mouse as declared in nk/nuklear.h:1981
type Mouse C.struct_nk_mouse

// ConfigStackStyleItemElement as declared in nk/nuklear.h:2849
type ConfigStackStyleItemElement C.struct_nk_config_stack_style_item_element

// UserFont as declared in nk/nuklear.h:430
type UserFont struct {
	Userdata       Handle
	Height         float32
	Width          TextWidthF
	Query          QueryFontGlyphF
	Texture        Handle
	ref738ce62e    *C.struct_nk_user_font
	allocs738ce62e interface{}
}

// ConfigStackFloat as declared in nk/nuklear.h:2858
type ConfigStackFloat C.struct_nk_config_stack_float

// CommandPolygonFilled as declared in nk/nuklear.h:1898
type CommandPolygonFilled C.struct_nk_command_polygon_filled

// StyleCombo as declared in nk/nuklear.h:443
type StyleCombo C.struct_nk_style_combo

// StyleWindow as declared in nk/nuklear.h:446
type StyleWindow C.struct_nk_style_window

// PopupBuffer as declared in nk/nuklear.h:2664
type PopupBuffer C.struct_nk_popup_buffer

// StyleProperty as declared in nk/nuklear.h:441
type StyleProperty C.struct_nk_style_property

// Keyboard as declared in nk/nuklear.h:1996
type Keyboard C.struct_nk_keyboard

// RowLayout as declared in nk/nuklear.h:2650
type RowLayout C.struct_nk_row_layout

// CommandRectFilled as declared in nk/nuklear.h:1823
type CommandRectFilled C.struct_nk_command_rect_filled

// ConfigStackFloatElement as declared in nk/nuklear.h:2850
type ConfigStackFloatElement C.struct_nk_config_stack_float_element

// Scroll as declared in nk/nuklear.h:459
type Scroll C.struct_nk_scroll

// PopupState as declared in nk/nuklear.h:2721
type PopupState C.struct_nk_popup_state

// CommandCurve as declared in nk/nuklear.h:1805
type CommandCurve C.struct_nk_command_curve

// StyleItemData as declared in nk/nuklear.h:2181
const sizeofStyleItemData = unsafe.Sizeof(C.union_nk_style_item_data{})

type StyleItemData [sizeofStyleItemData]byte

// TextEdit as declared in nk/nuklear.h:428
type TextEdit C.struct_nk_text_edit

// EditState as declared in nk/nuklear.h:2732
type EditState C.struct_nk_edit_state

// DrawNullTexture as declared in nk/nuklear.h:487
type DrawNullTexture C.struct_nk_draw_null_texture

// StyleSlider as declared in nk/nuklear.h:2288
type StyleSlider C.struct_nk_style_slider

// Pool as declared in nk/nuklear.h:2906
type Pool C.struct_nk_pool

// ConfigStackButtonBehavior as declared in nk/nuklear.h:2863
type ConfigStackButtonBehavior C.struct_nk_config_stack_button_behavior

// ConfigurationStacks as declared in nk/nuklear.h:2865
type ConfigurationStacks C.struct_nk_configuration_stacks

// PageElement as declared in nk/nuklear.h:2894
type PageElement C.struct_nk_page_element

// FontAtlas as declared in nk/nuklear.h:1396
type FontAtlas C.struct_nk_font_atlas

// CommandTriangleFilled as declared in nk/nuklear.h:1850
type CommandTriangleFilled C.struct_nk_command_triangle_filled

// CommandText as declared in nk/nuklear.h:1921
type CommandText C.struct_nk_command_text

// StyleToggle as declared in nk/nuklear.h:435
type StyleToggle C.struct_nk_style_toggle

// FontGlyph as declared in nk/nuklear.h:1372
type FontGlyph C.struct_nk_font_glyph

// ConfigStackVec2Element as declared in nk/nuklear.h:2851
type ConfigStackVec2Element C.struct_nk_config_stack_vec2_element

// Vec2i as declared in nk/nuklear.h:452
type Vec2i C.struct_nk_vec2i

// ConfigStackVec2 as declared in nk/nuklear.h:2859
type ConfigStackVec2 C.struct_nk_config_stack_vec2

// Color as declared in nk/nuklear.h:449
type Color C.struct_nk_color

// StyleButton as declared in nk/nuklear.h:434
type StyleButton C.struct_nk_style_button

// ListView as declared in nk/nuklear.h:504
type ListView C.struct_nk_list_view

// ConfigStackButtonBehaviorElement as declared in nk/nuklear.h:2855
type ConfigStackButtonBehaviorElement C.struct_nk_config_stack_button_behavior_element

// StyleTab as declared in nk/nuklear.h:444
type StyleTab C.struct_nk_style_tab

// Rect as declared in nk/nuklear.h:453
type Rect C.struct_nk_rect

// Allocator as declared in nk/nuklear.h:423
type Allocator C.struct_nk_allocator

// Cursor as declared in nk/nuklear.h:458
type Cursor C.struct_nk_cursor

// StyleEdit as declared in nk/nuklear.h:440
type StyleEdit C.struct_nk_style_edit

// Memory as declared in nk/nuklear.h:1504
type Memory C.struct_nk_memory

// CommandRectMultiColor as declared in nk/nuklear.h:1831
type CommandRectMultiColor C.struct_nk_command_rect_multi_color

// ConfigStackUserFontElement as declared in nk/nuklear.h:2854
type ConfigStackUserFontElement struct {
	Address        [][]UserFont
	OldValue       []UserFont
	ref5572630c    *C.struct_nk_config_stack_user_font_element
	allocs5572630c interface{}
}

// CommandPolyline as declared in nk/nuklear.h:1905
type CommandPolyline C.struct_nk_command_polyline

// Recti as declared in nk/nuklear.h:454
type Recti C.struct_nk_recti

// Colorf as declared in nk/nuklear.h:450
type Colorf C.struct_nk_colorf

// Clipboard as declared in nk/nuklear.h:1634
type Clipboard C.struct_nk_clipboard

// UserFontGlyph as declared in nk/nuklear.h:1285
type UserFontGlyph struct {
	Uv             [2]Vec2
	Offset         Vec2
	Width          float32
	Height         float32
	Xadvance       float32
	ref4a84b297    *C.struct_nk_user_font_glyph
	allocs4a84b297 interface{}
}
