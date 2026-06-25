package gl

// Auto generated, do not edit.

import (
	gl "github.com/go-gl/gl/v3.0/gles2"
)

const (
	GLProgramInfoType_LogLength GLProgramInfoType = gl.INFO_LOG_LENGTH

	GLShaderInfoType_CompileStatus GLShaderInfoType = gl.COMPILE_STATUS
	GLShaderInfoType_LogLength     GLShaderInfoType = gl.INFO_LOG_LENGTH

	GLShaderType_Vertex   = gl.VERTEX_SHADER
	GLShaderType_Fragment = gl.FRAGMENT_SHADER

	GLBindBufferTarget_ArrayBuffer        GLBindBufferTarget = gl.ARRAY_BUFFER
	GLBindBufferTarget_ElementArrayBuffer GLBindBufferTarget = gl.ELEMENT_ARRAY_BUFFER

	GLPrimitiveType_Float32      GLPrimitiveType = gl.FLOAT
	GLPrimitiveType_UnsignedByte GLPrimitiveType = gl.UNSIGNED_BYTE
	GLPrimitiveType_UnsignedInt  GLPrimitiveType = gl.UNSIGNED_INT

	GLDrawArraysMode_Triangles     GLDrawArraysMode = gl.TRIANGLES
	GLDrawArraysMode_Triangles_Fan GLDrawArraysMode = gl.TRIANGLE_FAN
	GLDrawArraysMode_Line_Loop     GLDrawArraysMode = gl.LINE_LOOP

	GLBufferDataUsage_StaticDraw  GLBufferDataUsage = gl.STATIC_DRAW
	GLBufferDataUsage_DynamicDraw GLBufferDataUsage = gl.DYNAMIC_DRAW

	GLClearMask_ColorBuffer   GLClearMask = gl.COLOR_BUFFER_BIT
	GLClearMask_DepthBuffer   GLClearMask = gl.DEPTH_BUFFER_BIT
	GLClearMask_StencilBuffer GLClearMask = gl.STENCIL_BUFFER_BIT

	GLBindTextureTarget_Texture2D  GLBindTextureTarget   = gl.TEXTURE_2D
	GLActiveTextureTarget_Texture0 GLActiveTextureTarget = gl.TEXTURE0

	// GLTextureParamterName
	GLTextureParamterName_Texture_Wrap_T     GLTextureParameterName = gl.TEXTURE_WRAP_T
	GLTextureParamterName_Texture_Wrap_S     GLTextureParameterName = gl.TEXTURE_WRAP_S
	GLTextureParamterName_Texture_Min_Filter GLTextureParameterName = gl.TEXTURE_MIN_FILTER
	GLTextureParamterName_Texture_Mag_Filter GLTextureParameterName = gl.TEXTURE_MAG_FILTER

	// GLTextureParamterWrapS
	GLTextureParamterWrapS_Repeat GLTextureParamterWrapS = gl.REPEAT
	// GLTextureParamterWrapT
	GLTextureParamterWrapT_Repeat      GLTextureParamterWrapT     = gl.REPEAT
	GLTextureParamter_MinFilter_Linear GLTextureParamterMinFilter = gl.LINEAR
	GLTextureParamter_MagFilter_Linear GLTextureParamterMagFilter = gl.LINEAR

	// Texture Internal Format.
	GLInternalColorFormat_RGBA GLInternalColorFormat = gl.RGBA

	// Image Color Format.
	GLSourceImageColorFormat_RGBA GLSourceImageColorFormat = gl.RGBA

	GLCapability_Blend               GLCapability    = gl.BLEND
	GLCapability_DEPTH_TEST          GLCapability    = gl.DEPTH_TEST
	GLCapability_StencilTest         GLCapability    = gl.STENCIL_TEST
	GLCapability_CULL_FACE           GLCapability    = gl.CULL_FACE
	GLBlendFunction_SrcAlpha         GLBlendFunction = gl.SRC_ALPHA
	GLBlendFunction_OneMinusSrcAlpha GLBlendFunction = gl.ONE_MINUS_SRC_ALPHA

	// GLStencilFunc
	GLStencilFunc_Equal  GLStencilFuncEnum = gl.EQUAL
	GLStencilFunc_Always GLStencilFuncEnum = gl.ALWAYS

	// GLStencilOp
	GLStencilOp_Keep GLStencilOpEnum = gl.KEEP
	GLStencilOp_Incr GLStencilOpEnum = gl.INCR
)
