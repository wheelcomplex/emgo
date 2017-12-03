package eve

// DL provides a convenient way to write Display List commands. Every command is
// a function call, so for better performance or lower RAM usage, use raw Writer
// with many Display List commands in array / slice.
type DL struct {
	w Writer
}

// AlphaFunc sets the alpha test function.
func (dl DL) AlphaFunc(fun, ref byte) {
	dl.w.wr32(ALPHA_FUNC | uint32(fun)<<8 | uint32(ref))
}

// Begin begins drawing a graphics primitive.
func (dl DL) Begin(prim byte) {
	dl.w.wr32(BEGIN | uint32(prim))
}

// BitmapHandle selscts the bitmap handle.
func (dl DL) BitmapHandle(handle byte) {
	dl.w.wr32(BITMAP_HANDLE | uint32(handle))
}

// BitmapLayout sets the bitmap memory format and layout for the current handle.
func (dl DL) BitmapLayout(format byte, linestride, height int) {
	dl.w.wr32(BITMAP_LAYOUT | uint32(format)<<19 | uint32(linestride)&1023<<9 |
		uint32(height)&511)
	if linestride > 1023 || height > 511 {
		// BUG?: Does BITMAP_LAYOUT zeros bits set by previous BITMAP_LAYOUT_H?
		l := uint32(linestride) >> 10 & 3
		h := uint32(height) >> 9 & 3
		dl.w.wr32(BITMAP_LAYOUT_H | l<<2 | h)
	}
}

// BitmapSize sets the screen drawing of bitmaps for the current handle.
func (dl DL) BitmapSize(options byte, width, height int) {
	w := uint32(width) & 511
	h := uint32(height) & 511
	dl.w.wr32(BITMAP_SIZE | uint32(options)<<18 | w<<9 | h)
	if width > 511 || height > 511 {
		// BUG?: Does BITMAP_SIZE clears bits set by previous BITMAP_SIZE_H?
		w = uint32(width) >> 9 & 3
		h = uint32(height) >> 9 & 3
		dl.w.wr32(BITMAP_SIZE_H | w<<2 | h)
	}
}

// BitmapSource sets the source address of bitmap data in graphics memory RAM_G.
func (dl DL) BitmapSource(addr int) {
	dl.w.wr32(BITMAP_SOURCE | uint32(addr)&0x3FFFFF)
}

// BitmapTransA sets the A coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformA(a int) {
	dl.w.wr32(BITMAP_TRANSFORM_A | uint32(a)&0x1FFFF)
}

// BitmapTransformB sets the B coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformB(b int) {
	dl.w.wr32(BITMAP_TRANSFORM_B | uint32(b)&0x1FFFF)
}

// BitmapTransformC sets the C coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformC(c int) {
	dl.w.wr32(BITMAP_TRANSFORM_C | uint32(c)&0x1FFFF)
}

// BitmapTransformD sets the D coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformD(d int) {
	dl.w.wr32(BITMAP_TRANSFORM_D | uint32(d)&0x1FFFF)
}

// BitmapTransformE sets the E coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformE(e int) {
	dl.w.wr32(BITMAP_TRANSFORM_E | uint32(e)&0x1FFFF)
}

// BitmapTransformF sets the F coefficient of the bitmap transform matrix.
func (dl DL) BitmapTransformF(f int) {
	dl.w.wr32(BITMAP_TRANSFORM_F | uint32(f)&0x1FFFF)
}

// BlendFunc configures pixel arithmetic.
func (dl DL) BlendFunc(src, dst byte) {
	dl.w.wr32(BLEND_FUNC | uint32(src)<<3 | uint32(dst))
}

// Call executes a sequence of commands at another location in the display list.
func (dl DL) Call(dest int) {
	dl.w.wr32(CALL | uint32(dest)&0xFFFF)
}

// Cell sets the bitmap cell number for the Vertex2F command.
func (dl DL) Cell(cell byte) {
	dl.w.wr32(CELL | uint32(cell))
}

// Clear clears buffers to preset values.
func (dl DL) Clear(cst byte) {
	dl.w.wr32(CLEAR | uint32(cst))
}

// ClearColorA sets the clear value for the alpha channel.
func (dl DL) ClearColorA(alpha byte) {
	dl.w.wr32(CLEAR_COLOR_A | uint32(alpha))
}

// ClearColorRGB sets the clear values for red, green and blue channels.
func (dl DL) ClearColorRGB(red, green, blue byte) {
	dl.w.wr32(CLEAR_COLOR_RGB | uint32(red)<<16 | uint32(blue)<<8 | uint32(green))
}

// ClearStencil sets the clear value for the stencil buffer.
func (dl DL) ClearStencil(s byte) {
	dl.w.wr32(CLEAR_STENCIL | uint32(s))
}

// ClearTag sets the clear value for the stencil buffer.
func (dl DL) ClearTag(t byte) {
	dl.w.wr32(CLEAR_TAG | uint32(t))
}

// ColorA sets the current color alpha.
func (dl DL) ColorA(alpha byte) {
	dl.w.wr32(COLOR_A | uint32(alpha))
}

// ColorMask enables or disables writing of color components.
func (dl DL) ColorMask(rgba byte) {
	dl.w.wr32(COLOR_MASK | uint32(rgba))
}

// ColorRGB sets the current color red, green and blue.
func (dl DL) ColorRGB(red, green, blue byte) {
	dl.w.wr32(COLOR_RGB | uint32(red)<<16 | uint32(blue)<<8 | uint32(green))
}

// Display ends the display list (following command will be ignored).
func (dl DL) Display() {
	dl.w.wr32(DISPLAY)
}

// End ends drawing a graphics primitive.
func (dl DL) End() {
	dl.w.wr32(END)
}

// Jump executes commands at another location in the display list.
func (dl DL) Jump(dest int) {
	dl.w.wr32(JUMP | uint32(dest)&0xFFFF)
}

// LineWidth sets the width of lines to be drawn with primitive LINES in 1/16
// pixel precision.
func (dl DL) LineWidth(width int) {
	dl.w.wr32(LINE_WIDTH | uint32(width)&0xFFF)
}

// Macro executes a single command from a macro register.
func (dl DL) Macro(m byte) {
	dl.w.wr32(MACRO | uint32(m))
}

// Nop does nothing.
func (dl DL) Nop() {
	dl.w.wr32(NOP)
}

// PaletteSource sets the base address of the palette (EVE2).
func (dl DL) PaletteSource(addr int) {
	dl.w.wr32(PALETTE_SOURCE | uint32(addr)&0x3FFFFF)
}

// PointSize sets the radius of points.
func (dl DL) PointSize(size int) {
	dl.w.wr32(POINT_SIZE | uint32(size)&0x1FFF)
}

// RestoreContext restores the current graphics context from the context stack.
func (dl DL) RestoreContext() {
	dl.w.wr32(RESTORE_CONTEXT)
}

// Return returns from a previous CALL command.
func (dl DL) Return() {
	dl.w.wr32(RETURN)
}

// SaveContext pushes the current graphics context on the context stack.
func (dl DL) SaveContext() {
	dl.w.wr32(SAVE_CONTEXT)
}

// ScissorSize sets the size of the scissor clip rectangle.
func (dl DL) ScissorSize(width, height int) {
	dl.w.wr32(SCISSOR_SIZE | uint32(width)&0xFFF<<12 | uint32(height)&0xFFF)
}

// ScissorXY sets the size of the scissor clip rectangle.
func (dl DL) ScissorXY(x, y int) {
	dl.w.wr32(SCISSOR_XY | uint32(x)&0x7FF<<11 | uint32(y)&0x7FF)
}

// StencilFunc sets function and reference value for stencil testing.
func (dl DL) StencilFunc(fun, ref, mask byte) {
	dl.w.wr32(STENCIL_FUNC | uint32(fun)<<16 | uint32(ref)<<8 | uint32(mask))
}

// StencilMask controls the writing of individual bits in the stencil planes.
func (dl DL) StencilMask(mask byte) {
	dl.w.wr32(STENCIL_MASK | uint32(mask))
}

// StencilOp sets stencil test actions.
func (dl DL) StencilOp(sfail, spass byte) {
	dl.w.wr32(STENCIL_OP | uint32(sfail)<<3 | uint32(spass))
}

// Tag attaches the tag value for the following graphics objects drawn on the
// screen. The initial tag buffer value is 255.
func (dl DL) Tag(t byte) {
	dl.w.wr32(TAG | uint32(t))
}

// TagMask controls the writing of the tag buffer.
func (dl DL) TagMask(mask byte) {
	dl.w.wr32(TAG_MASK | uint32(mask))
}

// Vertex2F starts the operation of graphics primitives at the specified screen
// coordinate, in the pixel precision set by VertexFormat (default: 1/16 pixel).
func (dl DL) Vertex2F(x, y int) {
	dl.w.wr32(VERTEX2F | uint32(x)&0x7FFF<<15 | uint32(y)&0x7FFF)
}

// Vertex2II starts the operation of graphics primitive at the specified
// coordinates in pixel precision.
func (dl DL) Vertex2II(x, y int, handle, cell byte) {
	dl.w.wr32(VERTEX2II | uint32(x)&511<<21 | uint32(y)&511<<12 |
		uint32(handle)<<7 | uint32(cell))
}

// VertexFormat sets the precision of Vertex2F coordinates.
func (dl DL) VertexFormat(frac uint) {
	dl.w.wr32(VERTEX_FORMAT | uint32(frac)&7)
}
