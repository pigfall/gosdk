package gl

// Auto generated, do not edit.

// Shader type: vertex or fragment.
type GLShaderType uint32

// GL shader object.
type GLShader uint32

// GL program.
type GLProgram uint32

// GL buffer object. Created from GLGenBuffer().
type GLBuffer uint32

// VAO. Created from GLGenVertexArray()
type GLVertexArray uint32

// GL buffer target. Eg: ARRAY_BUFFER
type GLBindBufferTarget uint32

// GL primitive data type. Eg: float
type GLPrimitiveType uint32

// The param name for GLGetProgramiv()
type GLProgramInfoType uint32

// The param name for GLGetShaderiv()
type GLShaderInfoType uint32

type GLDrawArraysMode uint32

type GLBufferDataUsage uint32

type GLClearMask uint32

type GLTexture uint32

type GLBindTextureTarget uint32
type GLActiveTextureTarget uint32

type GLInternalColorFormat int32

type GLSourceImageColorFormat uint32

type GLTextureParameterName uint32

type GLTextureParamterWrapS int32
type GLTextureParamterWrapT int32

// For GLEnable.
type GLCapability uint32

// Minify
type GLTextureParamterMinFilter int32

// Magnify
type GLTextureParamterMagFilter int32

// For GLBlendFunc.
type GLBlendFunction uint32

// For GLStencilFunc
type GLStencilFuncEnum uint32

// For GLStencilOp
type GLStencilOpEnum uint32

func (p GLProgram) ToGLType() uint32 {
	return uint32(p)
}

func (p GLProgramInfoType) ToGLType() uint32 {
	return uint32(p)
}

func (s GLShader) ToGLType() uint32 {
	return uint32(s)
}

func (s GLShaderType) ToGLType() uint32 {
	return uint32(s)
}

func (s GLShaderInfoType) ToGLType() uint32 {
	return uint32(s)
}

func (b GLBuffer) ToGLType() uint32 {
	return uint32(b)
}

func (b GLBindBufferTarget) ToGLType() uint32 {
	return uint32(b)
}

func (p GLBufferDataUsage) ToGLType() uint32 {
	return uint32(p)
}

func (p GLDrawArraysMode) ToGLType() uint32 {
	return uint32(p)
}

func (p GLClearMask) ToGLType() uint32 {
	return uint32(p)
}

func (t GLBindTextureTarget) ToGLType() uint32 {
	return uint32(t)
}

func (t GLActiveTextureTarget) ToGLType() uint32 {
	return uint32(t)
}

func (t GLTexture) ToGLType() uint32 {
	return uint32(t)
}

func (f GLInternalColorFormat) ToGLType() int32 {
	return int32(f)
}

func (f GLSourceImageColorFormat) ToGLType() uint32 {
	return uint32(f)
}

func (t GLPrimitiveType) ToGLType() uint32 {
	return uint32(t)
}

func (t GLTextureParameterName) ToGLType() uint32 {
	return uint32(t)
}

func (t GLTextureParamterWrapS) ToGLType() int32 {
	return int32(t)
}

func (t GLTextureParamterWrapT) ToGLType() int32 {
	return int32(t)
}

func (t GLTextureParamterMagFilter) ToGLType() int32 {
	return int32(t)
}

func (t GLTextureParamterMinFilter) ToGLType() int32 {
	return int32(t)
}

func (c GLCapability) ToGLType() uint32 {
	return uint32(c)
}

func (f GLBlendFunction) ToGLType() uint32 {
	return uint32(f)
}

func (f GLStencilFuncEnum) ToGLType() uint32 {
	return uint32(f)
}

func (f GLStencilOpEnum) ToGLType() uint32 {
	return uint32(f)
}

func (a GLVertexArray) ToGLType() uint32 {
	return uint32(a)
}
