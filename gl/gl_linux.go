package gl

// Auto generated, do not edit.

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

  
	gl "github.com/go-gl/gl/v3.0/gles2"
  
	"github.com/pigfall/gosdk/math"
)

func Init() error {
	return gl.Init()
}

func GLCreateProgram() (GLProgram, error) {
	program := gl.CreateProgram()
	if program == 0 {
		return 0, fmt.Errorf("create gl program failed")
	}

	return GLProgram(program), nil
}

func GLLinkProgram(program GLProgram, vertexShader GLShader, fragmentShader GLShader) error {
	prog := program.ToGLType()
	gl.AttachShader(prog, vertexShader.ToGLType())
	gl.AttachShader(prog, fragmentShader.ToGLType())
	gl.LinkProgram(prog)

	var linkStatus int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &linkStatus)
	if linkStatus == gl.TRUE {
		return nil
	}

	return fmt.Errorf("link gl program error: %s", GLGetProgramInfoLog(program))
}

func GLGetProgramiv(program GLProgram, paramName GLProgramInfoType) int32 {
	var out int32
	gl.GetProgramiv(program.ToGLType(), paramName.ToGLType(), &out)
	return out
}

func GLGetProgramInfoLog(program GLProgram) string {
	logLen := GLGetProgramiv(program, GLProgramInfoType_LogLength)
	if logLen == 0 {
		return ""
	}
	buffer := make([]byte, logLen)
	var actualLogLen int32
	gl.GetProgramInfoLog(program.ToGLType(), logLen, &actualLogLen, &buffer[0])

	return string(buffer[:actualLogLen])
}

func GLCreateShader(shaderType GLShaderType) (GLShader, error) {
	shader := gl.CreateShader(shaderType.ToGLType())
	if shader == 0 {
		return 0, fmt.Errorf("create shader failed")
	}

	return GLShader(shader), nil
}

func GLCompileProgram(vertexShaderSource string, fragmentShaderSource string) (vertexShader GLShader, fragmentShader GLShader, program GLProgram, err error) {
	vertexShader, err = GLCreateShader(GLShaderType_Vertex)
	if err != nil {
		err = fmt.Errorf("create vertex shader error: %w", err)
		return
	}
	onErrF := []func(){func() { gl.DeleteShader(vertexShader.ToGLType()) }}
	defer func() {
		if err != nil {
			for _, f := range onErrF {
				f()
			}
		}
	}()

	fragmentShader, err = GLCreateShader(GLShaderType_Fragment)
	if err != nil {
		gl.DeleteShader(vertexShader.ToGLType())
		err = fmt.Errorf("create fragment shader error: %w", err)
		return
	}
	onErrF = append(onErrF, func() { gl.DeleteShader(fragmentShader.ToGLType()) })

	if err = GLCompileShader(vertexShader, vertexShaderSource); err != nil {
		err = fmt.Errorf("compile vertex shader error: %w", err)
		return
	}

	if err = GLCompileShader(fragmentShader, fragmentShaderSource); err != nil {
		err = fmt.Errorf("compile fragment shader error: %w", err)
		return
	}

	program, err = GLCreateProgram()
	if err != nil {
		return
	}
	onErrF = append(onErrF, func() { gl.DeleteProgram(program.ToGLType()) })

	if err = GLLinkProgram(program, vertexShader, fragmentShader); err != nil {
		return
	}

	return
}

func GLCompileShader(shader GLShader, source string) error {
	GLShaderSource(shader, source)
	gl.CompileShader(shader.ToGLType())
	if GLGetShaderiv(shader, GLShaderInfoType_CompileStatus) != gl.TRUE { // Compile failed
		return fmt.Errorf("compile shader error: %s", GLGetShaderInfoLog(shader))
	}
	return nil
}

func GLShaderSource(shader GLShader, source string) {
	var length = int32(len(source))
	csource, free := gl.Strs(source)
	defer free()
	gl.ShaderSource(uint32(shader), 1, csource, &length)
}

func GLGetShaderiv(shader GLShader, paramType GLShaderInfoType) int32 {
	var out int32
	gl.GetShaderiv(shader.ToGLType(), paramType.ToGLType(), &out)
	return out
}

func GLGetShaderInfoLog(shader GLShader) string {
	logLength := GLGetShaderiv(shader, GLShaderInfoType_LogLength)
	if logLength == 0 {
		return ""
	}
	buffer := make([]byte, logLength)
	var outLogLengh int32
	gl.GetShaderInfoLog(shader.ToGLType(), logLength, &outLogLengh, &buffer[0])

	return string(buffer[:outLogLengh])
}

// GLGenBuffer returns a new GLBuffer object and an error if the buffer cannot be created.
func GLGenBuffer() (GLBuffer, error) {
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	if buffer == 0 {
		return 0, fmt.Errorf("create gl buffer failed")
	}

	return GLBuffer(buffer), nil
}

func MustGenBuffer()GLBuffer{
  buffer,err:=GLGenBuffer()
  if err != nil{
    panic(err)
  }
  return buffer
}

func GLBindBuffer(target GLBindBufferTarget, buffer GLBuffer) {
	gl.BindBuffer(target.ToGLType(), buffer.ToGLType())
}

func GLVertexAttribPointer(index uint32, vertexDimensions int32, primitiveType GLPrimitiveType, stride int32, offset int32) {
	gl.VertexAttribPointerWithOffset(uint32(index), vertexDimensions, gl.FLOAT, false, int32(stride), uintptr(offset))
}

func GLEnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}

func GLDrawArrays(mode GLDrawArraysMode, first, count int32) {
	gl.DrawArrays(mode.ToGLType(), first, count)
}

func GLBufferData(target GLBindBufferTarget, size int, data []byte, usage GLBufferDataUsage) {
	var dataPtr unsafe.Pointer
	if len(data) == 0 {
		dataPtr = unsafe.Pointer(nil)
	} else {
		dataPtr = unsafe.Pointer(&data[0])
	}
	gl.BufferData(target.ToGLType(), size, dataPtr, usage.ToGLType())
}

func GLBufferDataWithPointer(target GLBindBufferTarget, size int, dataPtr unsafe.Pointer, usage GLBufferDataUsage) {
	gl.BufferData(target.ToGLType(), size, dataPtr, usage.ToGLType())
}

func GLClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func GLClear(mask GLClearMask) {
	gl.Clear(mask.ToGLType())
}

func GLUseProgram(program GLProgram) {
	gl.UseProgram(program.ToGLType())
}

func GLBufferSubData(target GLBindBufferTarget, offset int, data []byte) {
	var dataPtr unsafe.Pointer
	if len(data) == 0 {
		dataPtr = unsafe.Pointer(nil)
	} else {
		dataPtr = unsafe.Pointer(&data[0])
	}
	gl.BufferSubData(target.ToGLType(), offset, len(data), dataPtr)
}

func GLGetUniformLocation(program GLProgram, name string) (int32, error) {
	if !strings.HasSuffix(name, "\x00") {
		name += "\x00"
	}
	location := gl.GetUniformLocation(program.ToGLType(), gl.Str(name))
	if location < 0 {
		return 0, fmt.Errorf("get uniform location failed")
	}

	return location, nil
}

func GLUniformMatrix4fv(location int32, matrix *math.Matrix4) {
	gl.UniformMatrix4fv(location, 1, false, &matrix.Values[0])
}

func GLUniform1i(location int32,value int32){
  gl.Uniform1i(location,value)
}

func GLUniform4f(location int32,v1,v2,v3,v4 float32){
  gl.Uniform4f(location,v1,v2,v3,v4)
}

func GLViewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func GLGenTexture() (GLTexture, error) {
	var texture uint32
	gl.GenTextures(1, &texture)
	if texture <= 0 {
		return 0, fmt.Errorf("create gl texture failed")
	}

	return GLTexture(texture), nil
}

func GLBindTexture(target GLBindTextureTarget, texture GLTexture) {
	gl.BindTexture(target.ToGLType(), texture.ToGLType())
}

func GLActiveTexture(target GLActiveTextureTarget){
  gl.ActiveTexture(target.ToGLType())
}

func GLTexImage2D(target GLBindTextureTarget, level int32, internalFormat GLInternalColorFormat, width, height int32, border int32, sourceImageFormat GLSourceImageColorFormat, pixelDataType GLPrimitiveType, pixels []byte) {
	gl.TexImage2D(
		target.ToGLType(),
		level,
		internalFormat.ToGLType(),
		width,
		height,
		border,
		sourceImageFormat.ToGLType(),
		pixelDataType.ToGLType(),
		unsafe.Pointer(&pixels[0]),
	)
}

func GLTexParameteri(target GLBindTextureTarget, paramName GLTextureParameterName, paramValue int32) {
	gl.TexParameteri(target.ToGLType(), paramName.ToGLType(), paramValue)
}

func GLTexParamterWrapS(target GLBindTextureTarget, value GLTextureParamterWrapS) {
	GLTexParameteri(target, GLTextureParamterName_Texture_Wrap_S, value.ToGLType())
}

func GLTexParamterWrapT(target GLBindTextureTarget, value GLTextureParamterWrapT) {
	GLTexParameteri(target, GLTextureParamterName_Texture_Wrap_S, value.ToGLType())
}

func GLTexParamterMinFilter(target GLBindTextureTarget, value GLTextureParamterMinFilter) {
	GLTexParameteri(target, GLTextureParamterName_Texture_Min_Filter, value.ToGLType())
}

func GLTexParamterMagFilter(target GLBindTextureTarget, value GLTextureParamterMagFilter) {
	GLTexParameteri(target, GLTextureParamterName_Texture_Mag_Filter, value.ToGLType())
}

func GLEnable(cap GLCapability) {
	gl.Enable(cap.ToGLType())
}

func GLDisable(cap GLCapability) {
	gl.Disable(cap.ToGLType())
}

func GLBlendFunc(s, d GLBlendFunction) {
	gl.BlendFunc(s.ToGLType(), d.ToGLType())
}

func GLDrawElements(mode GLDrawArraysMode, count int32, offset int32) {
	gl.DrawElementsWithOffset(mode.ToGLType(), count, gl.UNSIGNED_INT, uintptr(offset))
}

func GLStencilMask(v uint32) {
	gl.StencilMask(v)
}

func GLStencilFunc(f GLStencilFuncEnum, ref int32, mask uint32) {
	gl.StencilFunc(f.ToGLType(), ref, mask)
}

func GLStencilOp(fail, zfail, zpass GLStencilOpEnum) {
	gl.StencilOp(fail.ToGLType(), zfail.ToGLType(), zpass.ToGLType())
}

func GLColorMask(r, g, b, a bool) {
	gl.ColorMask(r, g, b, a)
}

// Create a vao.
func GLGenVertexArray() (GLVertexArray, error) {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	if vao > 0 {
		return GLVertexArray(vao), nil
	}

	return 0, errors.New("gen gl vertex array failed")
}

// Delete vertex buffer object.
func GLDeleteBuffer(buffer GLBuffer) {
	b := buffer.ToGLType()
	gl.DeleteBuffers(1, &b)
}

// Delete vertex array object.
func GLDeleteArray(vao GLVertexArray) {
	b := vao.ToGLType()
	gl.DeleteVertexArrays(1, &b)
}

func GLBindVertexArray(vao GLVertexArray) {
	gl.BindVertexArray(vao.ToGLType())
}

func GLVersion() string{
  return gl.GoStr(gl.GetString(gl.VERSION))
}

func GLShaderLanguageVersion() string {
  return gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
}

func GLGetError() error {
	code := gl.GetError()
	if code != gl.NO_ERROR {
		return fmt.Errorf("gl error code: %v", code)
	}

	return nil
}

func GLDeleteTexture(texture GLTexture) {
	t := texture.ToGLType()
	gl.DeleteTextures(1, &t)
}

func GLCleanStencil(value int32){
  gl.ClearStencil(value)
}
