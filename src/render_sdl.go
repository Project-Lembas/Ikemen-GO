package main

import (
	"fmt"
	"os"
	"sync"

	sdl "github.com/veandco/go-sdl2/sdl"
)

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

// ------------------------------------------------------------------
// Renderer

type Renderer struct {
	runningMutex sync.Mutex

	sdlrenderer *sdl.Renderer
}

//go:embed shaders/sprite.vert.glsl
var vertShader string

//go:embed shaders/sprite.frag.glsl
var fragShader string

//go:embed shaders/ident.vert.glsl
var identVertShader string

//go:embed shaders/ident.frag.glsl
var identFragShader string

// Render initialization.
// Creates the default shaders, the framebuffer and enables MSAA.
func (r *Renderer) Init(window *Window) {
	var err error
	fmt.Fprint(os.Stdout, "sdl render init\n")

	sdl.Do(func() {
		r.sdlrenderer, err = sdl.CreateRenderer(window.Window, -1, sdl.RENDERER_SOFTWARE)
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}

	sdl.Do(func() {
		r.sdlrenderer.Clear()
	})

	running := true
	for running {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			}

			r.sdlrenderer.Clear()
			r.sdlrenderer.SetDrawColor(0, 0, 0, 0x20)
			r.sdlrenderer.FillRect(&sdl.Rect{0, 0, int32(window.w), int32(window.h)})
		})

		//wg := sync.WaitGroup{}
		sdl.Do(func() {
			r.sdlrenderer.Present()
			sdl.Delay(1000 / 60)
		})
	}
	/*
		sys.errLog.Printf("Using OpenGL %v (%v)",
			gl.GetString(gl.VERSION), gl.GetString(gl.RENDERER))

		r.postShaderSelect = make([]*ShaderProgram, 1+len(sys.externalShaderList))

		// Data buffers for rendering
		postVertData := f32.Bytes(binary.LittleEndian, -1, -1, 1, -1, -1, 1, 1, 1)
		r.postVertBuffer = gl.CreateBuffer()
		gl.BindBuffer(gl.ARRAY_BUFFER, r.postVertBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, postVertData, gl.STATIC_DRAW)

		r.vertexBuffer = gl.CreateBuffer()

		// Sprite shader
		r.spriteShader = newShaderProgram(vertShader, fragShader, "Main Shader")
		r.spriteShader.RegisterUniforms("modelview", "projection", "x1x2x4x3",
			"alpha", "tint", "mask", "neg", "gray", "add", "mult", "isFlat", "isRgba", "isTrapez")
		r.spriteShader.RegisterTextures("pal", "tex")

		// Compile postprocessing shaders

		// Calculate total amount of shaders loaded.
		r.postShaderSelect = make([]*ShaderProgram, 1+len(sys.externalShaderList))

		// Ident shader (no postprocessing)
		r.postShaderSelect[0] = newShaderProgram(identVertShader, identFragShader, "Identity Postprocess")
		r.postShaderSelect[0].RegisterUniforms("Texture", "TextureSize")

		// External Shaders
		for i := 0; i < len(sys.externalShaderList); i++ {
			r.postShaderSelect[1+i] = newShaderProgram(sys.externalShaders[0][i],
				sys.externalShaders[1][i], fmt.Sprintf("Postprocess Shader #%v", i+1))
			r.postShaderSelect[1+i].RegisterUniforms("Texture", "TextureSize")
		}

		if sys.multisampleAntialiasing {
			gl.Enable(gl.MULTISAMPLE)
		}

		gl.ActiveTexture(gl.TEXTURE0)
		r.fbo_texture = gl.CreateTexture()

		if sys.multisampleAntialiasing {
			gl.BindTexture(gl.TEXTURE_2D_MULTISAMPLE, r.fbo_texture)
		} else {
			gl.BindTexture(gl.TEXTURE_2D, r.fbo_texture)
		}

		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

		if sys.multisampleAntialiasing {
			gl.TexImage2DMultisample(gl.TEXTURE_2D_MULTISAMPLE, 16, gl.RGBA, int(sys.scrrect[2]), int(sys.scrrect[3]), false)
		} else {
			gl.TexImage2D(gl.TEXTURE_2D, 0, int(sys.scrrect[2]), int(sys.scrrect[3]), gl.RGBA, gl.UNSIGNED_BYTE, nil)
		}

		gl.BindTexture(gl.TEXTURE_2D, gl.NoTexture)

		if sys.multisampleAntialiasing {
			r.fbo_f_texture = newTexture(sys.scrrect[2], sys.scrrect[3], 32, false)
			r.fbo_f_texture.SetData(nil)
		} else {
			r.rbo_depth = gl.CreateRenderbuffer()
			gl.BindRenderbuffer(gl.RENDERBUFFER, r.rbo_depth)
			gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT16, int(sys.scrrect[2]), int(sys.scrrect[3]))
			gl.BindRenderbuffer(gl.RENDERBUFFER, gl.NoRenderbuffer)
		}

		r.fbo = gl.CreateFramebuffer()
		gl.BindFramebuffer(gl.FRAMEBUFFER, r.fbo)

		if sys.multisampleAntialiasing {
			gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D_MULTISAMPLE, r.fbo_texture, 0)

			r.fbo_f = gl.CreateFramebuffer()
			gl.BindFramebuffer(gl.FRAMEBUFFER, r.fbo_f)
			gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, r.fbo_f_texture.handle, 0)
		} else {
			gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, r.fbo_texture, 0)
			gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, r.rbo_depth)
		}

		if status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER); status != gl.FRAMEBUFFER_COMPLETE {
			sys.errLog.Printf("framebuffer create failed: 0x%x", status)
		}

		gl.BindFramebuffer(gl.FRAMEBUFFER, gl.NoFramebuffer)
	*/
}
