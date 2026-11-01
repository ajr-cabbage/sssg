package initialize

const HTMLTEMPLATE string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ Title }}</title>
  <link rel="stylesheet" href="/assets/index.css">
</head>
<body>
  <header class="title-bar">
      <div class="title-bar__nav">
          <a href="/">{{ Site Name }}</a>
      </div>
  </header>
  <div class="layout">
    <aside class="sidebar">
      <ul class="sidebar__list">{{ Sidebar }}</ul>
    </aside>
    <main class="content">
      {{ Content }}
    </main>
  </div>

  <footer class="footer">
    <p>&copy; 2026 Small Shitty Software Co. All rights reserved.</p>
  </footer>

</body>
</html>
`

const CSSDEFAULT string = `/* Shitty Default Theme */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

:root {
  /* Solarized palette */
  --sz-bg: #58594f;
  --sz-surface: #41423b;
  --sz-overlay: #586e75;
  --sz-text: #bfa1ff;
  --sz-muted: #9983c9;
  --sz-red: #dc322f;
  --sz-green: #859900;
  --sz-yellow: #b58900;
  --sz-blue: #268bd2;
  --sz-violet: #6c71c4;
  --sz-cyan: #2aa198;

  --color-bg: var(--sz-bg);
  --color-surface: var(--sz-surface);
  --color-overlay: var(--sz-overlay);
  --color-text: var(--sz-text);
  --color-text-muted: var(--sz-muted);
  --color-accent: var(--sz-blue);
  --color-accent-alt: var(--sz-violet);
  --color-border: var(--sz-overlay);
  --title-bar-height: 56px;
}

body {
  font-family: 'Zed Mono', 'Lilex', 'IBM Plex Mono', monospace;
  background-color: var(--color-bg);
  color: var(--color-text);
  line-height: 1.6;
}

a {
  color: var(--color-accent);
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

/* Title Bar */
.title-bar {
  position: sticky;
  top: 0;
  z-index: 100;
  height: var(--title-bar-height);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 1.5rem;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.35);
}

.title-bar__title {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--color-text);
}

.title-bar__nav {
  font-size: 1.25rem;
  display: flex;
  gap: 1.25rem;
}

.title-bar__nav a {
  color: var(--color-text-muted);
  font-weight: 500;
}

.title-bar__nav a:hover {
  color: var(--color-accent);
  text-decoration: none;
}

/* Layout: sidebar + content */
.layout {
  display: flex;
  align-items: flex-start;
  max-width: 1100px;
  margin: 1.5rem auto;
  gap: 1.5rem;
  padding: 0 1rem;
}

/* Sidebar */
.sidebar {
  position: sticky;
  top: 0; /* sits below the sticky title bar */
  flex: 0 0 250px;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  padding: 1rem;
  overflow-y: visible;
}

.sidebar__heading {
  font-size: 1rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  color: var(--color-text-muted);
  margin-bottom: 0.75rem;
}

.sidebar__list {
  list-style: none;
}

.sidebar__list li {
  padding: 0.25rem 0.25rem;
}

.sidebar__list a {
  display: block;
  padding: 0.25rem 0.25rem;
  border-radius: 4px;
  color: var(--color-text);
}

.sidebar__list a:hover {
  background-color: var(--color-overlay);
  color: var(--color-accent);
  text-decoration: none;
}

/* Main content */
.content {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.content__section {
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: 1.25rem 1.5rem;
}

.content__section h2 {
  font-size: 1.15rem;
  margin-bottom: 0.5rem;
  color: var(--color-accent-alt);
}

/* Footer */
.footer {
  text-align: center;
  padding: 1.5rem;
  color: var(--color-text-muted);
  font-size: 0.875rem;
}

/* Responsive: stack sidebar above content on narrow screens */
@media (max-width: 720px) {
  .layout {
    flex-direction: column;
  }

  .sidebar {
    position: static;
    flex: none;
    width: 100%;
    max-height: none;
  }
}
`
