# 🎨 Shader Sketchbook

A sandbox for experimenting with real-time 2D shaders using [Ebitengine](https://ebiten.org/) and the [Kage](https://ebiten.org/documents/shader.html) shader language. This project supports animated fragment shaders, runtime switching, and interactive input-driven effects.

---

## ✨ Features

* 🔁 Live shader switching (`Tab`)
* ⏱️ Time-based animations (`Time` uniform)
* 🎮 Input-responsive shaders (mouse & key toggles)
* 📦 Embeds all `.kage` shaders via `embed.FS`
* 🧪 Experimental playground for visual effects
* 🖥️ Fullscreen-compatible render pipeline

---

## 🕹 Controls

| Key     | Action                                 |
| ------- | -------------------------------------- |
| `Tab`   | Cycle through loaded shaders           |
| `Space` | Toggle uniform `Toggle` (1.0 / 0.0)    |
| Mouse   | Tracked via `Mouse` uniform (`[x, y]`) |

---

## 🔧 Shader Uniforms Provided

```glsl
var (
    Time   float     // seconds since start
    Mouse  vec2      // pixel coordinates
    Toggle float     // 0.0 or 1.0 based on key press
)
```

---

## 🧪 Adding Shaders

1. Save a `.kage` file in `assets/shaders/`
2. Rebuild the project — it will be auto-discovered at runtime
3. Use `image/color` effects, math animations, or input interactions

Example:

```glsl
package main

var Time float

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
    pulse := sin(Time * 2.0)
    return vec4(texCoord.x * pulse, texCoord.y * pulse, 0.5, 1.0)
}
```

---

## 🚀 Getting Started

```bash
go run main.go
```

Requires Go 1.20+ and Ebitengine.

---

## ✅ Completed Milestones

* [x] Shader switching system
* [x] Time-driven animation support
* [x] Mouse & key uniforms
* [x] Shader embedding via `embed.FS`
* [x] Shader testing loop and debug overlay

---