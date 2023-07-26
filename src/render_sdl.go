package main

// ------------------------------------------------------------------
// Texture

type Texture struct {
	width  int32
	height int32
	depth  int32
	filter bool
	//handle gl.Texture
}

// Generate a new texture name
func newTexture(width, height, depth int32, filter bool) (t *Texture) {
	/*t = &Texture{width, height, depth, filter, gl.CreateTexture()}
	runtime.SetFinalizer(t, func(t *Texture) {
		sys.mainThreadTask <- func() {
			gl.DeleteTexture(t.handle)
		}
	})*/
	return
}

// Bind a texture and upload texel data to it
func (t *Texture) SetData(data []byte) {
	/*
		var interp int = gl.NEAREST
		if t.filter {
			interp = gl.LINEAR
		}

		format := InternalFormatLUT[Max(t.depth, 8)]

		gl.BindTexture(gl.TEXTURE_2D, t.handle)
		gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
		gl.TexImage2D(gl.TEXTURE_2D, 0, int(t.width), int(t.height), format, gl.UNSIGNED_BYTE, data)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, interp)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, interp)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	*/
}

// Return whether texture has a valid handle
func (t *Texture) IsValid() bool {
	//return t.handle.IsValid()
	return true
}
