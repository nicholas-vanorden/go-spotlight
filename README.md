# go-spotlight

A simple app written in **Go** using the **Fyne** UI framework, for adding highlights to images.

---

- Load an image (JPG or PNG)
- Draw yellow highlights on top of the image
- Save the highlighted image back to disk

The goal of this project is to provide a lightweight, easy-to-understand example of a Go + Fyne graphics application, while also being genuinely useful for quick image markup.

---

## Features

- 🖼 Load JPG and PNG images
- ✏️ Draw freeform yellow highlights
- 💾 Save the edited image
- 🧩 Built with Go and Fyne
- 🖥 Cross-platform (macOS, Windows, Linux)

---

## Requirements

- Go 1.21+ (recommended)
- Fyne v2

---

## Installation

Clone the repository:

```bash
git clone https://github.com/your-username/go-spotlight.git
cd go-spotlight
```

Install dependencies:

```bash
go mod tidy
```

Run the app:

```bash
go run .
```

---

## Usage

1. Launch the app
2. Click **Open Image** and select a JPG or PNG file
3. Use your mouse to draw yellow highlights over the image
4. Click **Save Image** to export the highlighted version

---

## Built With

- [Go](https://golang.org/)
- [Fyne](https://fyne.io/) – Cross-platform GUI framework for Go

---

## Limitations / Notes

- Highlight color is currently fixed to yellow
- No undo/redo support (yet)
- Intended as a simple utility and learning project, not a full image editor

---

## Roadmap (Ideas)

- Adjustable highlight color and thickness
- Undo / redo support
- Zoom and pan controls
- Export to additional formats

---

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

---

If you find this useful or are learning Go + Fyne, ⭐️ the repo!
