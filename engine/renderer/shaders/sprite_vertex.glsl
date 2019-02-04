//
// Vertex shader for sprites
//

precision highp float;

#include <attributes>

// Input uniforms
uniform mat4 MVP;

#include <material>

// Outputs for fragment shader
varying vec3 Color;
varying vec2 FragTexcoord;

void main() {

    // Applies transformation to vertex position
    gl_Position = MVP * vec4(VertexPosition, 1.0);

    // Outputs color
    Color = MatDiffuseColor;

    // Flips texture coordinate Y if requested.
    vec2 texcoord = VertexTexcoord;
#if MAT_TEXTURES>0
    if (MatTexFlipY[0]) {
        texcoord.y = 1.0 - texcoord.y;
    }
#endif
    FragTexcoord = texcoord;
}

