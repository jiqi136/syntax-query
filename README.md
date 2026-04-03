
---

# Programming Language - Syntax Query V1.2

 [ **English** ]| [ [**中文版**](https://github.com/jiqi136/syntax-query/blob/main/Chinese.md) ] 

**Green**, **Free**, **Open-source**, **No external network connection**, **Local**
---

Supported platforms: Win7 / Win10 / Win11


**How to use the program**

0. **Move to the syntax directory**: Move the `Syntax Query Program` into the folder containing your programming syntax files.

1. **When starting the program**: The program will automatically scan its current folder and up to two subfolder levels deep, searching for all text file names (e.g., `.txt`, `.html`, etc.).

2. **Search for a term**: Enter the word you want to search for in the input box. The program will list all text files containing that word.

3. **Copy file contents**: Click on any file name. The program will automatically copy the entire content of that file to your computer's clipboard. Then, simply press the "Paste" shortcut in your programming editor (e.g., code-writing software) to paste the content.

![Image](https://github.com/jiqi136/syntax-query/blob/main/add/Tu1E.png?raw=true)

---
4. **Customize search scope**: There is another input box at the bottom of the program where you can add additional file extensions to search for (e.g., `.css`, `.py`). The program will then include those file types in its search.

5. **⭐ Frequently Used Terms sidebar on the right**: This area automatically records the terms you've searched for. The more often you use a term, the higher it ranks (sorted from most to least frequent). The most commonly used terms will be right there when you open the program, making it more convenient with repeated use.

---

## 📥Program & Public Source Code [Download Links]:

[github Download ]()


[LanDisk Download https://wwbrl.lanzoum.com/b0fqae41g](https://wwbrl.lanzoum.com/b0fqae41g)  
Password: `1122`




---
This program was built using a desktop GUI framework, completed in a 24-hour sprint.


# PC-Gui: Born for AI, a lightweight desktop GUI framework with native support for Deepseek-like real-time typewriter streaming output! 🎉




> 💡 **Core Philosophy: Rapid Development · Minimal Size · Native Performance · Helping You Build Quality Tools Users Are Willing to Pay For**


On the desktop, user demand for efficient, practical tools has never waned, and the willingness to pay remains strong.
PC-Gui aims to help developers quickly respond to this market need, using the simplest and most stable technologies to build compact yet powerful commercial desktop applications.

---

### Core Tech Stack

Abandoning complex dependencies and bloated third-party GUI libraries, returning to the essence of programming: **building desktop applications with a backend mindset**.
By using a stable Go backend to provide web services and drive a flexible web frontend, we achieve unparalleled lightweight performance.

| Component | Technical Details |
| :--- | :--- |
| **Backend Service** | Go language, using the standard library `net/http` to provide local web services. |
| **Frontend Interface** | Standard web technologies: HTML, JavaScript, CSS. |
| **Data Storage** | Locally encrypted SQLite database, lightweight and reliable. |





### Core Advantages & Multi-solution Comparison

| Category | ✅ PC-Gui Advantages | ⚠️ Other Solutions Comparison |
| :--- | :--- | :--- |
| **🚀 Zero-dependency Execution** | **Backend in Go** enables rapid development, strong typing for easy maintenance; cross-compilation produces a single executable file, requiring no runtime or library installation from users – just double-click to run. | ⚠️ Requires users to pre-install and configure complex environments and dependencies like WebView2, Python, C++, Node.js, etc. |
| **🎨 Interface Technology (HTML)** | **HTML** frontend interfaces can be quickly generated using numerous templates and AI tools, achieving not only high efficiency but also beautiful, modern visual styles with ease. | Traditional GUI libraries often have outdated interfaces that are difficult to customize. |
| **AI Streaming Output** | Simple asynchronous handling is all it takes to achieve streaming output of AI content, enhancing user experience. | Implementing streaming output usually requires complex callbacks or multi-threading. |
| **Markdown Rendering** | Perfectly renders Markdown format returned by AI, with syntax highlighting support for various languages. | Chatbox, Cherry, etc., have relatively basic Markdown rendering and code highlighting effects. |
| **Single-file Deployment** | Using `embed` from the Go standard library, you can package all static resources (like images, CSS, etc.) directly into a single executable file, and reuse the HTML service for direct access. | Relies on bloated dependencies: requires external packaging tools, resulting in large binary sizes or scattered files, making deployment complex. |
| ... | ... | ...|

<a href="https://github.com/jiqi136/PC-Gui" target="_blank">GUI Framework Public Source Code Download (github): https://github.com/jiqi136/PC-Gui</a>
