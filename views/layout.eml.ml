let render body =
  <html>
  <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="Rawley.xyz, Rawley Fowler's blog and personal website">
  <title>rawley.xyz</title>
  <link href="/static/index.css" rel="stylesheet">
  </head>
  <body>
  <header>
  <h1>
  <a href="/">Rawley.xyz</a>
  </h1>
  <nav>
  <li>
  <a href="/blog">blog</a>
  </li>
  <li>
  <a href="/resume">resume</a>
  </li>
  <li>
  <a href="/philosophy">philosophy</a>
  </li>
  <li>
  <a href="/web-ring">web ring</a>
  </li>
  <li>
  <a href="/bookshelf">bookshelf</a>
  </li>

  </nav>
  </header>
  <main>
  <%s! body %>
  </main>
  <footer>
  <hr>
  <p>
  Copyright &copy; Rawley Fowler 2022
  </p>
  <p>
  <b>Disclaimer</b>: All opinions on this site, are that of my own. They do not reflect the opinions of
  any of my employers; past, present, or future.
  </p>
  </footer>
  </body>
  </html>
